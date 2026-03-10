# PSI Node Condition Experiment Report
Date: Sat Mar  7 01:28:11 UTC 2026
Node: 127.0.0.1

## Memory PSI Timeline
**Attempt 1:**
```text
some avg10=0.84 avg60=0.65 avg300=0.17 total=1657359
full avg10=0.76 avg60=0.62 avg300=0.16 total=1601142
```
**Attempt 2:**
```text
some avg10=0.56 avg60=0.61 avg300=0.16 total=1668889
full avg10=0.51 avg60=0.58 avg300=0.16 total=1611877
```
## Condition Triggered Successfully!
```json
{
  "lastHeartbeatTime": "2026-03-07T01:28:12Z",
  "lastTransitionTime": "2026-03-07T01:28:12Z",
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
  "timeAdded": "2026-03-07T01:28:13Z"
}
```
## Validation Pod Scheduling Events
```text
10s         Warning   FailedScheduling   pod/scheduling-test-pod   0/1 nodes are available: 1 node(s) had untolerated taint(s). no new claims to deallocate, preemption: 0/1 nodes are available: 1 Preemption is not helpful for scheduling.
```
