//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalMetadataSpec) DeepCopyInto(out *AdditionalMetadataSpec) {
	*out = *in
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalAnnotations != nil {
		in, out := &in.AdditionalAnnotations, &out.AdditionalAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalMetadataSpec.
func (in *AdditionalMetadataSpec) DeepCopy() *AdditionalMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(AdditionalMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalRoleBindingsSpec) DeepCopyInto(out *AdditionalRoleBindingsSpec) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]v1.Subject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalRoleBindingsSpec.
func (in *AdditionalRoleBindingsSpec) DeepCopy() *AdditionalRoleBindingsSpec {
	if in == nil {
		return nil
	}
	out := new(AdditionalRoleBindingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedListSpec) DeepCopyInto(out *AllowedListSpec) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedListSpec.
func (in *AllowedListSpec) DeepCopy() *AllowedListSpec {
	if in == nil {
		return nil
	}
	out := new(AllowedListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapsuleConfiguration) DeepCopyInto(out *CapsuleConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapsuleConfiguration.
func (in *CapsuleConfiguration) DeepCopy() *CapsuleConfiguration {
	if in == nil {
		return nil
	}
	out := new(CapsuleConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CapsuleConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapsuleConfigurationList) DeepCopyInto(out *CapsuleConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CapsuleConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapsuleConfigurationList.
func (in *CapsuleConfigurationList) DeepCopy() *CapsuleConfigurationList {
	if in == nil {
		return nil
	}
	out := new(CapsuleConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CapsuleConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapsuleConfigurationSpec) DeepCopyInto(out *CapsuleConfigurationSpec) {
	*out = *in
	if in.UserGroups != nil {
		in, out := &in.UserGroups, &out.UserGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapsuleConfigurationSpec.
func (in *CapsuleConfigurationSpec) DeepCopy() *CapsuleConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(CapsuleConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalServiceIPsSpec) DeepCopyInto(out *ExternalServiceIPsSpec) {
	*out = *in
	if in.Allowed != nil {
		in, out := &in.Allowed, &out.Allowed
		*out = make([]AllowedIP, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalServiceIPsSpec.
func (in *ExternalServiceIPsSpec) DeepCopy() *ExternalServiceIPsSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalServiceIPsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerSpec) DeepCopyInto(out *OwnerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerSpec.
func (in *OwnerSpec) DeepCopy() *OwnerSpec {
	if in == nil {
		return nil
	}
	out := new(OwnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenant) DeepCopyInto(out *Tenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenant.
func (in *Tenant) DeepCopy() *Tenant {
	if in == nil {
		return nil
	}
	out := new(Tenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantList) DeepCopyInto(out *TenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantList.
func (in *TenantList) DeepCopy() *TenantList {
	if in == nil {
		return nil
	}
	out := new(TenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantSpec) DeepCopyInto(out *TenantSpec) {
	*out = *in
	out.Owner = in.Owner
	if in.NamespaceQuota != nil {
		in, out := &in.NamespaceQuota, &out.NamespaceQuota
		*out = new(int32)
		**out = **in
	}
	if in.NamespacesMetadata != nil {
		in, out := &in.NamespacesMetadata, &out.NamespacesMetadata
		*out = new(AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ServicesMetadata != nil {
		in, out := &in.ServicesMetadata, &out.ServicesMetadata
		*out = new(AdditionalMetadataSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IngressClasses != nil {
		in, out := &in.IngressClasses, &out.IngressClasses
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IngressHostnames != nil {
		in, out := &in.IngressHostnames, &out.IngressHostnames
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerRegistries != nil {
		in, out := &in.ContainerRegistries, &out.ContainerRegistries
		*out = new(AllowedListSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkPolicies != nil {
		in, out := &in.NetworkPolicies, &out.NetworkPolicies
		*out = make([]networkingv1.NetworkPolicySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LimitRanges != nil {
		in, out := &in.LimitRanges, &out.LimitRanges
		*out = make([]corev1.LimitRangeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceQuota != nil {
		in, out := &in.ResourceQuota, &out.ResourceQuota
		*out = make([]corev1.ResourceQuotaSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalRoleBindings != nil {
		in, out := &in.AdditionalRoleBindings, &out.AdditionalRoleBindings
		*out = make([]AdditionalRoleBindingsSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExternalServiceIPs != nil {
		in, out := &in.ExternalServiceIPs, &out.ExternalServiceIPs
		*out = new(ExternalServiceIPsSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantSpec.
func (in *TenantSpec) DeepCopy() *TenantSpec {
	if in == nil {
		return nil
	}
	out := new(TenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantStatus) DeepCopyInto(out *TenantStatus) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantStatus.
func (in *TenantStatus) DeepCopy() *TenantStatus {
	if in == nil {
		return nil
	}
	out := new(TenantStatus)
	in.DeepCopyInto(out)
	return out
}
