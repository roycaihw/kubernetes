#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"
PHASE="${1:-1}"
THRESHOLD="${2:-0.6}"

echo "Generating remote verification payload for Phase ${PHASE}..."
cat << 'EOF' > /tmp/04-measure-timing-remote.sh
#!/bin/bash
set -euo pipefail
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
NODE_NAME=$(kubectl get nodes -o jsonpath='{.items[0].metadata.name}')

PHASE="$1"
THRESHOLD="$2"
REPORT_FILE="exp2-phase${PHASE}-threshold${THRESHOLD}-report.md"

echo "# Experiment 2: Phase ${PHASE} (Threshold ${THRESHOLD})" > "$REPORT_FILE"

if [ "$PHASE" == "1" ]; then
  echo "## Phase 1: False Positive Validation" >> "$REPORT_FILE"
  echo "Polling for 60 seconds to ensure condition never fires under safe load..."
  for i in {1..60}; do
    if kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "SystemMemoryContentionPressure" and .status == "True")' > /dev/null 2>&1; then
      echo "FAIL: Node artificially asserted SystemMemoryContentionPressure=True! Threshold $THRESHOLD is too sensitive." | tee -a "$REPORT_FILE"
      exit 1
    fi
    sleep 1
  done
  echo "SUCCESS: Node condition remained stably False under safe load. Threshold $THRESHOLD passed Phase 1." | tee -a "$REPORT_FILE"

elif [ "$PHASE" == "2" ] || [ "$PHASE" == "3" ]; then
  CONDITION_TYPE="SystemMemoryContentionPressure"
  if [ "$PHASE" == "3" ]; then
    CONDITION_TYPE="NodeKubepodsMemoryContentionPressure"
  fi

  echo "## Phase $PHASE: Lead Time Measurement for $CONDITION_TYPE" >> "$REPORT_FILE"
  echo "Polling for condition $CONDITION_TYPE..."
  
  START_TIME=$(date +%s)
  TRIGGERED=false
  for i in {1..300}; do
    if kubectl get node "$NODE_NAME" -o json | jq -e ".status.conditions[] | select(.type == \"$CONDITION_TYPE\" and .status == \"True\")" > /dev/null 2>&1; then
      TRIGGER_TIME=$(date +%s)
      DIFF=$((TRIGGER_TIME - START_TIME))
      echo "SUCCESS: Condition $CONDITION_TYPE fired at $DIFF seconds from workload start!" | tee -a "$REPORT_FILE"
      TRIGGERED=true
      break
    fi
    sleep 1
  done
  
  if [ "$TRIGGERED" == "false" ]; then
    echo "FAIL: Condition did not fire within 5 minutes under memory leak!" | tee -a "$REPORT_FILE"
    exit 1
  fi
fi
EOF

gcloud compute scp /tmp/04-measure-timing-remote.sh "$VM_NAME:/tmp/04-measure-timing-remote.sh" --zone="$ZONE"
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="bash /tmp/04-measure-timing-remote.sh $PHASE $THRESHOLD"
gcloud compute scp "$VM_NAME:exp2-phase${PHASE}-threshold${THRESHOLD}-report.md" "./exp2-phase${PHASE}-threshold${THRESHOLD}-report.md" --zone="$ZONE"
echo "Metrics saved to ./exp2-phase${PHASE}-threshold${THRESHOLD}-report.md"
