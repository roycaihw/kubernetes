#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"

echo "[1/2] Verifying cluster API server is reachable..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
  export PATH=\$PATH:\$(pwd)/kubernetes/_output/local/bin/linux/amd64/
  
  echo 'Waiting for node to become Ready...'
  for i in {1..30}; do
    if kubectl get nodes | grep -q ' Ready'; then
      echo 'Node is Ready!'
      break
    fi
    sleep 5
  done
"

echo "[2/2] Deploying workloads..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
  export PATH=\$PATH:\$(pwd)/kubernetes/_output/local/bin/linux/amd64/

  echo 'Deploying baseline NGINX deployment...'
  kubectl create deployment baseline-nginx --image=nginx || true

  echo 'Deploying memory stressor pod...'
  cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: memory-stressor
spec:
  containers:
  - name: stress-ng
    image: alexeiled/stress-ng
    args: [\"--vm\", \"2\", \"--vm-bytes\", \"90%\", \"--vm-method\", \"all\", \"--verify\", \"-t\", \"10m\"]
EOF
  echo 'Workloads deployed. Watch node conditions natively via SSH using:'
  echo 'gcloud compute ssh $VM_NAME --zone=$ZONE --command=\"export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig && export PATH=\\\$PATH:\\\$(pwd)/kubernetes/_output/local/bin/linux/amd64/ && watch -n 1 \\\"kubectl describe node | grep -A 5 \'SystemMemoryContentionPressure\'\\\"\"'
"
