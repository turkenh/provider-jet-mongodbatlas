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

type IndexObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IndexID *string `json:"indexId,omitempty" tf:"index_id,omitempty"`
}

type IndexParameters struct {

	// +kubebuilder:validation:Required
	Analyzer *string `json:"analyzer" tf:"analyzer,omitempty"`

	// +kubebuilder:validation:Optional
	Analyzers *string `json:"analyzers,omitempty" tf:"analyzers,omitempty"`

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Required
	CollectionName *string `json:"collectionName" tf:"collection_name,omitempty"`

	// +kubebuilder:validation:Required
	Database *string `json:"database" tf:"database,omitempty"`

	// +kubebuilder:validation:Optional
	MappingsDynamic *bool `json:"mappingsDynamic,omitempty" tf:"mappings_dynamic,omitempty"`

	// +kubebuilder:validation:Optional
	MappingsFields *string `json:"mappingsFields,omitempty" tf:"mappings_fields,omitempty"`

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" tf:"search_analyzer,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	Synonyms []SynonymsParameters `json:"synonyms,omitempty" tf:"synonyms,omitempty"`
}

type SynonymsObservation struct {
}

type SynonymsParameters struct {

	// +kubebuilder:validation:Required
	Analyzer *string `json:"analyzer" tf:"analyzer,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SourceCollection *string `json:"sourceCollection" tf:"source_collection,omitempty"`
}

// IndexSpec defines the desired state of Index
type IndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IndexParameters `json:"forProvider"`
}

// IndexStatus defines the observed state of Index.
type IndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IndexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Index is the Schema for the Indexs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type Index struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IndexSpec   `json:"spec"`
	Status            IndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IndexList contains a list of Indexs
type IndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Index `json:"items"`
}

// Repository type metadata.
var (
	Index_Kind             = "Index"
	Index_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Index_Kind}.String()
	Index_KindAPIVersion   = Index_Kind + "." + CRDGroupVersion.String()
	Index_GroupVersionKind = CRDGroupVersion.WithKind(Index_Kind)
)

func init() {
	SchemeBuilder.Register(&Index{}, &IndexList{})
}
