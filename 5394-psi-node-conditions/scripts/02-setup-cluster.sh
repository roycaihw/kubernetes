#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"
REPO_URL="https://github.com/roycaihw/kubernetes.git"
BRANCH_NAME="${1:-psi-dev-136}"

echo "[1/3] Installing dependencies..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  sudo apt-get update && \
  sudo DEBIAN_FRONTEND=noninteractive apt-get install -y build-essential git jq make wget docker.io && \
  sudo usermod -aG docker \$USER && \
  wget -q https://go.dev/dl/go1.22.4.linux-amd64.tar.gz && \
  sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz && \
  if ! grep -q '/usr/local/go/bin' ~/.bashrc; then
      echo 'export PATH=\$PATH:/usr/local/go/bin' >> ~/.bashrc
  fi
"

echo "[2/3] Cloning repository and checking out branch: $BRANCH_NAME..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  if [ ! -d \"kubernetes\" ]; then
    git clone $REPO_URL kubernetes
  fi
  cd kubernetes && \
  git fetch --all && \
  git checkout $BRANCH_NAME
"

echo "[3/3] Starting local-up-cluster.sh in the background..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  source ~/.bashrc
  cd kubernetes
  export FEATURE_GATES='PSINodeCondition=true'
  export KUBELET_FLAGS='--feature-gates=PSINodeCondition=true'
  nohup hack/local-up-cluster.sh > cluster.log 2>&1 &
"
echo "Cluster is now provisioning in the background on $VM_NAME."
echo "You can check the logs by running: gcloud compute ssh $VM_NAME --zone=$ZONE --command='tail -f ~/kubernetes/cluster.log'"
