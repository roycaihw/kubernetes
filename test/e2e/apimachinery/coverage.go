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
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/golang/glog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructuredv1 "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/kubernetes/test/e2e/framework"
)

var (
	patch, _ = json.Marshal(jsonpatch.Patch{})
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
			testResource(f, r, "")
		})
	}
})

func testResource(f *framework.Framework, r *resource, parentName string) {
	gvr := schema.GroupVersionResource{Group: r.group, Version: r.version, Resource: r.name}
	client, err := f.ClientPool.ClientForGroupVersionResource(gvr)
	Expect(err).NotTo(HaveOccurred(), "failed to create dynamic client for resource %v", r)
	apiResource := metav1.APIResource{Name: gvr.Resource, Namespaced: r.namespaced}
	unstruct := r.dumpResourceYAML()
	if parentName == "" {
		parentName = getResourceName(unstruct)
	}

	// Iterate through verbs in serial, skip verbs that don't exist
	err = r.listResource(f, client, apiResource)
	Expect(err).ToNot(HaveOccurred(), "failed to list resource %v", r)
	err = r.createResource(f, client, apiResource, unstruct)
	Expect(err).ToNot(HaveOccurred(), "failed to create resource %v", r)
	unstructGot, err := r.getResource(f, client, apiResource, parentName)
	Expect(err).ToNot(HaveOccurred(), "failed to get resource %v", r)
	err = r.updateResource(f, client, apiResource, unstructGot)
	Expect(err).ToNot(HaveOccurred(), "failed to update resource %v", r)
	err = r.watchResource(f, client, apiResource)
	Expect(err).ToNot(HaveOccurred(), "failed to watch resource %v", r)
	err = r.patchResource(f, client, apiResource, parentName)
	Expect(err).ToNot(HaveOccurred(), "failed to patch resource %v", r)

	for _, sr := range r.subresources {
		testResource(f, sr, parentName)
	}

	// Delete resource after all subresources tested
	err = r.deleteResource(f, client, apiResource, parentName)
	Expect(err).ToNot(HaveOccurred(), "failed to delete resource %v", r)
	err = r.deleteCollectionResource(f, client, apiResource)
	Expect(err).ToNot(HaveOccurred(), "failed to deletecollection resource %v", r)
}

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

		err = parseResourceRecord(tables, record)
		if err != nil {
			return nil, fmt.Errorf("failed to parse resource record %v: %v", record, err)
		}
	}

	return tables, nil
}

func parseResourceRecord(rMap resourceMap, record []string) error {
	// Resource record should be in format of: group,version,name,namespaced,verb
	if len(record) != 5 {
		return fmt.Errorf("unexpected resource record length: %v, want: 5, got: %d", record, len(record))
	}

	namespaced, err := strconv.ParseBool(record[3])
	if err != nil {
		return fmt.Errorf("failed to parse resource (%v) namespaced property: %v", record, err)
	}

	if !strings.Contains(record[2], "/") {
		// If record is a resource
		m := resourceMeta{
			group:   record[0],
			version: record[1],
			name:    record[2],
		}
		if rMap[m] == nil {
			rMap[m] = new(resource)
			*rMap[m] = resource{
				group:        record[0],
				version:      record[1],
				name:         record[2],
				namespaced:   namespaced,
				verbs:        []string{record[4]},
				subresources: resourceMap{},
			}
		} else {
			rMap[m].verbs = append(rMap[m].verbs, record[4])
		}
	} else {
		// If record is a subresource
		parent := resourceMeta{
			group:   record[0],
			version: record[1],
			name:    strings.Split(record[2], "/")[0],
		}
		// Enforce preparing resource before subresource in test data
		if rMap[parent] == nil {
			return fmt.Errorf("parent resource not found for subresource: %v", record)
		}

		m := resourceMeta{
			group:   record[0],
			version: record[1],
			name:    record[2],
		}
		if rMap[parent].subresources[m] == nil {
			rMap[parent].subresources[m] = new(resource)
			*rMap[parent].subresources[m] = resource{
				group:        record[0],
				version:      record[1],
				name:         record[2],
				namespaced:   namespaced,
				verbs:        []string{record[4]},
				subresources: resourceMap{},
			}
		} else {
			rMap[parent].subresources[m].verbs = append(rMap[parent].subresources[m].verbs, record[4])
		}
	}

	return nil
}

func hasVerb(verbs []string, verb string) bool {
	for _, v := range verbs {
		if v == verb {
			return true
		}
	}
	return false
}

func (r *resource) dumpResourceYAML() *unstructuredv1.Unstructured {
	unstruct := &unstructuredv1.Unstructured{}
	if !hasVerb(r.verbs, "create") {
		return unstruct
	}

	// TODO (roycaihw): use framework reporoot or bindata
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
	// TODO (roycaihw): error handling
	return unstruct.Object["metadata"].(map[string]interface{})["name"].(string)
}

func (r *resource) listResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource) error {
	if !hasVerb(r.verbs, "list") {
		return nil
	}
	target := fmt.Sprintf("list resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)

	_, err := client.Resource(&apiResource, f.Namespace.Name).List(metav1.ListOptions{})
	return err
}

func (r *resource) createResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource, unstruct *unstructuredv1.Unstructured) error {
	if !hasVerb(r.verbs, "create") {
		return nil
	}
	target := fmt.Sprintf("create resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)

	_, err := client.Resource(&apiResource, f.Namespace.Name).Create(unstruct)
	return err
}

func (r *resource) updateResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource, unstruct *unstructuredv1.Unstructured) error {
	if !hasVerb(r.verbs, "update") {
		return nil
	}
	target := fmt.Sprintf("update resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)

	_, err := client.Resource(&apiResource, f.Namespace.Name).Update(unstruct)
	return err
}

func (r *resource) getResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource, name string) (*unstructuredv1.Unstructured, error) {
	if !hasVerb(r.verbs, "get") {
		return &unstructuredv1.Unstructured{}, nil
	}
	target := fmt.Sprintf("get resource (g: %s, v: %s, r: %s) name: %s", r.group, r.version, r.name, name)
	By(target)

	unstructGot, err := client.Resource(&apiResource, f.Namespace.Name).Get(name, metav1.GetOptions{})
	return unstructGot, err
}

func (r *resource) patchResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource, name string) error {
	if !hasVerb(r.verbs, "patch") {
		return nil
	}
	target := fmt.Sprintf("patch resource (g: %s, v: %s, r: %s) name: %s", r.group, r.version, r.name, name)
	By(target)

	_, err := client.Resource(&apiResource, f.Namespace.Name).Patch(name, types.JSONPatchType, patch)
	return err
}

func (r *resource) watchResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource) error {
	if !hasVerb(r.verbs, "watch") {
		return nil
	}
	target := fmt.Sprintf("watch resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)

	_, err := client.Resource(&apiResource, f.Namespace.Name).Watch(metav1.ListOptions{})
	return err
}

func (r *resource) deleteResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource, name string) error {
	if !hasVerb(r.verbs, "delete") {
		return nil
	}
	target := fmt.Sprintf("delete resource (g: %s, v: %s, r: %s) name: %s", r.group, r.version, r.name, name)
	By(target)

	err := client.Resource(&apiResource, f.Namespace.Name).Delete(name, &metav1.DeleteOptions{})
	return err
}

func (r *resource) deleteCollectionResource(f *framework.Framework, client dynamic.Interface, apiResource metav1.APIResource) error {
	if !hasVerb(r.verbs, "deletecollection") {
		return nil
	}
	target := fmt.Sprintf("deletecollection resource (g: %s, v: %s, r: %s)", r.group, r.version, r.name)
	By(target)

	err := client.Resource(&apiResource, f.Namespace.Name).DeleteCollection(&metav1.DeleteOptions{}, metav1.ListOptions{})
	return err
}
