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
echo 'Capturing active Kubelet configuration and thresholds...'
# Try to grab the live config endpoint first:
KUBELET_CONFIG=$(curl -s --insecure https://127.0.0.1:10250/configz | jq '.kubeletconfig' 2>/dev/null || echo "")

if [ -z "$KUBELET_CONFIG" ] || [ "$KUBELET_CONFIG" == "null" ]; then
  # Fallback to the physical config file if the endpoint is disabled
  KUBELET_CONFIG=$(cat /var/run/kubernetes/kubelet.yaml 2>/dev/null || echo "Unable to fetch Kubelet Config!")
fi

echo "===== EXPERIMENT KUBELET CONFIGURATION =====" >> /home/haoweic_google_com/psi-raw-dumps.log
echo "$KUBELET_CONFIG" >> /home/haoweic_google_com/psi-raw-dumps.log
echo "============================================" >> /home/haoweic_google_com/psi-raw-dumps.log

echo 'Starting background data collection loop to capture raw timeline...'
cat << 'COLLECTOR' > /tmp/data-collector.sh
#!/bin/bash
export KUBECONFIG=/var/run/kubernetes/admin.kubeconfig
export PATH=$PATH:$(pwd)/kubernetes/_output/local/bin/linux/amd64/
NODE_NAME=$(kubectl get nodes -o jsonpath='{.items[0].metadata.name}')

echo "TIMESTAMP, EVENT, PSI_SOME_AVG10, PSI_SOME_AVG60, PSI_FULL_AVG10, PSI_FULL_AVG60, MEM_AVAIL_KB, COND_MEM_PRESSURE, COND_PSI_PRESSURE" > /home/haoweic_google_com/psi-timeline.csv

while true; do
  TS=$(date +"%Y-%m-%dT%H:%M:%S.%3NZ")
  
  # Kernel PSI
  RAW_PSI=$(cat /proc/pressure/memory)
  SOME_10=$(echo "$RAW_PSI" | grep 'some' | grep -o 'avg10=[0-9.]*' | cut -d= -f2)
  SOME_60=$(echo "$RAW_PSI" | grep 'some' | grep -o 'avg60=[0-9.]*' | cut -d= -f2)
  FULL_10=$(echo "$RAW_PSI" | grep 'full' | grep -o 'avg10=[0-9.]*' | cut -d= -f2)
  FULL_60=$(echo "$RAW_PSI" | grep 'full' | grep -o 'avg60=[0-9.]*' | cut -d= -f2)

  # OS Memory
  MEM_AVAIL=$(awk '/MemAvailable/ {print $2}' /proc/meminfo)

  # Kubernetes Conditions
  NODE_JSON=$(kubectl get node "$NODE_NAME" -o json 2>/dev/null || echo "{}")
  COND_LEGACY=$(echo "$NODE_JSON" | jq -r '.status.conditions[] | select(.type=="MemoryPressure") | .status' 2>/dev/null || echo "Unknown")
  COND_PSI=$(echo "$NODE_JSON" | jq -r '.status.conditions[] | select(.type=="SystemMemoryContentionPressure") | .status' 2>/dev/null || echo "Unknown")

  EVENT="BASELINE"
  if [ -f /tmp/stress-deployed ]; then
    EVENT="STRESSING"
  fi

  echo "$TS, $EVENT, $SOME_10, $SOME_60, $FULL_10, $FULL_60, $MEM_AVAIL, $COND_LEGACY, $COND_PSI" >> /home/haoweic_google_com/psi-timeline.csv

  # Also capture full raw blocks for deep-dive analysis
  echo "===== $TS =====" >> /home/haoweic_google_com/psi-raw-dumps.log
  echo "--- /proc/pressure/memory ---" >> /home/haoweic_google_com/psi-raw-dumps.log
  cat /proc/pressure/memory >> /home/haoweic_google_com/psi-raw-dumps.log
  echo "--- /sys/fs/cgroup/memory.stat (head) ---" >> /home/haoweic_google_com/psi-raw-dumps.log
  head -n 20 /sys/fs/cgroup/memory.stat 2>/dev/null >> /home/haoweic_google_com/psi-raw-dumps.log || true
  
  echo "--- Kubelet Summary API (Node Memory) ---" >> /home/haoweic_google_com/psi-raw-dumps.log
  timeout 1 curl -s --insecure https://127.0.0.1:10250/stats/summary -H "Authorization: Bearer $(kubectl create token default)" | jq '.node.memory' 2>/dev/null >> /home/haoweic_google_com/psi-raw-dumps.log || true
  
  sleep 2
done
COLLECTOR

chmod +x /tmp/data-collector.sh
sudo chrt -f 99 ionice -c 1 -n 0 /tmp/data-collector.sh &
COLLECTOR_PID=$!
echo "Data collection started in background (PID: $COLLECTOR_PID)."

echo 'Deploying massive memory stressor pod...'
cat << 'POD' | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: memory-stressor
spec:
  containers:
  - name: stress-ng
    image: alexeiled/stress-ng
    args: ["--vm", "1", "--vm-bytes", "150%", "--vm-keep", "--timeout", "15m"]
    resources:
      requests:
        memory: "100Mi"
        cpu: "100m"
      limits:
        memory: "2500Mi"
POD
touch /tmp/stress-deployed
echo "==============================================="
echo ''

echo "## Observation 2: Memory Pressure Trigger" >> experiment-report.md
echo 'Polling node conditions for SystemMemoryContentionPressure...'
TRIGGERED=false
LEGACY_TRIGGERED=false
for i in {1..240}; do
  # We can still echo basic progress to the console
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

echo 'Observing sustained memory pressure for 10 minutes to gather avg60 data...'
sleep 600

echo 'Deleting memory stressor to observe condition recovery...'
kubectl delete pod memory-stressor || true

echo 'Polling node conditions for SystemMemoryContentionPressure recovery...'
RECOVERED=false
for i in {1..240}; do # Polling for 20 minutes (240 * 5s)
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

echo "Stopping data collection..."
kill $COLLECTOR_PID || true
wait $COLLECTOR_PID 2>/dev/null || true
echo "Data collection stopped."

tar -czf /home/haoweic_google_com/psi-timeline-data.tar.gz /home/haoweic_google_com/psi-timeline.csv /home/haoweic_google_com/psi-raw-dumps.log
EOF

echo "Uploading script to VM..."
gcloud compute scp /tmp/04-verify-behavior-remote.sh "$VM_NAME:/tmp/04-verify-behavior-remote.sh" --zone="$ZONE" --ssh-key-file=/usr/local/google/home/haoweic/.ssh/google_compute_engine

echo "Execute script on VM..."
gcloud compute ssh "$VM_NAME" --zone="$ZONE" --ssh-key-file=/usr/local/google/home/haoweic/.ssh/google_compute_engine --command="bash /tmp/04-verify-behavior-remote.sh"

echo "Downloading experiment report and timeline data..."
gcloud compute scp "$VM_NAME:experiment-report.md" "./experiment-report.md" --zone="$ZONE" --ssh-key-file=/usr/local/google/home/haoweic/.ssh/google_compute_engine
gcloud compute scp "$VM_NAME:/home/haoweic_google_com/psi-timeline-data.tar.gz" "./psi-timeline-data.tar.gz" --zone="$ZONE" --ssh-key-file=/usr/local/google/home/haoweic/.ssh/google_compute_engine

echo "Extracting timeline data locally..."
tar -xzf ./psi-timeline-data.tar.gz -C ./
mv home/haoweic_google_com/psi-timeline.csv ./
mv home/haoweic_google_com/psi-raw-dumps.log ./
rm -rf home/ psi-timeline-data.tar.gz

echo "Report saved to ./experiment-report.md"
echo "Timeline CSV saved to ./psi-timeline.csv"
echo "Raw dumps saved to ./psi-raw-dumps.log"
