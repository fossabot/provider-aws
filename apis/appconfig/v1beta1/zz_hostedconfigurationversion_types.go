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

type HostedConfigurationVersionInitParameters struct {

	// Standard MIME type describing the format of the configuration content. For more information, see Content-Type.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Description of the configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type HostedConfigurationVersionObservation struct {

	// Application ID.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// ARN of the AppConfig  hosted configuration version.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration profile ID.
	ConfigurationProfileID *string `json:"configurationProfileId,omitempty" tf:"configuration_profile_id,omitempty"`

	// Standard MIME type describing the format of the configuration content. For more information, see Content-Type.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Description of the configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// AppConfig application ID, configuration profile ID, and version number separated by a slash (/).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Version number of the hosted configuration.
	VersionNumber *float64 `json:"versionNumber,omitempty" tf:"version_number,omitempty"`
}

type HostedConfigurationVersionParameters struct {

	// Application ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Configuration profile ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.ConfigurationProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("configuration_profile_id",true)
	// +kubebuilder:validation:Optional
	ConfigurationProfileID *string `json:"configurationProfileId,omitempty" tf:"configuration_profile_id,omitempty"`

	// Reference to a ConfigurationProfile in appconfig to populate configurationProfileId.
	// +kubebuilder:validation:Optional
	ConfigurationProfileIDRef *v1.Reference `json:"configurationProfileIdRef,omitempty" tf:"-"`

	// Selector for a ConfigurationProfile in appconfig to populate configurationProfileId.
	// +kubebuilder:validation:Optional
	ConfigurationProfileIDSelector *v1.Selector `json:"configurationProfileIdSelector,omitempty" tf:"-"`

	// Content of the configuration or the configuration data.
	// +kubebuilder:validation:Optional
	ContentSecretRef v1.SecretKeySelector `json:"contentSecretRef" tf:"-"`

	// Standard MIME type describing the format of the configuration content. For more information, see Content-Type.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Description of the configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// HostedConfigurationVersionSpec defines the desired state of HostedConfigurationVersion
type HostedConfigurationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedConfigurationVersionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider HostedConfigurationVersionInitParameters `json:"initProvider,omitempty"`
}

// HostedConfigurationVersionStatus defines the observed state of HostedConfigurationVersion.
type HostedConfigurationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedConfigurationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedConfigurationVersion is the Schema for the HostedConfigurationVersions API. Provides an AppConfig Hosted Configuration Version resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HostedConfigurationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.contentSecretRef)",message="contentSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.contentType) || has(self.initProvider.contentType)",message="contentType is a required parameter"
	Spec   HostedConfigurationVersionSpec   `json:"spec"`
	Status HostedConfigurationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedConfigurationVersionList contains a list of HostedConfigurationVersions
type HostedConfigurationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedConfigurationVersion `json:"items"`
}

// Repository type metadata.
var (
	HostedConfigurationVersion_Kind             = "HostedConfigurationVersion"
	HostedConfigurationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedConfigurationVersion_Kind}.String()
	HostedConfigurationVersion_KindAPIVersion   = HostedConfigurationVersion_Kind + "." + CRDGroupVersion.String()
	HostedConfigurationVersion_GroupVersionKind = CRDGroupVersion.WithKind(HostedConfigurationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedConfigurationVersion{}, &HostedConfigurationVersionList{})
}
