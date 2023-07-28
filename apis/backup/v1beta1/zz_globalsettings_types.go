/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GlobalSettingsInitParameters struct {

	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings map[string]*string `json:"globalSettings,omitempty" tf:"global_settings,omitempty"`
}

type GlobalSettingsObservation struct {

	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings map[string]*string `json:"globalSettings,omitempty" tf:"global_settings,omitempty"`

	// The AWS Account ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GlobalSettingsParameters struct {

	// A list of resources along with the opt-in preferences for the account.
	// +kubebuilder:validation:Optional
	GlobalSettings map[string]*string `json:"globalSettings,omitempty" tf:"global_settings,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// GlobalSettingsSpec defines the desired state of GlobalSettings
type GlobalSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalSettingsParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider GlobalSettingsInitParameters `json:"initProvider,omitempty"`
}

// GlobalSettingsStatus defines the observed state of GlobalSettings.
type GlobalSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalSettings is the Schema for the GlobalSettingss API. Provides an AWS Backup Global Settings resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GlobalSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.globalSettings) || has(self.initProvider.globalSettings)",message="globalSettings is a required parameter"
	Spec   GlobalSettingsSpec   `json:"spec"`
	Status GlobalSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalSettingsList contains a list of GlobalSettingss
type GlobalSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalSettings `json:"items"`
}

// Repository type metadata.
var (
	GlobalSettings_Kind             = "GlobalSettings"
	GlobalSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalSettings_Kind}.String()
	GlobalSettings_KindAPIVersion   = GlobalSettings_Kind + "." + CRDGroupVersion.String()
	GlobalSettings_GroupVersionKind = CRDGroupVersion.WithKind(GlobalSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalSettings{}, &GlobalSettingsList{})
}
