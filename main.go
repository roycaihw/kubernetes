package main

import (
	"flag"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	config.ContentConfig.ContentType = "application/vnd.kubernetes.protobuf"
	if err != nil {
		panic(err)
	}

	// create the clientset
	apiExtensionClient, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	_, err = apiExtensionClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bars.stable.example.com",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "stable.example.com",
			Version: "v1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "Bar",
				Plural: "bars",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	if err := apiExtensionClient.ApiextensionsV1beta1().CustomResourceDefinitions().Delete("bars.stable.example.com", &metav1.DeleteOptions{}); err != nil {
		panic(err)
	}
}
