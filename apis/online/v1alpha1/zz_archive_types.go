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

type ArchiveObservation struct {
	ArchiveID *string `json:"archiveId,omitempty" tf:"archive_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ArchiveParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Required
	CollName *string `json:"collName" tf:"coll_name,omitempty"`

	// +kubebuilder:validation:Required
	Criteria []CriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// +kubebuilder:validation:Required
	DBName *string `json:"dbName" tf:"db_name,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionFields []PartitionFieldsParameters `json:"partitionFields,omitempty" tf:"partition_fields,omitempty"`

	// +kubebuilder:validation:Optional
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	SyncCreation *bool `json:"syncCreation,omitempty" tf:"sync_creation,omitempty"`
}

type CriteriaObservation struct {
}

type CriteriaParameters struct {

	// +kubebuilder:validation:Optional
	DateField *string `json:"dateField,omitempty" tf:"date_field,omitempty"`

	// +kubebuilder:validation:Optional
	DateFormat *string `json:"dateFormat,omitempty" tf:"date_format,omitempty"`

	// +kubebuilder:validation:Optional
	ExpireAfterDays *int64 `json:"expireAfterDays,omitempty" tf:"expire_after_days,omitempty"`

	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PartitionFieldsObservation struct {
	FieldType *string `json:"fieldType,omitempty" tf:"field_type,omitempty"`
}

type PartitionFieldsParameters struct {

	// +kubebuilder:validation:Required
	FieldName *string `json:"fieldName" tf:"field_name,omitempty"`

	// +kubebuilder:validation:Required
	Order *int64 `json:"order" tf:"order,omitempty"`
}

// ArchiveSpec defines the desired state of Archive
type ArchiveSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ArchiveParameters `json:"forProvider"`
}

// ArchiveStatus defines the observed state of Archive.
type ArchiveStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ArchiveObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Archive is the Schema for the Archives API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type Archive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArchiveSpec   `json:"spec"`
	Status            ArchiveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveList contains a list of Archives
type ArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Archive `json:"items"`
}

// Repository type metadata.
var (
	Archive_Kind             = "Archive"
	Archive_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Archive_Kind}.String()
	Archive_KindAPIVersion   = Archive_Kind + "." + CRDGroupVersion.String()
	Archive_GroupVersionKind = CRDGroupVersion.WithKind(Archive_Kind)
)

func init() {
	SchemeBuilder.Register(&Archive{}, &ArchiveList{})
}
