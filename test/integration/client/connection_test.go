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

package client

import (
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	kubeapiservertesting "k8s.io/kubernetes/cmd/kube-apiserver/app/testing"
	utildbus "k8s.io/kubernetes/pkg/util/dbus"
	"k8s.io/kubernetes/pkg/util/iptables"
	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
	"k8s.io/kubernetes/test/integration/framework"
	utilsexec "k8s.io/utils/exec"
)

func TestReconnection(t *testing.T) {
	// TODO: skip if OS is not linux
	server := kubeapiservertesting.StartTestServerOrDie(t, nil, nil, framework.SharedEtcd())
	defer server.TearDownFn()

	// server.ClientConfig.Timeout = 10 * time.Second
	execer := utilsexec.New()
	dbus := utildbus.New()
	protocol := utiliptables.ProtocolIpv4
	protocol6 := utiliptables.ProtocolIpv6
	iptInterface := utiliptables.New(execer, dbus, protocol)
	ip6tInterface := utiliptables.New(execer, dbus, protocol6)
	client := clientset.NewForConfigOrDie(server.ClientConfig).CoreV1().Endpoints("default")
	klog.Errorf(">>> done creating client")

	w, err := client.Watch(metav1.ListOptions{ResourceVersion: "0"})
	if err != nil {
		t.Fatalf("failed to watch pods: %v", err)
	}

	connections, err := exec.Command("ss", "-t").Output()
	if err != nil {
		t.Fatalf("failed ot list tcp connections")
	}

	u, err := url.Parse(server.ClientConfig.Host)
	if err != nil {
		t.Fatalf("failed to parse url: %s", server.ClientConfig.Host)
	}
	// klog.Errorf("--- target: %v", u.Host)
	// klog.Errorf("--- connections: %v", string(connections))
	pattern := regexp.MustCompile(fmt.Sprintf(`%s[\t ]+[0-9.]+:([0-9]+)`, u.Host))
	// pattern := regexp.MustCompile(fmt.Sprintf(`%s`, u.Host))
	ports := pattern.FindStringSubmatch(string(connections))
	sport := ports[1]
	// klog.Errorf("--- ports: %v", ports)

	defer func() {
		klog.Errorf(">>> deleting rule")
		if err := iptInterface.DeleteRule(iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--sport", sport); err != nil {
			t.Fatalf("%v", err)
		}
		if err := ip6tInterface.DeleteRule(iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--sport", sport); err != nil {
			t.Fatalf("%v", err)
		}
		klog.Errorf(">>> done deleting rule")
	}()
	klog.Errorf(">>> adding rule")
	if _, err := iptInterface.EnsureRule(iptables.Prepend, iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--sport", sport); err != nil {
		t.Fatalf("%v", err)
	}
	if _, err := ip6tInterface.EnsureRule(iptables.Prepend, iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--sport", sport); err != nil {
		t.Fatalf("%v", err)
	}
	klog.Errorf(">>> done adding rule")
	go func() {
		for i := 0; i < 3; i++ {
			if _, err := client.Create(&v1.Endpoints{
				ObjectMeta: metav1.ObjectMeta{
					Name: fmt.Sprintf("configmaps-%d", i),
				},
			}); err != nil {
				t.Fatalf("failed to list pods: %v", err)
			}
		}
	}()
	klog.Errorf(">>> done listing")
	stopTimer := time.NewTimer(15 * time.Second)
	defer stopTimer.Stop()
	count := 0
	for {
		select {
		case actual, ok := <-w.ResultChan():
			if ok {
				count++
				klog.Errorf("Got : %v %v", actual.Type, actual.Object)
			} else {
				t.Fatalf("Watch closed unexpectedly")
			}
		case <-stopTimer.C:
			if count != 3 {
				klog.Errorf("--- count: %v", count)
				t.Fatalf("timeout and not matching")
			} else {
				klog.Errorf("--- count: %v", count)
				t.Fatalf("timeout reached")
			}
		}
	}
}

// DropNetwork changes iptables to drop network traffic from host to given destination
// func DropNetwork(to string) error {
// 	if err := iptables.Append(); err != nil {
// 		// if err := exec.Command("iptables", "-I OUTPUT", "-p tcp", "--jump DROP", fmt.Sprintf("-d %s", to)).Run(); err != nil {
// 		return fmt.Errorf("failed to drop network traffic to %s using iptables: %v", to, err)
// 	}
// 	return nil
// }
//
// // UndropNetwork changes iptables to stop dropping network traffic from host to given destination
// func UndropNetwork(to string) error {
// 	// cmd := exec.Command("iptables", "-D OUTPUT", "-p tcp", "--jump DROP", fmt.Sprintf("-d %s", to))
// 	// TODO(roycaihw): rephrase
// 	// Undrop command may fail if the rule has never been created.
// 	// In such case we just lose 30 seconds, but the cluster is healthy.
// 	// But if the rule had been created and removing it failed, the node is broken and
// 	// not coming back. Subsequent tests will run or fewer nodes (some of the tests
// 	// may fail). Manual intervention is required in such case (recreating the
// 	// cluster solves the problem too).
// 	if err := wait.Poll(time.Millisecond*100, time.Second*30, func() (bool, error) {
// 		// TODO(roycaihw): revisit error handling
// 		if err := cmd.Run(); err == nil {
// 			return true, nil
// 		}
// 		return false, nil
// 	}); err != nil {
// 		return fmt.Errorf("failed to stop dropping network traffic to %s using iptables: %v", to, err)
// 	}
// 	return nil
// }
