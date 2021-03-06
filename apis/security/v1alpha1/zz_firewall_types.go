/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FirewallParameters struct {

	// Instance identifier
	// +kubebuilder:validation:Required
	InstanceID *float64 `json:"instanceId" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// Naming descripton e.g. 'Default'
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IP address together with netmask to allow acces
	// +kubebuilder:validation:Required
	IP *string `json:"ip" tf:"ip,omitempty"`

	// Custom ports between 0 - 65554
	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	// Pre-defined services 'AMQP', 'AMQPS', 'HTTPS', 'MQTT', 'MQTTS', 'STOMP', 'STOMPS', 'STREAM', 'STREAM_SSL'
	// +kubebuilder:validation:Optional
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudamqpjet}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec"`
	Status            FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
