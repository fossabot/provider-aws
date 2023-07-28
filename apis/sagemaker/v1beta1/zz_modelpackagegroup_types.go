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

type ModelPackageGroupInitParameters struct {

	// A description for the model group.
	ModelPackageGroupDescription *string `json:"modelPackageGroupDescription,omitempty" tf:"model_package_group_description,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ModelPackageGroupObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Model Package Group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the Model Package Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A description for the model group.
	ModelPackageGroupDescription *string `json:"modelPackageGroupDescription,omitempty" tf:"model_package_group_description,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ModelPackageGroupParameters struct {

	// A description for the model group.
	// +kubebuilder:validation:Optional
	ModelPackageGroupDescription *string `json:"modelPackageGroupDescription,omitempty" tf:"model_package_group_description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ModelPackageGroupSpec defines the desired state of ModelPackageGroup
type ModelPackageGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModelPackageGroupParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ModelPackageGroupInitParameters `json:"initProvider,omitempty"`
}

// ModelPackageGroupStatus defines the observed state of ModelPackageGroup.
type ModelPackageGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModelPackageGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ModelPackageGroup is the Schema for the ModelPackageGroups API. Provides a SageMaker Model Package Group resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ModelPackageGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelPackageGroupSpec   `json:"spec"`
	Status            ModelPackageGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelPackageGroupList contains a list of ModelPackageGroups
type ModelPackageGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelPackageGroup `json:"items"`
}

// Repository type metadata.
var (
	ModelPackageGroup_Kind             = "ModelPackageGroup"
	ModelPackageGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ModelPackageGroup_Kind}.String()
	ModelPackageGroup_KindAPIVersion   = ModelPackageGroup_Kind + "." + CRDGroupVersion.String()
	ModelPackageGroup_GroupVersionKind = CRDGroupVersion.WithKind(ModelPackageGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ModelPackageGroup{}, &ModelPackageGroupList{})
}
