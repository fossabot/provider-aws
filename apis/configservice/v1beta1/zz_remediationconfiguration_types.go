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

type ExecutionControlsInitParameters struct {

	// Configuration block for SSM controls. See below.
	SsmControls []SsmControlsInitParameters `json:"ssmControls,omitempty" tf:"ssm_controls,omitempty"`
}

type ExecutionControlsObservation struct {

	// Configuration block for SSM controls. See below.
	SsmControls []SsmControlsObservation `json:"ssmControls,omitempty" tf:"ssm_controls,omitempty"`
}

type ExecutionControlsParameters struct {

	// Configuration block for SSM controls. See below.
	// +kubebuilder:validation:Optional
	SsmControls []SsmControlsParameters `json:"ssmControls,omitempty" tf:"ssm_controls,omitempty"`
}

type ParameterInitParameters struct {

	// Name of the attribute.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value is dynamic and changes at run-time.
	ResourceValue *string `json:"resourceValue,omitempty" tf:"resource_value,omitempty"`

	// Value is static and does not change at run-time.
	StaticValue *string `json:"staticValue,omitempty" tf:"static_value,omitempty"`

	// List of static values.
	StaticValues []*string `json:"staticValues,omitempty" tf:"static_values,omitempty"`
}

type ParameterObservation struct {

	// Name of the attribute.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value is dynamic and changes at run-time.
	ResourceValue *string `json:"resourceValue,omitempty" tf:"resource_value,omitempty"`

	// Value is static and does not change at run-time.
	StaticValue *string `json:"staticValue,omitempty" tf:"static_value,omitempty"`

	// List of static values.
	StaticValues []*string `json:"staticValues,omitempty" tf:"static_values,omitempty"`
}

type ParameterParameters struct {

	// Name of the attribute.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Value is dynamic and changes at run-time.
	// +kubebuilder:validation:Optional
	ResourceValue *string `json:"resourceValue,omitempty" tf:"resource_value,omitempty"`

	// Value is static and does not change at run-time.
	// +kubebuilder:validation:Optional
	StaticValue *string `json:"staticValue,omitempty" tf:"static_value,omitempty"`

	// List of static values.
	// +kubebuilder:validation:Optional
	StaticValues []*string `json:"staticValues,omitempty" tf:"static_values,omitempty"`
}

type RemediationConfigurationInitParameters struct {

	// Remediation is triggered automatically if true.
	Automatic *bool `json:"automatic,omitempty" tf:"automatic,omitempty"`

	// Configuration block for execution controls. See below.
	ExecutionControls []ExecutionControlsInitParameters `json:"executionControls,omitempty" tf:"execution_controls,omitempty"`

	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts *float64 `json:"maximumAutomaticAttempts,omitempty" tf:"maximum_automatic_attempts,omitempty"`

	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameter []ParameterInitParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Type of resource.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds *float64 `json:"retryAttemptSeconds,omitempty" tf:"retry_attempt_seconds,omitempty"`

	// Target ID is the name of the public document.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Type of the target. Target executes remediation. For example, SSM document.
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// Version of the target. For example, version of the SSM document
	TargetVersion *string `json:"targetVersion,omitempty" tf:"target_version,omitempty"`
}

type RemediationConfigurationObservation struct {

	// ARN of the Config Remediation Configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Remediation is triggered automatically if true.
	Automatic *bool `json:"automatic,omitempty" tf:"automatic,omitempty"`

	// Configuration block for execution controls. See below.
	ExecutionControls []ExecutionControlsObservation `json:"executionControls,omitempty" tf:"execution_controls,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts *float64 `json:"maximumAutomaticAttempts,omitempty" tf:"maximum_automatic_attempts,omitempty"`

	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Type of resource.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds *float64 `json:"retryAttemptSeconds,omitempty" tf:"retry_attempt_seconds,omitempty"`

	// Target ID is the name of the public document.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Type of the target. Target executes remediation. For example, SSM document.
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// Version of the target. For example, version of the SSM document
	TargetVersion *string `json:"targetVersion,omitempty" tf:"target_version,omitempty"`
}

type RemediationConfigurationParameters struct {

	// Remediation is triggered automatically if true.
	// +kubebuilder:validation:Optional
	Automatic *bool `json:"automatic,omitempty" tf:"automatic,omitempty"`

	// Configuration block for execution controls. See below.
	// +kubebuilder:validation:Optional
	ExecutionControls []ExecutionControlsParameters `json:"executionControls,omitempty" tf:"execution_controls,omitempty"`

	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	// +kubebuilder:validation:Optional
	MaximumAutomaticAttempts *float64 `json:"maximumAutomaticAttempts,omitempty" tf:"maximum_automatic_attempts,omitempty"`

	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Type of resource.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	// +kubebuilder:validation:Optional
	RetryAttemptSeconds *float64 `json:"retryAttemptSeconds,omitempty" tf:"retry_attempt_seconds,omitempty"`

	// Target ID is the name of the public document.
	// +kubebuilder:validation:Optional
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Type of the target. Target executes remediation. For example, SSM document.
	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// Version of the target. For example, version of the SSM document
	// +kubebuilder:validation:Optional
	TargetVersion *string `json:"targetVersion,omitempty" tf:"target_version,omitempty"`
}

type SsmControlsInitParameters struct {

	// Maximum percentage of remediation actions allowed to run in parallel on the non-compliant resources for that specific rule. The default value is 10%.
	ConcurrentExecutionRatePercentage *float64 `json:"concurrentExecutionRatePercentage,omitempty" tf:"concurrent_execution_rate_percentage,omitempty"`

	// Percentage of errors that are allowed before SSM stops running automations on non-compliant resources for that specific rule. The default is 50%.
	ErrorPercentage *float64 `json:"errorPercentage,omitempty" tf:"error_percentage,omitempty"`
}

type SsmControlsObservation struct {

	// Maximum percentage of remediation actions allowed to run in parallel on the non-compliant resources for that specific rule. The default value is 10%.
	ConcurrentExecutionRatePercentage *float64 `json:"concurrentExecutionRatePercentage,omitempty" tf:"concurrent_execution_rate_percentage,omitempty"`

	// Percentage of errors that are allowed before SSM stops running automations on non-compliant resources for that specific rule. The default is 50%.
	ErrorPercentage *float64 `json:"errorPercentage,omitempty" tf:"error_percentage,omitempty"`
}

type SsmControlsParameters struct {

	// Maximum percentage of remediation actions allowed to run in parallel on the non-compliant resources for that specific rule. The default value is 10%.
	// +kubebuilder:validation:Optional
	ConcurrentExecutionRatePercentage *float64 `json:"concurrentExecutionRatePercentage,omitempty" tf:"concurrent_execution_rate_percentage,omitempty"`

	// Percentage of errors that are allowed before SSM stops running automations on non-compliant resources for that specific rule. The default is 50%.
	// +kubebuilder:validation:Optional
	ErrorPercentage *float64 `json:"errorPercentage,omitempty" tf:"error_percentage,omitempty"`
}

// RemediationConfigurationSpec defines the desired state of RemediationConfiguration
type RemediationConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RemediationConfigurationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider RemediationConfigurationInitParameters `json:"initProvider,omitempty"`
}

// RemediationConfigurationStatus defines the observed state of RemediationConfiguration.
type RemediationConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RemediationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RemediationConfiguration is the Schema for the RemediationConfigurations API. Provides an AWS Config Remediation Configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RemediationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetId) || has(self.initProvider.targetId)",message="targetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetType) || has(self.initProvider.targetType)",message="targetType is a required parameter"
	Spec   RemediationConfigurationSpec   `json:"spec"`
	Status RemediationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RemediationConfigurationList contains a list of RemediationConfigurations
type RemediationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RemediationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	RemediationConfiguration_Kind             = "RemediationConfiguration"
	RemediationConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RemediationConfiguration_Kind}.String()
	RemediationConfiguration_KindAPIVersion   = RemediationConfiguration_Kind + "." + CRDGroupVersion.String()
	RemediationConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(RemediationConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&RemediationConfiguration{}, &RemediationConfigurationList{})
}
