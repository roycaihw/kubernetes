/*
Copyright 2017 The Kubernetes Authors.

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

// Package webhook delegates admission checks to dynamically configured webhooks.
package webhook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"sync"

	"github.com/golang/glog"
	lru "github.com/hashicorp/golang-lru"

	admissionv1alpha1 "k8s.io/api/admission/v1alpha1"
	"k8s.io/api/admissionregistration/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/configuration"
	genericadmissioninit "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
)

const (
	// Name of admission plug-in
	PluginName       = "GenericAdmissionWebhook"
	defaultCacheSize = 200
)

var (
	ErrNeedServiceOrURL = errors.New("webhook configuration must have either service or URL")
)

type ErrCallingWebhook struct {
	WebhookName string
	Reason      error
}

func (e *ErrCallingWebhook) Error() string {
	if e.Reason != nil {
		return fmt.Sprintf("failed calling admission webhook %q: %v", e.WebhookName, e.Reason)
	}
	return fmt.Sprintf("failed calling admission webhook %q; no further details available", e.WebhookName)
}

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register(PluginName, func(configFile io.Reader) (admission.Interface, error) {
		plugin, err := NewGenericAdmissionWebhook(configFile)
		if err != nil {
			return nil, err
		}

		return plugin, nil
	})
}

// WebhookSource can list dynamic webhook plugins.
type WebhookSource interface {
	Run(stopCh <-chan struct{})
	Webhooks() (*v1alpha1.ValidatingWebhookConfiguration, error)
}

// NewGenericAdmissionWebhook returns a generic admission webhook plugin.
func NewGenericAdmissionWebhook(configFile io.Reader) (*GenericAdmissionWebhook, error) {
	kubeconfigFile := ""
	if configFile != nil {
		// TODO: move this to a versioned configuration file format
		var config AdmissionConfig
		d := yaml.NewYAMLOrJSONDecoder(configFile, 4096)
		err := d.Decode(&config)
		if err != nil {
			return nil, err
		}
		kubeconfigFile = config.KubeConfigFile
	}
	authInfoResolver, err := newDefaultAuthenticationInfoResolver(kubeconfigFile)
	if err != nil {
		return nil, err
	}

	cache, err := lru.New(defaultCacheSize)
	if err != nil {
		return nil, err
	}

	return &GenericAdmissionWebhook{
		Handler: admission.NewHandler(
			admission.Connect,
			admission.Create,
			admission.Delete,
			admission.Update,
		),
		authInfoResolver: authInfoResolver,
		serviceResolver:  defaultServiceResolver{},
		cache:            cache,
	}, nil
}

// GenericAdmissionWebhook is an implementation of admission.Interface.
type GenericAdmissionWebhook struct {
	*admission.Handler
	hookSource           WebhookSource
	serviceResolver      ServiceResolver
	negotiatedSerializer runtime.NegotiatedSerializer
	namespaceLister      corelisters.NamespaceLister
	client               clientset.Interface

	authInfoResolver AuthenticationInfoResolver
	cache            *lru.Cache
}

// serviceResolver knows how to convert a service reference into an actual location.
type ServiceResolver interface {
	ResolveEndpoint(namespace, name string) (*url.URL, error)
}

var (
	_ = genericadmissioninit.WantsExternalKubeClientSet(&GenericAdmissionWebhook{})
)

// TODO find a better way wire this, but keep this pull small for now.
func (a *GenericAdmissionWebhook) SetAuthenticationInfoResolverWrapper(wrapper AuthenticationInfoResolverWrapper) {
	if wrapper != nil {
		a.authInfoResolver = wrapper(a.authInfoResolver)
	}
}

// SetServiceResolver sets a service resolver for the webhook admission plugin.
// Passing a nil resolver does not have an effect, instead a default one will be used.
func (a *GenericAdmissionWebhook) SetServiceResolver(sr ServiceResolver) {
	if sr != nil {
		a.serviceResolver = sr
	}
}

// SetScheme sets a serializer(NegotiatedSerializer) which is derived from the scheme
func (a *GenericAdmissionWebhook) SetScheme(scheme *runtime.Scheme) {
	if scheme != nil {
		a.negotiatedSerializer = serializer.NegotiatedSerializerWrapper(runtime.SerializerInfo{
			Serializer: serializer.NewCodecFactory(scheme).LegacyCodec(admissionv1alpha1.SchemeGroupVersion),
		})
	}
}

// WantsExternalKubeClientSet defines a function which sets external ClientSet for admission plugins that need it
func (a *GenericAdmissionWebhook) SetExternalKubeClientSet(client clientset.Interface) {
	a.client = client
	a.hookSource = configuration.NewValidatingWebhookConfigurationManager(client.AdmissionregistrationV1alpha1().ValidatingWebhookConfigurations())
}

// SetExternalKubeInformerFactory implements the WantsExternalKubeInformerFactory interface.
func (a *GenericAdmissionWebhook) SetExternalKubeInformerFactory(f informers.SharedInformerFactory) {
	namespaceInformer := f.Core().V1().Namespaces()
	a.namespaceLister = namespaceInformer.Lister()
	a.SetReadyFunc(namespaceInformer.Informer().HasSynced)
}

// ValidateInitialization implements the InitializationValidator interface.
func (a *GenericAdmissionWebhook) ValidateInitialization() error {
	if a.hookSource == nil {
		return fmt.Errorf("the GenericAdmissionWebhook admission plugin requires a Kubernetes client to be provided")
	}
	if a.negotiatedSerializer == nil {
		return fmt.Errorf("the GenericAdmissionWebhook admission plugin requires a runtime.Scheme to be provided to derive a serializer")
	}
	if a.namespaceLister == nil {
		return fmt.Errorf("the GenericAdmissionWebhook admission plugin requires a namespaceLister")
	}
	go a.hookSource.Run(wait.NeverStop)
	return nil
}

func (a *GenericAdmissionWebhook) loadConfiguration(attr admission.Attributes) (*v1alpha1.ValidatingWebhookConfiguration, error) {
	hookConfig, err := a.hookSource.Webhooks()
	// if Webhook configuration is disabled, fail open
	if err == configuration.ErrDisabled {
		return &v1alpha1.ValidatingWebhookConfiguration{}, nil
	}
	if err != nil {
		e := apierrors.NewServerTimeout(attr.GetResource().GroupResource(), string(attr.GetOperation()), 1)
		e.ErrStatus.Message = fmt.Sprintf("Unable to refresh the Webhook configuration: %v", err)
		e.ErrStatus.Reason = "LoadingConfiguration"
		e.ErrStatus.Details.Causes = append(e.ErrStatus.Details.Causes, metav1.StatusCause{
			Type:    "ValidatingWebhookConfigurationFailure",
			Message: "An error has occurred while refreshing the ValidatingWebhook configuration, no resources can be created/updated/deleted/connected until a refresh succeeds.",
		})
		return nil, e
	}
	return hookConfig, nil
}

// Admit makes an admission decision based on the request attributes.
func (a *GenericAdmissionWebhook) Admit(attr admission.Attributes) error {
	hookConfig, err := a.loadConfiguration(attr)
	if err != nil {
		return err
	}
	hooks := hookConfig.Webhooks
	ctx := context.TODO()

	errCh := make(chan error, len(hooks))
	wg := sync.WaitGroup{}
	wg.Add(len(hooks))
	for i := range hooks {
		go func(hook *v1alpha1.Webhook) {
			defer wg.Done()

			err := a.callHook(ctx, hook, attr)
			if err == nil {
				return
			}

			ignoreClientCallFailures := hook.FailurePolicy != nil && *hook.FailurePolicy == v1alpha1.Ignore
			if callErr, ok := err.(*ErrCallingWebhook); ok {
				if ignoreClientCallFailures {
					glog.Warningf("Failed calling webhook, failing open %v: %v", hook.Name, callErr)
					utilruntime.HandleError(callErr)
					// Since we are failing open to begin with, we do not send an error down the channel
					return
				}

				glog.Warningf("Failed calling webhook, failing closed %v: %v", hook.Name, err)
				errCh <- apierrors.NewInternalError(err)
				return
			}

			glog.Warningf("rejected by webhook %q: %#v", hook.Name, err)
			errCh <- err
		}(&hooks[i])
	}
	wg.Wait()
	close(errCh)

	var errs []error
	for e := range errCh {
		errs = append(errs, e)
	}
	if len(errs) == 0 {
		return nil
	}
	if len(errs) > 1 {
		for i := 1; i < len(errs); i++ {
			// TODO: merge status errors; until then, just return the first one.
			utilruntime.HandleError(errs[i])
		}
	}
	return errs[0]
}

func (a *GenericAdmissionWebhook) getNamespaceLabels(attr admission.Attributes) (map[string]string, error) {
	// If the request itself is creating or updating a namespace, then get the
	// labels from attr.Object, because namespaceLister doesn't have the latest
	// namespace yet.
	//
	// However, if the request is deleting a namespace, then get the label from
	// the namespace in the namespaceLister, because a delete request is not
	// going to change the object, and attr.Object will be a DeleteOptions
	// rather than a namespace object.
	if attr.GetResource().Resource == "namespaces" &&
		len(attr.GetSubresource()) == 0 &&
		(attr.GetOperation() == admission.Create || attr.GetOperation() == admission.Update) {
		accessor, err := meta.Accessor(attr.GetObject())
		if err != nil {
			return nil, err
		}
		return accessor.GetLabels(), nil
	}

	namespaceName := attr.GetNamespace()
	namespace, err := a.namespaceLister.Get(namespaceName)
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, err
	}
	if apierrors.IsNotFound(err) {
		// in case of latency in our caches, make a call direct to storage to verify that it truly exists or not
		namespace, err = a.client.Core().Namespaces().Get(namespaceName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
	}
	return namespace.Labels, nil
}

// whether the request is exempted by the webhook because of the
// namespaceSelector of the webhook.
func (a *GenericAdmissionWebhook) exemptedByNamespaceSelector(h *v1alpha1.Webhook, attr admission.Attributes) (bool, error) {
	namespaceName := attr.GetNamespace()
	if len(namespaceName) == 0 && attr.GetResource().Resource != "namespaces" {
		// If the request is about a cluster scoped resource, and it is not a
		// namespace, it is exempted from all webhooks for now.
		// TODO: figure out a way selective exempt cluster scoped resources.
		// Also update the comment in types.go
		return true, nil
	}
	namespaceLabels, err := a.getNamespaceLabels(attr)
	if apierrors.IsNotFound(err) {
		return false, err
	}
	if err != nil {
		return false, apierrors.NewInternalError(err)
	}
	// TODO: adding an LRU cache to cache the translation
	selector, err := metav1.LabelSelectorAsSelector(h.NamespaceSelector)
	if err != nil {
		return false, apierrors.NewInternalError(err)
	}
	return !selector.Matches(labels.Set(namespaceLabels)), nil
}

func (a *GenericAdmissionWebhook) callHook(ctx context.Context, h *v1alpha1.Webhook, attr admission.Attributes) error {
	excluded, err := a.exemptedByNamespaceSelector(h, attr)
	if err != nil {
		return err
	}
	if excluded {
		return nil
	}
	matches := false
	for _, r := range h.Rules {
		m := RuleMatcher{Rule: r, Attr: attr}
		if m.Matches() {
			matches = true
			break
		}
	}
	if !matches {
		return nil
	}

	// Make the webhook request
	request := createAdmissionReview(attr)
	client, err := a.hookClient(h)
	if err != nil {
		return &ErrCallingWebhook{WebhookName: h.Name, Reason: err}
	}
	response := &admissionv1alpha1.AdmissionReview{}
	if err := client.Post().Context(ctx).Body(&request).Do().Into(response); err != nil {
		return &ErrCallingWebhook{WebhookName: h.Name, Reason: err}
	}

	if response.Status.Allowed {
		return nil
	}

	return toStatusErr(h.Name, response.Status.Result)
}

// toStatusErr returns a StatusError with information about the webhook controller
func toStatusErr(name string, result *metav1.Status) *apierrors.StatusError {
	deniedBy := fmt.Sprintf("admission webhook %q denied the request", name)
	const noExp = "without explanation"

	if result == nil {
		result = &metav1.Status{Status: metav1.StatusFailure}
	}

	switch {
	case len(result.Message) > 0:
		result.Message = fmt.Sprintf("%s: %s", deniedBy, result.Message)
	case len(result.Reason) > 0:
		result.Message = fmt.Sprintf("%s: %s", deniedBy, result.Reason)
	default:
		result.Message = fmt.Sprintf("%s %s", deniedBy, noExp)
	}

	return &apierrors.StatusError{
		ErrStatus: *result,
	}
}

func (a *GenericAdmissionWebhook) hookClient(h *v1alpha1.Webhook) (*rest.RESTClient, error) {
	cacheKey, err := json.Marshal(h.ClientConfig)
	if err != nil {
		return nil, err
	}
	if client, ok := a.cache.Get(string(cacheKey)); ok {
		return client.(*rest.RESTClient), nil
	}

	complete := func(cfg *rest.Config) (*rest.RESTClient, error) {
		cfg.TLSClientConfig.CAData = h.ClientConfig.CABundle
		cfg.ContentConfig.NegotiatedSerializer = a.negotiatedSerializer
		cfg.ContentConfig.ContentType = runtime.ContentTypeJSON
		client, err := rest.UnversionedRESTClientFor(cfg)
		if err == nil {
			a.cache.Add(string(cacheKey), client)
		}
		return client, err
	}

	if svc := h.ClientConfig.Service; svc != nil {
		serverName := svc.Name + "." + svc.Namespace + ".svc"
		restConfig, err := a.authInfoResolver.ClientConfigFor(serverName)
		if err != nil {
			return nil, err
		}
		cfg := rest.CopyConfig(restConfig)
		host := serverName + ":443"
		cfg.Host = "https://" + host
		if svc.Path != nil {
			cfg.APIPath = *svc.Path
		}
		cfg.TLSClientConfig.ServerName = serverName

		delegateDialer := cfg.Dial
		if delegateDialer == nil {
			delegateDialer = net.Dial
		}
		cfg.Dial = func(network, addr string) (net.Conn, error) {
			if addr == host {
				u, err := a.serviceResolver.ResolveEndpoint(svc.Namespace, svc.Name)
				if err != nil {
					return nil, err
				}
				addr = u.Host
			}
			return delegateDialer(network, addr)
		}

		return complete(cfg)
	}

	if h.ClientConfig.URL == nil {
		return nil, &ErrCallingWebhook{WebhookName: h.Name, Reason: ErrNeedServiceOrURL}
	}

	u, err := url.Parse(*h.ClientConfig.URL)
	if err != nil {
		return nil, &ErrCallingWebhook{WebhookName: h.Name, Reason: fmt.Errorf("Unparsable URL: %v", err)}
	}

	restConfig, err := a.authInfoResolver.ClientConfigFor(u.Host)
	if err != nil {
		return nil, err
	}

	cfg := rest.CopyConfig(restConfig)
	cfg.Host = u.Host
	cfg.APIPath = u.Path

	return complete(cfg)
}
