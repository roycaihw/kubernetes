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

echo "Waiting for Kubelet to fully stabilize..."
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
  # NOTE ON PSI MEMORY STALLS AND SWAPLESS ENVIRONMENTS:
  # There is no off-the-shelf way to generate Memory PSI (stalls) on a stock, swapless
  # Linux system. If you just allocate massive amounts of pure memory (e.g. without 
  # any backing disk), the kernel will instantly panic and invoke the OOM Killer when
  # physical RAM is full. An instant OOM kill evaluates to a 0% PSI stall because 
  # the process simply dies instead of waiting/stalling for memory to become available.
  # 
  # To generate Kubelet's required `avg60 > 0.9` stall natively, the kernel must be 
  # able to pause the application and thrash the disk (via Page Cache Eviction or 
  # Swapping). This is why Experiment 1 requires the explicit creation of a 4G 
  # /swapfile to safely bottleneck the container into a memory stall, preventing 
  # premature OOM termination and allowing the metrics to surface accurately.
  - name: stress
    image: alexeiled/stress-ng
    args: ["--vm", "1", "--vm-bytes", "17G", "--vm-hang", "0"]
POD
echo "Workloads deployed."
EOF

echo "Uploading deployment script to VM..."
gcloud compute scp /tmp/03-deploy-workloads-remote.sh "$VM_NAME:/tmp/03-deploy-workloads-remote.sh" --zone="$ZONE"

echo "Executing deployment script on VM..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="bash /tmp/03-deploy-workloads-remote.sh"
