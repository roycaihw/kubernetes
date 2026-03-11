# Appendix A: Swap "Magic" vs. Death by Thrashing

When evaluating memory pressure, a common question is: *Why is Kubelet reporting a PSI-based memory condition when Swap is enabled? Shouldn't Swap safely handle bursts of memory demand? Is PSI giving a false alarm?*

The answer is that PSI is detecting a devastating, real issue: **Thrashing**. Here is the critical difference between Swap operating safely and Swap paralyzing the node.

### Scenario A: Swap Doing Its Magic (No PSI Alarm)
Imagine you have a node with 16GB of RAM and a 4GB Swapfile.
You deploy a massive Java application that requests 17GB of memory. It loads up, but 4GB of its data is just old, stale log buffers that it only touches once a day. 
1. The kernel quietly pushes that 4GB of stale data into the Swapfile.
2. The Java app comfortably runs its hot, active code inside the 16GB of fast physical RAM.
3. Because the active code is in physical RAM, the CPU never has to wait for the disk. It just executes math at full speed.
4. **The Result:** The node has 17GB allocated, but it is lightning fast. **Memory PSI stays at `0.0` (0% stall).** Kubelet correctly leaves the node alone. Swap did its magic perfectly.

### Scenario B: Thrashing (The PSI Alarm Triggers)
Now look at our `stress-ng` workload experiment. `stress-ng` doesn't just allocate 17GB of memory and let it sit idle. It violently reads and writes every single byte of that 17GB in an infinite loop!
1. The 16GB of physical RAM is full. The remaining 1GB is sitting in the Swapfile.
2. `stress-ng` tries to read the 1GB of data from the Swapfile.
3. To load that 1GB into RAM, the kernel has to violently evict 1GB of *currently active data* out of fast RAM and write it to the slow hard drive. 
4. A millisecond later, `stress-ng` loops around and needs that evicted data back. The kernel has to violently swap it back again.

This is called **Thrashing**. The CPU is no longer doing math; it is sitting completely idle (`stalled`), waiting for the spinning rust hard drive to constantly fetch the next page of memory.

### Why PSI is the Perfect Signal
Because the CPU in Scenario B is completely frozen waiting for the disk 90% of the time, the **`full avg60` Memory PSI spikes to `0.9`**. Kubelet's condition triggers and correctly declares the node is paralyzed.

PSI elegantly bridges the gap:
*   In Scenario A (stale data swapped out), PSI = 0.0. Kubelet ignores the Swap usage. Swap did exactly what it was designed to do.
*   In Scenario B (active data thrashing), PSI = 0.9. Kubelet sees the node is paralyzed and taints it. 

By measuring the *runnable CPU stall time* rather than just *total capacity boundaries*, PSI avoids penalizing safe Swap usage while accurately screaming for help the millisecond that Swap usage starts paralyzing the system's performance.
