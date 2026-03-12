#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"

echo "[1/2] Creating GCE VM: $VM_NAME in $ZONE..."
gcloud compute instances create "$VM_NAME" \
    --machine-type=n2-standard-4 \
    --image-family=ubuntu-2204-lts \
    --image-project=ubuntu-os-cloud \
    --boot-disk-size=100GB \
    --zone="$ZONE"

echo "[2/2] Waiting for SSH to become available..."
for i in {1..15}; do
  if gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="echo 'SSH is ready.'"; then
    sleep 10
    echo "Configuring 4G Native Swapfile to gracefully bottleneck PSI without instant OOM invocation..."
    gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="sudo fallocate -l 4G /swapfile && sudo chmod 600 /swapfile && sudo mkswap /swapfile && sudo swapon /swapfile"
    echo "VM provisioning complete! Proceed to 02-setup-cluster.sh"
    exit 0
  fi
  echo "SSH not ready yet, retrying in 5 seconds..."
  sleep 5
done

echo "Failed to connect to VM via SSH."
exit 1
