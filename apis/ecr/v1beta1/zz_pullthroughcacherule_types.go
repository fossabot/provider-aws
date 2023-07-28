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

type PullThroughCacheRuleInitParameters struct {

	// The repository name prefix to use when caching images from the source registry.
	EcrRepositoryPrefix *string `json:"ecrRepositoryPrefix,omitempty" tf:"ecr_repository_prefix,omitempty"`

	// The registry URL of the upstream public registry to use as the source.
	UpstreamRegistryURL *string `json:"upstreamRegistryUrl,omitempty" tf:"upstream_registry_url,omitempty"`
}

type PullThroughCacheRuleObservation struct {

	// The repository name prefix to use when caching images from the source registry.
	EcrRepositoryPrefix *string `json:"ecrRepositoryPrefix,omitempty" tf:"ecr_repository_prefix,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The registry ID where the repository was created.
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// The registry URL of the upstream public registry to use as the source.
	UpstreamRegistryURL *string `json:"upstreamRegistryUrl,omitempty" tf:"upstream_registry_url,omitempty"`
}

type PullThroughCacheRuleParameters struct {

	// The repository name prefix to use when caching images from the source registry.
	// +kubebuilder:validation:Optional
	EcrRepositoryPrefix *string `json:"ecrRepositoryPrefix,omitempty" tf:"ecr_repository_prefix,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The registry URL of the upstream public registry to use as the source.
	// +kubebuilder:validation:Optional
	UpstreamRegistryURL *string `json:"upstreamRegistryUrl,omitempty" tf:"upstream_registry_url,omitempty"`
}

// PullThroughCacheRuleSpec defines the desired state of PullThroughCacheRule
type PullThroughCacheRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PullThroughCacheRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider PullThroughCacheRuleInitParameters `json:"initProvider,omitempty"`
}

// PullThroughCacheRuleStatus defines the observed state of PullThroughCacheRule.
type PullThroughCacheRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PullThroughCacheRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PullThroughCacheRule is the Schema for the PullThroughCacheRules API. Provides an Elastic Container Registry Pull Through Cache Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PullThroughCacheRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ecrRepositoryPrefix) || has(self.initProvider.ecrRepositoryPrefix)",message="ecrRepositoryPrefix is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.upstreamRegistryUrl) || has(self.initProvider.upstreamRegistryUrl)",message="upstreamRegistryUrl is a required parameter"
	Spec   PullThroughCacheRuleSpec   `json:"spec"`
	Status PullThroughCacheRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PullThroughCacheRuleList contains a list of PullThroughCacheRules
type PullThroughCacheRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PullThroughCacheRule `json:"items"`
}

// Repository type metadata.
var (
	PullThroughCacheRule_Kind             = "PullThroughCacheRule"
	PullThroughCacheRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PullThroughCacheRule_Kind}.String()
	PullThroughCacheRule_KindAPIVersion   = PullThroughCacheRule_Kind + "." + CRDGroupVersion.String()
	PullThroughCacheRule_GroupVersionKind = CRDGroupVersion.WithKind(PullThroughCacheRule_Kind)
)

func init() {
	SchemeBuilder.Register(&PullThroughCacheRule{}, &PullThroughCacheRuleList{})
}
