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

package apiserverleasegc

import (
	"context"
	"testing"
	"time"

	coordinationv1 "k8s.io/api/coordination/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/features"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	kubeapiservertesting "k8s.io/kubernetes/cmd/kube-apiserver/app/testing"
	"k8s.io/kubernetes/pkg/controller/apiserverleasegc"
	"k8s.io/kubernetes/test/integration/framework"
	"k8s.io/utils/pointer"
)

const (
	testLeaseName string = "apiserver-lease-test"
	// TODO(roycaihw): change to namespaceAPIServerLease afer #95533 is merged
	namespaceAPIServerLease string = "kube-apiserver-lease"
)

func TestLeaseGarbageCollection(t *testing.T) {
	// TODO(roycaihw): enable the test after #95533 is merged
	if true {
		return
	}
	defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.APIServerIdentity, true)()
	result := kubeapiservertesting.StartTestServerOrDie(t, nil, nil, framework.SharedEtcd())
	defer result.TearDownFn()

	kubeclient, err := kubernetes.NewForConfig(result.ClientConfig)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	informers := informers.NewSharedInformerFactory(kubeclient, time.Second)
	leaseInformer := informers.Coordination().V1().Leases()
	controller := apiserverleasegc.NewAPIServerLeaseGC(kubeclient, leaseInformer, time.Hour)

	stopCh := make(chan struct{})
	defer close(stopCh)
	go leaseInformer.Informer().Run(stopCh)
	go controller.Run(stopCh)

	t.Logf(`Waiting the kube-apiserver lease Namespace to be created`)
	if err := wait.PollImmediate(500*time.Millisecond, 10*time.Second, func() (bool, error) {
		if _, err := kubeclient.CoreV1().Namespaces().Get(context.TODO(), namespaceAPIServerLease, metav1.GetOptions{}); err != nil {
			return false, nil
		}
		return true, nil
	}); err != nil {
		t.Fatalf("Failed to see the kube-apiserver lease Namespace: %v", err)
	}

	_, err = kubeclient.CoordinationV1().Leases(namespaceAPIServerLease).Create(context.TODO(), leaseWithRenewTime(time.Now().Add(-2*time.Hour)), metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Unexpected error creating expired lease: %v", err)
	}
	t.Logf(`Waiting the expired lease to be garbage collected`)
	if err := wait.PollImmediate(500*time.Millisecond, 10*time.Second, func() (bool, error) {
		_, err := kubeclient.CoordinationV1().Leases(namespaceAPIServerLease).Get(context.TODO(), testLeaseName, metav1.GetOptions{})
		if err == nil {
			return false, nil
		}
		if apierrors.IsNotFound(err) {
			return true, nil
		}
		return false, err
	}); err != nil {
		t.Fatalf("Failed to see the expired lease garbage collected: %v", err)
	}

	_, err = kubeclient.CoordinationV1().Leases(namespaceAPIServerLease).Create(context.TODO(), leaseWithRenewTime(time.Now().Add(-2*time.Minute)), metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("Unexpected error creating valid lease: %v", err)
	}
	t.Logf(`Making sure the valid lease is not garbage collected`)
	if err := wait.PollImmediate(500*time.Millisecond, 10*time.Second, func() (bool, error) {
		_, err := kubeclient.CoordinationV1().Leases(namespaceAPIServerLease).Get(context.TODO(), testLeaseName, metav1.GetOptions{})
		if err != nil && apierrors.IsNotFound(err) {
			return true, nil
		}
		return false, nil
	}); err == nil {
		t.Fatalf("Unexpected valid lease getting garbage collected")
	}
	_, err = kubeclient.CoordinationV1().Leases(namespaceAPIServerLease).Get(context.TODO(), testLeaseName, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to retrieve valid lease: %v", err)
	}
	err = kubeclient.CoordinationV1().Leases(namespaceAPIServerLease).Delete(context.TODO(), testLeaseName, metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to clean up valid lease: %v", err)
	}
}

func leaseWithRenewTime(t time.Time) *coordinationv1.Lease {
	return &coordinationv1.Lease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testLeaseName,
			Namespace: namespaceAPIServerLease,
		},
		Spec: coordinationv1.LeaseSpec{
			HolderIdentity:       pointer.StringPtr(testLeaseName),
			LeaseDurationSeconds: pointer.Int32Ptr(3600),
			AcquireTime:          &metav1.MicroTime{Time: t},
			RenewTime:            &metav1.MicroTime{Time: t},
		},
	}
}
