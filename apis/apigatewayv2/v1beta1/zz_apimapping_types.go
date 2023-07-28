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

type APIMappingInitParameters struct {

	// The API mapping key. Refer to REST API, HTTP API or WebSocket API.
	APIMappingKey *string `json:"apiMappingKey,omitempty" tf:"api_mapping_key,omitempty"`
}

type APIMappingObservation struct {

	// API identifier.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// The API mapping key. Refer to REST API, HTTP API or WebSocket API.
	APIMappingKey *string `json:"apiMappingKey,omitempty" tf:"api_mapping_key,omitempty"`

	// Domain name. Use the aws_apigatewayv2_domain_name resource to configure a domain name.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// API mapping identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// API stage. Use the aws_apigatewayv2_stage resource to configure an API stage.
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`
}

type APIMappingParameters struct {

	// API identifier.
	// +crossplane:generate:reference:type=API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a API to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// The API mapping key. Refer to REST API, HTTP API or WebSocket API.
	// +kubebuilder:validation:Optional
	APIMappingKey *string `json:"apiMappingKey,omitempty" tf:"api_mapping_key,omitempty"`

	// Domain name. Use the aws_apigatewayv2_domain_name resource to configure a domain name.
	// +crossplane:generate:reference:type=DomainName
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a DomainName to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a DomainName to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// API stage. Use the aws_apigatewayv2_stage resource to configure an API stage.
	// +crossplane:generate:reference:type=Stage
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// Reference to a Stage to populate stage.
	// +kubebuilder:validation:Optional
	StageRef *v1.Reference `json:"stageRef,omitempty" tf:"-"`

	// Selector for a Stage to populate stage.
	// +kubebuilder:validation:Optional
	StageSelector *v1.Selector `json:"stageSelector,omitempty" tf:"-"`
}

// APIMappingSpec defines the desired state of APIMapping
type APIMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIMappingParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider APIMappingInitParameters `json:"initProvider,omitempty"`
}

// APIMappingStatus defines the observed state of APIMapping.
type APIMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIMapping is the Schema for the APIMappings API. Manages an Amazon API Gateway Version 2 API mapping.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type APIMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIMappingSpec   `json:"spec"`
	Status            APIMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIMappingList contains a list of APIMappings
type APIMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIMapping `json:"items"`
}

// Repository type metadata.
var (
	APIMapping_Kind             = "APIMapping"
	APIMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIMapping_Kind}.String()
	APIMapping_KindAPIVersion   = APIMapping_Kind + "." + CRDGroupVersion.String()
	APIMapping_GroupVersionKind = CRDGroupVersion.WithKind(APIMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&APIMapping{}, &APIMappingList{})
}
