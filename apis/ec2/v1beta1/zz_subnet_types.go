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

type SubnetInitParameters_2 struct {

	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is false
	AssignIPv6AddressOnCreation *bool `json:"assignIpv6AddressOnCreation,omitempty" tf:"assign_ipv6_address_on_creation,omitempty"`

	// AZ for the subnet.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use availability_zone instead.
	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	// The IPv4 CIDR block for the subnet.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The customer owned IPv4 address pool. Typically used with the map_customer_owned_ip_on_launch argument. The outpost_arn argument must be specified when configured.
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: false.
	EnableDns64 *bool `json:"enableDns64,omitempty" tf:"enable_dns64,omitempty"`

	// Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
	EnableLniAtDeviceIndex *float64 `json:"enableLniAtDeviceIndex,omitempty" tf:"enable_lni_at_device_index,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: false.
	EnableResourceNameDNSARecordOnLaunch *bool `json:"enableResourceNameDnsARecordOnLaunch,omitempty" tf:"enable_resource_name_dns_a_record_on_launch,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: false.
	EnableResourceNameDNSAaaaRecordOnLaunch *bool `json:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty" tf:"enable_resource_name_dns_aaaa_record_on_launch,omitempty"`

	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// Indicates whether to create an IPv6-only subnet. Default: false.
	IPv6Native *bool `json:"ipv6Native,omitempty" tf:"ipv6_native,omitempty"`

	// Specify true to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The customer_owned_ipv4_pool and outpost_arn arguments must be specified when set to true. Default is false.
	MapCustomerOwnedIPOnLaunch *bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" tf:"map_customer_owned_ip_on_launch,omitempty"`

	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is false.
	MapPublicIPOnLaunch *bool `json:"mapPublicIpOnLaunch,omitempty" tf:"map_public_ip_on_launch,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: ip-name, resource-name.
	PrivateDNSHostnameTypeOnLaunch *string `json:"privateDnsHostnameTypeOnLaunch,omitempty" tf:"private_dns_hostname_type_on_launch,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SubnetObservation_2 struct {

	// The ARN of the subnet.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is false
	AssignIPv6AddressOnCreation *bool `json:"assignIpv6AddressOnCreation,omitempty" tf:"assign_ipv6_address_on_creation,omitempty"`

	// AZ for the subnet.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use availability_zone instead.
	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	// The IPv4 CIDR block for the subnet.
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The customer owned IPv4 address pool. Typically used with the map_customer_owned_ip_on_launch argument. The outpost_arn argument must be specified when configured.
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: false.
	EnableDns64 *bool `json:"enableDns64,omitempty" tf:"enable_dns64,omitempty"`

	// Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
	EnableLniAtDeviceIndex *float64 `json:"enableLniAtDeviceIndex,omitempty" tf:"enable_lni_at_device_index,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: false.
	EnableResourceNameDNSARecordOnLaunch *bool `json:"enableResourceNameDnsARecordOnLaunch,omitempty" tf:"enable_resource_name_dns_a_record_on_launch,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: false.
	EnableResourceNameDNSAaaaRecordOnLaunch *bool `json:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty" tf:"enable_resource_name_dns_aaaa_record_on_launch,omitempty"`

	// The ID of the subnet
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// The association ID for the IPv6 CIDR block.
	IPv6CidrBlockAssociationID *string `json:"ipv6CidrBlockAssociationId,omitempty" tf:"ipv6_cidr_block_association_id,omitempty"`

	// Indicates whether to create an IPv6-only subnet. Default: false.
	IPv6Native *bool `json:"ipv6Native,omitempty" tf:"ipv6_native,omitempty"`

	// Specify true to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The customer_owned_ipv4_pool and outpost_arn arguments must be specified when set to true. Default is false.
	MapCustomerOwnedIPOnLaunch *bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" tf:"map_customer_owned_ip_on_launch,omitempty"`

	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is false.
	MapPublicIPOnLaunch *bool `json:"mapPublicIpOnLaunch,omitempty" tf:"map_public_ip_on_launch,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// The ID of the AWS account that owns the subnet.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: ip-name, resource-name.
	PrivateDNSHostnameTypeOnLaunch *string `json:"privateDnsHostnameTypeOnLaunch,omitempty" tf:"private_dns_hostname_type_on_launch,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VPC ID.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type SubnetParameters_2 struct {

	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is false
	// +kubebuilder:validation:Optional
	AssignIPv6AddressOnCreation *bool `json:"assignIpv6AddressOnCreation,omitempty" tf:"assign_ipv6_address_on_creation,omitempty"`

	// AZ for the subnet.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// AZ ID of the subnet. This argument is not supported in all regions or partitions. If necessary, use availability_zone instead.
	// +kubebuilder:validation:Optional
	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	// The IPv4 CIDR block for the subnet.
	// +kubebuilder:validation:Optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	// The customer owned IPv4 address pool. Typically used with the map_customer_owned_ip_on_launch argument. The outpost_arn argument must be specified when configured.
	// +kubebuilder:validation:Optional
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool,omitempty"`

	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. Default: false.
	// +kubebuilder:validation:Optional
	EnableDns64 *bool `json:"enableDns64,omitempty" tf:"enable_dns64,omitempty"`

	// Indicates the device position for local network interfaces in this subnet. For example, 1 indicates local network interfaces in this subnet are the secondary network interface (eth1). A local network interface cannot be the primary network interface (eth0).
	// +kubebuilder:validation:Optional
	EnableLniAtDeviceIndex *float64 `json:"enableLniAtDeviceIndex,omitempty" tf:"enable_lni_at_device_index,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records. Default: false.
	// +kubebuilder:validation:Optional
	EnableResourceNameDNSARecordOnLaunch *bool `json:"enableResourceNameDnsARecordOnLaunch,omitempty" tf:"enable_resource_name_dns_a_record_on_launch,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. Default: false.
	// +kubebuilder:validation:Optional
	EnableResourceNameDNSAaaaRecordOnLaunch *bool `json:"enableResourceNameDnsAaaaRecordOnLaunch,omitempty" tf:"enable_resource_name_dns_aaaa_record_on_launch,omitempty"`

	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// Indicates whether to create an IPv6-only subnet. Default: false.
	// +kubebuilder:validation:Optional
	IPv6Native *bool `json:"ipv6Native,omitempty" tf:"ipv6_native,omitempty"`

	// Specify true to indicate that network interfaces created in the subnet should be assigned a customer owned IP address. The customer_owned_ipv4_pool and outpost_arn arguments must be specified when set to true. Default is false.
	// +kubebuilder:validation:Optional
	MapCustomerOwnedIPOnLaunch *bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" tf:"map_customer_owned_ip_on_launch,omitempty"`

	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is false.
	// +kubebuilder:validation:Optional
	MapPublicIPOnLaunch *bool `json:"mapPublicIpOnLaunch,omitempty" tf:"map_public_ip_on_launch,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost.
	// +kubebuilder:validation:Optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// The type of hostnames to assign to instances in the subnet at launch. For IPv6-only subnets, an instance DNS name must be based on the instance ID. For dual-stack and IPv4-only subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: ip-name, resource-name.
	// +kubebuilder:validation:Optional
	PrivateDNSHostnameTypeOnLaunch *string `json:"privateDnsHostnameTypeOnLaunch,omitempty" tf:"private_dns_hostname_type_on_launch,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters_2 `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SubnetInitParameters_2 `json:"initProvider,omitempty"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet is the Schema for the Subnets API. Provides an VPC subnet resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec"`
	Status            SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
