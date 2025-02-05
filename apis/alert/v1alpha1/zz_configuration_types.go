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

type ConfigurationObservation struct {
	AlertConfigurationID *string `json:"alertConfigurationId,omitempty" tf:"alert_configuration_id,omitempty"`

	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type ConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	EventType *string `json:"eventType" tf:"event_type,omitempty"`

	// +kubebuilder:validation:Optional
	Matcher []MatcherParameters `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// +kubebuilder:validation:Optional
	MetricThreshold map[string]*string `json:"metricThreshold,omitempty" tf:"metric_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	MetricThresholdConfig []MetricThresholdConfigParameters `json:"metricThresholdConfig,omitempty" tf:"metric_threshold_config,omitempty"`

	// +kubebuilder:validation:Required
	Notification []NotificationParameters `json:"notification" tf:"notification,omitempty"`

	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// +kubebuilder:validation:Optional
	Threshold map[string]*string `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Optional
	ThresholdConfig []ThresholdConfigParameters `json:"thresholdConfig,omitempty" tf:"threshold_config,omitempty"`
}

type MatcherObservation struct {
}

type MatcherParameters struct {

	// +kubebuilder:validation:Optional
	FieldName *string `json:"fieldName,omitempty" tf:"field_name,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MetricThresholdConfigObservation struct {
}

type MetricThresholdConfigParameters struct {

	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Optional
	Units *string `json:"units,omitempty" tf:"units,omitempty"`
}

type NotificationObservation struct {
	TeamName *string `json:"teamName,omitempty" tf:"team_name,omitempty"`
}

type NotificationParameters struct {

	// +kubebuilder:validation:Optional
	APITokenSecretRef *v1.SecretKeySelector `json:"apiTokenSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ChannelName *string `json:"channelName,omitempty" tf:"channel_name,omitempty"`

	// +kubebuilder:validation:Optional
	DatadogAPIKeySecretRef *v1.SecretKeySelector `json:"datadogApiKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DatadogRegion *string `json:"datadogRegion,omitempty" tf:"datadog_region,omitempty"`

	// +kubebuilder:validation:Optional
	DelayMin *int64 `json:"delayMin,omitempty" tf:"delay_min,omitempty"`

	// +kubebuilder:validation:Optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// +kubebuilder:validation:Optional
	EmailEnabled *bool `json:"emailEnabled,omitempty" tf:"email_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	FlowName *string `json:"flowName,omitempty" tf:"flow_name,omitempty"`

	// +kubebuilder:validation:Optional
	FlowdockAPITokenSecretRef *v1.SecretKeySelector `json:"flowdockApiTokenSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IntervalMin *int64 `json:"intervalMin,omitempty" tf:"interval_min,omitempty"`

	// +kubebuilder:validation:Optional
	MobileNumber *string `json:"mobileNumber,omitempty" tf:"mobile_number,omitempty"`

	// +kubebuilder:validation:Optional
	OpsGenieAPIKeySecretRef *v1.SecretKeySelector `json:"opsGenieApiKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	OpsGenieRegion *string `json:"opsGenieRegion,omitempty" tf:"ops_genie_region,omitempty"`

	// +kubebuilder:validation:Optional
	OrgName *string `json:"orgName,omitempty" tf:"org_name,omitempty"`

	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// +kubebuilder:validation:Optional
	SMSEnabled *bool `json:"smsEnabled,omitempty" tf:"sms_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceKeySecretRef *v1.SecretKeySelector `json:"serviceKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// +kubebuilder:validation:Optional
	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// +kubebuilder:validation:Optional
	VictorOpsAPIKeySecretRef *v1.SecretKeySelector `json:"victorOpsApiKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VictorOpsRoutingKeySecretRef *v1.SecretKeySelector `json:"victorOpsRoutingKeySecretRef,omitempty" tf:"-"`
}

type ThresholdConfigObservation struct {
}

type ThresholdConfigParameters struct {

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Optional
	Units *string `json:"units,omitempty" tf:"units,omitempty"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationParameters `json:"forProvider"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Configuration is the Schema for the Configurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,mongodbatlasjet}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigurationSpec   `json:"spec"`
	Status            ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	Configuration_Kind             = "Configuration"
	Configuration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Configuration_Kind}.String()
	Configuration_KindAPIVersion   = Configuration_Kind + "." + CRDGroupVersion.String()
	Configuration_GroupVersionKind = CRDGroupVersion.WithKind(Configuration_Kind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
