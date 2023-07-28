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

type XSSMatchSetInitParameters struct {

	// The name or description of the SizeConstraintSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XSSMatchTuples []XSSMatchTuplesInitParameters `json:"xssMatchTuples,omitempty" tf:"xss_match_tuples,omitempty"`
}

type XSSMatchSetObservation struct {

	// Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF XssMatchSet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description of the SizeConstraintSet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XSSMatchTuples []XSSMatchTuplesObservation `json:"xssMatchTuples,omitempty" tf:"xss_match_tuples,omitempty"`
}

type XSSMatchSetParameters struct {

	// The name or description of the SizeConstraintSet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	// +kubebuilder:validation:Optional
	XSSMatchTuples []XSSMatchTuplesParameters `json:"xssMatchTuples,omitempty" tf:"xss_match_tuples,omitempty"`
}

type XSSMatchTuplesFieldToMatchInitParameters struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type XSSMatchTuplesFieldToMatchObservation struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type XSSMatchTuplesFieldToMatchParameters struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type XSSMatchTuplesInitParameters struct {

	// Specifies where in a web request to look for cross-site scripting attacks.
	FieldToMatch []XSSMatchTuplesFieldToMatchInitParameters `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on target_string before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type XSSMatchTuplesObservation struct {

	// Specifies where in a web request to look for cross-site scripting attacks.
	FieldToMatch []XSSMatchTuplesFieldToMatchObservation `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on target_string before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type XSSMatchTuplesParameters struct {

	// Specifies where in a web request to look for cross-site scripting attacks.
	// +kubebuilder:validation:Optional
	FieldToMatch []XSSMatchTuplesFieldToMatchParameters `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on target_string before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Optional
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

// XSSMatchSetSpec defines the desired state of XSSMatchSet
type XSSMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     XSSMatchSetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider XSSMatchSetInitParameters `json:"initProvider,omitempty"`
}

// XSSMatchSetStatus defines the observed state of XSSMatchSet.
type XSSMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        XSSMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// XSSMatchSet is the Schema for the XSSMatchSets API. Provides a AWS WAF XssMatchSet resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type XSSMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   XSSMatchSetSpec   `json:"spec"`
	Status XSSMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XSSMatchSetList contains a list of XSSMatchSets
type XSSMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XSSMatchSet `json:"items"`
}

// Repository type metadata.
var (
	XSSMatchSet_Kind             = "XSSMatchSet"
	XSSMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: XSSMatchSet_Kind}.String()
	XSSMatchSet_KindAPIVersion   = XSSMatchSet_Kind + "." + CRDGroupVersion.String()
	XSSMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(XSSMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&XSSMatchSet{}, &XSSMatchSetList{})
}
