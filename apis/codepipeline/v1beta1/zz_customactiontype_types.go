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

type ConfigurationPropertyInitParameters struct {

	// The description of the action configuration property.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the configuration property is a key.
	Key *bool `json:"key,omitempty" tf:"key,omitempty"`

	// The name of the action configuration property.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates that the property will be used in conjunction with PollForJobs.
	Queryable *bool `json:"queryable,omitempty" tf:"queryable,omitempty"`

	// Whether the configuration property is a required value.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// Whether the configuration property is secret.
	Secret *bool `json:"secret,omitempty" tf:"secret,omitempty"`

	// The type of the configuration property. Valid values: String, Number, Boolean
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigurationPropertyObservation struct {

	// The description of the action configuration property.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the configuration property is a key.
	Key *bool `json:"key,omitempty" tf:"key,omitempty"`

	// The name of the action configuration property.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates that the property will be used in conjunction with PollForJobs.
	Queryable *bool `json:"queryable,omitempty" tf:"queryable,omitempty"`

	// Whether the configuration property is a required value.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// Whether the configuration property is secret.
	Secret *bool `json:"secret,omitempty" tf:"secret,omitempty"`

	// The type of the configuration property. Valid values: String, Number, Boolean
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConfigurationPropertyParameters struct {

	// The description of the action configuration property.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether the configuration property is a key.
	// +kubebuilder:validation:Optional
	Key *bool `json:"key,omitempty" tf:"key,omitempty"`

	// The name of the action configuration property.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates that the property will be used in conjunction with PollForJobs.
	// +kubebuilder:validation:Optional
	Queryable *bool `json:"queryable,omitempty" tf:"queryable,omitempty"`

	// Whether the configuration property is a required value.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// Whether the configuration property is secret.
	// +kubebuilder:validation:Optional
	Secret *bool `json:"secret,omitempty" tf:"secret,omitempty"`

	// The type of the configuration property. Valid values: String, Number, Boolean
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CustomActionTypeInitParameters struct {

	// The category of the custom action. Valid values: Source, Build, Deploy, Test, Invoke, Approval
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The configuration properties for the custom action. Max 10 items.
	ConfigurationProperty []ConfigurationPropertyInitParameters `json:"configurationProperty,omitempty" tf:"configuration_property,omitempty"`

	// The details of the input artifact for the action.
	InputArtifactDetails []InputArtifactDetailsInitParameters `json:"inputArtifactDetails,omitempty" tf:"input_artifact_details,omitempty"`

	// The details of the output artifact of the action.
	OutputArtifactDetails []OutputArtifactDetailsInitParameters `json:"outputArtifactDetails,omitempty" tf:"output_artifact_details,omitempty"`

	// The provider of the service used in the custom action
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The settings for an action type.
	Settings []SettingsInitParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version identifier of the custom action.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CustomActionTypeObservation struct {

	// The action ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The category of the custom action. Valid values: Source, Build, Deploy, Test, Invoke, Approval
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The configuration properties for the custom action. Max 10 items.
	ConfigurationProperty []ConfigurationPropertyObservation `json:"configurationProperty,omitempty" tf:"configuration_property,omitempty"`

	// Composed of category, provider and version
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The details of the input artifact for the action.
	InputArtifactDetails []InputArtifactDetailsObservation `json:"inputArtifactDetails,omitempty" tf:"input_artifact_details,omitempty"`

	// The details of the output artifact of the action.
	OutputArtifactDetails []OutputArtifactDetailsObservation `json:"outputArtifactDetails,omitempty" tf:"output_artifact_details,omitempty"`

	// The creator of the action being called.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The provider of the service used in the custom action
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The settings for an action type.
	Settings []SettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The version identifier of the custom action.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CustomActionTypeParameters struct {

	// The category of the custom action. Valid values: Source, Build, Deploy, Test, Invoke, Approval
	// +kubebuilder:validation:Optional
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// The configuration properties for the custom action. Max 10 items.
	// +kubebuilder:validation:Optional
	ConfigurationProperty []ConfigurationPropertyParameters `json:"configurationProperty,omitempty" tf:"configuration_property,omitempty"`

	// The details of the input artifact for the action.
	// +kubebuilder:validation:Optional
	InputArtifactDetails []InputArtifactDetailsParameters `json:"inputArtifactDetails,omitempty" tf:"input_artifact_details,omitempty"`

	// The details of the output artifact of the action.
	// +kubebuilder:validation:Optional
	OutputArtifactDetails []OutputArtifactDetailsParameters `json:"outputArtifactDetails,omitempty" tf:"output_artifact_details,omitempty"`

	// The provider of the service used in the custom action
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The settings for an action type.
	// +kubebuilder:validation:Optional
	Settings []SettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The version identifier of the custom action.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InputArtifactDetailsInitParameters struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	MaximumCount *float64 `json:"maximumCount,omitempty" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	MinimumCount *float64 `json:"minimumCount,omitempty" tf:"minimum_count,omitempty"`
}

type InputArtifactDetailsObservation struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	MaximumCount *float64 `json:"maximumCount,omitempty" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	MinimumCount *float64 `json:"minimumCount,omitempty" tf:"minimum_count,omitempty"`
}

type InputArtifactDetailsParameters struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Optional
	MaximumCount *float64 `json:"maximumCount,omitempty" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Optional
	MinimumCount *float64 `json:"minimumCount,omitempty" tf:"minimum_count,omitempty"`
}

type OutputArtifactDetailsInitParameters struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	MaximumCount *float64 `json:"maximumCount,omitempty" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	MinimumCount *float64 `json:"minimumCount,omitempty" tf:"minimum_count,omitempty"`
}

type OutputArtifactDetailsObservation struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	MaximumCount *float64 `json:"maximumCount,omitempty" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	MinimumCount *float64 `json:"minimumCount,omitempty" tf:"minimum_count,omitempty"`
}

type OutputArtifactDetailsParameters struct {

	// The maximum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Optional
	MaximumCount *float64 `json:"maximumCount,omitempty" tf:"maximum_count,omitempty"`

	// The minimum number of artifacts allowed for the action type. Min: 0, Max: 5
	// +kubebuilder:validation:Optional
	MinimumCount *float64 `json:"minimumCount,omitempty" tf:"minimum_count,omitempty"`
}

type SettingsInitParameters struct {

	// The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system.
	EntityURLTemplate *string `json:"entityUrlTemplate,omitempty" tf:"entity_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system.
	ExecutionURLTemplate *string `json:"executionUrlTemplate,omitempty" tf:"execution_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	RevisionURLTemplate *string `json:"revisionUrlTemplate,omitempty" tf:"revision_url_template,omitempty"`

	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	ThirdPartyConfigurationURL *string `json:"thirdPartyConfigurationUrl,omitempty" tf:"third_party_configuration_url,omitempty"`
}

type SettingsObservation struct {

	// The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system.
	EntityURLTemplate *string `json:"entityUrlTemplate,omitempty" tf:"entity_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system.
	ExecutionURLTemplate *string `json:"executionUrlTemplate,omitempty" tf:"execution_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	RevisionURLTemplate *string `json:"revisionUrlTemplate,omitempty" tf:"revision_url_template,omitempty"`

	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	ThirdPartyConfigurationURL *string `json:"thirdPartyConfigurationUrl,omitempty" tf:"third_party_configuration_url,omitempty"`
}

type SettingsParameters struct {

	// The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system.
	// +kubebuilder:validation:Optional
	EntityURLTemplate *string `json:"entityUrlTemplate,omitempty" tf:"entity_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system.
	// +kubebuilder:validation:Optional
	ExecutionURLTemplate *string `json:"executionUrlTemplate,omitempty" tf:"execution_url_template,omitempty"`

	// The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action.
	// +kubebuilder:validation:Optional
	RevisionURLTemplate *string `json:"revisionUrlTemplate,omitempty" tf:"revision_url_template,omitempty"`

	// The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
	// +kubebuilder:validation:Optional
	ThirdPartyConfigurationURL *string `json:"thirdPartyConfigurationUrl,omitempty" tf:"third_party_configuration_url,omitempty"`
}

// CustomActionTypeSpec defines the desired state of CustomActionType
type CustomActionTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomActionTypeParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CustomActionTypeInitParameters `json:"initProvider,omitempty"`
}

// CustomActionTypeStatus defines the observed state of CustomActionType.
type CustomActionTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomActionTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomActionType is the Schema for the CustomActionTypes API. Provides a CodePipeline CustomActionType.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CustomActionType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.category) || has(self.initProvider.category)",message="category is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.inputArtifactDetails) || has(self.initProvider.inputArtifactDetails)",message="inputArtifactDetails is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.outputArtifactDetails) || has(self.initProvider.outputArtifactDetails)",message="outputArtifactDetails is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerName) || has(self.initProvider.providerName)",message="providerName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || has(self.initProvider.version)",message="version is a required parameter"
	Spec   CustomActionTypeSpec   `json:"spec"`
	Status CustomActionTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomActionTypeList contains a list of CustomActionTypes
type CustomActionTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomActionType `json:"items"`
}

// Repository type metadata.
var (
	CustomActionType_Kind             = "CustomActionType"
	CustomActionType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomActionType_Kind}.String()
	CustomActionType_KindAPIVersion   = CustomActionType_Kind + "." + CRDGroupVersion.String()
	CustomActionType_GroupVersionKind = CRDGroupVersion.WithKind(CustomActionType_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomActionType{}, &CustomActionTypeList{})
}
