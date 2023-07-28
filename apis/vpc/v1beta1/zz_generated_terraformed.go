/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this NetworkPerformanceMetricSubscription
func (mg *NetworkPerformanceMetricSubscription) GetTerraformResourceType() string {
	return "aws_vpc_network_performance_metric_subscription"
}

// GetConnectionDetailsMapping for this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetInitParameters for this NetworkPerformanceMetricSubscription
func (tr *NetworkPerformanceMetricSubscription) SetInitParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.InitProvider)
}

// LateInitialize this NetworkPerformanceMetricSubscription using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NetworkPerformanceMetricSubscription) LateInitialize(attrs []byte) (bool, error) {
	params := &NetworkPerformanceMetricSubscriptionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NetworkPerformanceMetricSubscription) GetTerraformSchemaVersion() int {
	return 0
}
