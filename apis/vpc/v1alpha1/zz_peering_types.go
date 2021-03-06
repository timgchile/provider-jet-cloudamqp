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

type PeeringObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type PeeringParameters struct {

	// Instance identifier
	// +kubebuilder:validation:Required
	InstanceID *float64 `json:"instanceId" tf:"instance_id,omitempty"`

	// VPC peering identifier
	// +kubebuilder:validation:Required
	PeeringID *string `json:"peeringId" tf:"peering_id,omitempty"`
}

// PeeringSpec defines the desired state of Peering
type PeeringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PeeringParameters `json:"forProvider"`
}

// PeeringStatus defines the observed state of Peering.
type PeeringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Peering is the Schema for the Peerings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudamqpjet}
type Peering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PeeringSpec   `json:"spec"`
	Status            PeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringList contains a list of Peerings
type PeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Peering `json:"items"`
}

// Repository type metadata.
var (
	Peering_Kind             = "Peering"
	Peering_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Peering_Kind}.String()
	Peering_KindAPIVersion   = Peering_Kind + "." + CRDGroupVersion.String()
	Peering_GroupVersionKind = CRDGroupVersion.WithKind(Peering_Kind)
)

func init() {
	SchemeBuilder.Register(&Peering{}, &PeeringList{})
}
