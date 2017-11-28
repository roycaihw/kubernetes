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

type apiResource struct {
	group      string
	version    string
	name       string
	namespaced bool
	verbs      []string
}

var _ = SIGDescribe("Coverage", func() {

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
	tables, err := readTables()
	if err != nil {
		glog.Fatalf("failed to read tables: %v", err)
	}

	for _, table := range tables {
		itTable := table
		rule := fmt.Sprintf("should be able to support expected CRUD operations for resource (g: %s, v: %s, r: %s)", itTable.group, itTable.version, itTable.name)
		It(rule, func() {
			listResource(f, itTable)
		})
	}
})

func readTables() ([]apiResource, error) {
	tables := make([]apiResource, 0)

	resourcesFile, err := os.Open(filepath.Join("/usr/local/google/home/haoweic/Projects/k8s-p1/src/k8s.io/kubernetes/test/e2e/apimachinery/testdata", "resources.csv"))
	if err != nil {
		return nil, fmt.Errorf("failed to open api resources file: %v", err)
	}
	defer resourcesFile.Close()

	scanner := bufio.NewScanner(resourcesFile)
	for scanner.Scan() {
		var r apiResource
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) < 4 {
			return nil, fmt.Errorf("unexpected resource format: %s", scanner.Text())
		}
		groupVersion := strings.Split(parts[0], "/")
		if len(groupVersion) == 1 {
			r.version = groupVersion[0]
		} else if len(groupVersion) == 2 {
			r.group = groupVersion[0]
			r.version = groupVersion[1]
		} else {
			return nil, fmt.Errorf("unexpected resource group version format: %s", parts[0])
		}
		r.name = parts[1]
		namespaced, err := strconv.ParseBool(parts[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse resource namespaced property: %v", err)
		}
		r.namespaced = namespaced
		r.verbs = parts[3:]
		tables = append(tables, r)
	}
	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to scan api resources file: %v", err)
	}

	return tables, nil
}

func hasVerb(verbs []string, verb string) bool {
	for _, v := range verbs {
		if v == verb {
			return true
		}
	}
	return false
}

func createResource(f *framework.Framework, r apiResource) {
}

func listResource(f *framework.Framework, r apiResource) {
	if !hasVerb(r.verbs, "list") {
		return
	}
	target := fmt.Sprintf("list resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)
	gvr := schema.GroupVersionResource{Group: r.group, Version: r.version, Resource: r.name}
	client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
	Expect(err).NotTo(HaveOccurred(), "failed to create dynamic client for resource %v", r)
	resource := metav1.APIResource{Name: gvr.Resource, Namespaced: r.namespaced}
	_, err = client.Resource(&resource, f.Namespace.Name).List(metav1.ListOptions{})
	Expect(err).ToNot(HaveOccurred(), "failed to list resource %v", r)
}
