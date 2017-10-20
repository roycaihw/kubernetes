package model

import "k8s.io/api/core/v1"

// MinimalPod is a minimal pod.
// +k8s:openapi-gen=true
type MinimalPod struct {
	Name string `json:"name"`

	v1.PodSpec
}

// Car is a simple car model.
// +k8s:openapi-gen=true
type Car struct {
	Color    string
	Capacity int
	// +k8s:openapi-gen=false
	HiddenFeature string
}
