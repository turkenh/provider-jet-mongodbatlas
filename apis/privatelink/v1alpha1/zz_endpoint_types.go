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

type EndpointObservation struct {
	EndpointGroupNames []*string `json:"endpointGroupNames,omitempty" tf:"endpoint_group_names,omitempty"`

	EndpointServiceName *string `json:"endpointServiceName,omitempty" tf:"endpoint_service_name,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InterfaceEndpoints []*string `json:"interfaceEndpoints,omitempty" tf:"interface_endpoints,omitempty"`

	PrivateEndpoints []*string `json:"privateEndpoints,omitempty" tf:"private_endpoints,omitempty"`

	PrivateLinkID *string `json:"privateLinkId,omitempty" tf:"private_link_id,omitempty"`

	PrivateLinkServiceName *string `json:"privateLinkServiceName,omitempty" tf:"private_link_service_name,omitempty"`

	PrivateLinkServiceResourceID *string `json:"privateLinkServiceResourceId,omitempty" tf:"private_link_service_resource_id,omitempty"`

	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	ServiceAttachmentNames []*string `json:"serviceAttachmentNames,omitempty" tf:"service_attachment_names,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EndpointParameters struct {

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Required
	ProviderName *string `json:"providerName" tf:"provider_name,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointParameters `json:"forProvider"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Endpoint is the Schema for the Endpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec"`
	Status            EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

// Repository type metadata.
var (
	Endpoint_Kind             = "Endpoint"
	Endpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Endpoint_Kind}.String()
	Endpoint_KindAPIVersion   = Endpoint_Kind + "." + CRDGroupVersion.String()
	Endpoint_GroupVersionKind = CRDGroupVersion.WithKind(Endpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
