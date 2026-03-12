#!/bin/bash
set -euo pipefail

# This script runs the full Phase 2 (System Lead Time) and Phase 3 (Kubepods Lead Time)
# memory test matrix across all 5 candidate PSI threshold defaults.

THRESHOLDS=("0.7" "0.8" "0.9" "0.95")

# Ensure clean slate
./scripts/shared/99-cleanup-vm.sh || true

for t in "${THRESHOLDS[@]}"; do
  echo "========================================================="
  echo "Starting Phase 2 for Threshold Candidate: $t"
  echo "========================================================="
  
  ./scripts/shared/01-provision-vm.sh
  ./scripts/exp2-tune-thresholds/02-setup-cluster.sh psi-dev-136 "$t"
  
  echo "-> Pre-Flight Phase 2: Deploying System Memory Leak..."
  ./scripts/exp2-tune-thresholds/03-deploy-workloads.sh 2
  
  echo "-> Measuring Phase 2 trigger timing..."
  ./scripts/exp2-tune-thresholds/04-measure-timing.sh 2 "$t" || echo "WARNING: Measurement failed for Phase 2 at $t"

  echo "-> Phase 2 Complete. Tearing down VM to ensure clean slate for Phase 3..."
  ./scripts/shared/99-cleanup-vm.sh || true

  echo "========================================================="
  echo "Starting Phase 3 for Threshold Candidate: $t"
  echo "========================================================="

  echo "-> Provisioning fresh GCE Node for Phase 3..."
  ./scripts/shared/01-provision-vm.sh
  ./scripts/exp2-tune-thresholds/02-setup-cluster.sh psi-dev-136 "$t"

  echo "-> Pre-Flight Phase 3: Deploying Kubepods Memory Leak..."
  ./scripts/exp2-tune-thresholds/03-deploy-workloads.sh 3
  
  echo "-> Measuring Phase 3 trigger timing..."
  ./scripts/exp2-tune-thresholds/04-measure-timing.sh 3 "$t" || echo "WARNING: Measurement failed for Phase 3 at $t"

  echo "-> Sweeping Node for next loop..."
  ./scripts/shared/99-cleanup-vm.sh || true
done

echo "========================================================="
echo "Full Experiment 2 Matrix Completed!"
echo "========================================================="
