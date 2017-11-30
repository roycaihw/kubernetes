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
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang/glog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructuredv1 "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/kubernetes/test/e2e/framework"
)

type resourceMap map[resourceMeta]*resource

type resourceMeta struct {
	group   string
	version string
	name    string
}

type resource struct {
	group        string
	version      string
	name         string
	namespaced   bool
	verbs        []string
	subresources resourceMap
}

var _ = SIGDescribe("Coverage", func() {
	f := framework.NewDefaultFramework("coverage")
	tables, err := readTables()
	if err != nil {
		glog.Fatalf("failed to read tables: %v", err)
	}

	for _, table := range tables {
		r := table
		rule := fmt.Sprintf("should be able to support expected CRUD operations for resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
		It(rule, func() {
			gvr := schema.GroupVersionResource{Group: r.group, Version: r.version, Resource: r.name}
			client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
			Expect(err).NotTo(HaveOccurred(), "failed to create dynamic client for resource %v", r)
			resource := metav1.APIResource{Name: gvr.Resource, Namespaced: r.namespaced}
			unstruct := r.dumpResourceYAML()
			name := getResourceName(unstruct)

			// Iterate through verbs in serial, skip verbs that don't exist
			err = r.listResource(f, client, resource)
			Expect(err).ToNot(HaveOccurred(), "failed to list resource %v", r)
			err = r.createResource(f, client, resource, unstruct)
			Expect(err).ToNot(HaveOccurred(), "failed to create resource %v", r)
			err = r.getResource(f, client, resource, name)
			Expect(err).ToNot(HaveOccurred(), "failed to get resource %v", r)
			err = r.deleteResource(f, client, resource, name)
			Expect(err).ToNot(HaveOccurred(), "failed to delete resource %v", r)
		})
	}
})

func readTables() (resourceMap, error) {
	tables := resourceMap{}

	// TODO (roycaihw): use framework reporoot or bindata
	resourcesFile, err := os.Open(filepath.Join("/usr/local/google/home/haoweic/Projects/k8s-p1/src/k8s.io/kubernetes/test/e2e/apimachinery/testdata", "resources.csv"))
	if err != nil {
		return nil, fmt.Errorf("failed to open api resources file: %v", err)
	}
	defer resourcesFile.Close()

	reader := csv.NewReader(bufio.NewReader(resourcesFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to read api resources file: %v", err)
		}

		// Resource record in format: group,version,name,namespaced,verb
		if len(record) != 5 {
			return nil, fmt.Errorf("unexpected resource record length: %v, want: 5, got: %d", record, len(record))
		}

		namespaced, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil, fmt.Errorf("failed to parse resource (%v) namespaced property: %v", record, err)
		}

		m := resourceMeta{
			group:   record[0],
			version: record[1],
			name:    record[2],
		}
		if tables[m] == nil {
			tables[m] = new(resource)
			*tables[m] = resource{
				group:      record[0],
				version:    record[1],
				name:       record[2],
				namespaced: namespaced,
				verbs:      []string{record[4]},
			}
		} else {
			tables[m].verbs = append(tables[m].verbs, record[4])
		}

		// if !strings.Contains(record[2], "/") {
		// 	// If record is a resource
		// 	m.name = record[2]
		// 	tables[m].namespaced = namespaced
		// 	if tables[m].verbs == nil {
		// 		tables[m].verbs = make([]string, 0)
		// 	}
		// 	tables[m].verbs = append(tables[m].verbs, record[4])
		// } else {
		// 	// If record is a subresource
		// 	m.name = strings.Split(record[2], "/")[0]
		// 	subname := strings.Split(record[2], "/")[1]
		// 	if tables[m].subresources == nil {
		// 		tables[m].subresources = make(map[string]subresourceProperty)
		// 	}
		// 	tables[m].subresources[subname].namespaced = namespaced
		// 	if tables[m].subresources[subname].verbs == nil {
		// 		tables[m].subresources[subname].verbs = make([]string, 0)
		// 	}
		// 	tables[m].subresources[subname].verbs = append(tables[m].subresources[subname].verbs, record[4])
		// }
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

func (r resource) dumpResourceYAML() *unstructuredv1.Unstructured {
	unstruct := &unstructuredv1.Unstructured{}
	if !hasVerb(r.verbs, "create") {
		return unstruct
	}

	// TODO (roycaihw): figure out using framework reporoot or bindata
	yamlFile, err := os.Open(filepath.Join("/usr/local/google/home/haoweic/Projects/k8s-p1/src/k8s.io/kubernetes/test/e2e/apimachinery/testdata", r.name+".yaml"))
	defer yamlFile.Close()
	Expect(err).ToNot(HaveOccurred(), "failed to open yaml file for resource %v", r)

	reader := yaml.NewYAMLReader(bufio.NewReader(yamlFile))
	yamlResource, err := reader.Read()
	Expect(err).ToNot(HaveOccurred(), "failed to read yaml file for resource %v", r)

	jsonResource, err := yaml.ToJSON(yamlResource)
	Expect(err).ToNot(HaveOccurred(), "failed to convert yaml to json for resource %v", r)

	err = unstruct.UnmarshalJSON(jsonResource)
	Expect(err).ToNot(HaveOccurred(), "failed to unmarshal json for resource %v", r)

	return unstruct
}

func getResourceName(unstruct *unstructuredv1.Unstructured) string {
	if len(unstruct.Object) == 0 {
		return ""
	}
	return unstruct.Object["metadata"].(map[string]interface{})["name"].(string)
}

func (r resource) listResource(f *framework.Framework, client dynamic.Interface, resource metav1.APIResource) error {
	if !hasVerb(r.verbs, "list") {
		return nil
	}
	target := fmt.Sprintf("list resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)

	_, err := client.Resource(&resource, f.Namespace.Name).List(metav1.ListOptions{})
	return err
}

func (r resource) createResource(f *framework.Framework, client dynamic.Interface, resource metav1.APIResource, unstruct *unstructuredv1.Unstructured) error {
	if !hasVerb(r.verbs, "create") {
		return nil
	}
	target := fmt.Sprintf("create resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)

	_, err := client.Resource(&resource, f.Namespace.Name).Create(unstruct)
	return err
}

func (r resource) getResource(f *framework.Framework, client dynamic.Interface, resource metav1.APIResource, name string) error {
	if !hasVerb(r.verbs, "get") {
		return nil
	}
	target := fmt.Sprintf("get resource (g: %s, v: %s, r: %s) name: %s", r.group, r.version, r.name, name)
	By(target)

	_, err := client.Resource(&resource, f.Namespace.Name).Get(name, metav1.GetOptions{})
	return err
}

func (r resource) deleteResource(f *framework.Framework, client dynamic.Interface, resource metav1.APIResource, name string) error {
	if !hasVerb(r.verbs, "delete") {
		return nil
	}
	target := fmt.Sprintf("delete resource (g: %s, v: %s, r: %s) name: %s", r.group, r.version, r.name, name)
	By(target)

	err := client.Resource(&resource, f.Namespace.Name).Delete(name, &metav1.DeleteOptions{})
	return err
}
