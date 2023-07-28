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

type BucketAnalyticsConfigurationFilterInitParameters struct {

	// Object prefix for filtering.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketAnalyticsConfigurationFilterObservation struct {

	// Object prefix for filtering.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketAnalyticsConfigurationFilterParameters struct {

	// Object prefix for filtering.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BucketAnalyticsConfigurationInitParameters struct {

	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter []BucketAnalyticsConfigurationFilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier of the analytics configuration for the bucket.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis []StorageClassAnalysisInitParameters `json:"storageClassAnalysis,omitempty" tf:"storage_class_analysis,omitempty"`
}

type BucketAnalyticsConfigurationObservation struct {

	// Name of the bucket this analytics configuration is associated with.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter []BucketAnalyticsConfigurationFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique identifier of the analytics configuration for the bucket.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis []StorageClassAnalysisObservation `json:"storageClassAnalysis,omitempty" tf:"storage_class_analysis,omitempty"`
}

type BucketAnalyticsConfigurationParameters struct {

	// Name of the bucket this analytics configuration is associated with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	// +kubebuilder:validation:Optional
	Filter []BucketAnalyticsConfigurationFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier of the analytics configuration for the bucket.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration for the analytics data export (documented below).
	// +kubebuilder:validation:Optional
	StorageClassAnalysis []StorageClassAnalysisParameters `json:"storageClassAnalysis,omitempty" tf:"storage_class_analysis,omitempty"`
}

type DataExportDestinationInitParameters struct {

	// Analytics data export currently only supports an S3 bucket destination (documented below).
	S3BucketDestination []S3BucketDestinationInitParameters `json:"s3BucketDestination,omitempty" tf:"s3_bucket_destination,omitempty"`
}

type DataExportDestinationObservation struct {

	// Analytics data export currently only supports an S3 bucket destination (documented below).
	S3BucketDestination []S3BucketDestinationObservation `json:"s3BucketDestination,omitempty" tf:"s3_bucket_destination,omitempty"`
}

type DataExportDestinationParameters struct {

	// Analytics data export currently only supports an S3 bucket destination (documented below).
	// +kubebuilder:validation:Optional
	S3BucketDestination []S3BucketDestinationParameters `json:"s3BucketDestination,omitempty" tf:"s3_bucket_destination,omitempty"`
}

type DataExportInitParameters struct {

	// Specifies the destination for the exported analytics data (documented below).
	Destination []DataExportDestinationInitParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Schema version of exported analytics data. Allowed values: V_1. Default value: V_1.
	OutputSchemaVersion *string `json:"outputSchemaVersion,omitempty" tf:"output_schema_version,omitempty"`
}

type DataExportObservation struct {

	// Specifies the destination for the exported analytics data (documented below).
	Destination []DataExportDestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// Schema version of exported analytics data. Allowed values: V_1. Default value: V_1.
	OutputSchemaVersion *string `json:"outputSchemaVersion,omitempty" tf:"output_schema_version,omitempty"`
}

type DataExportParameters struct {

	// Specifies the destination for the exported analytics data (documented below).
	// +kubebuilder:validation:Optional
	Destination []DataExportDestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Schema version of exported analytics data. Allowed values: V_1. Default value: V_1.
	// +kubebuilder:validation:Optional
	OutputSchemaVersion *string `json:"outputSchemaVersion,omitempty" tf:"output_schema_version,omitempty"`
}

type S3BucketDestinationInitParameters struct {

	// Account ID that owns the destination bucket.
	BucketAccountID *string `json:"bucketAccountId,omitempty" tf:"bucket_account_id,omitempty"`

	// Output format of exported analytics data. Allowed values: CSV. Default value: CSV.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Object prefix for filtering.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3BucketDestinationObservation struct {

	// Account ID that owns the destination bucket.
	BucketAccountID *string `json:"bucketAccountId,omitempty" tf:"bucket_account_id,omitempty"`

	// ARN of the destination bucket.
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Output format of exported analytics data. Allowed values: CSV. Default value: CSV.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Object prefix for filtering.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3BucketDestinationParameters struct {

	// Account ID that owns the destination bucket.
	// +kubebuilder:validation:Optional
	BucketAccountID *string `json:"bucketAccountId,omitempty" tf:"bucket_account_id,omitempty"`

	// ARN of the destination bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// Output format of exported analytics data. Allowed values: CSV. Default value: CSV.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Object prefix for filtering.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type StorageClassAnalysisInitParameters struct {

	// Data export configuration (documented below).
	DataExport []DataExportInitParameters `json:"dataExport,omitempty" tf:"data_export,omitempty"`
}

type StorageClassAnalysisObservation struct {

	// Data export configuration (documented below).
	DataExport []DataExportObservation `json:"dataExport,omitempty" tf:"data_export,omitempty"`
}

type StorageClassAnalysisParameters struct {

	// Data export configuration (documented below).
	// +kubebuilder:validation:Optional
	DataExport []DataExportParameters `json:"dataExport,omitempty" tf:"data_export,omitempty"`
}

// BucketAnalyticsConfigurationSpec defines the desired state of BucketAnalyticsConfiguration
type BucketAnalyticsConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketAnalyticsConfigurationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider BucketAnalyticsConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketAnalyticsConfigurationStatus defines the observed state of BucketAnalyticsConfiguration.
type BucketAnalyticsConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketAnalyticsConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAnalyticsConfiguration is the Schema for the BucketAnalyticsConfigurations API. Provides a S3 bucket analytics configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketAnalyticsConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   BucketAnalyticsConfigurationSpec   `json:"spec"`
	Status BucketAnalyticsConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAnalyticsConfigurationList contains a list of BucketAnalyticsConfigurations
type BucketAnalyticsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAnalyticsConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketAnalyticsConfiguration_Kind             = "BucketAnalyticsConfiguration"
	BucketAnalyticsConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketAnalyticsConfiguration_Kind}.String()
	BucketAnalyticsConfiguration_KindAPIVersion   = BucketAnalyticsConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketAnalyticsConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketAnalyticsConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketAnalyticsConfiguration{}, &BucketAnalyticsConfigurationList{})
}
