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

type BucketReplicationConfigurationInitParameters struct {

	// List of configuration blocks describing the rules managing the replication. See below.
	Rule []BucketReplicationConfigurationRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketReplicationConfigurationObservation struct {

	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// S3 source bucket name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// List of configuration blocks describing the rules managing the replication. See below.
	Rule []BucketReplicationConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketReplicationConfigurationParameters struct {

	// Name of the source S3 bucket you want Amazon S3 to monitor.
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

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of the IAM role for Amazon S3 to assume when replicating the objects.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// List of configuration blocks describing the rules managing the replication. See below.
	// +kubebuilder:validation:Optional
	Rule []BucketReplicationConfigurationRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Token to allow replication to be enabled on an Object Lock-enabled bucket. You must contact AWS support for the bucket's "Object Lock token".
	// For more details, see Using S3 Object Lock with replication.
	// +kubebuilder:validation:Optional
	TokenSecretRef *v1.SecretKeySelector `json:"tokenSecretRef,omitempty" tf:"-"`
}

type BucketReplicationConfigurationRuleFilterInitParameters struct {

	// Configuration block for specifying rule filters. This element is required only if you specify more than one filter. See and below for more details.
	And []FilterAndInitParameters `json:"and,omitempty" tf:"and,omitempty"`

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Configuration block for specifying a tag key and value. See below.
	Tag []FilterTagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BucketReplicationConfigurationRuleFilterObservation struct {

	// Configuration block for specifying rule filters. This element is required only if you specify more than one filter. See and below for more details.
	And []FilterAndObservation `json:"and,omitempty" tf:"and,omitempty"`

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Configuration block for specifying a tag key and value. See below.
	Tag []FilterTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BucketReplicationConfigurationRuleFilterParameters struct {

	// Configuration block for specifying rule filters. This element is required only if you specify more than one filter. See and below for more details.
	// +kubebuilder:validation:Optional
	And []FilterAndParameters `json:"and,omitempty" tf:"and,omitempty"`

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Configuration block for specifying a tag key and value. See below.
	// +kubebuilder:validation:Optional
	Tag []FilterTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BucketReplicationConfigurationRuleInitParameters struct {

	// Whether delete markers are replicated. This argument is only valid with V2 replication configurations (i.e., when filter is used)documented below.
	DeleteMarkerReplication []DeleteMarkerReplicationInitParameters `json:"deleteMarkerReplication,omitempty" tf:"delete_marker_replication,omitempty"`

	// Specifies the destination for the rule. See below.
	Destination []RuleDestinationInitParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Replicate existing objects in the source bucket according to the rule configurations. See below.
	ExistingObjectReplication []ExistingObjectReplicationInitParameters `json:"existingObjectReplication,omitempty" tf:"existing_object_replication,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies. See below. If not specified, the rule will default to using prefix.
	Filter []BucketReplicationConfigurationRuleFilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Priority associated with the rule. Priority should only be set if filter is configured. If not provided, defaults to 0. Priority must be unique between multiple rules.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies special object selection criteria. See below.
	SourceSelectionCriteria []RuleSourceSelectionCriteriaInitParameters `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`

	// Status of the rule. Either "Enabled" or "Disabled". The rule is ignored if status is not "Enabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BucketReplicationConfigurationRuleObservation struct {

	// Whether delete markers are replicated. This argument is only valid with V2 replication configurations (i.e., when filter is used)documented below.
	DeleteMarkerReplication []DeleteMarkerReplicationObservation `json:"deleteMarkerReplication,omitempty" tf:"delete_marker_replication,omitempty"`

	// Specifies the destination for the rule. See below.
	Destination []RuleDestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// Replicate existing objects in the source bucket according to the rule configurations. See below.
	ExistingObjectReplication []ExistingObjectReplicationObservation `json:"existingObjectReplication,omitempty" tf:"existing_object_replication,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies. See below. If not specified, the rule will default to using prefix.
	Filter []BucketReplicationConfigurationRuleFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Priority associated with the rule. Priority should only be set if filter is configured. If not provided, defaults to 0. Priority must be unique between multiple rules.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies special object selection criteria. See below.
	SourceSelectionCriteria []RuleSourceSelectionCriteriaObservation `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`

	// Status of the rule. Either "Enabled" or "Disabled". The rule is ignored if status is not "Enabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BucketReplicationConfigurationRuleParameters struct {

	// Whether delete markers are replicated. This argument is only valid with V2 replication configurations (i.e., when filter is used)documented below.
	// +kubebuilder:validation:Optional
	DeleteMarkerReplication []DeleteMarkerReplicationParameters `json:"deleteMarkerReplication,omitempty" tf:"delete_marker_replication,omitempty"`

	// Specifies the destination for the rule. See below.
	// +kubebuilder:validation:Optional
	Destination []RuleDestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Replicate existing objects in the source bucket according to the rule configurations. See below.
	// +kubebuilder:validation:Optional
	ExistingObjectReplication []ExistingObjectReplicationParameters `json:"existingObjectReplication,omitempty" tf:"existing_object_replication,omitempty"`

	// Filter that identifies subset of objects to which the replication rule applies. See below. If not specified, the rule will default to using prefix.
	// +kubebuilder:validation:Optional
	Filter []BucketReplicationConfigurationRuleFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Priority associated with the rule. Priority should only be set if filter is configured. If not provided, defaults to 0. Priority must be unique between multiple rules.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Specifies special object selection criteria. See below.
	// +kubebuilder:validation:Optional
	SourceSelectionCriteria []RuleSourceSelectionCriteriaParameters `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`

	// Status of the rule. Either "Enabled" or "Disabled". The rule is ignored if status is not "Enabled".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DeleteMarkerReplicationInitParameters struct {

	// Whether delete markers should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DeleteMarkerReplicationObservation struct {

	// Whether delete markers should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DeleteMarkerReplicationParameters struct {

	// Whether delete markers should be replicated. Either "Enabled" or "Disabled".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DestinationAccessControlTranslationInitParameters struct {

	// Specifies the replica ownership. For default and valid values, see PUT bucket replication in the Amazon S3 API Reference. Valid values: Destination.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type DestinationAccessControlTranslationObservation struct {

	// Specifies the replica ownership. For default and valid values, see PUT bucket replication in the Amazon S3 API Reference. Valid values: Destination.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type DestinationAccessControlTranslationParameters struct {

	// Specifies the replica ownership. For default and valid values, see PUT bucket replication in the Amazon S3 API Reference. Valid values: Destination.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type DestinationMetricsInitParameters struct {

	// Configuration block that specifies the time threshold for emitting the s3:Replication:OperationMissedThreshold event. See below.
	EventThreshold []EventThresholdInitParameters `json:"eventThreshold,omitempty" tf:"event_threshold,omitempty"`

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DestinationMetricsObservation struct {

	// Configuration block that specifies the time threshold for emitting the s3:Replication:OperationMissedThreshold event. See below.
	EventThreshold []EventThresholdObservation `json:"eventThreshold,omitempty" tf:"event_threshold,omitempty"`

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DestinationMetricsParameters struct {

	// Configuration block that specifies the time threshold for emitting the s3:Replication:OperationMissedThreshold event. See below.
	// +kubebuilder:validation:Optional
	EventThreshold []EventThresholdParameters `json:"eventThreshold,omitempty" tf:"event_threshold,omitempty"`

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DestinationReplicationTimeInitParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Configuration block specifying the time by which replication should be complete for all objects and operations on objects. See below.
	Time []TimeInitParameters `json:"time,omitempty" tf:"time,omitempty"`
}

type DestinationReplicationTimeObservation struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Configuration block specifying the time by which replication should be complete for all objects and operations on objects. See below.
	Time []TimeObservation `json:"time,omitempty" tf:"time,omitempty"`
}

type DestinationReplicationTimeParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Configuration block specifying the time by which replication should be complete for all objects and operations on objects. See below.
	// +kubebuilder:validation:Optional
	Time []TimeParameters `json:"time,omitempty" tf:"time,omitempty"`
}

type EncryptionConfigurationInitParameters struct {

	// ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket.
	ReplicaKMSKeyID *string `json:"replicaKmsKeyId,omitempty" tf:"replica_kms_key_id,omitempty"`
}

type EncryptionConfigurationObservation struct {

	// ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket.
	ReplicaKMSKeyID *string `json:"replicaKmsKeyId,omitempty" tf:"replica_kms_key_id,omitempty"`
}

type EncryptionConfigurationParameters struct {

	// ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket.
	// +kubebuilder:validation:Optional
	ReplicaKMSKeyID *string `json:"replicaKmsKeyId,omitempty" tf:"replica_kms_key_id,omitempty"`
}

type EventThresholdInitParameters struct {

	// Time in minutes. Valid values: 15.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type EventThresholdObservation struct {

	// Time in minutes. Valid values: 15.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type EventThresholdParameters struct {

	// Time in minutes. Valid values: 15.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type ExistingObjectReplicationInitParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ExistingObjectReplicationObservation struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ExistingObjectReplicationParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type FilterAndInitParameters struct {

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Map of tags (key and value pairs) that identifies a subset of objects to which the rule applies. The rule applies only to objects having all the tags in its tagset.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FilterAndObservation struct {

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Map of tags (key and value pairs) that identifies a subset of objects to which the rule applies. The rule applies only to objects having all the tags in its tagset.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FilterAndParameters struct {

	// Object key name prefix identifying one or more objects to which the rule applies. Must be less than or equal to 1024 characters in length. Defaults to an empty string ("") if filter is not specified.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// Map of tags (key and value pairs) that identifies a subset of objects to which the rule applies. The rule applies only to objects having all the tags in its tagset.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FilterTagInitParameters struct {

	// Name of the object key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of the tag.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterTagObservation struct {

	// Name of the object key.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of the tag.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FilterTagParameters struct {

	// Name of the object key.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of the tag.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ReplicaModificationsInitParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ReplicaModificationsObservation struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ReplicaModificationsParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RuleDestinationInitParameters struct {

	// Configuration block that specifies the overrides to use for object owners on replication. See below. Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object. Must be used in conjunction with account owner override configuration.
	AccessControlTranslation []DestinationAccessControlTranslationInitParameters `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`

	// Account ID to specify the replica ownership. Must be used in conjunction with access_control_translation override configuration.
	Account *string `json:"account,omitempty" tf:"account,omitempty"`

	// Configuration block that provides information about encryption. See below. If source_selection_criteria is specified, you must specify this element.
	EncryptionConfiguration []EncryptionConfigurationInitParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// Configuration block that specifies replication metrics-related settings enabling replication metrics and events. See below.
	Metrics []DestinationMetricsInitParameters `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Configuration block that specifies S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. See below. Replication Time Control must be used in conjunction with metrics.
	ReplicationTime []DestinationReplicationTimeInitParameters `json:"replicationTime,omitempty" tf:"replication_time,omitempty"`

	// The storage class used to store the object. By default, Amazon S3 uses the storage class of the source object to create the object replica.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleDestinationObservation struct {

	// Configuration block that specifies the overrides to use for object owners on replication. See below. Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object. Must be used in conjunction with account owner override configuration.
	AccessControlTranslation []DestinationAccessControlTranslationObservation `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`

	// Account ID to specify the replica ownership. Must be used in conjunction with access_control_translation override configuration.
	Account *string `json:"account,omitempty" tf:"account,omitempty"`

	// ARN of the bucket where you want Amazon S3 to store the results.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Configuration block that provides information about encryption. See below. If source_selection_criteria is specified, you must specify this element.
	EncryptionConfiguration []EncryptionConfigurationObservation `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// Configuration block that specifies replication metrics-related settings enabling replication metrics and events. See below.
	Metrics []DestinationMetricsObservation `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Configuration block that specifies S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. See below. Replication Time Control must be used in conjunction with metrics.
	ReplicationTime []DestinationReplicationTimeObservation `json:"replicationTime,omitempty" tf:"replication_time,omitempty"`

	// The storage class used to store the object. By default, Amazon S3 uses the storage class of the source object to create the object replica.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleDestinationParameters struct {

	// Configuration block that specifies the overrides to use for object owners on replication. See below. Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object. Must be used in conjunction with account owner override configuration.
	// +kubebuilder:validation:Optional
	AccessControlTranslation []DestinationAccessControlTranslationParameters `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`

	// Account ID to specify the replica ownership. Must be used in conjunction with access_control_translation override configuration.
	// +kubebuilder:validation:Optional
	Account *string `json:"account,omitempty" tf:"account,omitempty"`

	// ARN of the bucket where you want Amazon S3 to store the results.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Configuration block that provides information about encryption. See below. If source_selection_criteria is specified, you must specify this element.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// Configuration block that specifies replication metrics-related settings enabling replication metrics and events. See below.
	// +kubebuilder:validation:Optional
	Metrics []DestinationMetricsParameters `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// Configuration block that specifies S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. See below. Replication Time Control must be used in conjunction with metrics.
	// +kubebuilder:validation:Optional
	ReplicationTime []DestinationReplicationTimeParameters `json:"replicationTime,omitempty" tf:"replication_time,omitempty"`

	// The storage class used to store the object. By default, Amazon S3 uses the storage class of the source object to create the object replica.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type RuleSourceSelectionCriteriaInitParameters struct {

	// Configuration block that you can specify for selections for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when filter is specified), you can specify this element and set the status to Enabled to replicate modifications on replicas.
	ReplicaModifications []ReplicaModificationsInitParameters `json:"replicaModifications,omitempty" tf:"replica_modifications,omitempty"`

	// Configuration block for filter information for the selection of Amazon S3 objects encrypted with AWS KMS. If specified, replica_kms_key_id in destination encryption_configuration must be specified as well.
	SseKMSEncryptedObjects []SourceSelectionCriteriaSseKMSEncryptedObjectsInitParameters `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type RuleSourceSelectionCriteriaObservation struct {

	// Configuration block that you can specify for selections for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when filter is specified), you can specify this element and set the status to Enabled to replicate modifications on replicas.
	ReplicaModifications []ReplicaModificationsObservation `json:"replicaModifications,omitempty" tf:"replica_modifications,omitempty"`

	// Configuration block for filter information for the selection of Amazon S3 objects encrypted with AWS KMS. If specified, replica_kms_key_id in destination encryption_configuration must be specified as well.
	SseKMSEncryptedObjects []SourceSelectionCriteriaSseKMSEncryptedObjectsObservation `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type RuleSourceSelectionCriteriaParameters struct {

	// Configuration block that you can specify for selections for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when filter is specified), you can specify this element and set the status to Enabled to replicate modifications on replicas.
	// +kubebuilder:validation:Optional
	ReplicaModifications []ReplicaModificationsParameters `json:"replicaModifications,omitempty" tf:"replica_modifications,omitempty"`

	// Configuration block for filter information for the selection of Amazon S3 objects encrypted with AWS KMS. If specified, replica_kms_key_id in destination encryption_configuration must be specified as well.
	// +kubebuilder:validation:Optional
	SseKMSEncryptedObjects []SourceSelectionCriteriaSseKMSEncryptedObjectsParameters `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type SourceSelectionCriteriaSseKMSEncryptedObjectsInitParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SourceSelectionCriteriaSseKMSEncryptedObjectsObservation struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SourceSelectionCriteriaSseKMSEncryptedObjectsParameters struct {

	// Whether the existing objects should be replicated. Either "Enabled" or "Disabled".
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type TimeInitParameters struct {

	// Time in minutes. Valid values: 15.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type TimeObservation struct {

	// Time in minutes. Valid values: 15.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type TimeParameters struct {

	// Time in minutes. Valid values: 15.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

// BucketReplicationConfigurationSpec defines the desired state of BucketReplicationConfiguration
type BucketReplicationConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketReplicationConfigurationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider BucketReplicationConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketReplicationConfigurationStatus defines the observed state of BucketReplicationConfiguration.
type BucketReplicationConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketReplicationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketReplicationConfiguration is the Schema for the BucketReplicationConfigurations API. Provides a S3 bucket replication configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketReplicationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || has(self.initProvider.rule)",message="rule is a required parameter"
	Spec   BucketReplicationConfigurationSpec   `json:"spec"`
	Status BucketReplicationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketReplicationConfigurationList contains a list of BucketReplicationConfigurations
type BucketReplicationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketReplicationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketReplicationConfiguration_Kind             = "BucketReplicationConfiguration"
	BucketReplicationConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketReplicationConfiguration_Kind}.String()
	BucketReplicationConfiguration_KindAPIVersion   = BucketReplicationConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketReplicationConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketReplicationConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketReplicationConfiguration{}, &BucketReplicationConfigurationList{})
}
