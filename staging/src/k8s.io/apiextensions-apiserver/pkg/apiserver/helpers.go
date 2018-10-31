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

package apiserver

import (
	"fmt"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
)

// getCRDSchemaForVersionOrDie returns the validation schema for given version in given CRD.
func getCRDSchemaForVersionOrDie(crd *apiextensions.CustomResourceDefinition, version string) *apiextensions.CustomResourceValidation {
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
func getCRDSubresourcesForVersionOrDie(crd *apiextensions.CustomResourceDefinition, version string) *apiextensions.CustomResourceSubresources {
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
func getCRDColumnsForVersionOrDie(crd *apiextensions.CustomResourceDefinition, version string) []apiextensions.CustomResourceColumnDefinition {
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
func hasPerVersionSchema(versions []apiextensions.CustomResourceDefinitionVersion) bool {
	for _, v := range versions {
		if v.Schema != nil {
			return true
		}
	}
	return false
}

// hasPerVersionSubresources returns true if a CRD uses per-version subresources.
func hasPerVersionSubresources(versions []apiextensions.CustomResourceDefinitionVersion) bool {
	for _, v := range versions {
		if v.Subresources != nil {
			return true
		}
	}
	return false
}

// hasPerVersionColumns returns true if a CRD uses per-version columns.
func hasPerVersionColumns(versions []apiextensions.CustomResourceDefinitionVersion) bool {
	for _, v := range versions {
		if len(v.AdditionalPrinterColumns) > 0 {
			return true
		}
	}
	return false
}
