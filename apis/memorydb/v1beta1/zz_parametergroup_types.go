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

type ParameterGroupInitParameters struct {

	// Description for the parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The engine version that the parameter group can be used with.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParameterGroupObservation struct {

	// The ARN of the parameter group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description for the parameter group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The engine version that the parameter group can be used with.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Same as name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ParameterGroupParameters struct {

	// Description for the parameter group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The engine version that the parameter group can be used with.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Set of MemoryDB parameters to apply. Any parameters not specified will fall back to their family defaults. Detailed below.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParameterInitParameters struct {

	// The name of the parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterObservation struct {

	// The name of the parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// The name of the parameter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the parameter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// ParameterGroupSpec defines the desired state of ParameterGroup
type ParameterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ParameterGroupParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ParameterGroupInitParameters `json:"initProvider,omitempty"`
}

// ParameterGroupStatus defines the observed state of ParameterGroup.
type ParameterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroup is the Schema for the ParameterGroups API. Provides a MemoryDB Parameter Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.family) || has(self.initProvider.family)",message="family is a required parameter"
	Spec   ParameterGroupSpec   `json:"spec"`
	Status ParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroupList contains a list of ParameterGroups
type ParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	ParameterGroup_Kind             = "ParameterGroup"
	ParameterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ParameterGroup_Kind}.String()
	ParameterGroup_KindAPIVersion   = ParameterGroup_Kind + "." + CRDGroupVersion.String()
	ParameterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ParameterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ParameterGroup{}, &ParameterGroupList{})
}
