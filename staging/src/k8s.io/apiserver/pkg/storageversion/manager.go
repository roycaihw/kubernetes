/*
Copyright 2020 The Kubernetes Authors.

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

package storageversion

import (
	"fmt"
	"sync"
	"sync/atomic"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	apiserverclientset "k8s.io/apiserver/pkg/client/clientset_generated/clientset"
	"k8s.io/client-go/rest"
	_ "k8s.io/component-base/metrics/prometheus/workqueue" // for workqueue metric registration
	"k8s.io/klog/v2"
)

// ResourceInfo contains the information to register the resource to the
// storage version API.
type ResourceInfo struct {
	Resource metav1.APIResource
	// We use a standalone Group instead of reusing the Resource.Group
	// because Resource.Group is often omitted, see the comment on
	// Resource.Group for why it's omitted.
	Group                    string
	EncodingVersion          string
	DecodableVersions        []string
	EquivalentResourceMapper runtime.EquivalentResourceRegistry
}

// Manager records the resources whose StorageVersions need updates, and provides a method to update those StorageVersions.
type Manager interface {
	// AddResourceInfo adds ResourceInfo to the manager.
	AddResourceInfo(resources ...*ResourceInfo)
	// MarkCompletedResourceInfo marks updated ResourceInfo as completed.
	MarkCompletedResourceInfo(r *ResourceInfo)
	// RecordErrorResourceInfo records latest error updating ResourceInfo.
	RecordErrorResourceInfo(r *ResourceInfo, err error)
	// UpdatesPending returns if the StorageVersion of a resource is still wait to be updated.
	UpdatesPending(gv schema.GroupResource) bool

	// UpdateStorageVersions updates the StorageVersions.
	UpdateStorageVersions(kubeAPIServerClientConfig *rest.Config, apiserverID string)
	// Completed returns if updating StorageVersions has completed.
	Completed() bool
}

var _ Manager = &DefaultManager{}

// NewDefaultManager creates a new DefaultManager.
func NewDefaultManager() *DefaultManager {
	s := &DefaultManager{}
	s.completed.Store(false)
	s.groupResources = make(map[schema.GroupResource]*resourceStatus)
	s.resources = make(map[*ResourceInfo]struct{})
	return s
}

// AddResourceInfo adds ResourceInfo to the manager.
// This is not thread-safe. It is expected to be called when the apiserver is installing the endpoints, which is done serially.
func (s *DefaultManager) AddResourceInfo(resources ...*ResourceInfo) {
	for _, r := range resources {
		s.resources[r] = struct{}{}
		s.addGroupResourceFor(r)
	}
}

func (s *DefaultManager) addGroupResourceFor(r *ResourceInfo) {
	gvrs := r.EquivalentResourceMapper.EquivalentResourcesFor(schema.GroupVersionResource{
		Group:    r.Group,
		Resource: r.Resource.Name,
	}, "")
	for _, gvr := range gvrs {
		s.groupResources[gvr.GroupResource()] = &resourceStatus{resourceInfo: r}
	}
}

// MarkCompletedResourceInfo marks updated ResourceInfo as completed.
// It is not safe to call this function concurrently with AddResourceInfo.
func (s *DefaultManager) MarkCompletedResourceInfo(r *ResourceInfo) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.resources, r)
	s.markCompletedGroupResourceFor(r)
}

func (s *DefaultManager) markCompletedGroupResourceFor(r *ResourceInfo) {
	gvrs := r.EquivalentResourceMapper.EquivalentResourcesFor(schema.GroupVersionResource{
		Group:    r.Group,
		Resource: r.Resource.Name,
	}, "")
	for _, gvr := range gvrs {
		s.markCompletedGroupResource(gvr.GroupResource())
	}
}

func (s *DefaultManager) markCompletedGroupResource(gr schema.GroupResource) {
	if _, ok := s.groupResources[gr]; !ok {
		return
	}
	s.groupResources[gr].done = true
	s.groupResources[gr].lastErr = nil
}

// RecordErrorResourceInfo records latest error updating ResourceInfo.
// It is not safe to call this function concurrently with AddResourceInfo.
func (s *DefaultManager) RecordErrorResourceInfo(r *ResourceInfo, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.resources, r)
	s.recordErrorGroupResourceFor(r, err)
}

func (s *DefaultManager) recordErrorGroupResourceFor(r *ResourceInfo, err error) {
	gvrs := r.EquivalentResourceMapper.EquivalentResourcesFor(schema.GroupVersionResource{
		Group:    r.Group,
		Resource: r.Resource.Name,
	}, "")
	for _, gvr := range gvrs {
		s.recordErrorGroupResource(gvr.GroupResource(), err)
	}
}

func (s *DefaultManager) recordErrorGroupResource(gr schema.GroupResource, err error) {
	if _, ok := s.groupResources[gr]; !ok {
		return
	}
	s.groupResources[gr].lastErr = err
}

// UpdatesPending returns if the StorageVersion of a resource is still wait to be updated.
func (s *DefaultManager) UpdatesPending(gv schema.GroupResource) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.groupResources[gv]; !ok {
		return false
	}
	return !s.groupResources[gv].done
}

// DefaultManager indicates if the aggregator, kube-apiserver, and the
// apiextensions-apiserver have completed reporting their storage versions.
type DefaultManager struct {
	completed atomic.Value

	mu             sync.RWMutex
	resources      map[*ResourceInfo]struct{}
	groupResources map[schema.GroupResource]*resourceStatus
}

type resourceStatus struct {
	done         bool
	lastErr      error
	resourceInfo *ResourceInfo
}

// setComplete marks the completion of updating StorageVersions. No write requests need to be blocked anymore.
func (s *DefaultManager) setComplete() {
	s.completed.Store(true)
}

// Completed returns if updating StorageVersions has completed.
func (s *DefaultManager) Completed() bool {
	return s.completed.Load().(bool)
}

func decodableVersions(e runtime.EquivalentResourceRegistry, group string, resource string) []string {
	var versions []string
	decodingGVRs := e.EquivalentResourcesFor(schema.GroupVersionResource{
		Group:    group,
		Resource: resource,
	}, "")
	for _, v := range decodingGVRs {
		versions = append(versions, v.GroupVersion().String())
	}
	return versions
}

// UpdateStorageVersions updates the StorageVersions. If the updates are
// successful, following calls to Completed() returns true.
func (s *DefaultManager) UpdateStorageVersions(kubeAPIServerClientConfig *rest.Config, serverID string) {
	cfg := rest.AddUserAgent(kubeAPIServerClientConfig, "system:kube-apiserver")
	clientset, err := apiserverclientset.NewForConfig(cfg)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to get clientset: %v", err))
		return
	}
	sc := clientset.InternalV1alpha1().StorageVersions()

	s.mu.RLock()
	resources := s.resources
	s.mu.RUnlock()
	hasFailure := false
	for r := range resources {
		r.DecodableVersions = decodableVersions(r.EquivalentResourceMapper, r.Group, r.Resource.Name)
		if err := updateStorageVersionFor(sc, serverID, r.Group+"."+r.Resource.Name, r.EncodingVersion, r.DecodableVersions); err != nil {
			utilruntime.HandleError(fmt.Errorf("failed to update storage version for %v", r.Resource.Name))
			s.RecordErrorResourceInfo(r, err)
			hasFailure = true
			continue
		}
		klog.V(2).Infof("successfully updated storage version for %v", r.Resource.Name)
		s.MarkCompletedResourceInfo(r)
	}
	if hasFailure {
		return
	}
	klog.V(2).Infof("storage version updates complete")
	s.setComplete()
}
