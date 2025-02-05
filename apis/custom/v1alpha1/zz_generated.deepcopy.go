// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionsObservation) DeepCopyInto(out *ActionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionsObservation.
func (in *ActionsObservation) DeepCopy() *ActionsObservation {
	if in == nil {
		return nil
	}
	out := new(ActionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionsParameters) DeepCopyInto(out *ActionsParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionsParameters.
func (in *ActionsParameters) DeepCopy() *ActionsParameters {
	if in == nil {
		return nil
	}
	out := new(ActionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBRole) DeepCopyInto(out *DBRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBRole.
func (in *DBRole) DeepCopy() *DBRole {
	if in == nil {
		return nil
	}
	out := new(DBRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBRoleList) DeepCopyInto(out *DBRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DBRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBRoleList.
func (in *DBRoleList) DeepCopy() *DBRoleList {
	if in == nil {
		return nil
	}
	out := new(DBRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DBRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBRoleObservation) DeepCopyInto(out *DBRoleObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBRoleObservation.
func (in *DBRoleObservation) DeepCopy() *DBRoleObservation {
	if in == nil {
		return nil
	}
	out := new(DBRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBRoleParameters) DeepCopyInto(out *DBRoleParameters) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]ActionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InheritedRoles != nil {
		in, out := &in.InheritedRoles, &out.InheritedRoles
		*out = make([]InheritedRolesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBRoleParameters.
func (in *DBRoleParameters) DeepCopy() *DBRoleParameters {
	if in == nil {
		return nil
	}
	out := new(DBRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBRoleSpec) DeepCopyInto(out *DBRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBRoleSpec.
func (in *DBRoleSpec) DeepCopy() *DBRoleSpec {
	if in == nil {
		return nil
	}
	out := new(DBRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBRoleStatus) DeepCopyInto(out *DBRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBRoleStatus.
func (in *DBRoleStatus) DeepCopy() *DBRoleStatus {
	if in == nil {
		return nil
	}
	out := new(DBRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigurationClusterAws) DeepCopyInto(out *DNSConfigurationClusterAws) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigurationClusterAws.
func (in *DNSConfigurationClusterAws) DeepCopy() *DNSConfigurationClusterAws {
	if in == nil {
		return nil
	}
	out := new(DNSConfigurationClusterAws)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSConfigurationClusterAws) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigurationClusterAwsList) DeepCopyInto(out *DNSConfigurationClusterAwsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSConfigurationClusterAws, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigurationClusterAwsList.
func (in *DNSConfigurationClusterAwsList) DeepCopy() *DNSConfigurationClusterAwsList {
	if in == nil {
		return nil
	}
	out := new(DNSConfigurationClusterAwsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSConfigurationClusterAwsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigurationClusterAwsObservation) DeepCopyInto(out *DNSConfigurationClusterAwsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigurationClusterAwsObservation.
func (in *DNSConfigurationClusterAwsObservation) DeepCopy() *DNSConfigurationClusterAwsObservation {
	if in == nil {
		return nil
	}
	out := new(DNSConfigurationClusterAwsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigurationClusterAwsParameters) DeepCopyInto(out *DNSConfigurationClusterAwsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigurationClusterAwsParameters.
func (in *DNSConfigurationClusterAwsParameters) DeepCopy() *DNSConfigurationClusterAwsParameters {
	if in == nil {
		return nil
	}
	out := new(DNSConfigurationClusterAwsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigurationClusterAwsSpec) DeepCopyInto(out *DNSConfigurationClusterAwsSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigurationClusterAwsSpec.
func (in *DNSConfigurationClusterAwsSpec) DeepCopy() *DNSConfigurationClusterAwsSpec {
	if in == nil {
		return nil
	}
	out := new(DNSConfigurationClusterAwsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigurationClusterAwsStatus) DeepCopyInto(out *DNSConfigurationClusterAwsStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigurationClusterAwsStatus.
func (in *DNSConfigurationClusterAwsStatus) DeepCopy() *DNSConfigurationClusterAwsStatus {
	if in == nil {
		return nil
	}
	out := new(DNSConfigurationClusterAwsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InheritedRolesObservation) DeepCopyInto(out *InheritedRolesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InheritedRolesObservation.
func (in *InheritedRolesObservation) DeepCopy() *InheritedRolesObservation {
	if in == nil {
		return nil
	}
	out := new(InheritedRolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InheritedRolesParameters) DeepCopyInto(out *InheritedRolesParameters) {
	*out = *in
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InheritedRolesParameters.
func (in *InheritedRolesParameters) DeepCopy() *InheritedRolesParameters {
	if in == nil {
		return nil
	}
	out := new(InheritedRolesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesObservation) DeepCopyInto(out *ResourcesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesObservation.
func (in *ResourcesObservation) DeepCopy() *ResourcesObservation {
	if in == nil {
		return nil
	}
	out := new(ResourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesParameters) DeepCopyInto(out *ResourcesParameters) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(bool)
		**out = **in
	}
	if in.CollectionName != nil {
		in, out := &in.CollectionName, &out.CollectionName
		*out = new(string)
		**out = **in
	}
	if in.DatabaseName != nil {
		in, out := &in.DatabaseName, &out.DatabaseName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesParameters.
func (in *ResourcesParameters) DeepCopy() *ResourcesParameters {
	if in == nil {
		return nil
	}
	out := new(ResourcesParameters)
	in.DeepCopyInto(out)
	return out
}
