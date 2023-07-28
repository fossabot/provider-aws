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

type CatalogTargetInitParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// The ARN of the dead-letter SQS queue.
	DlqEventQueueArn *string `json:"dlqEventQueueArn,omitempty" tf:"dlq_event_queue_arn,omitempty"`

	// The ARN of the SQS queue to receive S3 notifications from.
	EventQueueArn *string `json:"eventQueueArn,omitempty" tf:"event_queue_arn,omitempty"`

	// A list of catalog tables to be synchronized.
	Tables []*string `json:"tables,omitempty" tf:"tables,omitempty"`
}

type CatalogTargetObservation struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Glue database where results are written.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The ARN of the dead-letter SQS queue.
	DlqEventQueueArn *string `json:"dlqEventQueueArn,omitempty" tf:"dlq_event_queue_arn,omitempty"`

	// The ARN of the SQS queue to receive S3 notifications from.
	EventQueueArn *string `json:"eventQueueArn,omitempty" tf:"event_queue_arn,omitempty"`

	// A list of catalog tables to be synchronized.
	Tables []*string `json:"tables,omitempty" tf:"tables,omitempty"`
}

type CatalogTargetParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	// +kubebuilder:validation:Optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Glue database where results are written.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.CatalogDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// The ARN of the dead-letter SQS queue.
	// +kubebuilder:validation:Optional
	DlqEventQueueArn *string `json:"dlqEventQueueArn,omitempty" tf:"dlq_event_queue_arn,omitempty"`

	// The ARN of the SQS queue to receive S3 notifications from.
	// +kubebuilder:validation:Optional
	EventQueueArn *string `json:"eventQueueArn,omitempty" tf:"event_queue_arn,omitempty"`

	// A list of catalog tables to be synchronized.
	// +kubebuilder:validation:Optional
	Tables []*string `json:"tables,omitempty" tf:"tables,omitempty"`
}

type CrawlerInitParameters struct {
	CatalogTarget []CatalogTargetInitParameters `json:"catalogTarget,omitempty" tf:"catalog_target,omitempty"`

	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers []*string `json:"classifiers,omitempty" tf:"classifiers,omitempty"`

	// JSON string of configuration information. For more details see Setting Crawler Configuration Options.
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	DeltaTarget []DeltaTargetInitParameters `json:"deltaTarget,omitempty" tf:"delta_target,omitempty"`

	// Description of the crawler.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of nested DynamoDB target arguments. See Dynamodb Target below.
	DynamodbTarget []DynamodbTargetInitParameters `json:"dynamodbTarget,omitempty" tf:"dynamodb_target,omitempty"`

	// List of nested JBDC target arguments. See JDBC Target below.
	JdbcTarget []JdbcTargetInitParameters `json:"jdbcTarget,omitempty" tf:"jdbc_target,omitempty"`

	// Specifies Lake Formation configuration settings for the crawler. See Lake Formation Configuration below.
	LakeFormationConfiguration []LakeFormationConfigurationInitParameters `json:"lakeFormationConfiguration,omitempty" tf:"lake_formation_configuration,omitempty"`

	// Specifies data lineage configuration settings for the crawler. See Lineage Configuration below.
	LineageConfiguration []LineageConfigurationInitParameters `json:"lineageConfiguration,omitempty" tf:"lineage_configuration,omitempty"`

	// List nested MongoDB target arguments. See MongoDB Target below.
	MongodbTarget []MongodbTargetInitParameters `json:"mongodbTarget,omitempty" tf:"mongodb_target,omitempty"`

	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.. See Recrawl Policy below.
	RecrawlPolicy []RecrawlPolicyInitParameters `json:"recrawlPolicy,omitempty" tf:"recrawl_policy,omitempty"`

	// List nested Amazon S3 target arguments. See S3 Target below.
	S3Target []S3TargetInitParameters `json:"s3Target,omitempty" tf:"s3_target,omitempty"`

	// Based Schedules for Jobs and Crawlers. For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *).
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Policy for the crawler's update and deletion behavior. See Schema Change Policy below.
	SchemaChangePolicy []SchemaChangePolicyInitParameters `json:"schemaChangePolicy,omitempty" tf:"schema_change_policy,omitempty"`

	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// The table prefix used for catalog tables that are created.
	TablePrefix *string `json:"tablePrefix,omitempty" tf:"table_prefix,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CrawlerObservation struct {

	// The ARN of the crawler
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CatalogTarget []CatalogTargetObservation `json:"catalogTarget,omitempty" tf:"catalog_target,omitempty"`

	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	Classifiers []*string `json:"classifiers,omitempty" tf:"classifiers,omitempty"`

	// JSON string of configuration information. For more details see Setting Crawler Configuration Options.
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Glue database where results are written.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	DeltaTarget []DeltaTargetObservation `json:"deltaTarget,omitempty" tf:"delta_target,omitempty"`

	// Description of the crawler.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of nested DynamoDB target arguments. See Dynamodb Target below.
	DynamodbTarget []DynamodbTargetObservation `json:"dynamodbTarget,omitempty" tf:"dynamodb_target,omitempty"`

	// Crawler name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of nested JBDC target arguments. See JDBC Target below.
	JdbcTarget []JdbcTargetObservation `json:"jdbcTarget,omitempty" tf:"jdbc_target,omitempty"`

	// Specifies Lake Formation configuration settings for the crawler. See Lake Formation Configuration below.
	LakeFormationConfiguration []LakeFormationConfigurationObservation `json:"lakeFormationConfiguration,omitempty" tf:"lake_formation_configuration,omitempty"`

	// Specifies data lineage configuration settings for the crawler. See Lineage Configuration below.
	LineageConfiguration []LineageConfigurationObservation `json:"lineageConfiguration,omitempty" tf:"lineage_configuration,omitempty"`

	// List nested MongoDB target arguments. See MongoDB Target below.
	MongodbTarget []MongodbTargetObservation `json:"mongodbTarget,omitempty" tf:"mongodb_target,omitempty"`

	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.. See Recrawl Policy below.
	RecrawlPolicy []RecrawlPolicyObservation `json:"recrawlPolicy,omitempty" tf:"recrawl_policy,omitempty"`

	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// List nested Amazon S3 target arguments. See S3 Target below.
	S3Target []S3TargetObservation `json:"s3Target,omitempty" tf:"s3_target,omitempty"`

	// Based Schedules for Jobs and Crawlers. For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *).
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Policy for the crawler's update and deletion behavior. See Schema Change Policy below.
	SchemaChangePolicy []SchemaChangePolicyObservation `json:"schemaChangePolicy,omitempty" tf:"schema_change_policy,omitempty"`

	// The name of Security Configuration to be used by the crawler
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// The table prefix used for catalog tables that are created.
	TablePrefix *string `json:"tablePrefix,omitempty" tf:"table_prefix,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CrawlerParameters struct {

	// +kubebuilder:validation:Optional
	CatalogTarget []CatalogTargetParameters `json:"catalogTarget,omitempty" tf:"catalog_target,omitempty"`

	// List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
	// +kubebuilder:validation:Optional
	Classifiers []*string `json:"classifiers,omitempty" tf:"classifiers,omitempty"`

	// JSON string of configuration information. For more details see Setting Crawler Configuration Options.
	// +kubebuilder:validation:Optional
	Configuration *string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// Glue database where results are written.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.CatalogDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeltaTarget []DeltaTargetParameters `json:"deltaTarget,omitempty" tf:"delta_target,omitempty"`

	// Description of the crawler.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of nested DynamoDB target arguments. See Dynamodb Target below.
	// +kubebuilder:validation:Optional
	DynamodbTarget []DynamodbTargetParameters `json:"dynamodbTarget,omitempty" tf:"dynamodb_target,omitempty"`

	// List of nested JBDC target arguments. See JDBC Target below.
	// +kubebuilder:validation:Optional
	JdbcTarget []JdbcTargetParameters `json:"jdbcTarget,omitempty" tf:"jdbc_target,omitempty"`

	// Specifies Lake Formation configuration settings for the crawler. See Lake Formation Configuration below.
	// +kubebuilder:validation:Optional
	LakeFormationConfiguration []LakeFormationConfigurationParameters `json:"lakeFormationConfiguration,omitempty" tf:"lake_formation_configuration,omitempty"`

	// Specifies data lineage configuration settings for the crawler. See Lineage Configuration below.
	// +kubebuilder:validation:Optional
	LineageConfiguration []LineageConfigurationParameters `json:"lineageConfiguration,omitempty" tf:"lineage_configuration,omitempty"`

	// List nested MongoDB target arguments. See MongoDB Target below.
	// +kubebuilder:validation:Optional
	MongodbTarget []MongodbTargetParameters `json:"mongodbTarget,omitempty" tf:"mongodb_target,omitempty"`

	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.. See Recrawl Policy below.
	// +kubebuilder:validation:Optional
	RecrawlPolicy []RecrawlPolicyParameters `json:"recrawlPolicy,omitempty" tf:"recrawl_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
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

	// List nested Amazon S3 target arguments. See S3 Target below.
	// +kubebuilder:validation:Optional
	S3Target []S3TargetParameters `json:"s3Target,omitempty" tf:"s3_target,omitempty"`

	// Based Schedules for Jobs and Crawlers. For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *).
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Policy for the crawler's update and deletion behavior. See Schema Change Policy below.
	// +kubebuilder:validation:Optional
	SchemaChangePolicy []SchemaChangePolicyParameters `json:"schemaChangePolicy,omitempty" tf:"schema_change_policy,omitempty"`

	// The name of Security Configuration to be used by the crawler
	// +kubebuilder:validation:Optional
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`

	// The table prefix used for catalog tables that are created.
	// +kubebuilder:validation:Optional
	TablePrefix *string `json:"tablePrefix,omitempty" tf:"table_prefix,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DeltaTargetInitParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.
	CreateNativeDeltaTable *bool `json:"createNativeDeltaTable,omitempty" tf:"create_native_delta_table,omitempty"`

	// A list of the Amazon S3 paths to the Delta tables.
	DeltaTables []*string `json:"deltaTables,omitempty" tf:"delta_tables,omitempty"`

	// Specifies whether to write the manifest files to the Delta table path.
	WriteManifest *bool `json:"writeManifest,omitempty" tf:"write_manifest,omitempty"`
}

type DeltaTargetObservation struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.
	CreateNativeDeltaTable *bool `json:"createNativeDeltaTable,omitempty" tf:"create_native_delta_table,omitempty"`

	// A list of the Amazon S3 paths to the Delta tables.
	DeltaTables []*string `json:"deltaTables,omitempty" tf:"delta_tables,omitempty"`

	// Specifies whether to write the manifest files to the Delta table path.
	WriteManifest *bool `json:"writeManifest,omitempty" tf:"write_manifest,omitempty"`
}

type DeltaTargetParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	// +kubebuilder:validation:Optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.
	// +kubebuilder:validation:Optional
	CreateNativeDeltaTable *bool `json:"createNativeDeltaTable,omitempty" tf:"create_native_delta_table,omitempty"`

	// A list of the Amazon S3 paths to the Delta tables.
	// +kubebuilder:validation:Optional
	DeltaTables []*string `json:"deltaTables,omitempty" tf:"delta_tables,omitempty"`

	// Specifies whether to write the manifest files to the Delta table path.
	// +kubebuilder:validation:Optional
	WriteManifest *bool `json:"writeManifest,omitempty" tf:"write_manifest,omitempty"`
}

type DynamodbTargetInitParameters struct {

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table.  defaults to true.
	ScanAll *bool `json:"scanAll,omitempty" tf:"scan_all,omitempty"`

	// The percentage of the configured read capacity units to use by the AWS Glue crawler. The valid values are null or a value between 0.1 to 1.5.
	ScanRate *float64 `json:"scanRate,omitempty" tf:"scan_rate,omitempty"`
}

type DynamodbTargetObservation struct {

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table.  defaults to true.
	ScanAll *bool `json:"scanAll,omitempty" tf:"scan_all,omitempty"`

	// The percentage of the configured read capacity units to use by the AWS Glue crawler. The valid values are null or a value between 0.1 to 1.5.
	ScanRate *float64 `json:"scanRate,omitempty" tf:"scan_rate,omitempty"`
}

type DynamodbTargetParameters struct {

	// The name of the DynamoDB table to crawl.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table.  defaults to true.
	// +kubebuilder:validation:Optional
	ScanAll *bool `json:"scanAll,omitempty" tf:"scan_all,omitempty"`

	// The percentage of the configured read capacity units to use by the AWS Glue crawler. The valid values are null or a value between 0.1 to 1.5.
	// +kubebuilder:validation:Optional
	ScanRate *float64 `json:"scanRate,omitempty" tf:"scan_rate,omitempty"`
}

type JdbcTargetInitParameters struct {

	// Specify a value of RAWTYPES or COMMENTS to enable additional metadata intable responses. RAWTYPES provides the native-level datatype. COMMENTS provides comments associated with a column or table in the database.
	EnableAdditionalMetadata []*string `json:"enableAdditionalMetadata,omitempty" tf:"enable_additional_metadata,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type JdbcTargetObservation struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Specify a value of RAWTYPES or COMMENTS to enable additional metadata intable responses. RAWTYPES provides the native-level datatype. COMMENTS provides comments associated with a column or table in the database.
	EnableAdditionalMetadata []*string `json:"enableAdditionalMetadata,omitempty" tf:"enable_additional_metadata,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type JdbcTargetParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.Connection
	// +kubebuilder:validation:Optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Reference to a Connection in glue to populate connectionName.
	// +kubebuilder:validation:Optional
	ConnectionNameRef *v1.Reference `json:"connectionNameRef,omitempty" tf:"-"`

	// Selector for a Connection in glue to populate connectionName.
	// +kubebuilder:validation:Optional
	ConnectionNameSelector *v1.Selector `json:"connectionNameSelector,omitempty" tf:"-"`

	// Specify a value of RAWTYPES or COMMENTS to enable additional metadata intable responses. RAWTYPES provides the native-level datatype. COMMENTS provides comments associated with a column or table in the database.
	// +kubebuilder:validation:Optional
	EnableAdditionalMetadata []*string `json:"enableAdditionalMetadata,omitempty" tf:"enable_additional_metadata,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	// +kubebuilder:validation:Optional
	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The name of the DynamoDB table to crawl.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type LakeFormationConfigurationInitParameters struct {

	// Required for cross account crawls. For same account crawls as the target data, this can omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Specifies whether to use Lake Formation credentials for the crawler instead of the IAM role credentials.
	UseLakeFormationCredentials *bool `json:"useLakeFormationCredentials,omitempty" tf:"use_lake_formation_credentials,omitempty"`
}

type LakeFormationConfigurationObservation struct {

	// Required for cross account crawls. For same account crawls as the target data, this can omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Specifies whether to use Lake Formation credentials for the crawler instead of the IAM role credentials.
	UseLakeFormationCredentials *bool `json:"useLakeFormationCredentials,omitempty" tf:"use_lake_formation_credentials,omitempty"`
}

type LakeFormationConfigurationParameters struct {

	// Required for cross account crawls. For same account crawls as the target data, this can omitted.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Specifies whether to use Lake Formation credentials for the crawler instead of the IAM role credentials.
	// +kubebuilder:validation:Optional
	UseLakeFormationCredentials *bool `json:"useLakeFormationCredentials,omitempty" tf:"use_lake_formation_credentials,omitempty"`
}

type LineageConfigurationInitParameters struct {

	// Specifies whether data lineage is enabled for the crawler. Valid values are: ENABLE and DISABLE. Default value is Disable.
	CrawlerLineageSettings *string `json:"crawlerLineageSettings,omitempty" tf:"crawler_lineage_settings,omitempty"`
}

type LineageConfigurationObservation struct {

	// Specifies whether data lineage is enabled for the crawler. Valid values are: ENABLE and DISABLE. Default value is Disable.
	CrawlerLineageSettings *string `json:"crawlerLineageSettings,omitempty" tf:"crawler_lineage_settings,omitempty"`
}

type LineageConfigurationParameters struct {

	// Specifies whether data lineage is enabled for the crawler. Valid values are: ENABLE and DISABLE. Default value is Disable.
	// +kubebuilder:validation:Optional
	CrawlerLineageSettings *string `json:"crawlerLineageSettings,omitempty" tf:"crawler_lineage_settings,omitempty"`
}

type MongodbTargetInitParameters struct {

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table.  defaults to true.
	ScanAll *bool `json:"scanAll,omitempty" tf:"scan_all,omitempty"`
}

type MongodbTargetObservation struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table.  defaults to true.
	ScanAll *bool `json:"scanAll,omitempty" tf:"scan_all,omitempty"`
}

type MongodbTargetParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.Connection
	// +kubebuilder:validation:Optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// Reference to a Connection in glue to populate connectionName.
	// +kubebuilder:validation:Optional
	ConnectionNameRef *v1.Reference `json:"connectionNameRef,omitempty" tf:"-"`

	// Selector for a Connection in glue to populate connectionName.
	// +kubebuilder:validation:Optional
	ConnectionNameSelector *v1.Selector `json:"connectionNameSelector,omitempty" tf:"-"`

	// The name of the DynamoDB table to crawl.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Indicates whether to scan all the records, or to sample rows from the table. Scanning all the records can take a long time when the table is not a high throughput table.  defaults to true.
	// +kubebuilder:validation:Optional
	ScanAll *bool `json:"scanAll,omitempty" tf:"scan_all,omitempty"`
}

type RecrawlPolicyInitParameters struct {

	// Specifies whether to crawl the entire dataset again, crawl only folders that were added since the last crawler run, or crawl what S3 notifies the crawler of via SQS. Valid Values are: CRAWL_EVENT_MODE, CRAWL_EVERYTHING and CRAWL_NEW_FOLDERS_ONLY. Default value is CRAWL_EVERYTHING.
	RecrawlBehavior *string `json:"recrawlBehavior,omitempty" tf:"recrawl_behavior,omitempty"`
}

type RecrawlPolicyObservation struct {

	// Specifies whether to crawl the entire dataset again, crawl only folders that were added since the last crawler run, or crawl what S3 notifies the crawler of via SQS. Valid Values are: CRAWL_EVENT_MODE, CRAWL_EVERYTHING and CRAWL_NEW_FOLDERS_ONLY. Default value is CRAWL_EVERYTHING.
	RecrawlBehavior *string `json:"recrawlBehavior,omitempty" tf:"recrawl_behavior,omitempty"`
}

type RecrawlPolicyParameters struct {

	// Specifies whether to crawl the entire dataset again, crawl only folders that were added since the last crawler run, or crawl what S3 notifies the crawler of via SQS. Valid Values are: CRAWL_EVENT_MODE, CRAWL_EVERYTHING and CRAWL_NEW_FOLDERS_ONLY. Default value is CRAWL_EVERYTHING.
	// +kubebuilder:validation:Optional
	RecrawlBehavior *string `json:"recrawlBehavior,omitempty" tf:"recrawl_behavior,omitempty"`
}

type S3TargetInitParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// The ARN of the dead-letter SQS queue.
	DlqEventQueueArn *string `json:"dlqEventQueueArn,omitempty" tf:"dlq_event_queue_arn,omitempty"`

	// The ARN of the SQS queue to receive S3 notifications from.
	EventQueueArn *string `json:"eventQueueArn,omitempty" tf:"event_queue_arn,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	SampleSize *float64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`
}

type S3TargetObservation struct {

	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// The ARN of the dead-letter SQS queue.
	DlqEventQueueArn *string `json:"dlqEventQueueArn,omitempty" tf:"dlq_event_queue_arn,omitempty"`

	// The ARN of the SQS queue to receive S3 notifications from.
	EventQueueArn *string `json:"eventQueueArn,omitempty" tf:"event_queue_arn,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The name of the DynamoDB table to crawl.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	SampleSize *float64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`
}

type S3TargetParameters struct {

	// The name of the connection to use to connect to the JDBC target.
	// +kubebuilder:validation:Optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// The ARN of the dead-letter SQS queue.
	// +kubebuilder:validation:Optional
	DlqEventQueueArn *string `json:"dlqEventQueueArn,omitempty" tf:"dlq_event_queue_arn,omitempty"`

	// The ARN of the SQS queue to receive S3 notifications from.
	// +kubebuilder:validation:Optional
	EventQueueArn *string `json:"eventQueueArn,omitempty" tf:"event_queue_arn,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	// +kubebuilder:validation:Optional
	Exclusions []*string `json:"exclusions,omitempty" tf:"exclusions,omitempty"`

	// The name of the DynamoDB table to crawl.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	// +kubebuilder:validation:Optional
	SampleSize *float64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`
}

type SchemaChangePolicyInitParameters struct {

	// The deletion behavior when the crawler finds a deleted object. Valid values: LOG, DELETE_FROM_DATABASE, or DEPRECATE_IN_DATABASE. Defaults to DEPRECATE_IN_DATABASE.
	DeleteBehavior *string `json:"deleteBehavior,omitempty" tf:"delete_behavior,omitempty"`

	// The update behavior when the crawler finds a changed schema. Valid values: LOG or UPDATE_IN_DATABASE. Defaults to UPDATE_IN_DATABASE.
	UpdateBehavior *string `json:"updateBehavior,omitempty" tf:"update_behavior,omitempty"`
}

type SchemaChangePolicyObservation struct {

	// The deletion behavior when the crawler finds a deleted object. Valid values: LOG, DELETE_FROM_DATABASE, or DEPRECATE_IN_DATABASE. Defaults to DEPRECATE_IN_DATABASE.
	DeleteBehavior *string `json:"deleteBehavior,omitempty" tf:"delete_behavior,omitempty"`

	// The update behavior when the crawler finds a changed schema. Valid values: LOG or UPDATE_IN_DATABASE. Defaults to UPDATE_IN_DATABASE.
	UpdateBehavior *string `json:"updateBehavior,omitempty" tf:"update_behavior,omitempty"`
}

type SchemaChangePolicyParameters struct {

	// The deletion behavior when the crawler finds a deleted object. Valid values: LOG, DELETE_FROM_DATABASE, or DEPRECATE_IN_DATABASE. Defaults to DEPRECATE_IN_DATABASE.
	// +kubebuilder:validation:Optional
	DeleteBehavior *string `json:"deleteBehavior,omitempty" tf:"delete_behavior,omitempty"`

	// The update behavior when the crawler finds a changed schema. Valid values: LOG or UPDATE_IN_DATABASE. Defaults to UPDATE_IN_DATABASE.
	// +kubebuilder:validation:Optional
	UpdateBehavior *string `json:"updateBehavior,omitempty" tf:"update_behavior,omitempty"`
}

// CrawlerSpec defines the desired state of Crawler
type CrawlerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CrawlerParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CrawlerInitParameters `json:"initProvider,omitempty"`
}

// CrawlerStatus defines the observed state of Crawler.
type CrawlerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CrawlerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Crawler is the Schema for the Crawlers API. Manages a Glue Crawler
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Crawler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CrawlerSpec   `json:"spec"`
	Status            CrawlerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CrawlerList contains a list of Crawlers
type CrawlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Crawler `json:"items"`
}

// Repository type metadata.
var (
	Crawler_Kind             = "Crawler"
	Crawler_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Crawler_Kind}.String()
	Crawler_KindAPIVersion   = Crawler_Kind + "." + CRDGroupVersion.String()
	Crawler_GroupVersionKind = CRDGroupVersion.WithKind(Crawler_Kind)
)

func init() {
	SchemeBuilder.Register(&Crawler{}, &CrawlerList{})
}
