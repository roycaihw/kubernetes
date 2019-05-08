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
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	kubeapiservertesting "k8s.io/kubernetes/cmd/kube-apiserver/app/testing"
	"k8s.io/kubernetes/test/integration/framework"
)

type MyTransport struct {
	rt http.RoundTripper
}

func (t *MyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.rt.RoundTrip(req)
	klog.Errorf(">>> type of interface %v", reflect.TypeOf(t.rt))
	klog.Errorf(">>> Response header %v", resp.Header)
	klog.Errorf(">>> Response trailer %v", resp.Trailer)
	klog.Errorf(">>> Response req %v", resp.Request)
	return resp, err
}

var target string

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	klog.Errorf("---target %v", target)
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

func getListenAddress() string {
	port := "1338"
	return ":" + port
}

func TestReconnection(t *testing.T) {
	result := kubeapiservertesting.StartTestServerOrDie(t, nil, nil, framework.SharedEtcd())
	defer result.TearDownFn()
	target = result.ClientConfig.Host
	result.ClientConfig.Host = "localhost:1338"

	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
	// result.ClientConfig.Timeout = 10 * time.Second
	// execer := exec.New()
	// dbus := utildbus.New()
	// protocol := utiliptables.ProtocolIpv4
	// protocol6 := utiliptables.ProtocolIpv6
	// iptInterface := utiliptables.New(execer, dbus, protocol)
	// ip6tInterface := utiliptables.New(execer, dbus, protocol6)
	klog.Errorf(">>> host: %v", result.ClientConfig.Host)
	klog.Errorf(">>> creating client")
	klog.Errorf(">>> nil wrap? %v", result.ClientConfig.WrapTransport == nil)
	result.ClientConfig.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return &MyTransport{rt: rt}
	}
	client := clientset.NewForConfigOrDie(result.ClientConfig).CoreV1().Endpoints("default")
	klog.Errorf(">>> done creating client")
	w, err := client.Watch(metav1.ListOptions{ResourceVersion: "0"})
	if err != nil {
		t.Fatalf("failed to watch pods: %v", err)
	}
	// u, err := url.Parse(result.ClientConfig.Host)
	// if err != nil {
	// 	t.Fatalf("failed to parse url: %s", result.ClientConfig.Host)
	// }
	// defer func() {
	// 	klog.Errorf(">>> deleting rule")
	// 	if err := iptInterface.DeleteRule(iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--dport", u.Port()); err != nil {
	// 		t.Fatalf("%v", err)
	// 	}
	// 	if err := ip6tInterface.DeleteRule(iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--dport", u.Port()); err != nil {
	// 		t.Fatalf("%v", err)
	// 	}
	// 	klog.Errorf(">>> done deleting rule")
	// }()
	// klog.Errorf(">>> adding rule")
	// if _, err := iptInterface.EnsureRule(iptables.Prepend, iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--dport", u.Port()); err != nil {
	// 	t.Fatalf("%v", err)
	// }
	// if _, err := ip6tInterface.EnsureRule(iptables.Prepend, iptables.TableFilter, iptables.ChainOutput, "-p", "tcp", "--jump", "DROP", "--dport", u.Port()); err != nil {
	// 	t.Fatalf("%v", err)
	// }
	// klog.Errorf(">>> done adding rule")
	for i := 0; i < 3; i++ {
		if _, err := client.Create(&v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("configmaps-%d", i),
			},
		}); err != nil {
			t.Fatalf("failed to list pods: %v", err)
		}
	}
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
			t.Fatalf("timeout reached")
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
