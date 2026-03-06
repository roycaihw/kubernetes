#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"

echo "Destroying GCE VM: $VM_NAME in zone $ZONE..."
gcloud compute instances delete "$VM_NAME" --zone="$ZONE" --quiet

echo "Cleanup complete."
