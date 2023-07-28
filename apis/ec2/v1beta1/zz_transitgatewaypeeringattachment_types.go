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

type TransitGatewayPeeringAttachmentInitParameters struct {

	// Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the AWS provider is currently connected to.
	PeerAccountID *string `json:"peerAccountId,omitempty" tf:"peer_account_id,omitempty"`

	// Region of EC2 Transit Gateway to peer with.
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TransitGatewayPeeringAttachmentObservation struct {

	// EC2 Transit Gateway Attachment identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the AWS provider is currently connected to.
	PeerAccountID *string `json:"peerAccountId,omitempty" tf:"peer_account_id,omitempty"`

	// Region of EC2 Transit Gateway to peer with.
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// Identifier of EC2 Transit Gateway to peer with.
	PeerTransitGatewayID *string `json:"peerTransitGatewayId,omitempty" tf:"peer_transit_gateway_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`
}

type TransitGatewayPeeringAttachmentParameters struct {

	// Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the AWS provider is currently connected to.
	// +kubebuilder:validation:Optional
	PeerAccountID *string `json:"peerAccountId,omitempty" tf:"peer_account_id,omitempty"`

	// Region of EC2 Transit Gateway to peer with.
	// +kubebuilder:validation:Optional
	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region,omitempty"`

	// Identifier of EC2 Transit Gateway to peer with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PeerTransitGatewayID *string `json:"peerTransitGatewayId,omitempty" tf:"peer_transit_gateway_id,omitempty"`

	// Reference to a TransitGateway in ec2 to populate peerTransitGatewayId.
	// +kubebuilder:validation:Optional
	PeerTransitGatewayIDRef *v1.Reference `json:"peerTransitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway in ec2 to populate peerTransitGatewayId.
	// +kubebuilder:validation:Optional
	PeerTransitGatewayIDSelector *v1.Selector `json:"peerTransitGatewayIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Identifier of EC2 Transit Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`
}

// TransitGatewayPeeringAttachmentSpec defines the desired state of TransitGatewayPeeringAttachment
type TransitGatewayPeeringAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayPeeringAttachmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider TransitGatewayPeeringAttachmentInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayPeeringAttachmentStatus defines the observed state of TransitGatewayPeeringAttachment.
type TransitGatewayPeeringAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayPeeringAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachment is the Schema for the TransitGatewayPeeringAttachments API. Manages an EC2 Transit Gateway Peering Attachment
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayPeeringAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.peerRegion) || has(self.initProvider.peerRegion)",message="peerRegion is a required parameter"
	Spec   TransitGatewayPeeringAttachmentSpec   `json:"spec"`
	Status TransitGatewayPeeringAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayPeeringAttachmentList contains a list of TransitGatewayPeeringAttachments
type TransitGatewayPeeringAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayPeeringAttachment `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayPeeringAttachment_Kind             = "TransitGatewayPeeringAttachment"
	TransitGatewayPeeringAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayPeeringAttachment_Kind}.String()
	TransitGatewayPeeringAttachment_KindAPIVersion   = TransitGatewayPeeringAttachment_Kind + "." + CRDGroupVersion.String()
	TransitGatewayPeeringAttachment_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayPeeringAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayPeeringAttachment{}, &TransitGatewayPeeringAttachmentList{})
}
