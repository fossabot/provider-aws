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

type DirectoryInitParameters struct {

	// The identifiers of the IP access control groups associated with the directory.
	IPGroupIds []*string `json:"ipGroupIds,omitempty" tf:"ip_group_ids,omitempty"`

	// service capabilities. Defined below.
	SelfServicePermissions []SelfServicePermissionsInitParameters `json:"selfServicePermissions,omitempty" tf:"self_service_permissions,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  Specifies which devices and operating systems users can use to access their WorkSpaces. Defined below.
	WorkspaceAccessProperties []WorkspaceAccessPropertiesInitParameters `json:"workspaceAccessProperties,omitempty" tf:"workspace_access_properties,omitempty"`

	// –  Default properties that are used for creating WorkSpaces. Defined below.
	WorkspaceCreationProperties []WorkspaceCreationPropertiesInitParameters `json:"workspaceCreationProperties,omitempty" tf:"workspace_creation_properties,omitempty"`
}

type DirectoryObservation struct {

	// The directory alias.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// The user name for the service account.
	CustomerUserName *string `json:"customerUserName,omitempty" tf:"customer_user_name,omitempty"`

	// The IP addresses of the DNS servers for the directory.
	DNSIPAddresses []*string `json:"dnsIpAddresses,omitempty" tf:"dns_ip_addresses,omitempty"`

	// The directory identifier for registration in WorkSpaces service.
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// The name of the directory.
	DirectoryName *string `json:"directoryName,omitempty" tf:"directory_name,omitempty"`

	// The directory type.
	DirectoryType *string `json:"directoryType,omitempty" tf:"directory_type,omitempty"`

	// The identifier of the IAM role. This is the role that allows Amazon WorkSpaces to make calls to other services, such as Amazon EC2, on your behalf.
	IAMRoleID *string `json:"iamRoleId,omitempty" tf:"iam_role_id,omitempty"`

	// The WorkSpaces directory identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identifiers of the IP access control groups associated with the directory.
	IPGroupIds []*string `json:"ipGroupIds,omitempty" tf:"ip_group_ids,omitempty"`

	// The registration code for the directory. This is the code that users enter in their Amazon WorkSpaces client application to connect to the directory.
	RegistrationCode *string `json:"registrationCode,omitempty" tf:"registration_code,omitempty"`

	// service capabilities. Defined below.
	SelfServicePermissions []SelfServicePermissionsObservation `json:"selfServicePermissions,omitempty" tf:"self_service_permissions,omitempty"`

	// The identifiers of the subnets where the directory resides.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// –  Specifies which devices and operating systems users can use to access their WorkSpaces. Defined below.
	WorkspaceAccessProperties []WorkspaceAccessPropertiesObservation `json:"workspaceAccessProperties,omitempty" tf:"workspace_access_properties,omitempty"`

	// –  Default properties that are used for creating WorkSpaces. Defined below.
	WorkspaceCreationProperties []WorkspaceCreationPropertiesObservation `json:"workspaceCreationProperties,omitempty" tf:"workspace_creation_properties,omitempty"`

	// The identifier of the security group that is assigned to new WorkSpaces.
	WorkspaceSecurityGroupID *string `json:"workspaceSecurityGroupId,omitempty" tf:"workspace_security_group_id,omitempty"`
}

type DirectoryParameters struct {

	// The directory identifier for registration in WorkSpaces service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// The identifiers of the IP access control groups associated with the directory.
	// +kubebuilder:validation:Optional
	IPGroupIds []*string `json:"ipGroupIds,omitempty" tf:"ip_group_ids,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// service capabilities. Defined below.
	// +kubebuilder:validation:Optional
	SelfServicePermissions []SelfServicePermissionsParameters `json:"selfServicePermissions,omitempty" tf:"self_service_permissions,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The identifiers of the subnets where the directory resides.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  Specifies which devices and operating systems users can use to access their WorkSpaces. Defined below.
	// +kubebuilder:validation:Optional
	WorkspaceAccessProperties []WorkspaceAccessPropertiesParameters `json:"workspaceAccessProperties,omitempty" tf:"workspace_access_properties,omitempty"`

	// –  Default properties that are used for creating WorkSpaces. Defined below.
	// +kubebuilder:validation:Optional
	WorkspaceCreationProperties []WorkspaceCreationPropertiesParameters `json:"workspaceCreationProperties,omitempty" tf:"workspace_creation_properties,omitempty"`
}

type SelfServicePermissionsInitParameters struct {

	// –  Whether WorkSpaces directory users can change the compute type (bundle) for their workspace. Default false.
	ChangeComputeType *bool `json:"changeComputeType,omitempty" tf:"change_compute_type,omitempty"`

	// –  Whether WorkSpaces directory users can increase the volume size of the drives on their workspace. Default false.
	IncreaseVolumeSize *bool `json:"increaseVolumeSize,omitempty" tf:"increase_volume_size,omitempty"`

	// –  Whether WorkSpaces directory users can rebuild the operating system of a workspace to its original state. Default false.
	RebuildWorkspace *bool `json:"rebuildWorkspace,omitempty" tf:"rebuild_workspace,omitempty"`

	// –  Whether WorkSpaces directory users can restart their workspace. Default true.
	RestartWorkspace *bool `json:"restartWorkspace,omitempty" tf:"restart_workspace,omitempty"`

	// –  Whether WorkSpaces directory users can switch the running mode of their workspace. Default false.
	SwitchRunningMode *bool `json:"switchRunningMode,omitempty" tf:"switch_running_mode,omitempty"`
}

type SelfServicePermissionsObservation struct {

	// –  Whether WorkSpaces directory users can change the compute type (bundle) for their workspace. Default false.
	ChangeComputeType *bool `json:"changeComputeType,omitempty" tf:"change_compute_type,omitempty"`

	// –  Whether WorkSpaces directory users can increase the volume size of the drives on their workspace. Default false.
	IncreaseVolumeSize *bool `json:"increaseVolumeSize,omitempty" tf:"increase_volume_size,omitempty"`

	// –  Whether WorkSpaces directory users can rebuild the operating system of a workspace to its original state. Default false.
	RebuildWorkspace *bool `json:"rebuildWorkspace,omitempty" tf:"rebuild_workspace,omitempty"`

	// –  Whether WorkSpaces directory users can restart their workspace. Default true.
	RestartWorkspace *bool `json:"restartWorkspace,omitempty" tf:"restart_workspace,omitempty"`

	// –  Whether WorkSpaces directory users can switch the running mode of their workspace. Default false.
	SwitchRunningMode *bool `json:"switchRunningMode,omitempty" tf:"switch_running_mode,omitempty"`
}

type SelfServicePermissionsParameters struct {

	// –  Whether WorkSpaces directory users can change the compute type (bundle) for their workspace. Default false.
	// +kubebuilder:validation:Optional
	ChangeComputeType *bool `json:"changeComputeType,omitempty" tf:"change_compute_type,omitempty"`

	// –  Whether WorkSpaces directory users can increase the volume size of the drives on their workspace. Default false.
	// +kubebuilder:validation:Optional
	IncreaseVolumeSize *bool `json:"increaseVolumeSize,omitempty" tf:"increase_volume_size,omitempty"`

	// –  Whether WorkSpaces directory users can rebuild the operating system of a workspace to its original state. Default false.
	// +kubebuilder:validation:Optional
	RebuildWorkspace *bool `json:"rebuildWorkspace,omitempty" tf:"rebuild_workspace,omitempty"`

	// –  Whether WorkSpaces directory users can restart their workspace. Default true.
	// +kubebuilder:validation:Optional
	RestartWorkspace *bool `json:"restartWorkspace,omitempty" tf:"restart_workspace,omitempty"`

	// –  Whether WorkSpaces directory users can switch the running mode of their workspace. Default false.
	// +kubebuilder:validation:Optional
	SwitchRunningMode *bool `json:"switchRunningMode,omitempty" tf:"switch_running_mode,omitempty"`
}

type WorkspaceAccessPropertiesInitParameters struct {

	// –  Indicates whether users can use Android devices to access their WorkSpaces.
	DeviceTypeAndroid *string `json:"deviceTypeAndroid,omitempty" tf:"device_type_android,omitempty"`

	// –  Indicates whether users can use Chromebooks to access their WorkSpaces.
	DeviceTypeChromeos *string `json:"deviceTypeChromeos,omitempty" tf:"device_type_chromeos,omitempty"`

	// –  Indicates whether users can use iOS devices to access their WorkSpaces.
	DeviceTypeIos *string `json:"deviceTypeIos,omitempty" tf:"device_type_ios,omitempty"`

	// –  Indicates whether users can use Linux clients to access their WorkSpaces.
	DeviceTypeLinux *string `json:"deviceTypeLinux,omitempty" tf:"device_type_linux,omitempty"`

	// –  Indicates whether users can use macOS clients to access their WorkSpaces.
	DeviceTypeOsx *string `json:"deviceTypeOsx,omitempty" tf:"device_type_osx,omitempty"`

	// –  Indicates whether users can access their WorkSpaces through a web browser.
	DeviceTypeWeb *string `json:"deviceTypeWeb,omitempty" tf:"device_type_web,omitempty"`

	// –  Indicates whether users can use Windows clients to access their WorkSpaces.
	DeviceTypeWindows *string `json:"deviceTypeWindows,omitempty" tf:"device_type_windows,omitempty"`

	// –  Indicates whether users can use zero client devices to access their WorkSpaces.
	DeviceTypeZeroclient *string `json:"deviceTypeZeroclient,omitempty" tf:"device_type_zeroclient,omitempty"`
}

type WorkspaceAccessPropertiesObservation struct {

	// –  Indicates whether users can use Android devices to access their WorkSpaces.
	DeviceTypeAndroid *string `json:"deviceTypeAndroid,omitempty" tf:"device_type_android,omitempty"`

	// –  Indicates whether users can use Chromebooks to access their WorkSpaces.
	DeviceTypeChromeos *string `json:"deviceTypeChromeos,omitempty" tf:"device_type_chromeos,omitempty"`

	// –  Indicates whether users can use iOS devices to access their WorkSpaces.
	DeviceTypeIos *string `json:"deviceTypeIos,omitempty" tf:"device_type_ios,omitempty"`

	// –  Indicates whether users can use Linux clients to access their WorkSpaces.
	DeviceTypeLinux *string `json:"deviceTypeLinux,omitempty" tf:"device_type_linux,omitempty"`

	// –  Indicates whether users can use macOS clients to access their WorkSpaces.
	DeviceTypeOsx *string `json:"deviceTypeOsx,omitempty" tf:"device_type_osx,omitempty"`

	// –  Indicates whether users can access their WorkSpaces through a web browser.
	DeviceTypeWeb *string `json:"deviceTypeWeb,omitempty" tf:"device_type_web,omitempty"`

	// –  Indicates whether users can use Windows clients to access their WorkSpaces.
	DeviceTypeWindows *string `json:"deviceTypeWindows,omitempty" tf:"device_type_windows,omitempty"`

	// –  Indicates whether users can use zero client devices to access their WorkSpaces.
	DeviceTypeZeroclient *string `json:"deviceTypeZeroclient,omitempty" tf:"device_type_zeroclient,omitempty"`
}

type WorkspaceAccessPropertiesParameters struct {

	// –  Indicates whether users can use Android devices to access their WorkSpaces.
	// +kubebuilder:validation:Optional
	DeviceTypeAndroid *string `json:"deviceTypeAndroid,omitempty" tf:"device_type_android,omitempty"`

	// –  Indicates whether users can use Chromebooks to access their WorkSpaces.
	// +kubebuilder:validation:Optional
	DeviceTypeChromeos *string `json:"deviceTypeChromeos,omitempty" tf:"device_type_chromeos,omitempty"`

	// –  Indicates whether users can use iOS devices to access their WorkSpaces.
	// +kubebuilder:validation:Optional
	DeviceTypeIos *string `json:"deviceTypeIos,omitempty" tf:"device_type_ios,omitempty"`

	// –  Indicates whether users can use Linux clients to access their WorkSpaces.
	// +kubebuilder:validation:Optional
	DeviceTypeLinux *string `json:"deviceTypeLinux,omitempty" tf:"device_type_linux,omitempty"`

	// –  Indicates whether users can use macOS clients to access their WorkSpaces.
	// +kubebuilder:validation:Optional
	DeviceTypeOsx *string `json:"deviceTypeOsx,omitempty" tf:"device_type_osx,omitempty"`

	// –  Indicates whether users can access their WorkSpaces through a web browser.
	// +kubebuilder:validation:Optional
	DeviceTypeWeb *string `json:"deviceTypeWeb,omitempty" tf:"device_type_web,omitempty"`

	// –  Indicates whether users can use Windows clients to access their WorkSpaces.
	// +kubebuilder:validation:Optional
	DeviceTypeWindows *string `json:"deviceTypeWindows,omitempty" tf:"device_type_windows,omitempty"`

	// –  Indicates whether users can use zero client devices to access their WorkSpaces.
	// +kubebuilder:validation:Optional
	DeviceTypeZeroclient *string `json:"deviceTypeZeroclient,omitempty" tf:"device_type_zeroclient,omitempty"`
}

type WorkspaceCreationPropertiesInitParameters struct {

	// –  The default organizational unit (OU) for your WorkSpace directories. Should conform "OU=<value>,DC=<value>,...,DC=<value>" pattern.
	DefaultOu *string `json:"defaultOu,omitempty" tf:"default_ou,omitempty"`

	// –  Indicates whether internet access is enabled for your WorkSpaces.
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty" tf:"enable_internet_access,omitempty"`

	// –  Indicates whether maintenance mode is enabled for your WorkSpaces. For more information, see WorkSpace Maintenance..
	EnableMaintenanceMode *bool `json:"enableMaintenanceMode,omitempty" tf:"enable_maintenance_mode,omitempty"`

	// –  Indicates whether users are local administrators of their WorkSpaces.
	UserEnabledAsLocalAdministrator *bool `json:"userEnabledAsLocalAdministrator,omitempty" tf:"user_enabled_as_local_administrator,omitempty"`
}

type WorkspaceCreationPropertiesObservation struct {

	// –  The identifier of your custom security group. Should relate to the same VPC, where workspaces reside in.
	CustomSecurityGroupID *string `json:"customSecurityGroupId,omitempty" tf:"custom_security_group_id,omitempty"`

	// –  The default organizational unit (OU) for your WorkSpace directories. Should conform "OU=<value>,DC=<value>,...,DC=<value>" pattern.
	DefaultOu *string `json:"defaultOu,omitempty" tf:"default_ou,omitempty"`

	// –  Indicates whether internet access is enabled for your WorkSpaces.
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty" tf:"enable_internet_access,omitempty"`

	// –  Indicates whether maintenance mode is enabled for your WorkSpaces. For more information, see WorkSpace Maintenance..
	EnableMaintenanceMode *bool `json:"enableMaintenanceMode,omitempty" tf:"enable_maintenance_mode,omitempty"`

	// –  Indicates whether users are local administrators of their WorkSpaces.
	UserEnabledAsLocalAdministrator *bool `json:"userEnabledAsLocalAdministrator,omitempty" tf:"user_enabled_as_local_administrator,omitempty"`
}

type WorkspaceCreationPropertiesParameters struct {

	// –  The identifier of your custom security group. Should relate to the same VPC, where workspaces reside in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CustomSecurityGroupID *string `json:"customSecurityGroupId,omitempty" tf:"custom_security_group_id,omitempty"`

	// Reference to a SecurityGroup in ec2 to populate customSecurityGroupId.
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIDRef *v1.Reference `json:"customSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in ec2 to populate customSecurityGroupId.
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIDSelector *v1.Selector `json:"customSecurityGroupIdSelector,omitempty" tf:"-"`

	// –  The default organizational unit (OU) for your WorkSpace directories. Should conform "OU=<value>,DC=<value>,...,DC=<value>" pattern.
	// +kubebuilder:validation:Optional
	DefaultOu *string `json:"defaultOu,omitempty" tf:"default_ou,omitempty"`

	// –  Indicates whether internet access is enabled for your WorkSpaces.
	// +kubebuilder:validation:Optional
	EnableInternetAccess *bool `json:"enableInternetAccess,omitempty" tf:"enable_internet_access,omitempty"`

	// –  Indicates whether maintenance mode is enabled for your WorkSpaces. For more information, see WorkSpace Maintenance..
	// +kubebuilder:validation:Optional
	EnableMaintenanceMode *bool `json:"enableMaintenanceMode,omitempty" tf:"enable_maintenance_mode,omitempty"`

	// –  Indicates whether users are local administrators of their WorkSpaces.
	// +kubebuilder:validation:Optional
	UserEnabledAsLocalAdministrator *bool `json:"userEnabledAsLocalAdministrator,omitempty" tf:"user_enabled_as_local_administrator,omitempty"`
}

// DirectorySpec defines the desired state of Directory
type DirectorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DirectoryParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DirectoryInitParameters `json:"initProvider,omitempty"`
}

// DirectoryStatus defines the observed state of Directory.
type DirectoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DirectoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Directory is the Schema for the Directorys API. Provides a WorkSpaces directory in AWS WorkSpaces Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Directory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectorySpec   `json:"spec"`
	Status            DirectoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryList contains a list of Directorys
type DirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Directory `json:"items"`
}

// Repository type metadata.
var (
	Directory_Kind             = "Directory"
	Directory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Directory_Kind}.String()
	Directory_KindAPIVersion   = Directory_Kind + "." + CRDGroupVersion.String()
	Directory_GroupVersionKind = CRDGroupVersion.WithKind(Directory_Kind)
)

func init() {
	SchemeBuilder.Register(&Directory{}, &DirectoryList{})
}
