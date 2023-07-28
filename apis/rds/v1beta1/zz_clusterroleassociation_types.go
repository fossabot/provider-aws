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

type ClusterRoleAssociationInitParameters struct {

	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the SupportedFeatureNames list returned by AWS CLI rds describe-db-engine-versions.
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name,omitempty"`
}

type ClusterRoleAssociationObservation struct {

	// DB Cluster Identifier to associate with the IAM Role.
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the SupportedFeatureNames list returned by AWS CLI rds describe-db-engine-versions.
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name,omitempty"`

	// DB Cluster Identifier and IAM Role ARN separated by a comma (,)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Cluster.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type ClusterRoleAssociationParameters struct {

	// DB Cluster Identifier to associate with the IAM Role.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty" tf:"db_cluster_identifier,omitempty"`

	// Reference to a Cluster in rds to populate dbClusterIdentifier.
	// +kubebuilder:validation:Optional
	DBClusterIdentifierRef *v1.Reference `json:"dbClusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in rds to populate dbClusterIdentifier.
	// +kubebuilder:validation:Optional
	DBClusterIdentifierSelector *v1.Selector `json:"dbClusterIdentifierSelector,omitempty" tf:"-"`

	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the SupportedFeatureNames list returned by AWS CLI rds describe-db-engine-versions.
	// +kubebuilder:validation:Optional
	FeatureName *string `json:"featureName,omitempty" tf:"feature_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// ClusterRoleAssociationSpec defines the desired state of ClusterRoleAssociation
type ClusterRoleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterRoleAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ClusterRoleAssociationInitParameters `json:"initProvider,omitempty"`
}

// ClusterRoleAssociationStatus defines the observed state of ClusterRoleAssociation.
type ClusterRoleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterRoleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterRoleAssociation is the Schema for the ClusterRoleAssociations API. Manages a RDS DB Cluster association with an IAM Role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterRoleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.featureName) || has(self.initProvider.featureName)",message="featureName is a required parameter"
	Spec   ClusterRoleAssociationSpec   `json:"spec"`
	Status ClusterRoleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterRoleAssociationList contains a list of ClusterRoleAssociations
type ClusterRoleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRoleAssociation `json:"items"`
}

// Repository type metadata.
var (
	ClusterRoleAssociation_Kind             = "ClusterRoleAssociation"
	ClusterRoleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterRoleAssociation_Kind}.String()
	ClusterRoleAssociation_KindAPIVersion   = ClusterRoleAssociation_Kind + "." + CRDGroupVersion.String()
	ClusterRoleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ClusterRoleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterRoleAssociation{}, &ClusterRoleAssociationList{})
}
