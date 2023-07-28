//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventFilterInitParameters) DeepCopyInto(out *EventFilterInitParameters) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventFilterInitParameters.
func (in *EventFilterInitParameters) DeepCopy() *EventFilterInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventFilterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventFilterObservation) DeepCopyInto(out *EventFilterObservation) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventFilterObservation.
func (in *EventFilterObservation) DeepCopy() *EventFilterObservation {
	if in == nil {
		return nil
	}
	out := new(EventFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventFilterParameters) DeepCopyInto(out *EventFilterParameters) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventFilterParameters.
func (in *EventFilterParameters) DeepCopy() *EventFilterParameters {
	if in == nil {
		return nil
	}
	out := new(EventFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventIntegration) DeepCopyInto(out *EventIntegration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventIntegration.
func (in *EventIntegration) DeepCopy() *EventIntegration {
	if in == nil {
		return nil
	}
	out := new(EventIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventIntegration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventIntegrationInitParameters) DeepCopyInto(out *EventIntegrationInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EventFilter != nil {
		in, out := &in.EventFilter, &out.EventFilter
		*out = make([]EventFilterInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventbridgeBus != nil {
		in, out := &in.EventbridgeBus, &out.EventbridgeBus
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventIntegrationInitParameters.
func (in *EventIntegrationInitParameters) DeepCopy() *EventIntegrationInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventIntegrationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventIntegrationList) DeepCopyInto(out *EventIntegrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventIntegration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventIntegrationList.
func (in *EventIntegrationList) DeepCopy() *EventIntegrationList {
	if in == nil {
		return nil
	}
	out := new(EventIntegrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventIntegrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventIntegrationObservation) DeepCopyInto(out *EventIntegrationObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EventFilter != nil {
		in, out := &in.EventFilter, &out.EventFilter
		*out = make([]EventFilterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventbridgeBus != nil {
		in, out := &in.EventbridgeBus, &out.EventbridgeBus
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventIntegrationObservation.
func (in *EventIntegrationObservation) DeepCopy() *EventIntegrationObservation {
	if in == nil {
		return nil
	}
	out := new(EventIntegrationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventIntegrationParameters) DeepCopyInto(out *EventIntegrationParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EventFilter != nil {
		in, out := &in.EventFilter, &out.EventFilter
		*out = make([]EventFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventbridgeBus != nil {
		in, out := &in.EventbridgeBus, &out.EventbridgeBus
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventIntegrationParameters.
func (in *EventIntegrationParameters) DeepCopy() *EventIntegrationParameters {
	if in == nil {
		return nil
	}
	out := new(EventIntegrationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventIntegrationSpec) DeepCopyInto(out *EventIntegrationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventIntegrationSpec.
func (in *EventIntegrationSpec) DeepCopy() *EventIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(EventIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventIntegrationStatus) DeepCopyInto(out *EventIntegrationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventIntegrationStatus.
func (in *EventIntegrationStatus) DeepCopy() *EventIntegrationStatus {
	if in == nil {
		return nil
	}
	out := new(EventIntegrationStatus)
	in.DeepCopyInto(out)
	return out
}
