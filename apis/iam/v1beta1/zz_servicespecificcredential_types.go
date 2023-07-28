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

type ServiceSpecificCredentialInitParameters struct {

	// The name of the AWS service that is to be associated with the credentials. The service you specify here is the only service that can be accessed using these credentials.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The status to be assigned to the service-specific credential. Valid values are Active and Inactive. Default value is Active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ServiceSpecificCredentialObservation struct {

	// The combination of service_name and user_name as such: service_name:user_name:service_specific_credential_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the AWS service that is to be associated with the credentials. The service you specify here is the only service that can be accessed using these credentials.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The unique identifier for the service-specific credential.
	ServiceSpecificCredentialID *string `json:"serviceSpecificCredentialId,omitempty" tf:"service_specific_credential_id,omitempty"`

	// The generated user name for the service-specific credential. This value is generated by combining the IAM user's name combined with the ID number of the AWS account, as in jane-at-123456789012, for example.
	ServiceUserName *string `json:"serviceUserName,omitempty" tf:"service_user_name,omitempty"`

	// The status to be assigned to the service-specific credential. Valid values are Active and Inactive. Default value is Active.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the IAM user that is to be associated with the credentials. The new service-specific credentials have the same permissions as the associated user except that they can be used only to access the specified service.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type ServiceSpecificCredentialParameters struct {

	// The name of the AWS service that is to be associated with the credentials. The service you specify here is the only service that can be accessed using these credentials.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The status to be assigned to the service-specific credential. Valid values are Active and Inactive. Default value is Active.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the IAM user that is to be associated with the credentials. The new service-specific credentials have the same permissions as the associated user except that they can be used only to access the specified service.
	// +crossplane:generate:reference:type=User
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Reference to a User to populate userName.
	// +kubebuilder:validation:Optional
	UserNameRef *v1.Reference `json:"userNameRef,omitempty" tf:"-"`

	// Selector for a User to populate userName.
	// +kubebuilder:validation:Optional
	UserNameSelector *v1.Selector `json:"userNameSelector,omitempty" tf:"-"`
}

// ServiceSpecificCredentialSpec defines the desired state of ServiceSpecificCredential
type ServiceSpecificCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceSpecificCredentialParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ServiceSpecificCredentialInitParameters `json:"initProvider,omitempty"`
}

// ServiceSpecificCredentialStatus defines the observed state of ServiceSpecificCredential.
type ServiceSpecificCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceSpecificCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceSpecificCredential is the Schema for the ServiceSpecificCredentials API. Provides an IAM Service Specific Credential.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServiceSpecificCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || has(self.initProvider.serviceName)",message="serviceName is a required parameter"
	Spec   ServiceSpecificCredentialSpec   `json:"spec"`
	Status ServiceSpecificCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceSpecificCredentialList contains a list of ServiceSpecificCredentials
type ServiceSpecificCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceSpecificCredential `json:"items"`
}

// Repository type metadata.
var (
	ServiceSpecificCredential_Kind             = "ServiceSpecificCredential"
	ServiceSpecificCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceSpecificCredential_Kind}.String()
	ServiceSpecificCredential_KindAPIVersion   = ServiceSpecificCredential_Kind + "." + CRDGroupVersion.String()
	ServiceSpecificCredential_GroupVersionKind = CRDGroupVersion.WithKind(ServiceSpecificCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceSpecificCredential{}, &ServiceSpecificCredentialList{})
}
