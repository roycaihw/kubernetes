#!/bin/bash
set -euo pipefail

ZONE="us-central1-a"
VM_NAME="psi-experiment-node"

echo "Generating remote verification payload..."
cat << 'EOF' > /tmp/04-verify-behavior-remote.sh
#!/bin/bash
set -euo pipefail

mkdir -p /tmp/bin
cp $(pwd)/kubernetes/_output/local/bin/linux/amd64/kubectl /tmp/bin/
export PATH=/tmp/bin:$PATH
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig

NODE_NAME=$(kubectl get nodes -o jsonpath='{.items[0].metadata.name}')
echo "Testing against node: $NODE_NAME"

# Initialize Report
echo "# PSI Node Condition Experiment Report" > experiment-report.md
echo "Date: $(date -u)" >> experiment-report.md
echo "Node: $NODE_NAME" >> experiment-report.md
echo "" >> experiment-report.md
echo "## Memory PSI Timeline" >> experiment-report.md

echo "## Observation 1: Normal Operation Stability" >> experiment-report.md
echo 'Polling node conditions for 30 seconds to ensure SystemMemoryContentionPressure remains False under baseline load...'
for i in {1..6}; do
  if kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "SystemMemoryContentionPressure" and .status == "True")' > /dev/null 2>&1; then
    echo "FAIL: Node asserted SystemMemoryContentionPressure=True under normal baseline load! (False Positive)"
    echo "FAIL: False positive triggered during baseline observation." >> experiment-report.md
    exit 1
  fi
  sleep 5
done
echo "SUCCESS: Node condition remained stably False under baseline load."
echo "Baseline stable. No false positives detected." >> experiment-report.md

echo ''
echo "==============================================="
echo 'Deploying massive memory stressor pod...'
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
echo "==============================================="
echo ''

echo "## Observation 2: Memory Pressure Trigger" >> experiment-report.md
echo 'Polling node conditions for SystemMemoryContentionPressure...'
TRIGGERED=false
LEGACY_TRIGGERED=false
for i in {1..120}; do
  echo "**Attempt $i:**" >> experiment-report.md
  echo '```text' >> experiment-report.md
  cat /proc/pressure/memory >> experiment-report.md
  echo '```' >> experiment-report.md

  # Capture OS memory metrics to understand why legacy MemoryPressure behaves the way it does
  MEM_TOTAL=$(awk '/MemTotal/ {print $2}' /proc/meminfo)
  MEM_AVAIL=$(awk '/MemAvailable/ {print $2}' /proc/meminfo)
  echo "OS Memory: ${MEM_AVAIL} kB available out of ${MEM_TOTAL} kB total" >> experiment-report.md

  if [ "$LEGACY_TRIGGERED" = false ] && kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "MemoryPressure" and .status == "True")' > /dev/null 2>&1; then
    echo "NOTICE: Legacy MemoryPressure triggered at attempt $i"
    LEGACY_TRIGGERED=true
    echo "Legacy MemoryPressure triggered at attempt $i" >> experiment-report.md
  fi

  if kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "SystemMemoryContentionPressure" and .status == "True")' > /dev/null 2>&1; then
    echo 'SUCCESS: SystemMemoryContentionPressure condition is True!'
    echo 'Waiting 10 seconds to verify condition stability (no flapping)...'
    sleep 10
    if kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "SystemMemoryContentionPressure" and .status == "True")' > /dev/null 2>&1; then
        echo 'SUCCESS: SystemMemoryContentionPressure condition is still True after 10s (stable).'
        TRIGGERED=true
        
        echo "## Condition Triggered Successfully!" >> experiment-report.md
        echo '```json' >> experiment-report.md
        kubectl get node "$NODE_NAME" -o json | jq '.status.conditions[] | select(.type == "SystemMemoryContentionPressure")' >> experiment-report.md
        echo '```' >> experiment-report.md
        break
    else
        echo 'WARNING: SystemMemoryContentionPressure flapped back to False within 10s!'
        echo 'WARNING: SystemMemoryContentionPressure flapped!' >> experiment-report.md
    fi
  fi
  sleep 5
done

if [ "$TRIGGERED" = false ]; then
  echo 'FAIL: Condition did not trigger stably within 10 minutes.'
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

echo 'Deleting memory stressor to observe condition recovery...'
kubectl delete pod memory-stressor || true

echo 'Polling node conditions for SystemMemoryContentionPressure recovery...'
RECOVERED=false
for i in {1..120}; do
  if kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "SystemMemoryContentionPressure" and .status == "False")' > /dev/null 2>&1; then
    echo 'SUCCESS: SystemMemoryContentionPressure condition is False!'
    echo 'Waiting 10 seconds to verify condition stability (no flapping)...'
    sleep 10
    if kubectl get node "$NODE_NAME" -o json | jq -e '.status.conditions[] | select(.type == "SystemMemoryContentionPressure" and .status == "False")' > /dev/null 2>&1; then
        echo 'SUCCESS: SystemMemoryContentionPressure condition is still False after 10s (stable recovery).'
        RECOVERED=true
        
        echo "## Condition Recovered Successfully!" >> experiment-report.md
        echo "Recovery observed after stressor deletion." >> experiment-report.md
        break
    else
        echo 'WARNING: SystemMemoryContentionPressure flapped back to True within 10s!'
    fi
  fi
  sleep 5
done

if [ "$RECOVERED" = false ]; then
  echo 'FAIL: Condition did not recover stably within 10 minutes.'
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
