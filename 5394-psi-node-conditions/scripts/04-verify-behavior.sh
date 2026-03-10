#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"

echo "Generating remote verification payload..."
cat << 'EOF' > /tmp/04-verify-behavior-remote.sh
#!/bin/bash
set -euo pipefail

export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
export PATH=$PATH:$(pwd)/kubernetes/_output/local/bin/linux/amd64/

NODE_NAME=$(kubectl get nodes -o jsonpath='{.items[0].metadata.name}')
echo "Testing against node: $NODE_NAME"

# Initialize Report
echo "# PSI Node Condition Experiment Report" > experiment-report.md
echo "Date: $(date -u)" >> experiment-report.md
echo "Node: $NODE_NAME" >> experiment-report.md
echo "" >> experiment-report.md
echo "## Memory PSI Timeline" >> experiment-report.md

echo 'Polling node conditions for SystemMemoryContentionPressure...'
TRIGGERED=false
for i in {1..120}; do
  echo "**Attempt $i:**" >> experiment-report.md
  echo '```text' >> experiment-report.md
  cat /proc/pressure/memory >> experiment-report.md
  echo '```' >> experiment-report.md

  if kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "SystemMemoryContentionPressure" and .status == "True")' > /dev/null 2>&1; then
    echo 'SUCCESS: SystemMemoryContentionPressure condition is True!'
    TRIGGERED=true
    
    echo "## Condition Triggered Successfully!" >> experiment-report.md
    echo '```json' >> experiment-report.md
    kubectl get node "$NODE_NAME" -o json | jq '.status.conditions[] | select(.type == "SystemMemoryContentionPressure")' >> experiment-report.md
    echo '```' >> experiment-report.md
    break
  fi
  sleep 5
done

if [ "$TRIGGERED" = false ]; then
  echo 'FAIL: Condition did not trigger within 10 minutes.'
  exit 1
fi

echo 'Verifying taint is applied...'
if kubectl get node "$NODE_NAME" -o json | jq -e '.spec.taints[] | select(.key == "node.kubernetes.io/memory-contention-pressure" and .effect == "NoSchedule")' > /dev/null 2>&1; then
  echo 'SUCCESS: Taint node.kubernetes.io/memory-contention-pressure=:NoSchedule is present.'
  echo "## Node Taints" >> experiment-report.md
  echo '```json' >> experiment-report.md
  kubectl get node "$NODE_NAME" -o json | jq '.spec.taints[]' >> experiment-report.md
  echo '```' >> experiment-report.md
else
  echo 'FAIL: Taint is missing!'
  exit 1
fi

echo 'Spawning test pod to verify scheduling behavior...'
cat <<POD | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: scheduling-test-pod
spec:
  containers:
  - name: test
    image: nginx
POD

echo 'Waiting for scheduling attempt...'
sleep 10

POD_STATUS=$(kubectl get pod scheduling-test-pod -o jsonpath='{.status.phase}')
POD_EVENTS=$(kubectl get events --field-selector involvedObject.name=scheduling-test-pod | grep FailedScheduling || true)

if [ "$POD_STATUS" == "Pending" ]; then
  echo 'SUCCESS: Pod is correctly stuck in Pending state.'
  if [ ! -z "$POD_EVENTS" ]; then
    echo "Confirmed scheduling failure reason:"
    echo "$POD_EVENTS"
    
    echo "## Validation Pod Scheduling Events" >> experiment-report.md
    echo '```text' >> experiment-report.md
    echo "$POD_EVENTS" >> experiment-report.md
    echo '```' >> experiment-report.md
  fi
else
  echo "FAIL: Pod bypassed the taint! Status: $POD_STATUS"
  exit 1
fi

echo ''
echo 'All Verification Checks Passed Successfully!'
EOF

echo "Uploading script to VM..."
gcloud compute scp /tmp/04-verify-behavior-remote.sh "$VM_NAME:/tmp/04-verify-behavior-remote.sh" --zone="$ZONE"

echo "Executing script on VM..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --command="bash /tmp/04-verify-behavior-remote.sh"

echo "Downloading experiment report..."
gcloud compute scp "$VM_NAME:experiment-report.md" "./experiment-report.md" --zone="$ZONE"
echo "Report saved to ./experiment-report.md"
