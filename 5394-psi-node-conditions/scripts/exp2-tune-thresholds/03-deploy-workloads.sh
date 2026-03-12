#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"
PHASE="${1:-1}"

echo "Generating remote workload payload for Phase ${PHASE}..."
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

echo "Cleaning up old workloads..."
kubectl delete pod --all --ignore-not-found || true

if [ "$1" == "1" ]; then
  echo 'Deploying Phase 1 (False Positive) safe, heavy CPU burst...'
  cat <<POD | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: false-positive-stressor
spec:
  containers:
  - name: stress
    image: alexeiled/stress-ng
    args: ["--cpu", "4", "--vm", "2", "--vm-bytes", "1G"]
POD

elif [ "$1" == "2" ]; then
  echo 'Deploying Phase 2 (Lead Time) slow memory leak...'
  cat <<POD | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: slow-leak-stressor
spec:
  containers:
  - name: stress
    image: alexeiled/stress-ng
    # Slowly scale up memory to gently hit swap, simulating an organic memory leak
    # to measure the exact lead warning time.
    args: ["--vm", "1", "--vm-bytes", "17G", "--vm-stride", "4K", "--vm-keep"]
POD

elif [ "$1" == "3" ]; then
  echo 'Deploying Phase 3 (Kubepods Isolated) slow memory leak...'
  cat <<POD | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: kubepods-leak-stressor
spec:
  containers:
  - name: stress
    image: alexeiled/stress-ng
    # Force organic thrashing exclusively bounded within the Pod's cgroup memory limits.
    args: ["--vm", "1", "--vm-bytes", "2500M", "--vm-stride", "4K", "--vm-keep"]
    resources:
      limits:
        memory: "2Gi"
POD
fi

echo "Workload for Phase $1 deployed."
EOF

echo "Uploading script to VM..."
gcloud compute scp /tmp/03-deploy-workloads-remote.sh "$VM_NAME:/tmp/03-deploy-workloads-remote.sh" --zone="$ZONE"

echo "Executing script on VM..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="bash /tmp/03-deploy-workloads-remote.sh $PHASE"
