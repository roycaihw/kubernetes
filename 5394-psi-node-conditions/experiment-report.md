# PSI Node Condition Experiment Report
Date: Tue Mar 10 21:20:47 UTC 2026
Node: 127.0.0.1

## Memory PSI Timeline
**Attempt 1:**
```text
some avg10=0.12 avg60=0.18 avg300=0.06 total=27053970
full avg10=0.10 avg60=0.16 avg300=0.05 total=22109003
```
OS Memory: 15346492 kB available out of 16374236 kB total
## Condition Triggered Successfully!
```json
{
  "lastHeartbeatTime": "2026-03-10T21:20:20Z",
  "lastTransitionTime": "2026-03-10T21:20:20Z",
  "message": "kubelet has SystemMemoryContentionPressure",
  "reason": "KubeletHasSystemMemoryContentionPressure",
  "status": "True",
  "type": "SystemMemoryContentionPressure"
}
```
## Node Taints
```json
{
  "effect": "NoSchedule",
  "key": "node.kubernetes.io/memory-contention-pressure",
  "timeAdded": "2026-03-10T21:20:20Z"
}
```
## Validation Pod Scheduling Events
```text
9m38s       Warning   FailedScheduling   pod/scheduling-test-pod   0/1 nodes are available: 1 node(s) had untolerated taint(s). no new claims to deallocate, preemption: 0/1 nodes are available: 1 Preemption is not helpful for scheduling.
10s         Warning   FailedScheduling   pod/scheduling-test-pod   0/1 nodes are available: 1 node(s) had untolerated taint(s). no new claims to deallocate, preemption: 0/1 nodes are available: 1 Preemption is not helpful for scheduling.
```
## Condition Recovered Successfully!
Recovery observed after stressor deletion.
