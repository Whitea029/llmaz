/*
Copyright 2025 The InftyAI Team.

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

package v1alpha1

import (
	v2 "k8s.io/api/autoscaling/v2"
)

// HPATriggerApplyConfiguration represents a declarative configuration of the HPATrigger type for use
// with apply.
type HPATriggerApplyConfiguration struct {
	Metrics  []v2.MetricSpec                     `json:"metrics,omitempty"`
	Behavior *v2.HorizontalPodAutoscalerBehavior `json:"behavior,omitempty"`
}

// HPATriggerApplyConfiguration constructs a declarative configuration of the HPATrigger type for use with
// apply.
func HPATrigger() *HPATriggerApplyConfiguration {
	return &HPATriggerApplyConfiguration{}
}

// WithMetrics adds the given value to the Metrics field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Metrics field.
func (b *HPATriggerApplyConfiguration) WithMetrics(values ...v2.MetricSpec) *HPATriggerApplyConfiguration {
	for i := range values {
		b.Metrics = append(b.Metrics, values[i])
	}
	return b
}

// WithBehavior sets the Behavior field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Behavior field is set to the value of the last call.
func (b *HPATriggerApplyConfiguration) WithBehavior(value v2.HorizontalPodAutoscalerBehavior) *HPATriggerApplyConfiguration {
	b.Behavior = &value
	return b
}
