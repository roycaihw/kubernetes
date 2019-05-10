/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apimachinery

import (
	"os/exec"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/util/iptables"
	"k8s.io/kubernetes/test/e2e/framework"

	"github.com/onsi/ginkgo"
)

// TODO(roycaihw): convert this into an integration test with safer iptables change
var _ = SIGDescribe("Client", func() {
	f := framework.NewDefaultFramework("client")
	ginkgo.It("should reconnect", func() {
		podClient := f.ClientSet.CoreV1().Pods(f.Namespace.Name)
		w, err := podClient.Watch(metav1.ListOptions{ResourceVersion: "0"})
		if err != nil {
			framework.Failf("failed to watch pods: %v", err)
		}
		masterAddresses := framework.GetAllMasterAddresses(f.ClientSet)
		ginkgo.By("drop network traffic to the master")
		for _, masterAddress := range masterAddresses {
			DropNetwork(masterAddress)
		}
		time.Sleep(15 * time.Second)
		ginkgo.By("stop dropping network traffic to the master")
		for _, masterAddress := range masterAddresses {
			UndropNetwork(masterAddress)
		}

		stopTimer := time.NewTimer(10 * time.Second)
		defer stopTimer.Stop()
		if _, err := podClient.List(metav1.ListOptions{ResourceVersion: "0"}); err != nil {
			framework.Failf("failed to list pods: %v", err)
		}
		for {
			select {
			case actual, ok := <-w.ResultChan():
				if ok {
					framework.Logf("Got : %v %v", actual.Type, actual.Object)
				} else {
					framework.Failf("Watch closed unexpectedly")
				}
			case <-stopTimer.C:
				framework.Failf("timeout reached")
			}
		}
	})
})

// DropNetwork changes iptables to drop network traffic from host to given destination
func DropNetwork(to string) {
	execer := exec.New()
	dbus := utildbus.New()
	protocol := utiliptables.ProtocolIpv4
	iptInterface := utiliptables.New(execer, dbus, protocol)
	if _, err := iptInterface.EnsureRule(iptables.Prepend, iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "-d", to); err != nil {
		framework.Failf("failed to drop network traffic to %s using iptables: %v", to, err)
	}
	// framework.Logf("drop network traffic from to %s", to)
	// err := exec.Command("iptables", "-I OUTPUT", "-p tcp", "--jump DROP", fmt.Sprintf("-d %s", to)).Run()
	// framework.Failf("failed to drop network traffic to %s using iptables: %v", to, err)
}

// UndropNetwork changes iptables to stop dropping network traffic from host to given destination
func UndropNetwork(to string) {
	execer := exec.New()
	dbus := utildbus.New()
	protocol := utiliptables.ProtocolIpv4
	iptInterface := utiliptables.New(execer, dbus, protocol)
	if err := iptInterface.DeleteRule(iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "-d", to); err != nil {
		framework.Failf("failed to stop dropping network traffic to %s using iptables: %v", to, err)
	}
	// framework.Logf("stop dropping network traffic to %s", to)
	// cmd := exec.Command("iptables", "-D OUTPUT", "-p tcp", "--jump DROP", fmt.Sprintf("-d %s", to))
	// // TODO(roycaihw): rephrase
	// // Undrop command may fail if the rule has never been created.
	// // In such case we just lose 30 seconds, but the cluster is healthy.
	// // But if the rule had been created and removing it failed, the node is broken and
	// // not coming back. Subsequent tests will run or fewer nodes (some of the tests
	// // may fail). Manual intervention is required in such case (recreating the
	// // cluster solves the problem too).
	// if err := wait.Poll(time.Millisecond*100, time.Second*30, func() (bool, error) {
	// 	// TODO(roycaihw): revisit error handling
	// 	if err := cmd.Run(); err == nil {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }); err != nil {
	// 	framework.Failf("failed to stop dropping network traffic to %s using iptables: %v", to, err)
	// }
}
