/*
Copyright 2014 The Kubernetes Authors.

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

package admission

import (
	"fmt"
	"strings"
	"sync"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation"
	auditinternal "k8s.io/apiserver/pkg/apis/audit"
	"k8s.io/apiserver/pkg/audit"
	"k8s.io/apiserver/pkg/authentication/user"
)

type attributesRecord struct {
	kind        schema.GroupVersionKind
	namespace   string
	name        string
	resource    schema.GroupVersionResource
	subresource string
	operation   Operation
	options     runtime.Object
	dryRun      bool
	object      runtime.Object
	oldObject   runtime.Object
	userInfo    user.Info

	// other elements are always accessed in single goroutine.
	// But ValidatingAdmissionWebhook add annotations concurrently.
	annotations     map[string]audit.Annotation
	annotationsLock sync.RWMutex
}

func NewAttributesRecord(object runtime.Object, oldObject runtime.Object, kind schema.GroupVersionKind, namespace, name string, resource schema.GroupVersionResource, subresource string, operation Operation, operationOptions runtime.Object, dryRun bool, userInfo user.Info) Attributes {
	return &attributesRecord{
		kind:        kind,
		namespace:   namespace,
		name:        name,
		resource:    resource,
		subresource: subresource,
		operation:   operation,
		options:     operationOptions,
		dryRun:      dryRun,
		object:      object,
		oldObject:   oldObject,
		userInfo:    userInfo,
	}
}

func (record *attributesRecord) GetKind() schema.GroupVersionKind {
	return record.kind
}

func (record *attributesRecord) GetNamespace() string {
	return record.namespace
}

func (record *attributesRecord) GetName() string {
	return record.name
}

func (record *attributesRecord) GetResource() schema.GroupVersionResource {
	return record.resource
}

func (record *attributesRecord) GetSubresource() string {
	return record.subresource
}

func (record *attributesRecord) GetOperation() Operation {
	return record.operation
}

func (record *attributesRecord) GetOperationOptions() runtime.Object {
	return record.options
}

func (record *attributesRecord) IsDryRun() bool {
	return record.dryRun
}

func (record *attributesRecord) GetObject() runtime.Object {
	return record.object
}

func (record *attributesRecord) GetOldObject() runtime.Object {
	return record.oldObject
}

func (record *attributesRecord) GetUserInfo() user.Info {
	return record.userInfo
}

// getAnnotations implements privateAnnotationsGetter.It's a private method used
// by WithAudit decorator.
func (record *attributesRecord) getAnnotations() map[string]audit.Annotation {
	record.annotationsLock.RLock()
	defer record.annotationsLock.RUnlock()

	if record.annotations == nil {
		return nil
	}
	cp := make(map[string]audit.Annotation, len(record.annotations))
	for key, value := range record.annotations {
		cp[key] = value
	}
	return cp
}

func (record *attributesRecord) AddAnnotation(key, value string) error {
	return record.AddAnnotationWithLevel(key, value, auditinternal.LevelMetadata)
}

func (record *attributesRecord) AddAnnotationWithLevel(key, value string, level auditinternal.Level) error {
	if err := checkKeyFormat(key); err != nil {
		return err
	}
	if level != auditinternal.LevelMetadata && level != auditinternal.LevelRequest {
		return fmt.Errorf("admission annotations are not allowed to be set at level other than Metadata and Request, key: %q, level: %s", key, level)
	}
	record.annotationsLock.Lock()
	defer record.annotationsLock.Unlock()

	if record.annotations == nil {
		record.annotations = make(map[string]audit.Annotation)
	}
	annotation := audit.Annotation{Level: level, Value: value}
	if v, ok := record.annotations[key]; ok && v != annotation {
		return fmt.Errorf("admission annotations are not allowd to be overwritten, key:%q, old value: %v, new value: %v", key, record.annotations[key], annotation)
	}
	record.annotations[key] = annotation
	return nil
}

func checkKeyFormat(key string) error {
	parts := strings.Split(key, "/")
	if len(parts) != 2 {
		return fmt.Errorf("annotation key has invalid format, the right format is a DNS subdomain prefix and '/' and key name. (e.g. 'podsecuritypolicy.admission.k8s.io/admit-policy')")
	}
	if msgs := validation.IsQualifiedName(key); len(msgs) != 0 {
		return fmt.Errorf("annotation key has invalid format %s. A qualified name like 'podsecuritypolicy.admission.k8s.io/admit-policy' is required.", strings.Join(msgs, ","))
	}
	return nil
}
