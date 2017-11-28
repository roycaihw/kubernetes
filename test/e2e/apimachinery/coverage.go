/*
Copyright 2017 The Kubernetes Authors.

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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/glog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/test/e2e/framework"
)

type resource struct {
	group      string
	version    string
	name       string
	namespaced bool
	verbs      []string
}

var _ = SIGDescribe("Coverage", func() {

	// It("should be able to list pods endpoint", func() {
	// 	ns := f.Namespace.Name

	// 	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	// 	client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
	// 	framework.ExpectNoError(err, "getting group version resources for dynamic client")
	// 	apiResource := metav1.APIResource{Name: gvr.Resource, Namespaced: true}
	// 	podList, err := client.Resource(&apiResource, ns).List(metav1.ListOptions{})

	// 	By("using client to send list request")
	// 	Expect(err).ToNot(HaveOccurred())
	// 	glog.Infof(">>>>>>pod list: %v", podList)
	// })

	// It("should be able to create pods endpoint", func() {
	// 	ns := f.Namespace.Name

	// 	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	// 	testPod := v1.Pod{
	// 		TypeMeta: metav1.TypeMeta{
	// 			Kind:       "Pod",
	// 			APIVersion: "v1",
	// 		},
	// 		ObjectMeta: metav1.ObjectMeta{Name: "testing-manual-pod"},
	// 		Spec: v1.PodSpec{
	// 			Containers: []v1.Container{
	// 				{
	// 					Name:  "kubernetes-pause",
	// 					Image: "gcr.io/google-containers/pause:2.0",
	// 				},
	// 			},
	// 		},
	// 	}
	// 	jsonPod, err := json.Marshal(testPod)
	// 	framework.ExpectNoError(err, "marshalling test-pod for create using dynamic client")

	// 	unstruct := &unstructuredv1.Unstructured{}
	// 	err = unstruct.UnmarshalJSON(jsonPod)
	// 	framework.ExpectNoError(err, "unmarshalling test-pod as unstructured for create using dynamic client")

	// 	client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
	// 	framework.ExpectNoError(err, "getting group version resources for dynamic client")

	// 	apiResource := metav1.APIResource{Name: gvr.Resource, Namespaced: true}
	// 	unstruct, err = client.Resource(&apiResource, ns).Create(unstruct)

	// 	By("using client to send create request")
	// 	// _, err := client.List(metav1.ListOptions{})
	// 	Expect(err).ToNot(HaveOccurred())

	// 	podList, err := client.Resource(&apiResource, ns).List(metav1.ListOptions{})
	// 	By("getting pod list after creation")
	// 	Expect(err).ToNot(HaveOccurred())
	// 	glog.Infof(">>>>>>pod list: %v", podList)
	// })

	// It("should be able to create pods endpoint from yaml file", func() {
	// 	ns := f.Namespace.Name

	// 	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	// 	By("reading yaml file data")
	// 	file, err := os.Open(filepath.Join(framework.TestContext.RepoRoot, "test/e2e/apimachinery/testdata", "pod.yaml"))
	// 	defer file.Close()

	// 	Expect(err).ToNot(HaveOccurred())
	// 	reader := yaml.NewYAMLReader(bufio.NewReader(file))
	// 	yamlPod, err := reader.Read()
	// 	Expect(err).ToNot(HaveOccurred())

	// 	jsonPod, err := yaml.ToJSON(yamlPod)
	// 	framework.ExpectNoError(err, "converting yaml to json")

	// 	unstruct := &unstructuredv1.Unstructured{}
	// 	err = unstruct.UnmarshalJSON(jsonPod)
	// 	framework.ExpectNoError(err, "unmarshalling test-pod as unstructured for create using dynamic client")

	// 	client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
	// 	framework.ExpectNoError(err, "getting group version resources for dynamic client")

	// 	apiResource := metav1.APIResource{Name: gvr.Resource, Namespaced: true}
	// 	unstruct, err = client.Resource(&apiResource, ns).Create(unstruct)

	// 	By("using client to send create request")
	// 	// _, err := client.List(metav1.ListOptions{})
	// 	Expect(err).ToNot(HaveOccurred())

	// 	podList, err := client.Resource(&apiResource, ns).List(metav1.ListOptions{})
	// 	By("getting pod list after creation")
	// 	Expect(err).ToNot(HaveOccurred())
	// 	glog.Infof(">>>>>>pod list: %v", podList)
	// })

	f := framework.NewDefaultFramework("coverage")

	tables := make([]resource, 0)

	glog.Infof("reporoot: %v\n", framework.TestContext.RepoRoot)
	surfaceFile, err := os.Open(filepath.Join("/usr/local/google/home/haoweic/Projects/k8s-p1/src/k8s.io/kubernetes/test/e2e/apimachinery/testdata", "resources.csv"))
	// Expect(err).ToNot(HaveOccurred(), "failed to read api surface file")
	if err != nil {
		glog.Infof("failed to read api surface file: %v", err)
	}
	defer surfaceFile.Close()

	scanner := bufio.NewScanner(surfaceFile)
	for scanner.Scan() {
		glog.Infof(">>>>%v\n", scanner.Text())
		parts := strings.Split(scanner.Text(), ",")
		var r resource
		r.group = strings.Split(parts[0], "/")[0]
		r.version = strings.Split(parts[0], "/")[1]
		r.name = parts[1]

		namespaced, _ := strconv.ParseBool(parts[2])
		// 	Expect(err).ToNot(HaveOccurred())
		r.namespaced = namespaced

		r.verbs = parts[3:]
		tables = append(tables, r)
		glog.Infof("table entry: %v\n", r)
	}
	// err = scanner.Err()
	// Expect(err).ToNot(HaveOccurred(), "failed to bufio scan on api surface file")

	glog.Infof("table size: %v", len(tables))

	for _, table := range tables {
		itTable := table
		for _, verb := range itTable.verbs {
			itVerb := verb
			if itVerb == "list" {
				rule := fmt.Sprintf("should be able to list resource %v %v %v", itTable.group, itTable.version, itTable.name)
				glog.Infof("%v", rule)
				It(rule, func() {
					ns := f.Namespace.Name
					glog.Infof("<<<<<group: %v, version: %v, resource: %v, namespaced: %v.\n", itTable.group, itTable.version, itTable.name, itTable.namespaced)

					gvr := schema.GroupVersionResource{Group: itTable.group, Version: itTable.version, Resource: itTable.name}
					client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
					framework.ExpectNoError(err, "getting group version resources for dynamic client")
					apiResource := metav1.APIResource{Name: gvr.Resource, Namespaced: itTable.namespaced}
					resourceList, err := client.Resource(&apiResource, ns).List(metav1.ListOptions{})

					By("using client to send list request")
					Expect(err).ToNot(HaveOccurred())
					glog.Infof(">>>>>>resource list: %v", resourceList)
				})
			}
		}
	}

	It("should be able to list cronjobs endpoint", func() {
		ns := f.Namespace.Name

		gvr := schema.GroupVersionResource{Group: "batch", Version: "v1beta1", Resource: "cronjobs"}
		client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
		framework.ExpectNoError(err, "getting group version resources for dynamic client")
		apiResource := metav1.APIResource{Name: gvr.Resource, Namespaced: true}
		_, err = client.Resource(&apiResource, ns).List(metav1.ListOptions{})

		// c := f.ClientSet
		// client := c.CoreV1().PodTemplates(ns)

		By("using client to send list request")
		// _, err := client.List(metav1.ListOptions{})
		Expect(err).ToNot(HaveOccurred())
	})
})
