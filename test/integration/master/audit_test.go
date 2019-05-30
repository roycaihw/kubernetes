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

package master

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"k8s.io/api/admission/v1beta1"
	admissionv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	auditinternal "k8s.io/apiserver/pkg/apis/audit"
	auditv1 "k8s.io/apiserver/pkg/apis/audit/v1"
	auditv1beta1 "k8s.io/apiserver/pkg/apis/audit/v1beta1"
	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/client-go/kubernetes"
	kubeapiservertesting "k8s.io/kubernetes/cmd/kube-apiserver/app/testing"
	"k8s.io/kubernetes/test/integration/framework"
	"k8s.io/kubernetes/test/utils"

	jsonpatch "github.com/evanphx/json-patch"
)

var (
	auditPolicyPattern = `
apiVersion: {version}
kind: Policy
rules:
  - level: {level}
    resources:
      - group: "" # core
        resources: ["configmaps"]

`
	namespace              = "default"
	watchTestTimeout int64 = 1
	watchOptions           = metav1.ListOptions{TimeoutSeconds: &watchTestTimeout}
	patch, _               = json.Marshal(jsonpatch.Patch{})
	auditTestUser          = "system:apiserver"
	versions               = map[string]schema.GroupVersion{
		"audit.k8s.io/v1":      auditv1.SchemeGroupVersion,
		"audit.k8s.io/v1beta1": auditv1beta1.SchemeGroupVersion,
	}

	expectedEvents = []utils.AuditEvent{
		{
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseComplete,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps", namespace),
			Verb:              "create",
			Code:              201,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     true,
			ResponseObject:    true,
			AuthorizeDecision: "allow",
		}, {
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseComplete,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps/audit-configmap", namespace),
			Verb:              "get",
			Code:              200,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     false,
			ResponseObject:    true,
			AuthorizeDecision: "allow",
		}, {
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseComplete,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps", namespace),
			Verb:              "list",
			Code:              200,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     false,
			ResponseObject:    true,
			AuthorizeDecision: "allow",
		}, {
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseStarted,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps?timeout=%ds&timeoutSeconds=%d&watch=true", namespace, watchTestTimeout, watchTestTimeout),
			Verb:              "watch",
			Code:              200,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     false,
			ResponseObject:    false,
			AuthorizeDecision: "allow",
		}, {
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseComplete,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps?timeout=%ds&timeoutSeconds=%d&watch=true", namespace, watchTestTimeout, watchTestTimeout),
			Verb:              "watch",
			Code:              200,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     false,
			ResponseObject:    false,
			AuthorizeDecision: "allow",
		}, {
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseComplete,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps/audit-configmap", namespace),
			Verb:              "update",
			Code:              200,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     true,
			ResponseObject:    true,
			AuthorizeDecision: "allow",
		}, {
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseComplete,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps/audit-configmap", namespace),
			Verb:              "patch",
			Code:              200,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     true,
			ResponseObject:    true,
			AuthorizeDecision: "allow",
		}, {
			Level:             auditinternal.LevelRequestResponse,
			Stage:             auditinternal.StageResponseComplete,
			RequestURI:        fmt.Sprintf("/api/v1/namespaces/%s/configmaps/audit-configmap", namespace),
			Verb:              "delete",
			Code:              200,
			User:              auditTestUser,
			Resource:          "configmaps",
			Namespace:         namespace,
			RequestObject:     true,
			ResponseObject:    true,
			AuthorizeDecision: "allow",
		},
	}
)

// TestAudit ensures that both v1beta1 and v1 version audit api could work.
func TestAudit(t *testing.T) {
	tcs := []struct {
		auditLevel            auditinternal.Level
		enableMutatingWebhook bool
	}{
		{
			auditLevel:            auditinternal.LevelRequestResponse,
			enableMutatingWebhook: false,
		},
		{
			auditLevel:            auditinternal.LevelMetadata,
			enableMutatingWebhook: true,
		},
		{
			auditLevel:            auditinternal.LevelRequest,
			enableMutatingWebhook: true,
		},
		{
			auditLevel:            auditinternal.LevelRequestResponse,
			enableMutatingWebhook: true,
		},
	}
	for version := range versions {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s.%s.%t", version, tc.auditLevel, tc.enableMutatingWebhook), func(t *testing.T) {
				testAudit(t, version, tc.auditLevel, tc.enableMutatingWebhook)
			})
		}
	}
}

func testAudit(t *testing.T, version string, level auditinternal.Level, enableMutatingWebhook bool) {
	var url string
	closeFunc := func() {}
	if enableMutatingWebhook {
		url, closeFunc = newWebhookServer(t)
	}
	defer closeFunc()

	// prepare audit policy file
	auditPolicy := strings.Replace(auditPolicyPattern, "{version}", version, 1)
	auditPolicy = strings.Replace(auditPolicy, "{level}", string(level), 1)
	policyFile, err := ioutil.TempFile("", "audit-policy.yaml")
	if err != nil {
		t.Fatalf("Failed to create audit policy file: %v", err)
	}
	defer os.Remove(policyFile.Name())
	if _, err := policyFile.Write([]byte(auditPolicy)); err != nil {
		t.Fatalf("Failed to write audit policy file: %v", err)
	}
	if err := policyFile.Close(); err != nil {
		t.Fatalf("Failed to close audit policy file: %v", err)
	}

	// prepare audit log file
	logFile, err := ioutil.TempFile("", "audit.log")
	if err != nil {
		t.Fatalf("Failed to create audit log file: %v", err)
	}
	defer os.Remove(logFile.Name())

	// start api server
	result := kubeapiservertesting.StartTestServerOrDie(t, nil,
		[]string{
			"--audit-policy-file", policyFile.Name(),
			"--audit-log-version", version,
			"--audit-log-mode", "blocking",
			"--audit-log-path", logFile.Name()},
		framework.SharedEtcd())
	defer result.TearDownFn()

	kubeclient, err := kubernetes.NewForConfig(result.ClientConfig)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if enableMutatingWebhook {
		if err := createV1beta1MutationWebhook(kubeclient, url+"/mutation"); err != nil {
			t.Fatal(err)
		}
	}

	var lastMissingReport string
	if err := wait.Poll(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		// perform configmap operations
		configMapOperations(t, kubeclient)

		// check for corresponding audit logs
		stream, err := os.Open(logFile.Name())
		if err != nil {
			return false, fmt.Errorf("Unexpected error: %v", err)
		}
		defer stream.Close()
		missingReport, err := utils.CheckAuditLines(stream, getExpectedEvents(level, enableMutatingWebhook), versions[version])
		if err != nil {
			return false, fmt.Errorf("Unexpected error: %v", err)
		}
		if len(missingReport.MissingEvents) > 0 {
			lastMissingReport = missingReport.String()
			return false, nil
		}
		return true, nil
	}); err != nil {
		t.Fatalf("failed to get expected events -- missingReport: %s, error: %v", lastMissingReport, err)
	}
}

func getExpectedEvents(level auditinternal.Level, enableMutatingWebhook bool) []utils.AuditEvent {
	if !enableMutatingWebhook {
		return expectedEvents
	}

	var webhookMutationAnnotations, webhookPatchAnnotations map[string]string
	var requestObject, responseObject bool
	if level.GreaterOrEqual(auditinternal.LevelMetadata) {
		webhookMutationAnnotations = map[string]string{}
		webhookMutationAnnotations[utils.MutationAuditAnnotationPrefix+"auditmutation.integration.test"] = "true"
	}
	if level.GreaterOrEqual(auditinternal.LevelRequest) {
		webhookPatchAnnotations = map[string]string{}
		webhookPatchAnnotations[utils.PatchAuditAnnotationPrefix+"auditmutation.integration.test"] = `[{"op":"add","path":"/foo","value":"test"}]`
		requestObject = true
	}
	if level.GreaterOrEqual(auditinternal.LevelRequestResponse) {
		responseObject = true
	}
	return []utils.AuditEvent{
		{
			Level:                      level,
			Stage:                      auditinternal.StageResponseComplete,
			RequestURI:                 fmt.Sprintf("/api/v1/namespaces/%s/configmaps", namespace),
			Verb:                       "create",
			Code:                       201,
			User:                       auditTestUser,
			Resource:                   "configmaps",
			Namespace:                  namespace,
			AuthorizeDecision:          "allow",
			RequestObject:              requestObject,
			ResponseObject:             responseObject,
			WebhookMutationAnnotations: webhookMutationAnnotations,
			WebhookPatchAnnotations:    webhookPatchAnnotations,
		}, {
			Level:                      level,
			Stage:                      auditinternal.StageResponseComplete,
			RequestURI:                 fmt.Sprintf("/api/v1/namespaces/%s/configmaps/audit-configmap", namespace),
			Verb:                       "update",
			Code:                       200,
			User:                       auditTestUser,
			Resource:                   "configmaps",
			Namespace:                  namespace,
			AuthorizeDecision:          "allow",
			RequestObject:              requestObject,
			ResponseObject:             responseObject,
			WebhookMutationAnnotations: webhookMutationAnnotations,
			WebhookPatchAnnotations:    webhookPatchAnnotations,
		},
	}
}

// configMapOperations is a set of known operations performed on the configmap type
// which correspond to the expected events.
// This is shared by the dynamic test
func configMapOperations(t *testing.T, kubeclient kubernetes.Interface) {
	// create, get, watch, update, patch, list and delete configmap.
	configMap := &apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "audit-configmap",
		},
		Data: map[string]string{
			"map-key": "map-value",
		},
	}

	_, err := kubeclient.CoreV1().ConfigMaps(namespace).Create(configMap)
	expectNoError(t, err, "failed to create audit-configmap")

	_, err = kubeclient.CoreV1().ConfigMaps(namespace).Get(configMap.Name, metav1.GetOptions{})
	expectNoError(t, err, "failed to get audit-configmap")

	configMapChan, err := kubeclient.CoreV1().ConfigMaps(namespace).Watch(watchOptions)
	expectNoError(t, err, "failed to create watch for config maps")
	for range configMapChan.ResultChan() {
		// Block until watchOptions.TimeoutSeconds expires.
		// If the test finishes before watchOptions.TimeoutSeconds expires, the watch audit
		// event at stage ResponseComplete will not be generated.
	}

	_, err = kubeclient.CoreV1().ConfigMaps(namespace).Update(configMap)
	expectNoError(t, err, "failed to update audit-configmap")

	_, err = kubeclient.CoreV1().ConfigMaps(namespace).Patch(configMap.Name, types.JSONPatchType, patch)
	expectNoError(t, err, "failed to patch configmap")

	_, err = kubeclient.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	expectNoError(t, err, "failed to list config maps")

	err = kubeclient.CoreV1().ConfigMaps(namespace).Delete(configMap.Name, &metav1.DeleteOptions{})
	expectNoError(t, err, "failed to delete audit-configmap")
}

func expectNoError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}

// admission webhook registration helpers
func newWebhookServer(t *testing.T) (string, func()) {
	// set up webhook server
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(localhostCert) {
		t.Fatal("Failed to append Cert from PEM")
	}
	cert, err := tls.X509KeyPair(localhostCert, localhostKey)
	if err != nil {
		t.Fatalf("Failed to build cert with error: %+v", err)
	}

	webhookMux := http.NewServeMux()
	webhookMux.Handle("/mutation", webhookHandler(t))
	webhookServer := httptest.NewUnstartedServer(webhookMux)
	webhookServer.TLS = &tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{cert},
	}
	webhookServer.StartTLS()
	return webhookServer.URL, webhookServer.Close
}

func webhookHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			t.Errorf("contentType=%s, expect application/json", contentType)
			return
		}

		review := v1beta1.AdmissionReview{}
		if err := json.Unmarshal(data, &review); err != nil {
			t.Errorf("Fail to deserialize object: %s with error: %v", string(data), err)
			http.Error(w, err.Error(), 400)
			return
		}
		gvk := schema.GroupVersionKind{Group: "admission.k8s.io", Version: "v1beta1", Kind: "AdmissionReview"}
		if review.GetObjectKind().GroupVersionKind() != gvk {
			t.Errorf("Invalid admission review kind: %#v", review.GetObjectKind().GroupVersionKind())
			http.Error(w, err.Error(), 400)
			return
		}
		if len(review.Request.Object.Raw) > 0 {
			u := &unstructured.Unstructured{Object: map[string]interface{}{}}
			if err := json.Unmarshal(review.Request.Object.Raw, u); err != nil {
				t.Errorf("Fail to deserialize object: %s with error: %v", string(review.Request.Object.Raw), err)
				http.Error(w, err.Error(), 400)
				return
			}
			review.Request.Object.Object = u
		}
		if len(review.Request.OldObject.Raw) > 0 {
			u := &unstructured.Unstructured{Object: map[string]interface{}{}}
			if err := json.Unmarshal(review.Request.OldObject.Raw, u); err != nil {
				t.Errorf("Fail to deserialize object: %s with error: %v", string(review.Request.OldObject.Raw), err)
				http.Error(w, err.Error(), 400)
				return
			}
			review.Request.OldObject.Object = u
		}

		review.Response = &v1beta1.AdmissionResponse{
			Allowed: true,
			UID:     review.Request.UID,
			Result:  &metav1.Status{Message: "admitted"},
		}
		review.Response.Patch = []byte(`[{"op":"add","path":"/foo","value":"test"}]`)
		jsonPatch := v1beta1.PatchTypeJSONPatch
		review.Response.PatchType = &jsonPatch

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(review); err != nil {
			t.Errorf("Marshal of response failed with error: %v", err)
		}
	})
}

func createV1beta1MutationWebhook(client clientset.Interface, endpoint string) error {
	fail := admissionv1beta1.Fail
	// Attaching Mutation webhook to API server
	_, err := client.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Create(&admissionv1beta1.MutatingWebhookConfiguration{
		ObjectMeta: metav1.ObjectMeta{Name: "auditmutation.integration.test"},
		Webhooks: []admissionv1beta1.Webhook{{
			Name: "auditmutation.integration.test",
			ClientConfig: admissionv1beta1.WebhookClientConfig{
				URL:      &endpoint,
				CABundle: localhostCert,
			},
			Rules: []admissionv1beta1.RuleWithOperations{{
				Operations: []admissionv1beta1.OperationType{admissionv1beta1.Create, admissionv1beta1.Update},
				Rule:       admissionv1beta1.Rule{APIGroups: []string{"*"}, APIVersions: []string{"*"}, Resources: []string{"*/*"}},
			}},
			FailurePolicy:           &fail,
			AdmissionReviewVersions: []string{"v1beta1"},
		}},
	})
	return err
}

// localhostCert was generated from crypto/tls/generate_cert.go with the following command:
//     go run generate_cert.go  --rsa-bits 512 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var localhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIIBjzCCATmgAwIBAgIRAKpi2WmTcFrVjxrl5n5YDUEwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzBcMA0GCSqGSIb3DQEBAQUAA0sAMEgC
QQC9fEbRszP3t14Gr4oahV7zFObBI4TfA5i7YnlMXeLinb7MnvT4bkfOJzE6zktn
59zP7UiHs3l4YOuqrjiwM413AgMBAAGjaDBmMA4GA1UdDwEB/wQEAwICpDATBgNV
HSUEDDAKBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MC4GA1UdEQQnMCWCC2V4
YW1wbGUuY29thwR/AAABhxAAAAAAAAAAAAAAAAAAAAABMA0GCSqGSIb3DQEBCwUA
A0EAUsVE6KMnza/ZbodLlyeMzdo7EM/5nb5ywyOxgIOCf0OOLHsPS9ueGLQX9HEG
//yjTXuhNcUugExIjM/AIwAZPQ==
-----END CERTIFICATE-----`)

// localhostKey is the private key for localhostCert.
var localhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBAL18RtGzM/e3XgavihqFXvMU5sEjhN8DmLtieUxd4uKdvsye9Phu
R84nMTrOS2fn3M/tSIezeXhg66quOLAzjXcCAwEAAQJBAKcRxH9wuglYLBdI/0OT
BLzfWPZCEw1vZmMR2FF1Fm8nkNOVDPleeVGTWoOEcYYlQbpTmkGSxJ6ya+hqRi6x
goECIQDx3+X49fwpL6B5qpJIJMyZBSCuMhH4B7JevhGGFENi3wIhAMiNJN5Q3UkL
IuSvv03kaPR5XVQ99/UeEetUgGvBcABpAiBJSBzVITIVCGkGc7d+RCf49KTCIklv
bGWObufAR8Ni4QIgWpILjW8dkGg8GOUZ0zaNA6Nvt6TIv2UWGJ4v5PoV98kCIQDx
rIiZs5QbKdycsv9gQJzwQAogC8o04X3Zz3dsoX+h4A==
-----END RSA PRIVATE KEY-----`)
