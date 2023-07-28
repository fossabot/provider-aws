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

type BucketServerSideEncryptionConfigurationInitParameters struct {

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rule []BucketServerSideEncryptionConfigurationRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketServerSideEncryptionConfigurationObservation struct {

	// ID (name) of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rule []BucketServerSideEncryptionConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketServerSideEncryptionConfigurationParameters struct {

	// ID (name) of the bucket.
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

	// Account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	// +kubebuilder:validation:Optional
	Rule []BucketServerSideEncryptionConfigurationRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketServerSideEncryptionConfigurationRuleInitParameters struct {

	// Single object for setting server-side encryption by default. See below.
	ApplyServerSideEncryptionByDefault []RuleApplyServerSideEncryptionByDefaultInitParameters `json:"applyServerSideEncryptionByDefault,omitempty" tf:"apply_server_side_encryption_by_default,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`
}

type BucketServerSideEncryptionConfigurationRuleObservation struct {

	// Single object for setting server-side encryption by default. See below.
	ApplyServerSideEncryptionByDefault []RuleApplyServerSideEncryptionByDefaultObservation `json:"applyServerSideEncryptionByDefault,omitempty" tf:"apply_server_side_encryption_by_default,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`
}

type BucketServerSideEncryptionConfigurationRuleParameters struct {

	// Single object for setting server-side encryption by default. See below.
	// +kubebuilder:validation:Optional
	ApplyServerSideEncryptionByDefault []RuleApplyServerSideEncryptionByDefaultParameters `json:"applyServerSideEncryptionByDefault,omitempty" tf:"apply_server_side_encryption_by_default,omitempty"`

	// Whether or not to use Amazon S3 Bucket Keys for SSE-KMS.
	// +kubebuilder:validation:Optional
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`
}

type RuleApplyServerSideEncryptionByDefaultInitParameters struct {

	// Server-side encryption algorithm to use. Valid values are AES256 and aws:kms
	SseAlgorithm *string `json:"sseAlgorithm,omitempty" tf:"sse_algorithm,omitempty"`
}

type RuleApplyServerSideEncryptionByDefaultObservation struct {

	// AWS KMS master key ID used for the SSE-KMS encryption. This can only be used when you set the value of sse_algorithm as aws:kms. The default aws/s3 AWS KMS master key is used if this element is absent while the sse_algorithm is aws:kms.
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// Server-side encryption algorithm to use. Valid values are AES256 and aws:kms
	SseAlgorithm *string `json:"sseAlgorithm,omitempty" tf:"sse_algorithm,omitempty"`
}

type RuleApplyServerSideEncryptionByDefaultParameters struct {

	// AWS KMS master key ID used for the SSE-KMS encryption. This can only be used when you set the value of sse_algorithm as aws:kms. The default aws/s3 AWS KMS master key is used if this element is absent while the sse_algorithm is aws:kms.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsMasterKeyId.
	// +kubebuilder:validation:Optional
	KMSMasterKeyIDRef *v1.Reference `json:"kmsMasterKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsMasterKeyId.
	// +kubebuilder:validation:Optional
	KMSMasterKeyIDSelector *v1.Selector `json:"kmsMasterKeyIdSelector,omitempty" tf:"-"`

	// Server-side encryption algorithm to use. Valid values are AES256 and aws:kms
	// +kubebuilder:validation:Optional
	SseAlgorithm *string `json:"sseAlgorithm,omitempty" tf:"sse_algorithm,omitempty"`
}

// BucketServerSideEncryptionConfigurationSpec defines the desired state of BucketServerSideEncryptionConfiguration
type BucketServerSideEncryptionConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketServerSideEncryptionConfigurationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider BucketServerSideEncryptionConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketServerSideEncryptionConfigurationStatus defines the observed state of BucketServerSideEncryptionConfiguration.
type BucketServerSideEncryptionConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketServerSideEncryptionConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketServerSideEncryptionConfiguration is the Schema for the BucketServerSideEncryptionConfigurations API. Provides a S3 bucket server-side encryption configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketServerSideEncryptionConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || has(self.initProvider.rule)",message="rule is a required parameter"
	Spec   BucketServerSideEncryptionConfigurationSpec   `json:"spec"`
	Status BucketServerSideEncryptionConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketServerSideEncryptionConfigurationList contains a list of BucketServerSideEncryptionConfigurations
type BucketServerSideEncryptionConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketServerSideEncryptionConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketServerSideEncryptionConfiguration_Kind             = "BucketServerSideEncryptionConfiguration"
	BucketServerSideEncryptionConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketServerSideEncryptionConfiguration_Kind}.String()
	BucketServerSideEncryptionConfiguration_KindAPIVersion   = BucketServerSideEncryptionConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketServerSideEncryptionConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketServerSideEncryptionConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketServerSideEncryptionConfiguration{}, &BucketServerSideEncryptionConfigurationList{})
}
