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

type AssociationInitParameters struct {
}

type AssociationObservation struct {

	// The license configuration ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the license configuration.
	LicenseConfigurationArn *string `json:"licenseConfigurationArn,omitempty" tf:"license_configuration_arn,omitempty"`

	// ARN of the resource associated with the license configuration.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type AssociationParameters struct {

	// ARN of the license configuration.
	// +crossplane:generate:reference:type=LicenseConfiguration
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	LicenseConfigurationArn *string `json:"licenseConfigurationArn,omitempty" tf:"license_configuration_arn,omitempty"`

	// Reference to a LicenseConfiguration to populate licenseConfigurationArn.
	// +kubebuilder:validation:Optional
	LicenseConfigurationArnRef *v1.Reference `json:"licenseConfigurationArnRef,omitempty" tf:"-"`

	// Selector for a LicenseConfiguration to populate licenseConfigurationArn.
	// +kubebuilder:validation:Optional
	LicenseConfigurationArnSelector *v1.Selector `json:"licenseConfigurationArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of the resource associated with the license configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Instance in ec2 to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Instance in ec2 to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

// AssociationSpec defines the desired state of Association
type AssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AssociationInitParameters `json:"initProvider,omitempty"`
}

// AssociationStatus defines the observed state of Association.
type AssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Association is the Schema for the Associations API. Provides a License Manager association resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Association struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AssociationSpec   `json:"spec"`
	Status            AssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssociationList contains a list of Associations
type AssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Association `json:"items"`
}

// Repository type metadata.
var (
	Association_Kind             = "Association"
	Association_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Association_Kind}.String()
	Association_KindAPIVersion   = Association_Kind + "." + CRDGroupVersion.String()
	Association_GroupVersionKind = CRDGroupVersion.WithKind(Association_Kind)
)

func init() {
	SchemeBuilder.Register(&Association{}, &AssociationList{})
}
