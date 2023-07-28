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

type DefaultRouteTableInitParameters struct {

	// List of virtual gateways for propagation.
	PropagatingVgws []*string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`

	// Configuration block of routes. Detailed below. This argument is processed in attribute-as-blocks mode. This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Route []RouteInitParameters `json:"route,omitempty" tf:"route,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultRouteTableObservation struct {

	// The ARN of the route table.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ID of the default route table.
	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	// ID of the route table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the AWS account that owns the route table.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// List of virtual gateways for propagation.
	PropagatingVgws []*string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`

	// Configuration block of routes. Detailed below. This argument is processed in attribute-as-blocks mode. This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	Route []RouteObservation `json:"route,omitempty" tf:"route,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// ID of the VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultRouteTableParameters struct {

	// ID of the default route table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("default_route_table_id",true)
	// +kubebuilder:validation:Optional
	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	// Reference to a VPC in ec2 to populate defaultRouteTableId.
	// +kubebuilder:validation:Optional
	DefaultRouteTableIDRef *v1.Reference `json:"defaultRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate defaultRouteTableId.
	// +kubebuilder:validation:Optional
	DefaultRouteTableIDSelector *v1.Selector `json:"defaultRouteTableIdSelector,omitempty" tf:"-"`

	// List of virtual gateways for propagation.
	// +kubebuilder:validation:Optional
	PropagatingVgws []*string `json:"propagatingVgws,omitempty" tf:"propagating_vgws,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block of routes. Detailed below. This argument is processed in attribute-as-blocks mode. This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	// +kubebuilder:validation:Optional
	Route []RouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteInitParameters struct {

	// The CIDR block of the route.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn"`

	// The ID of a managed prefix list destination of the route.
	DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id"`

	// The Ipv6 CIDR block of the route
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	// Identifier of an EC2 instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id"`

	// Identifier of a VPC NAT gateway.
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id"`

	// Identifier of an EC2 network interface.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id"`

	// Identifier of an EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id"`

	// Identifier of a VPC Endpoint. This route must be removed prior to VPC Endpoint deletion.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id"`

	// Identifier of a VPC peering connection.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id"`
}

type RouteObservation struct {

	// The CIDR block of the route.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

	// The ID of a managed prefix list destination of the route.
	DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id,omitempty"`

	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id,omitempty"`

	// Identifier of a VPC internet gateway or a virtual private gateway.
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// The Ipv6 CIDR block of the route
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// Identifier of an EC2 instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Identifier of a VPC NAT gateway.
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id,omitempty"`

	// Identifier of an EC2 network interface.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Identifier of an EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Identifier of a VPC Endpoint. This route must be removed prior to VPC Endpoint deletion.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// Identifier of a VPC peering connection.
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id,omitempty"`
}

type RouteParameters struct {

	// The CIDR block of the route.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	// The Amazon Resource Name (ARN) of a core network.
	// +kubebuilder:validation:Optional
	CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn"`

	// The ID of a managed prefix list destination of the route.
	// +kubebuilder:validation:Optional
	DestinationPrefixListID *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id"`

	// Identifier of a VPC Egress Only Internet Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EgressOnlyInternetGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	EgressOnlyGatewayID *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id"`

	// Reference to a EgressOnlyInternetGateway in ec2 to populate egressOnlyGatewayId.
	// +kubebuilder:validation:Optional
	EgressOnlyGatewayIDRef *v1.Reference `json:"egressOnlyGatewayIdRef,omitempty" tf:"-"`

	// Selector for a EgressOnlyInternetGateway in ec2 to populate egressOnlyGatewayId.
	// +kubebuilder:validation:Optional
	EgressOnlyGatewayIDSelector *v1.Selector `json:"egressOnlyGatewayIdSelector,omitempty" tf:"-"`

	// Identifier of a VPC internet gateway or a virtual private gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.InternetGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id"`

	// Reference to a InternetGateway in ec2 to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDRef *v1.Reference `json:"gatewayIdRef,omitempty" tf:"-"`

	// Selector for a InternetGateway in ec2 to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDSelector *v1.Selector `json:"gatewayIdSelector,omitempty" tf:"-"`

	// The Ipv6 CIDR block of the route
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	// Identifier of an EC2 instance.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id"`

	// Identifier of a VPC NAT gateway.
	// +kubebuilder:validation:Optional
	NATGatewayID *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id"`

	// Identifier of an EC2 network interface.
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id"`

	// Identifier of an EC2 Transit Gateway.
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id"`

	// Identifier of a VPC Endpoint. This route must be removed prior to VPC Endpoint deletion.
	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id"`

	// Identifier of a VPC peering connection.
	// +kubebuilder:validation:Optional
	VPCPeeringConnectionID *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id"`
}

// DefaultRouteTableSpec defines the desired state of DefaultRouteTable
type DefaultRouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultRouteTableParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DefaultRouteTableInitParameters `json:"initProvider,omitempty"`
}

// DefaultRouteTableStatus defines the observed state of DefaultRouteTable.
type DefaultRouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultRouteTable is the Schema for the DefaultRouteTables API. Provides a resource to manage a default route table of a VPC.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DefaultRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultRouteTableSpec   `json:"spec"`
	Status            DefaultRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultRouteTableList contains a list of DefaultRouteTables
type DefaultRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultRouteTable `json:"items"`
}

// Repository type metadata.
var (
	DefaultRouteTable_Kind             = "DefaultRouteTable"
	DefaultRouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultRouteTable_Kind}.String()
	DefaultRouteTable_KindAPIVersion   = DefaultRouteTable_Kind + "." + CRDGroupVersion.String()
	DefaultRouteTable_GroupVersionKind = CRDGroupVersion.WithKind(DefaultRouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultRouteTable{}, &DefaultRouteTableList{})
}
