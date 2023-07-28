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

type BandwidthInitParameters struct {

	// Download speed in Mbps.
	DownloadSpeed *float64 `json:"downloadSpeed,omitempty" tf:"download_speed,omitempty"`

	// Upload speed in Mbps.
	UploadSpeed *float64 `json:"uploadSpeed,omitempty" tf:"upload_speed,omitempty"`
}

type BandwidthObservation struct {

	// Download speed in Mbps.
	DownloadSpeed *float64 `json:"downloadSpeed,omitempty" tf:"download_speed,omitempty"`

	// Upload speed in Mbps.
	UploadSpeed *float64 `json:"uploadSpeed,omitempty" tf:"upload_speed,omitempty"`
}

type BandwidthParameters struct {

	// Download speed in Mbps.
	// +kubebuilder:validation:Optional
	DownloadSpeed *float64 `json:"downloadSpeed,omitempty" tf:"download_speed,omitempty"`

	// Upload speed in Mbps.
	// +kubebuilder:validation:Optional
	UploadSpeed *float64 `json:"uploadSpeed,omitempty" tf:"upload_speed,omitempty"`
}

type LinkInitParameters struct {

	// The upload speed and download speed in Mbps. Documented below.
	Bandwidth []BandwidthInitParameters `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// A description of the link.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The provider of the link.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the link.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LinkObservation struct {

	// Link Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The upload speed and download speed in Mbps. Documented below.
	Bandwidth []BandwidthObservation `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// A description of the link.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the global network.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The provider of the link.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The ID of the site.
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of the link.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LinkParameters struct {

	// The upload speed and download speed in Mbps. Documented below.
	// +kubebuilder:validation:Optional
	Bandwidth []BandwidthParameters `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// A description of the link.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the global network.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// The provider of the link.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the site.
	// +crossplane:generate:reference:type=Site
	// +kubebuilder:validation:Optional
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// Reference to a Site to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDRef *v1.Reference `json:"siteIdRef,omitempty" tf:"-"`

	// Selector for a Site to populate siteId.
	// +kubebuilder:validation:Optional
	SiteIDSelector *v1.Selector `json:"siteIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the link.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// LinkSpec defines the desired state of Link
type LinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider LinkInitParameters `json:"initProvider,omitempty"`
}

// LinkStatus defines the observed state of Link.
type LinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Link is the Schema for the Links API. Creates a link for a site.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Link struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bandwidth) || has(self.initProvider.bandwidth)",message="bandwidth is a required parameter"
	Spec   LinkSpec   `json:"spec"`
	Status LinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkList contains a list of Links
type LinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Link `json:"items"`
}

// Repository type metadata.
var (
	Link_Kind             = "Link"
	Link_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Link_Kind}.String()
	Link_KindAPIVersion   = Link_Kind + "." + CRDGroupVersion.String()
	Link_GroupVersionKind = CRDGroupVersion.WithKind(Link_Kind)
)

func init() {
	SchemeBuilder.Register(&Link{}, &LinkList{})
}
