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

type ContactPointType string

const (
	TelegramType ContactPointType = "Telegram"
	WebhookType  ContactPointType = "Webhook"
)

type ContactPointTelegramSpec struct {
	ChatId int `json:"chatId"`
}

type ContactPointWebhookSpec struct {
	Url        string `json:"url"`
	HeaderName string `json:"headerName"`
}

type ContactPointApiToken struct {
	SecretName string `json:"secretName"`
	Key        string `json:"key"`
}

type ContactPointSpec struct {
	Type         ContactPointType         `json:"type"`
	ApiToken     ContactPointApiToken     `json:"apiToken"`
	WebhookSpec  ContactPointWebhookSpec  `json:"webhookSpec,omitempty"`
	TelegramSpec ContactPointTelegramSpec `json:"telegramSpec,omitempty"`
}

type ContactPointStatus struct {
	Initialized bool `json:"initialized"`
	Ready       bool `json:"ready"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ContactPoint is the Schema for the contactpoints API.
type ContactPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContactPointSpec   `json:"spec,omitempty"`
	Status ContactPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContactPointList contains a list of ContactPoint.
type ContactPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContactPoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContactPoint{}, &ContactPointList{})
}
