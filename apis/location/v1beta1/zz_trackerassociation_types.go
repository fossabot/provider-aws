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

type TrackerAssociationInitParameters struct {
}

type TrackerAssociationObservation struct {

	// The Amazon Resource Name (ARN) for the geofence collection to be associated to tracker resource. Used when you need to specify a resource across all AWS.
	ConsumerArn *string `json:"consumerArn,omitempty" tf:"consumer_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the tracker resource to be associated with a geofence collection.
	TrackerName *string `json:"trackerName,omitempty" tf:"tracker_name,omitempty"`
}

type TrackerAssociationParameters struct {

	// The Amazon Resource Name (ARN) for the geofence collection to be associated to tracker resource. Used when you need to specify a resource across all AWS.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/location/v1beta1.GeofenceCollection
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("collection_arn",true)
	// +kubebuilder:validation:Optional
	ConsumerArn *string `json:"consumerArn,omitempty" tf:"consumer_arn,omitempty"`

	// Reference to a GeofenceCollection in location to populate consumerArn.
	// +kubebuilder:validation:Optional
	ConsumerArnRef *v1.Reference `json:"consumerArnRef,omitempty" tf:"-"`

	// Selector for a GeofenceCollection in location to populate consumerArn.
	// +kubebuilder:validation:Optional
	ConsumerArnSelector *v1.Selector `json:"consumerArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the tracker resource to be associated with a geofence collection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/location/v1beta1.Tracker
	// +kubebuilder:validation:Optional
	TrackerName *string `json:"trackerName,omitempty" tf:"tracker_name,omitempty"`

	// Reference to a Tracker in location to populate trackerName.
	// +kubebuilder:validation:Optional
	TrackerNameRef *v1.Reference `json:"trackerNameRef,omitempty" tf:"-"`

	// Selector for a Tracker in location to populate trackerName.
	// +kubebuilder:validation:Optional
	TrackerNameSelector *v1.Selector `json:"trackerNameSelector,omitempty" tf:"-"`
}

// TrackerAssociationSpec defines the desired state of TrackerAssociation
type TrackerAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrackerAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider TrackerAssociationInitParameters `json:"initProvider,omitempty"`
}

// TrackerAssociationStatus defines the observed state of TrackerAssociation.
type TrackerAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrackerAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrackerAssociation is the Schema for the TrackerAssociations API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TrackerAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrackerAssociationSpec   `json:"spec"`
	Status            TrackerAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrackerAssociationList contains a list of TrackerAssociations
type TrackerAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrackerAssociation `json:"items"`
}

// Repository type metadata.
var (
	TrackerAssociation_Kind             = "TrackerAssociation"
	TrackerAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrackerAssociation_Kind}.String()
	TrackerAssociation_KindAPIVersion   = TrackerAssociation_Kind + "." + CRDGroupVersion.String()
	TrackerAssociation_GroupVersionKind = CRDGroupVersion.WithKind(TrackerAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&TrackerAssociation{}, &TrackerAssociationList{})
}
