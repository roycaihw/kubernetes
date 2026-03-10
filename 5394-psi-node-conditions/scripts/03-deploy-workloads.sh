#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"

echo "Generating remote deployment payload..."
cat << 'EOF' > /tmp/03-deploy-workloads-remote.sh
#!/bin/bash
set -euo pipefail

export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
export PATH=$PATH:$(pwd)/kubernetes/_output/local/bin/linux/amd64/

echo "Waiting for kubectl compilation to finish..."
for i in {1..120}; do
  if command -v kubectl >/dev/null 2>&1; then
    break
  fi
  sleep 5
done

echo "Waiting for node to become Ready..."
NODE_READY=false
for i in {1..120}; do
  if kubectl get nodes 2>/dev/null | grep -q ' Ready'; then
    echo "Node is Ready!"
    NODE_READY=true
    break
  fi
  sleep 5
done

if [ "$NODE_READY" = false ]; then
  echo "ERROR: Node never became Ready. Cluster likely failed to start!"
  cat cluster.log || true
  exit 1
fi

echo "Verifying Kubelet feature gates are active..."
if ! ps -eo args | grep "[k]ubelet " | grep "PSINodeCondition=true" > /dev/null; then
  echo "ERROR: Kubelet does not have PSINodeCondition=true in its arguments!"
  exit 1
fi
echo "Kubelet is running with PSINodeCondition feature gate enabled."

echo "Lowering SystemMemoryContentionThreshold for automation verification..."
for f in /tmp/local-up-cluster.sh.*/kubelet.yaml; do
  # Avoid compounding appending if run multiple times
  sed -i '/systemMemoryContentionThreshold/d' $f
  echo 'systemMemoryContentionThreshold: 0.001' >> $f
done
KUBELET_CMD=$(ps -eo args | grep "[k]ubelet " | grep config | grep -v sudo | head -n 1)
sudo pkill -9 kubelet || true
sudo bash -c "nohup $KUBELET_CMD > /tmp/kubelet-restart.log 2>&1 &"
echo "Waiting for Kubelet to restart..."
sleep 15

echo "Deleting old stress pod..."
kubectl delete pod memory-stressor --ignore-not-found
kubectl delete pod scheduling-test-pod --ignore-not-found

echo "Deploying baseline NGINX deployment..."
kubectl create deployment baseline-nginx --image=nginx || true

echo "Deploying memory stressor pod..."
cat <<POD | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: memory-stressor
spec:
  containers:
  - name: stress
    image: polinux/stress
    command: ["stress"]
    args: ["--vm", "2", "--vm-bytes", "7G"]
POD
echo "Workloads deployed."
EOF

echo "Uploading deployment script to VM..."
gcloud compute scp /tmp/03-deploy-workloads-remote.sh "$VM_NAME:/tmp/03-deploy-workloads-remote.sh" --zone="$ZONE"

echo "Executing deployment script on VM..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="bash /tmp/03-deploy-workloads-remote.sh"
