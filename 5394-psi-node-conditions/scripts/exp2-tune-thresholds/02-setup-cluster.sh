#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"
REPO_URL="https://github.com/roycaihw/kubernetes.git"
BRANCH_NAME="${1:-psi-dev-136}"
THRESHOLD="${2:-0.6}"

echo "Verifying VM Kernel Prerequisites (cgroup v2 and PSI)..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  echo 'Checking for cgroup v2...'
  if ! stat -fc %T /sys/fs/cgroup/ | grep -q 'cgroup2fs'; then
    echo 'ERROR: cgroup v2 is NOT enabled on this VM! PSI requires cgroup v2.'
    exit 1
  fi
  echo 'cgroup v2 is active.'
  
  echo 'Checking for PSI kernel support...'
  if [ ! -f /proc/pressure/memory ]; then
    echo 'ERROR: PSI is NOT enabled in this kernel (/proc/pressure/memory missing)!'
    exit 1
  fi
  echo 'Kernel PSI support is active.'
"

echo "[1/3] Installing dependencies..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  sudo apt-get update && \
  sudo DEBIAN_FRONTEND=noninteractive apt-get install -y build-essential git jq make wget docker.io && \
  sudo usermod -aG docker \$USER && \
  wget -q -nc https://go.dev/dl/go1.22.4.linux-amd64.tar.gz && \
  sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz && \
  if ! grep -q '/usr/local/go/bin' ~/.bashrc; then
      echo 'export PATH=\$PATH:/usr/local/go/bin' >> ~/.bashrc
  fi
"

echo "[2/3] Cloning repository, checking out branch, and patching threshold..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  if [ ! -d \"kubernetes\" ]; then
    git clone $REPO_URL kubernetes
  fi
  cd kubernetes && \
  git fetch --all && \
  git checkout $BRANCH_NAME && \
  git checkout -- pkg/kubelet/apis/config/v1beta1/defaults.go && \
  
  echo 'Patching SystemMemoryContentionThreshold default to ${THRESHOLD}...' && \
  sed -i 's/ptr.To\[float64\](0.9)/ptr.To\[float64\](${THRESHOLD})/g' pkg/kubelet/apis/config/v1beta1/defaults.go
"

echo "[3/3] Starting local-up-cluster.sh with threshold $THRESHOLD in the background..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="
  source ~/.bashrc
  cd kubernetes
  ./hack/install-etcd.sh
  export PATH=\$PATH:\$(pwd)/third_party/etcd
  export KUBE_GIT_VERSION=v1.36.0
  export FEATURE_GATES='PSINodeCondition=true'
  export KUBELET_FLAGS='--feature-gates=PSINodeCondition=true --fail-swap-on=false'
  nohup hack/local-up-cluster.sh > cluster.log 2>&1 &
"
