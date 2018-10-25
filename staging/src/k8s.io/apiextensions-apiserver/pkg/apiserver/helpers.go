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

// getCRDSchemaForVersion returns the validation schema for given version in given CRD Spec.
func getCRDSchemaForVersion(crdSpec *apiextensions.CustomResourceDefinitionSpec, version string) (*apiextensions.CustomResourceValidation, error) {
	for _, v := range crdSpec.Versions {
		if version != v.Name {
			continue
		}
		if crdSpec.Validation != nil && v.Schema != nil {
			return nil, fmt.Errorf("malformed CustomResourceDefinitionSpec: top-level and per-version schemas must be mutual exclusive; spec: %v, version: %s", *crdSpec, version)
		}
		if v.Schema != nil {
			// For backwards compatibility with existing code path, we wrap the
			// OpenAPIV3Schema into a CustomResourceValidation struct
			return &apiextensions.CustomResourceValidation{
				OpenAPIV3Schema: v.Schema,
			}, nil
		}
		return crdSpec.Validation, nil
	}
	return nil, fmt.Errorf("version %s not found in CustomResourceDefinitionSpec: %v", version, *crdSpec)
}

// getCRDSubresourcesForVersion returns the subresources for given version in given CRD Spec.
func getCRDSubresourcesForVersion(crdSpec *apiextensions.CustomResourceDefinitionSpec, version string) (*apiextensions.CustomResourceSubresources, error) {
	for _, v := range crdSpec.Versions {
		if version != v.Name {
			continue
		}
		if crdSpec.Subresources != nil && v.Subresources != nil {
			return nil, fmt.Errorf("malformed CustomResourceDefinitionSpec: top-level and per-version subresources must be mutual exclusive; spec: %v, version: %s", *crdSpec, version)
		}
		if v.Subresources != nil {
			return v.Subresources, nil
		}
		return crdSpec.Subresources, nil
	}
	return nil, fmt.Errorf("version %s not found in CustomResourceDefinitionSpec: %v", version, *crdSpec)
}

// getCRDColumnsForVersion returns the columns for given version in given CRD Spec.
func getCRDColumnsForVersion(crdSpec *apiextensions.CustomResourceDefinitionSpec, version string) ([]apiextensions.CustomResourceColumnDefinition, error) {
	for _, v := range crdSpec.Versions {
		if version != v.Name {
			continue
		}
		if len(crdSpec.AdditionalPrinterColumns) > 0 && len(v.AdditionalPrinterColumns) > 0 {
			return nil, fmt.Errorf("malformed CustomResourceDefinitionSpec: top-level and per-version additionalPrinterColumns must be mutual exclusive; spec: %v, version: %s", *crdSpec, version)
		}
		if len(v.AdditionalPrinterColumns) > 0 {
			return v.AdditionalPrinterColumns, nil
		}
		return crdSpec.AdditionalPrinterColumns, nil
	}
	return nil, fmt.Errorf("version %s not found in CustomResourceDefinitionSpec: %v", version, *crdSpec)
}
