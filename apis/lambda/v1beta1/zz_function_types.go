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

type DeadLetterConfigInitParameters struct {

	// ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM role must be granted suitable access to write to the target object, which means allowing either the sns:Publish or sqs:SendMessage action on this ARN, depending on which service is targeted.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
}

type DeadLetterConfigObservation struct {

	// ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM role must be granted suitable access to write to the target object, which means allowing either the sns:Publish or sqs:SendMessage action on this ARN, depending on which service is targeted.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
}

type DeadLetterConfigParameters struct {

	// ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM role must be granted suitable access to write to the target object, which means allowing either the sns:Publish or sqs:SendMessage action on this ARN, depending on which service is targeted.
	// +kubebuilder:validation:Optional
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
}

type EnvironmentInitParameters struct {

	// Map of environment variables that are accessible from the function code during execution. If provided at least one key must be present.
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type EnvironmentObservation struct {

	// Map of environment variables that are accessible from the function code during execution. If provided at least one key must be present.
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type EnvironmentParameters struct {

	// Map of environment variables that are accessible from the function code during execution. If provided at least one key must be present.
	// +kubebuilder:validation:Optional
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type EphemeralStorageInitParameters struct {

	// The size of the Lambda function Ephemeral storage(/tmp) represented in MB. The minimum supported ephemeral_storage value defaults to 512MB and the maximum supported value is 10240MB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type EphemeralStorageObservation struct {

	// The size of the Lambda function Ephemeral storage(/tmp) represented in MB. The minimum supported ephemeral_storage value defaults to 512MB and the maximum supported value is 10240MB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type EphemeralStorageParameters struct {

	// The size of the Lambda function Ephemeral storage(/tmp) represented in MB. The minimum supported ephemeral_storage value defaults to 512MB and the maximum supported value is 10240MB.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type FileSystemConfigInitParameters struct {

	// Path where the function can access the file system, starting with /mnt/.
	LocalMountPath *string `json:"localMountPath,omitempty" tf:"local_mount_path,omitempty"`
}

type FileSystemConfigObservation struct {

	// Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Path where the function can access the file system, starting with /mnt/.
	LocalMountPath *string `json:"localMountPath,omitempty" tf:"local_mount_path,omitempty"`
}

type FileSystemConfigParameters struct {

	// Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/efs/v1beta1.AccessPoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a AccessPoint in efs to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a AccessPoint in efs to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// Path where the function can access the file system, starting with /mnt/.
	// +kubebuilder:validation:Optional
	LocalMountPath *string `json:"localMountPath,omitempty" tf:"local_mount_path,omitempty"`
}

type FunctionInitParameters struct {

	// Instruction set architecture for your Lambda function. Valid values are ["x86_64"] and ["arm64"]. Default is ["x86_64"]. Removing this attribute, function's architecture stay the same.
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
	CodeSigningConfigArn *string `json:"codeSigningConfigArn,omitempty" tf:"code_signing_config_arn,omitempty"`

	// Configuration block. Detailed below.
	DeadLetterConfig []DeadLetterConfigInitParameters `json:"deadLetterConfig,omitempty" tf:"dead_letter_config,omitempty"`

	// Description of what your Lambda Function does.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block. Detailed below.
	Environment []EnvironmentInitParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// The amount of Ephemeral storage(/tmp) to allocate for the Lambda Function in MB. This parameter is used to expand the total amount of Ephemeral storage available, beyond the default amount of 512MB. Detailed below.
	EphemeralStorage []EphemeralStorageInitParameters `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// Configuration block. Detailed below.
	FileSystemConfig []FileSystemConfigInitParameters `json:"fileSystemConfig,omitempty" tf:"file_system_config,omitempty"`

	// Function entrypoint in your code.
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	// Configuration block. Detailed below.
	ImageConfig []ImageConfigInitParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// ECR image URI containing the function's deployment package. Exactly one of filename, image_uri,  or s3_bucket must be specified.
	ImageURI *string `json:"imageUri,omitempty" tf:"image_uri,omitempty"`

	// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See Lambda Layers
	Layers []*string `json:"layers,omitempty" tf:"layers,omitempty"`

	// Amount of memory in MB your Lambda Function can use at runtime. Defaults to 128. See Limits
	MemorySize *float64 `json:"memorySize,omitempty" tf:"memory_size,omitempty"`

	// Lambda deployment package type. Valid values are Zip and Image. Defaults to Zip.
	PackageType *string `json:"packageType,omitempty" tf:"package_type,omitempty"`

	// Whether to publish creation/change as new Lambda Function Version. Defaults to false.
	Publish *bool `json:"publish,omitempty" tf:"publish,omitempty"`

	// Whether to replace the security groups on associated lambda network interfaces upon destruction. Removing these security groups from orphaned network interfaces can speed up security group deletion times by avoiding a dependency on AWS's internal cleanup operations. By default, the ENI security groups will be replaced with the default security group in the function's VPC. Set the replacement_security_group_ids attribute to use a custom list of security groups for replacement.
	ReplaceSecurityGroupsOnDestroy *bool `json:"replaceSecurityGroupsOnDestroy,omitempty" tf:"replace_security_groups_on_destroy,omitempty"`

	// Amount of reserved concurrent executions for this lambda function. A value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. Defaults to Unreserved Concurrency Limits -1. See Managing Concurrency
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions,omitempty" tf:"reserved_concurrent_executions,omitempty"`

	// Identifier of the function's runtime. See Runtimes for valid values.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// S3 key of an object containing the function's deployment package. When s3_bucket is set, s3_key is required.
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`

	// Object version containing the function's deployment package. Conflicts with filename and image_uri.
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`

	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Snap start settings block. Detailed below.
	SnapStart []SnapStartInitParameters `json:"snapStart,omitempty" tf:"snap_start,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key. The usual way to set this is filebase64sha256("file.11.12 and later) or base64sha256(file("file.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Amount of time your Lambda Function has to run in seconds. Defaults to 3. See Limits.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Configuration block. Detailed below.
	TracingConfig []TracingConfigInitParameters `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`

	// Configuration block. Detailed below.
	VPCConfig []VPCConfigInitParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type FunctionObservation struct {

	// Instruction set architecture for your Lambda function. Valid values are ["x86_64"] and ["arm64"]. Default is ["x86_64"]. Removing this attribute, function's architecture stay the same.
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// Amazon Resource Name (ARN) identifying your Lambda Function.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
	CodeSigningConfigArn *string `json:"codeSigningConfigArn,omitempty" tf:"code_signing_config_arn,omitempty"`

	// Configuration block. Detailed below.
	DeadLetterConfig []DeadLetterConfigObservation `json:"deadLetterConfig,omitempty" tf:"dead_letter_config,omitempty"`

	// Description of what your Lambda Function does.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block. Detailed below.
	Environment []EnvironmentObservation `json:"environment,omitempty" tf:"environment,omitempty"`

	// The amount of Ephemeral storage(/tmp) to allocate for the Lambda Function in MB. This parameter is used to expand the total amount of Ephemeral storage available, beyond the default amount of 512MB. Detailed below.
	EphemeralStorage []EphemeralStorageObservation `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// Configuration block. Detailed below.
	FileSystemConfig []FileSystemConfigObservation `json:"fileSystemConfig,omitempty" tf:"file_system_config,omitempty"`

	// Function entrypoint in your code.
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block. Detailed below.
	ImageConfig []ImageConfigObservation `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// ECR image URI containing the function's deployment package. Exactly one of filename, image_uri,  or s3_bucket must be specified.
	ImageURI *string `json:"imageUri,omitempty" tf:"image_uri,omitempty"`

	// ARN to be used for invoking Lambda Function from API Gateway - to be used in aws_api_gateway_integration's uri.
	InvokeArn *string `json:"invokeArn,omitempty" tf:"invoke_arn,omitempty"`

	// Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. To fix the perpetual difference, remove this configuration.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Date this resource was last modified.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See Lambda Layers
	Layers []*string `json:"layers,omitempty" tf:"layers,omitempty"`

	// Amount of memory in MB your Lambda Function can use at runtime. Defaults to 128. See Limits
	MemorySize *float64 `json:"memorySize,omitempty" tf:"memory_size,omitempty"`

	// Lambda deployment package type. Valid values are Zip and Image. Defaults to Zip.
	PackageType *string `json:"packageType,omitempty" tf:"package_type,omitempty"`

	// Whether to publish creation/change as new Lambda Function Version. Defaults to false.
	Publish *bool `json:"publish,omitempty" tf:"publish,omitempty"`

	// ARN identifying your Lambda Function Version (if versioning is enabled via publish = true).
	QualifiedArn *string `json:"qualifiedArn,omitempty" tf:"qualified_arn,omitempty"`

	// Qualified ARN (ARN with lambda version number) to be used for invoking Lambda Function from API Gateway - to be used in aws_api_gateway_integration's uri.
	QualifiedInvokeArn *string `json:"qualifiedInvokeArn,omitempty" tf:"qualified_invoke_arn,omitempty"`

	// Whether to replace the security groups on associated lambda network interfaces upon destruction. Removing these security groups from orphaned network interfaces can speed up security group deletion times by avoiding a dependency on AWS's internal cleanup operations. By default, the ENI security groups will be replaced with the default security group in the function's VPC. Set the replacement_security_group_ids attribute to use a custom list of security groups for replacement.
	ReplaceSecurityGroupsOnDestroy *bool `json:"replaceSecurityGroupsOnDestroy,omitempty" tf:"replace_security_groups_on_destroy,omitempty"`

	// List of security group IDs to assign to orphaned Lambda function network interfaces upon destruction. replace_security_groups_on_destroy must be set to true to use this attribute.
	ReplacementSecurityGroupIds []*string `json:"replacementSecurityGroupIds,omitempty" tf:"replacement_security_group_ids,omitempty"`

	// Amount of reserved concurrent executions for this lambda function. A value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. Defaults to Unreserved Concurrency Limits -1. See Managing Concurrency
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions,omitempty" tf:"reserved_concurrent_executions,omitempty"`

	// Amazon Resource Name (ARN) of the function's execution role. The role provides the function's identity and access to AWS services and resources.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Identifier of the function's runtime. See Runtimes for valid values.
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// S3 bucket location containing the function's deployment package. This bucket must reside in the same AWS region where you are creating the Lambda function. Exactly one of filename, image_uri, or s3_bucket must be specified. When s3_bucket is set, s3_key is required.
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// S3 key of an object containing the function's deployment package. When s3_bucket is set, s3_key is required.
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`

	// Object version containing the function's deployment package. Conflicts with filename and image_uri.
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`

	// ARN of the signing job.
	SigningJobArn *string `json:"signingJobArn,omitempty" tf:"signing_job_arn,omitempty"`

	// ARN of the signing profile version.
	SigningProfileVersionArn *string `json:"signingProfileVersionArn,omitempty" tf:"signing_profile_version_arn,omitempty"`

	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Snap start settings block. Detailed below.
	SnapStart []SnapStartObservation `json:"snapStart,omitempty" tf:"snap_start,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key. The usual way to set this is filebase64sha256("file.11.12 and later) or base64sha256(file("file.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`

	// Size in bytes of the function .zip file.
	SourceCodeSize *float64 `json:"sourceCodeSize,omitempty" tf:"source_code_size,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Amount of time your Lambda Function has to run in seconds. Defaults to 3. See Limits.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Configuration block. Detailed below.
	TracingConfig []TracingConfigObservation `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`

	// Configuration block. Detailed below.
	VPCConfig []VPCConfigObservation `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`

	// Latest published version of your Lambda Function.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FunctionParameters struct {

	// Instruction set architecture for your Lambda function. Valid values are ["x86_64"] and ["arm64"]. Default is ["x86_64"]. Removing this attribute, function's architecture stay the same.
	// +kubebuilder:validation:Optional
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
	// +kubebuilder:validation:Optional
	CodeSigningConfigArn *string `json:"codeSigningConfigArn,omitempty" tf:"code_signing_config_arn,omitempty"`

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	DeadLetterConfig []DeadLetterConfigParameters `json:"deadLetterConfig,omitempty" tf:"dead_letter_config,omitempty"`

	// Description of what your Lambda Function does.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	Environment []EnvironmentParameters `json:"environment,omitempty" tf:"environment,omitempty"`

	// The amount of Ephemeral storage(/tmp) to allocate for the Lambda Function in MB. This parameter is used to expand the total amount of Ephemeral storage available, beyond the default amount of 512MB. Detailed below.
	// +kubebuilder:validation:Optional
	EphemeralStorage []EphemeralStorageParameters `json:"ephemeralStorage,omitempty" tf:"ephemeral_storage,omitempty"`

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	FileSystemConfig []FileSystemConfigParameters `json:"fileSystemConfig,omitempty" tf:"file_system_config,omitempty"`

	// Function entrypoint in your code.
	// +kubebuilder:validation:Optional
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	ImageConfig []ImageConfigParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// ECR image URI containing the function's deployment package. Exactly one of filename, image_uri,  or s3_bucket must be specified.
	// +kubebuilder:validation:Optional
	ImageURI *string `json:"imageUri,omitempty" tf:"image_uri,omitempty"`

	// Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. To fix the perpetual difference, remove this configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See Lambda Layers
	// +kubebuilder:validation:Optional
	Layers []*string `json:"layers,omitempty" tf:"layers,omitempty"`

	// Amount of memory in MB your Lambda Function can use at runtime. Defaults to 128. See Limits
	// +kubebuilder:validation:Optional
	MemorySize *float64 `json:"memorySize,omitempty" tf:"memory_size,omitempty"`

	// Lambda deployment package type. Valid values are Zip and Image. Defaults to Zip.
	// +kubebuilder:validation:Optional
	PackageType *string `json:"packageType,omitempty" tf:"package_type,omitempty"`

	// Whether to publish creation/change as new Lambda Function Version. Defaults to false.
	// +kubebuilder:validation:Optional
	Publish *bool `json:"publish,omitempty" tf:"publish,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether to replace the security groups on associated lambda network interfaces upon destruction. Removing these security groups from orphaned network interfaces can speed up security group deletion times by avoiding a dependency on AWS's internal cleanup operations. By default, the ENI security groups will be replaced with the default security group in the function's VPC. Set the replacement_security_group_ids attribute to use a custom list of security groups for replacement.
	// +kubebuilder:validation:Optional
	ReplaceSecurityGroupsOnDestroy *bool `json:"replaceSecurityGroupsOnDestroy,omitempty" tf:"replace_security_groups_on_destroy,omitempty"`

	// References to SecurityGroup in ec2 to populate replacementSecurityGroupIds.
	// +kubebuilder:validation:Optional
	ReplacementSecurityGroupIDRefs []v1.Reference `json:"replacementSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate replacementSecurityGroupIds.
	// +kubebuilder:validation:Optional
	ReplacementSecurityGroupIDSelector *v1.Selector `json:"replacementSecurityGroupIdSelector,omitempty" tf:"-"`

	// List of security group IDs to assign to orphaned Lambda function network interfaces upon destruction. replace_security_groups_on_destroy must be set to true to use this attribute.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=ReplacementSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=ReplacementSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	ReplacementSecurityGroupIds []*string `json:"replacementSecurityGroupIds,omitempty" tf:"replacement_security_group_ids,omitempty"`

	// Amount of reserved concurrent executions for this lambda function. A value of 0 disables lambda from being triggered and -1 removes any concurrency limitations. Defaults to Unreserved Concurrency Limits -1. See Managing Concurrency
	// +kubebuilder:validation:Optional
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions,omitempty" tf:"reserved_concurrent_executions,omitempty"`

	// Amazon Resource Name (ARN) of the function's execution role. The role provides the function's identity and access to AWS services and resources.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// Identifier of the function's runtime. See Runtimes for valid values.
	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// S3 bucket location containing the function's deployment package. This bucket must reside in the same AWS region where you are creating the Lambda function. Exactly one of filename, image_uri, or s3_bucket must be specified. When s3_bucket is set, s3_key is required.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// Reference to a Bucket in s3 to populate s3Bucket.
	// +kubebuilder:validation:Optional
	S3BucketRef *v1.Reference `json:"s3BucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate s3Bucket.
	// +kubebuilder:validation:Optional
	S3BucketSelector *v1.Selector `json:"s3BucketSelector,omitempty" tf:"-"`

	// S3 key of an object containing the function's deployment package. When s3_bucket is set, s3_key is required.
	// +kubebuilder:validation:Optional
	S3Key *string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`

	// Object version containing the function's deployment package. Conflicts with filename and image_uri.
	// +kubebuilder:validation:Optional
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`

	// +kubebuilder:validation:Optional
	SkipDestroy *bool `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`

	// Snap start settings block. Detailed below.
	// +kubebuilder:validation:Optional
	SnapStart []SnapStartParameters `json:"snapStart,omitempty" tf:"snap_start,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either filename or s3_key. The usual way to set this is filebase64sha256("file.11.12 and later) or base64sha256(file("file.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
	// +kubebuilder:validation:Optional
	SourceCodeHash *string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Amount of time your Lambda Function has to run in seconds. Defaults to 3. See Limits.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	TracingConfig []TracingConfigParameters `json:"tracingConfig,omitempty" tf:"tracing_config,omitempty"`

	// Configuration block. Detailed below.
	// +kubebuilder:validation:Optional
	VPCConfig []VPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type ImageConfigInitParameters struct {

	// Parameters that you want to pass in with entry_point.
	Command []*string `json:"command,omitempty" tf:"command,omitempty"`

	// Entry point to your application, which is typically the location of the runtime executable.
	EntryPoint []*string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// Working directory.
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type ImageConfigObservation struct {

	// Parameters that you want to pass in with entry_point.
	Command []*string `json:"command,omitempty" tf:"command,omitempty"`

	// Entry point to your application, which is typically the location of the runtime executable.
	EntryPoint []*string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// Working directory.
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type ImageConfigParameters struct {

	// Parameters that you want to pass in with entry_point.
	// +kubebuilder:validation:Optional
	Command []*string `json:"command,omitempty" tf:"command,omitempty"`

	// Entry point to your application, which is typically the location of the runtime executable.
	// +kubebuilder:validation:Optional
	EntryPoint []*string `json:"entryPoint,omitempty" tf:"entry_point,omitempty"`

	// Working directory.
	// +kubebuilder:validation:Optional
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type SnapStartInitParameters struct {

	// Conditions where snap start is enabled. Valid values are PublishedVersions.
	ApplyOn *string `json:"applyOn,omitempty" tf:"apply_on,omitempty"`
}

type SnapStartObservation struct {

	// Conditions where snap start is enabled. Valid values are PublishedVersions.
	ApplyOn *string `json:"applyOn,omitempty" tf:"apply_on,omitempty"`

	// Optimization status of the snap start configuration. Valid values are On and Off.
	OptimizationStatus *string `json:"optimizationStatus,omitempty" tf:"optimization_status,omitempty"`
}

type SnapStartParameters struct {

	// Conditions where snap start is enabled. Valid values are PublishedVersions.
	// +kubebuilder:validation:Optional
	ApplyOn *string `json:"applyOn,omitempty" tf:"apply_on,omitempty"`
}

type TracingConfigInitParameters struct {

	// Whether to sample and trace a subset of incoming requests with AWS X-Ray. Valid values are PassThrough and Active. If PassThrough, Lambda will only trace the request from an upstream service if it contains a tracing header with "sampled=1". If Active, Lambda will respect any tracing header it receives from an upstream service. If no tracing header is received, Lambda will call X-Ray for a tracing decision.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type TracingConfigObservation struct {

	// Whether to sample and trace a subset of incoming requests with AWS X-Ray. Valid values are PassThrough and Active. If PassThrough, Lambda will only trace the request from an upstream service if it contains a tracing header with "sampled=1". If Active, Lambda will respect any tracing header it receives from an upstream service. If no tracing header is received, Lambda will call X-Ray for a tracing decision.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type TracingConfigParameters struct {

	// Whether to sample and trace a subset of incoming requests with AWS X-Ray. Valid values are PassThrough and Active. If PassThrough, Lambda will only trace the request from an upstream service if it contains a tracing header with "sampled=1". If Active, Lambda will respect any tracing header it receives from an upstream service. If no tracing header is received, Lambda will call X-Ray for a tracing decision.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type VPCConfigInitParameters struct {
}

type VPCConfigObservation struct {

	// List of security group IDs associated with the Lambda function.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// List of subnet IDs associated with the Lambda function.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// ID of the VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigParameters struct {

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// List of security group IDs associated with the Lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// List of subnet IDs associated with the Lambda function.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FunctionInitParameters `json:"initProvider,omitempty"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API. Provides a Lambda Function resource. Lambda allows you to trigger execution of code in response to events in AWS, enabling serverless backend solutions. The Lambda Function itself includes source code and runtime configuration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
