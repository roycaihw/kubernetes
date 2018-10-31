/*
Copyright 2018 The Kubernetes Authors.

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

package integration

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

func instantiateCustomResource(t *testing.T, instanceToCreate *unstructured.Unstructured, client dynamic.ResourceInterface, definition *apiextensionsv1beta1.CustomResourceDefinition) (*unstructured.Unstructured, error) {
	return instantiateVersionedCustomResource(t, instanceToCreate, client, definition, definition.Spec.Versions[0].Name)
}

func instantiateVersionedCustomResource(t *testing.T, instanceToCreate *unstructured.Unstructured, client dynamic.ResourceInterface, definition *apiextensionsv1beta1.CustomResourceDefinition, version string) (*unstructured.Unstructured, error) {
	createdInstance, err := client.Create(instanceToCreate, metav1.CreateOptions{})
	if err != nil {
		t.Logf("%#v", createdInstance)
		return nil, err
	}
	createdObjectMeta, err := meta.Accessor(createdInstance)
	if err != nil {
		t.Fatal(err)
	}
	// it should have a UUID
	if len(createdObjectMeta.GetUID()) == 0 {
		t.Errorf("missing uuid: %#v", createdInstance)
	}
	createdTypeMeta, err := meta.TypeAccessor(createdInstance)
	if err != nil {
		t.Fatal(err)
	}
	if e, a := definition.Spec.Group+"/"+version, createdTypeMeta.GetAPIVersion(); e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	if e, a := definition.Spec.Names.Kind, createdTypeMeta.GetKind(); e != a {
		t.Errorf("expected %v, got %v", e, a)
	}
	return createdInstance, nil
}

func newNamespacedCustomResourceVersionedClient(ns string, client dynamic.Interface, crd *apiextensionsv1beta1.CustomResourceDefinition, version string) dynamic.ResourceInterface {
	gvr := schema.GroupVersionResource{Group: crd.Spec.Group, Version: version, Resource: crd.Spec.Names.Plural}

	if crd.Spec.Scope != apiextensionsv1beta1.ClusterScoped {
		return client.Resource(gvr).Namespace(ns)
	}
	return client.Resource(gvr)
}

func newNamespacedCustomResourceClient(ns string, client dynamic.Interface, crd *apiextensionsv1beta1.CustomResourceDefinition) dynamic.ResourceInterface {
	return newNamespacedCustomResourceVersionedClient(ns, client, crd, crd.Spec.Versions[0].Name)
}

// updateCustomResourceDefinitionWithRetry updates a CRD, retrying up to 5 times on version conflict errors.
func updateCustomResourceDefinitionWithRetry(client clientset.Interface, name string, update func(*apiextensionsv1beta1.CustomResourceDefinition)) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	for i := 0; i < 5; i++ {
		crd, err := client.ApiextensionsV1beta1().CustomResourceDefinitions().Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to get CustomResourceDefinition %q: %v", name, err)
		}
		update(crd)
		crd, err = client.ApiextensionsV1beta1().CustomResourceDefinitions().Update(crd)
		if err == nil {
			return crd, nil
		}
		if !errors.IsConflict(err) {
			return nil, fmt.Errorf("failed to update CustomResourceDefinition %q: %v", name, err)
		}
	}
	return nil, fmt.Errorf("too many retries after conflicts updating CustomResourceDefinition %q", name)
}

// getCRDSchemaForVersionOrDie returns the validation schema for given version in given CRD.
func getCRDSchemaForVersionOrDie(crd *apiextensionsv1beta1.CustomResourceDefinition, version string) *apiextensionsv1beta1.CustomResourceValidation {
	if !hasPerVersionSchema(crd.Spec.Versions) {
		return crd.Spec.Validation
	}
	for _, v := range crd.Spec.Versions {
		if version != v.Name {
			continue
		}
		return v.Schema
	}
	// Given that we call HasServedCRDVersion when serving CRD, this should not happen.
	// We fail hard to catch any internal bug.
	panic(fmt.Errorf("version %s not found in CustomResourceDefinition: %v", version, crd.Name))
}

// getCRDSubresourcesForVersionOrDie returns the subresources for given version in given CRD.
func getCRDSubresourcesForVersionOrDie(crd *apiextensionsv1beta1.CustomResourceDefinition, version string) *apiextensionsv1beta1.CustomResourceSubresources {
	if !hasPerVersionSubresources(crd.Spec.Versions) {
		return crd.Spec.Subresources
	}
	for _, v := range crd.Spec.Versions {
		if version != v.Name {
			continue
		}
		return v.Subresources
	}
	// Given that we call HasServedCRDVersion when serving CRD, this should not happen.
	// We fail hard to catch any internal bug.
	panic(fmt.Errorf("version %s not found in CustomResourceDefinition: %v", version, crd.Name))
}

// getCRDColumnsForVersionOrDie returns the columns for given version in given CRD.
func getCRDColumnsForVersionOrDie(crd *apiextensionsv1beta1.CustomResourceDefinition, version string) []apiextensionsv1beta1.CustomResourceColumnDefinition {
	if !hasPerVersionColumns(crd.Spec.Versions) {
		return crd.Spec.AdditionalPrinterColumns
	}
	for _, v := range crd.Spec.Versions {
		if version != v.Name {
			continue
		}
		return v.AdditionalPrinterColumns
	}
	// Given that we call HasServedCRDVersion when serving CRD, this should not happen.
	// We fail hard to catch any internal bug.
	panic(fmt.Errorf("version %s not found in CustomResourceDefinition: %v", version, crd.Name))
}

// hasPerVersionSchema returns true if a CRD uses per-version schema.
func hasPerVersionSchema(versions []apiextensionsv1beta1.CustomResourceDefinitionVersion) bool {
	for _, v := range versions {
		if v.Schema != nil {
			return true
		}
	}
	return false
}

// hasPerVersionSubresources returns true if a CRD uses per-version subresources.
func hasPerVersionSubresources(versions []apiextensionsv1beta1.CustomResourceDefinitionVersion) bool {
	for _, v := range versions {
		if v.Subresources != nil {
			return true
		}
	}
	return false
}

// hasPerVersionColumns returns true if a CRD uses per-version columns.
func hasPerVersionColumns(versions []apiextensionsv1beta1.CustomResourceDefinitionVersion) bool {
	for _, v := range versions {
		if len(v.AdditionalPrinterColumns) > 0 {
			return true
		}
	}
	return false
}
