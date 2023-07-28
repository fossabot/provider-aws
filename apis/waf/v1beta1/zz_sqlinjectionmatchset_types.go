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

type SQLInjectionMatchSetInitParameters struct {

	// The name or description of the SQL Injection Match Set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SQLInjectionMatchTuples []SQLInjectionMatchTuplesInitParameters `json:"sqlInjectionMatchTuples,omitempty" tf:"sql_injection_match_tuples,omitempty"`
}

type SQLInjectionMatchSetObservation struct {

	// The ID of the WAF SQL Injection Match Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description of the SQL Injection Match Set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SQLInjectionMatchTuples []SQLInjectionMatchTuplesObservation `json:"sqlInjectionMatchTuples,omitempty" tf:"sql_injection_match_tuples,omitempty"`
}

type SQLInjectionMatchSetParameters struct {

	// The name or description of the SQL Injection Match Set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	// +kubebuilder:validation:Optional
	SQLInjectionMatchTuples []SQLInjectionMatchTuplesParameters `json:"sqlInjectionMatchTuples,omitempty" tf:"sql_injection_match_tuples,omitempty"`
}

type SQLInjectionMatchTuplesFieldToMatchInitParameters struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLInjectionMatchTuplesFieldToMatchObservation struct {

	// When type is HEADER, enter the name of the header that you want to search, e.g., User-Agent or Referer.
	// If type is any other value, omit this field.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g., HEADER, METHOD or BODY.
	// See docs
	// for all supported values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLInjectionMatchTuplesFieldToMatchParameters struct {

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

type SQLInjectionMatchTuplesInitParameters struct {

	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch []SQLInjectionMatchTuplesFieldToMatchInitParameters `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on field_to_match before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type SQLInjectionMatchTuplesObservation struct {

	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch []SQLInjectionMatchTuplesFieldToMatchObservation `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on field_to_match before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

type SQLInjectionMatchTuplesParameters struct {

	// Specifies where in a web request to look for snippets of malicious SQL code.
	// +kubebuilder:validation:Optional
	FieldToMatch []SQLInjectionMatchTuplesFieldToMatchParameters `json:"fieldToMatch,omitempty" tf:"field_to_match,omitempty"`

	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on field_to_match before inspecting a request for a match.
	// e.g., CMD_LINE, HTML_ENTITY_DECODE or NONE.
	// See docs
	// for all supported values.
	// +kubebuilder:validation:Optional
	TextTransformation *string `json:"textTransformation,omitempty" tf:"text_transformation,omitempty"`
}

// SQLInjectionMatchSetSpec defines the desired state of SQLInjectionMatchSet
type SQLInjectionMatchSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLInjectionMatchSetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SQLInjectionMatchSetInitParameters `json:"initProvider,omitempty"`
}

// SQLInjectionMatchSetStatus defines the observed state of SQLInjectionMatchSet.
type SQLInjectionMatchSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLInjectionMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLInjectionMatchSet is the Schema for the SQLInjectionMatchSets API. Provides a AWS WAF SQL Injection Match Set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SQLInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   SQLInjectionMatchSetSpec   `json:"spec"`
	Status SQLInjectionMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLInjectionMatchSetList contains a list of SQLInjectionMatchSets
type SQLInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLInjectionMatchSet `json:"items"`
}

// Repository type metadata.
var (
	SQLInjectionMatchSet_Kind             = "SQLInjectionMatchSet"
	SQLInjectionMatchSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLInjectionMatchSet_Kind}.String()
	SQLInjectionMatchSet_KindAPIVersion   = SQLInjectionMatchSet_Kind + "." + CRDGroupVersion.String()
	SQLInjectionMatchSet_GroupVersionKind = CRDGroupVersion.WithKind(SQLInjectionMatchSet_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLInjectionMatchSet{}, &SQLInjectionMatchSetList{})
}
