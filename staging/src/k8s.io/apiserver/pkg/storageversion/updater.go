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
	"context"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/apis/apiserverinternal/v1alpha1"
	"k8s.io/klog/v2"
)

// Client has the methods required to update the storage version.
type Client interface {
	Create(context.Context, *v1alpha1.StorageVersion, metav1.CreateOptions) (*v1alpha1.StorageVersion, error)
	Update(context.Context, *v1alpha1.StorageVersion, metav1.UpdateOptions) (*v1alpha1.StorageVersion, error)
	Get(context.Context, string, metav1.GetOptions) (*v1alpha1.StorageVersion, error)
}

func setCommonEncodingVersion(sv *v1alpha1.StorageVersion) {
	if len(sv.Status.StorageVersions) == 0 {
		return
	}
	firstVersion := sv.Status.StorageVersions[0].EncodingVersion
	agreed := true
	for _, ssv := range sv.Status.StorageVersions {
		if ssv.EncodingVersion != firstVersion {
			agreed = false
		}
	}
	if agreed {
		sv.Status.CommonEncodingVersion = &firstVersion
	} else {
		sv.Status.CommonEncodingVersion = nil
	}
}

// updateStorageVersionFor updates the storage version object for the resource.
// resource is of the format "<group>.<resource>".
// TODO: split the resource parameter to two.
func updateStorageVersionFor(c Client, apiserverID string, resource string, encodingVersion string, decodableVersions []string) error {
	retries := 3
	var retry int
	var err error
	for retry < retries {
		err = singleUpdate(c, apiserverID, resource, encodingVersion, decodableVersions)
		if err == nil {
			return nil
		}
		if apierrors.IsAlreadyExists(err) || apierrors.IsConflict(err) {
			time.Sleep(1 * time.Second)
			continue
		}
		if err != nil {
			klog.Errorf("retry %d, failed to update storage version for %s: %v", retry, resource, err)
			retry++
			time.Sleep(1 * time.Second)
		}
	}
	return err
}

func singleUpdate(c Client, apiserverID, resource, encodingVersion string, decodableVersions []string) error {
	shouldCreate := false
	sv, err := c.Get(context.TODO(), resource, metav1.GetOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}
	if apierrors.IsNotFound(err) {
		shouldCreate = true
		sv = &v1alpha1.StorageVersion{}
		sv.ObjectMeta.Name = resource
	}
	updatedSV := localUpdateStorageVersion(sv, apiserverID, encodingVersion, decodableVersions)
	if shouldCreate {
		_, err := c.Create(context.TODO(), updatedSV, metav1.CreateOptions{})
		return err
	}
	_, err = c.Update(context.TODO(), updatedSV, metav1.UpdateOptions{})
	return err
}

func localUpdateStorageVersion(sv *v1alpha1.StorageVersion, apiserverID, encodingVersion string, decodableVersions []string) *v1alpha1.StorageVersion {
	updatedSV := sv.DeepCopy()
	newSSV := v1alpha1.ServerStorageVersion{
		APIServerID:       apiserverID,
		EncodingVersion:   encodingVersion,
		DecodableVersions: decodableVersions,
	}
	foundSSV := false
	for i, ssv := range updatedSV.Status.StorageVersions {
		if ssv.APIServerID == apiserverID {
			updatedSV.Status.StorageVersions[i] = newSSV
			foundSSV = true
		}
	}
	if !foundSSV {
		updatedSV.Status.StorageVersions = append(updatedSV.Status.StorageVersions, newSSV)
	}
	setCommonEncodingVersion(updatedSV)
	return updatedSV
}
