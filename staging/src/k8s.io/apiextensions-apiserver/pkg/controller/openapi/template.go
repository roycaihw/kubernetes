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

package openapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-openapi/spec"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiserver/pkg/endpoints/openapi"
)

const (
	tmpGroup         = "tmpGroup.io"
	tmpGroupFriendly = "io.tmpGroup"
	tmpGroupTag      = "tmpGroupIo"

	tmpVersion  = "tmpVersion"
	tmpKind     = "tmpKind"
	tmpListKind = "tmpListKind"
	tmpPlural   = "tmpPlural"
)

// Templates define patterns of CRD as template based on scope and subresources
var Templates = []struct {
	Name   string
	Scope  apiextensions.ResourceScope
	Status bool
	Scale  bool
}{
	{
		Name:   "NamespacedTemplateWithStatusScale",
		Scope:  apiextensions.NamespaceScoped,
		Status: true,
		Scale:  true,
	},
	{
		Name:   "NamespacedTemplateWithStatus",
		Scope:  apiextensions.NamespaceScoped,
		Status: true,
		Scale:  false,
	},
	{
		Name:   "NamespacedTemplateWithScale",
		Scope:  apiextensions.NamespaceScoped,
		Status: false,
		Scale:  true,
	},
	{
		Name:   "NamespacedTemplate",
		Scope:  apiextensions.NamespaceScoped,
		Status: false,
		Scale:  false,
	},
	{
		Name:   "TemplateWithStatusScale",
		Scope:  apiextensions.ClusterScoped,
		Status: true,
		Scale:  true,
	},
	{
		Name:   "TemplateWithStatus",
		Scope:  apiextensions.ClusterScoped,
		Status: true,
		Scale:  false,
	},
	{
		Name:   "TemplateWithScale",
		Scope:  apiextensions.ClusterScoped,
		Status: false,
		Scale:  true,
	},
	{
		Name:   "Template",
		Scope:  apiextensions.ClusterScoped,
		Status: false,
		Scale:  false,
	},
}

func templateCRD() *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   tmpGroup,
			Version: tmpVersion,
			Names: apiextensions.CustomResourceDefinitionNames{
				Plural:   tmpPlural,
				Kind:     tmpKind,
				ListKind: tmpListKind,
			},
			Scope: apiextensions.ClusterScoped,
			Subresources: &apiextensions.CustomResourceSubresources{
				Status: &apiextensions.CustomResourceSubresourceStatus{},
				Scale:  &apiextensions.CustomResourceSubresourceScale{},
			},
		},
	}
}

// SwaggerTemplate builds CRD swagger template based on scope and subresources
func SwaggerTemplate(scope apiextensions.ResourceScope, status, scale bool) *spec.Swagger {
	crd := templateCRD()
	crd.Spec.Scope = scope
	if !status {
		crd.Spec.Subresources.Status = nil
	}
	if !scale {
		crd.Spec.Subresources.Scale = nil
	}
	swagger, _, err := BuildSwagger(crd, tmpVersion)
	if err != nil {
		// panic on programming error
		panic("unexpected error: failed to build swagger template")
	}
	return swagger
}

func buildSwaggerTemplating(crd *apiextensions.CustomResourceDefinition, version string) (*spec.Swagger, string, error) {
	return nil, "", nil
	// subresources, err := getSubresourcesForVersion(&crd.Spec, version)
	// if err != nil {
	// 	return nil, "", err
	// }
	// status := subresources != nil && subresources.Status != nil
	// scale := subresources != nil && subresources.Scale != nil

	// if crd.Spec.Scope == apiextensions.NamespaceScoped {
	// 	if status && scale {
	// 		return buildSwaggerWithTemplate(NamespacedTemplateWithStatusScale, crd, version)
	// 	}
	// 	if status {
	// 		return buildSwaggerWithTemplate(NamespacedTemplateWithStatus, crd, version)
	// 	}
	// 	if scale {
	// 		return buildSwaggerWithTemplate(NamespacedTemplateWithScale, crd, version)
	// 	}
	// 	return buildSwaggerWithTemplate(NamespacedTemplate, crd, version)
	// }
	// if status && scale {
	// 	return buildSwaggerWithTemplate(TemplateWithStatusScale, crd, version)
	// }
	// if status {
	// 	return buildSwaggerWithTemplate(TemplateWithStatus, crd, version)
	// }
	// if scale {
	// 	return buildSwaggerWithTemplate(TemplateWithScale, crd, version)
	// }
	// return buildSwaggerWithTemplate(Template, crd, version)
}

func buildSwaggerWithTemplate(template *spec.Swagger, crd *apiextensions.CustomResourceDefinition, version string) (*spec.Swagger, string, error) {
	templateJSON, err := json.Marshal(template)
	if err != nil {
		return nil, "", err
	}

	group := crd.Spec.Group
	groupFriendly := openapi.FriendlyName(group)
	groupTag := camelCase(group)
	kind := crd.Spec.Names.Kind
	listKind := crd.Spec.Names.ListKind
	plural := crd.Spec.Names.Plural

	templateStr := string(templateJSON)
	templateStr = strings.Replace(templateStr, tmpGroup, group, -1)
	templateStr = strings.Replace(templateStr, tmpGroupFriendly, groupFriendly, -1)
	templateStr = strings.Replace(templateStr, tmpGroupTag, groupTag, -1)
	templateStr = strings.Replace(templateStr, tmpVersion, version, -1)
	templateStr = strings.Replace(templateStr, tmpKind, kind, -1)
	templateStr = strings.Replace(templateStr, tmpListKind, listKind, -1)
	templateStr = strings.Replace(templateStr, tmpPlural, plural, -1)

	templateStr = strings.Replace(templateStr, strings.Title(tmpGroup), strings.Title(group), -1)
	templateStr = strings.Replace(templateStr, strings.Title(tmpGroupFriendly), strings.Title(groupFriendly), -1)
	templateStr = strings.Replace(templateStr, strings.Title(tmpGroupTag), strings.Title(groupTag), -1)
	templateStr = strings.Replace(templateStr, strings.Title(tmpVersion), strings.Title(version), -1)
	templateStr = strings.Replace(templateStr, strings.Title(tmpKind), strings.Title(kind), -1)
	templateStr = strings.Replace(templateStr, strings.Title(tmpListKind), strings.Title(listKind), -1)
	templateStr = strings.Replace(templateStr, strings.Title(tmpPlural), strings.Title(plural), -1)

	swagger := &spec.Swagger{}
	if err := json.Unmarshal([]byte(templateStr), swagger); err != nil {
		return nil, "", err
	}

	s, err := getSchemaForVersion(crd, version)
	if err != nil {
		return nil, "", err
	}
	if s == nil && s.OpenAPIV3Schema == nil {
		// Return if CRD doesn't have schema
		return swaggerWithETag(swagger)
	}

	// Fill in schema to template
	schema, err := ConvertJSONSchemaPropsToOpenAPIv2Schema(s.OpenAPIV3Schema)
	if err != nil {
		return nil, "", err
	}
	b := newBuilder(crd, version, schema)
	schema = b.buildKubeNative(schema)
	swagger.Definitions[fmt.Sprintf("%s.%s.%s", groupFriendly, version, kind)] = *schema

	return swaggerWithETag(swagger)
}

func camelCase(name string) string {
	parts := strings.Split(name, ".")
	for i := range parts {
		if i == 0 {
			continue
		}
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
