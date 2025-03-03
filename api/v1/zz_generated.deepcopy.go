//go:build !ignore_autogenerated

/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPoint) DeepCopyInto(out *ContactPoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPoint.
func (in *ContactPoint) DeepCopy() *ContactPoint {
	if in == nil {
		return nil
	}
	out := new(ContactPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContactPoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPointApiToken) DeepCopyInto(out *ContactPointApiToken) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPointApiToken.
func (in *ContactPointApiToken) DeepCopy() *ContactPointApiToken {
	if in == nil {
		return nil
	}
	out := new(ContactPointApiToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPointList) DeepCopyInto(out *ContactPointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContactPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPointList.
func (in *ContactPointList) DeepCopy() *ContactPointList {
	if in == nil {
		return nil
	}
	out := new(ContactPointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContactPointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPointSpec) DeepCopyInto(out *ContactPointSpec) {
	*out = *in
	out.ApiToken = in.ApiToken
	out.WebhookSpec = in.WebhookSpec
	out.TelegramSpec = in.TelegramSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPointSpec.
func (in *ContactPointSpec) DeepCopy() *ContactPointSpec {
	if in == nil {
		return nil
	}
	out := new(ContactPointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPointStatus) DeepCopyInto(out *ContactPointStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPointStatus.
func (in *ContactPointStatus) DeepCopy() *ContactPointStatus {
	if in == nil {
		return nil
	}
	out := new(ContactPointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPointTelegramSpec) DeepCopyInto(out *ContactPointTelegramSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPointTelegramSpec.
func (in *ContactPointTelegramSpec) DeepCopy() *ContactPointTelegramSpec {
	if in == nil {
		return nil
	}
	out := new(ContactPointTelegramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactPointWebhookSpec) DeepCopyInto(out *ContactPointWebhookSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactPointWebhookSpec.
func (in *ContactPointWebhookSpec) DeepCopy() *ContactPointWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(ContactPointWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackedField) DeepCopyInto(out *TrackedField) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackedField.
func (in *TrackedField) DeepCopy() *TrackedField {
	if in == nil {
		return nil
	}
	out := new(TrackedField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrackedField) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackedFieldList) DeepCopyInto(out *TrackedFieldList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrackedField, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackedFieldList.
func (in *TrackedFieldList) DeepCopy() *TrackedFieldList {
	if in == nil {
		return nil
	}
	out := new(TrackedFieldList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrackedFieldList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackedFieldSpec) DeepCopyInto(out *TrackedFieldSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackedFieldSpec.
func (in *TrackedFieldSpec) DeepCopy() *TrackedFieldSpec {
	if in == nil {
		return nil
	}
	out := new(TrackedFieldSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrackedFieldStatus) DeepCopyInto(out *TrackedFieldStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrackedFieldStatus.
func (in *TrackedFieldStatus) DeepCopy() *TrackedFieldStatus {
	if in == nil {
		return nil
	}
	out := new(TrackedFieldStatus)
	in.DeepCopyInto(out)
	return out
}
