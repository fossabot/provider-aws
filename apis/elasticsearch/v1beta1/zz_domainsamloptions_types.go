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

type DomainSAMLOptionsInitParameters struct {

	// The SAML authentication options for an AWS Elasticsearch Domain.
	SAMLOptions []SAMLOptionsInitParameters `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`
}

type DomainSAMLOptionsObservation struct {

	// The name of the domain the SAML options are associated with.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The SAML authentication options for an AWS Elasticsearch Domain.
	SAMLOptions []SAMLOptionsObservation `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`
}

type DomainSAMLOptionsParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The SAML authentication options for an AWS Elasticsearch Domain.
	// +kubebuilder:validation:Optional
	SAMLOptions []SAMLOptionsParameters `json:"samlOptions,omitempty" tf:"saml_options,omitempty"`
}

type IdpInitParameters struct {

	// The unique Entity ID of the application in SAML Identity Provider.
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The Metadata of the SAML application in xml format.
	MetadataContent *string `json:"metadataContent,omitempty" tf:"metadata_content,omitempty"`
}

type IdpObservation struct {

	// The unique Entity ID of the application in SAML Identity Provider.
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The Metadata of the SAML application in xml format.
	MetadataContent *string `json:"metadataContent,omitempty" tf:"metadata_content,omitempty"`
}

type IdpParameters struct {

	// The unique Entity ID of the application in SAML Identity Provider.
	// +kubebuilder:validation:Optional
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The Metadata of the SAML application in xml format.
	// +kubebuilder:validation:Optional
	MetadataContent *string `json:"metadataContent,omitempty" tf:"metadata_content,omitempty"`
}

type SAMLOptionsInitParameters struct {

	// Whether SAML authentication is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Information from your identity provider.
	Idp []IdpInitParameters `json:"idp,omitempty" tf:"idp,omitempty"`

	// This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role,omitempty"`

	// Element of the SAML assertion to use for backend roles. Default is roles.
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key,omitempty"`

	// Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
	SessionTimeoutMinutes *float64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes,omitempty"`

	// Custom SAML attribute to use for user names. Default is an empty string - "". This will cause Elasticsearch to use the NameID element of the Subject, which is the default location for name identifiers in the SAML specification.
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key,omitempty"`
}

type SAMLOptionsObservation struct {

	// Whether SAML authentication is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Information from your identity provider.
	Idp []IdpObservation `json:"idp,omitempty" tf:"idp,omitempty"`

	// This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role,omitempty"`

	// Element of the SAML assertion to use for backend roles. Default is roles.
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key,omitempty"`

	// Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
	SessionTimeoutMinutes *float64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes,omitempty"`

	// Custom SAML attribute to use for user names. Default is an empty string - "". This will cause Elasticsearch to use the NameID element of the Subject, which is the default location for name identifiers in the SAML specification.
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key,omitempty"`
}

type SAMLOptionsParameters struct {

	// Whether SAML authentication is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Information from your identity provider.
	// +kubebuilder:validation:Optional
	Idp []IdpParameters `json:"idp,omitempty" tf:"idp,omitempty"`

	// This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	// +kubebuilder:validation:Optional
	MasterBackendRole *string `json:"masterBackendRole,omitempty" tf:"master_backend_role,omitempty"`

	// This username from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	// +kubebuilder:validation:Optional
	MasterUserNameSecretRef *v1.SecretKeySelector `json:"masterUserNameSecretRef,omitempty" tf:"-"`

	// Element of the SAML assertion to use for backend roles. Default is roles.
	// +kubebuilder:validation:Optional
	RolesKey *string `json:"rolesKey,omitempty" tf:"roles_key,omitempty"`

	// Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
	// +kubebuilder:validation:Optional
	SessionTimeoutMinutes *float64 `json:"sessionTimeoutMinutes,omitempty" tf:"session_timeout_minutes,omitempty"`

	// Custom SAML attribute to use for user names. Default is an empty string - "". This will cause Elasticsearch to use the NameID element of the Subject, which is the default location for name identifiers in the SAML specification.
	// +kubebuilder:validation:Optional
	SubjectKey *string `json:"subjectKey,omitempty" tf:"subject_key,omitempty"`
}

// DomainSAMLOptionsSpec defines the desired state of DomainSAMLOptions
type DomainSAMLOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainSAMLOptionsParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DomainSAMLOptionsInitParameters `json:"initProvider,omitempty"`
}

// DomainSAMLOptionsStatus defines the observed state of DomainSAMLOptions.
type DomainSAMLOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainSAMLOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainSAMLOptions is the Schema for the DomainSAMLOptionss API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainSAMLOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSAMLOptionsSpec   `json:"spec"`
	Status            DomainSAMLOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainSAMLOptionsList contains a list of DomainSAMLOptionss
type DomainSAMLOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainSAMLOptions `json:"items"`
}

// Repository type metadata.
var (
	DomainSAMLOptions_Kind             = "DomainSAMLOptions"
	DomainSAMLOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainSAMLOptions_Kind}.String()
	DomainSAMLOptions_KindAPIVersion   = DomainSAMLOptions_Kind + "." + CRDGroupVersion.String()
	DomainSAMLOptions_GroupVersionKind = CRDGroupVersion.WithKind(DomainSAMLOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainSAMLOptions{}, &DomainSAMLOptionsList{})
}
