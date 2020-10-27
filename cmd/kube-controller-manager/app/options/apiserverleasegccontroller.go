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

package options

import (
	"github.com/spf13/pflag"

	apiserverleasegcconfig "k8s.io/kubernetes/pkg/controller/apiserverleasegc/config"
)

// APIServerLeaseGCControllerOptions holds the APIServerLeaseGCController options.
type APIServerLeaseGCControllerOptions struct {
	*apiserverleasegcconfig.APIServerLeaseGCControllerConfiguration
}

// AddFlags adds flags related to APIServerLeaseGCController for controller manager to the specified FlagSet.
func (o *APIServerLeaseGCControllerOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.DurationVar(&o.LeaseResyncPeriod.Duration, "lease-resync-period", o.LeaseResyncPeriod.Duration,
		"Period for lease garbage collector syncing the leases.")
}

// ApplyTo fills up APIServerLeaseGCController config with options.
func (o *APIServerLeaseGCControllerOptions) ApplyTo(cfg *apiserverleasegcconfig.APIServerLeaseGCControllerConfiguration) error {
	if o == nil {
		return nil
	}

	cfg.LeaseResyncPeriod = o.LeaseResyncPeriod

	return nil
}

// Validate checks validation of APIServerLeaseGCController.
func (o *APIServerLeaseGCControllerOptions) Validate() []error {
	if o == nil {
		return nil
	}

	errs := []error{}
	return errs
}
