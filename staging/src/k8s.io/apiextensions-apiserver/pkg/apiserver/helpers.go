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

// getCRDSchemaForVersion returns the validation schema for given version in given CRD.
func getCRDSchemaForVersion(crd *apiextensions.CustomResourceDefinition, version string) (*apiextensions.CustomResourceValidation, error) {
	for _, v := range crd.Spec.Versions {
		if version != v.Name {
			continue
		}
		if hasPerVersionSchema(crd.Spec.Versions) {
			return v.Schema, nil
		}
		return crd.Spec.Validation, nil
	}
	return nil, fmt.Errorf("version %s not found in CustomResourceDefinition: %v", version, crd.Name)
}

// getCRDSubresourcesForVersion returns the subresources for given version in given CRD.
func getCRDSubresourcesForVersion(crd *apiextensions.CustomResourceDefinition, version string) (*apiextensions.CustomResourceSubresources, error) {
	for _, v := range crd.Spec.Versions {
		if version != v.Name {
			continue
		}
		if hasPerVersionSubresources(crd.Spec.Versions) {
			return v.Subresources, nil
		}
		return crd.Spec.Subresources, nil
	}
	return nil, fmt.Errorf("version %s not found in CustomResourceDefinition: %v", version, crd.Name)
}

// getCRDColumnsForVersion returns the columns for given version in given CRD.
func getCRDColumnsForVersion(crd *apiextensions.CustomResourceDefinition, version string) ([]apiextensions.CustomResourceColumnDefinition, error) {
	for _, v := range crd.Spec.Versions {
		if version != v.Name {
			continue
		}
		if hasPerVersionColumns(crd.Spec.Versions) {
			return v.AdditionalPrinterColumns, nil
		}
		return crd.Spec.AdditionalPrinterColumns, nil
	}
	return nil, fmt.Errorf("version %s not found in CustomResourceDefinition: %v", version, crd.Name)
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
