/*
Copyright 2023.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	leaderworkersetv1 "sigs.k8s.io/lws/api/leaderworkerset/v1"
)

// NetworkConfigApplyConfiguration represents a declarative configuration of the NetworkConfig type for use
// with apply.
type NetworkConfigApplyConfiguration struct {
	SubdomainPolicy *leaderworkersetv1.SubdomainPolicy `json:"subdomainPolicy,omitempty"`
}

// NetworkConfigApplyConfiguration constructs a declarative configuration of the NetworkConfig type for use with
// apply.
func NetworkConfig() *NetworkConfigApplyConfiguration {
	return &NetworkConfigApplyConfiguration{}
}

// WithSubdomainPolicy sets the SubdomainPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubdomainPolicy field is set to the value of the last call.
func (b *NetworkConfigApplyConfiguration) WithSubdomainPolicy(value leaderworkersetv1.SubdomainPolicy) *NetworkConfigApplyConfiguration {
	b.SubdomainPolicy = &value
	return b
}
