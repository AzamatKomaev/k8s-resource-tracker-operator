/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ActionType string
type TargetKind string

const (
	ResourceCreated ActionType = "Created"
	ResourceUpdated ActionType = "Updated"
	ResourceDeleted ActionType = "Deleted"

	DeploymentKind  TargetKind = "Deployment"
	StatefulSetKind TargetKind = "StatefulSet"
	ServiceKind     TargetKind = "Service"
)

type ResourceTarget struct {
	Kind      TargetKind `json:"kind"`
	Namespace string     `json:"namespace"`
	Name      string     `json:"name"`
}

type TrackedFieldSpec struct {
	Target       ResourceTarget `json:"target"`
	ContactPoint string         `json:"contactPoint"`
	Field        string         `json:"field"`
}

type TrackedFieldStatus struct {
	Time   *metav1.Time `json:"time"`
	Value  string       `json:"value,omitempty"`
	Action ActionType   `json:"action"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TrackedField is the Schema for the trackedfields API.

type TrackedField struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrackedFieldSpec     `json:"spec"`
	Status []TrackedFieldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrackedFieldList contains a list of TrackedField.
type TrackedFieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrackedField `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrackedField{}, &TrackedFieldList{})
}
