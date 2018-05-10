// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Cluster, InType: reflect.TypeOf(&Cluster{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ClusterCondition, InType: reflect.TypeOf(&ClusterCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ClusterList, InType: reflect.TypeOf(&ClusterList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ClusterSpec, InType: reflect.TypeOf(&ClusterSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ClusterStatus, InType: reflect.TypeOf(&ClusterStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ServerAddressByClientCIDR, InType: reflect.TypeOf(&ServerAddressByClientCIDR{})},
	)
}

func DeepCopy_v1beta1_Cluster(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Cluster)
		out := out.(*Cluster)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_ClusterSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_ClusterStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_ClusterCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterCondition)
		out := out.(*ClusterCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_v1beta1_ClusterList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterList)
		out := out.(*ClusterList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Cluster, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_Cluster(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ClusterSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterSpec)
		out := out.(*ClusterSpec)
		if in.ServerAddressByClientCIDRs != nil {
			in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
			*out = make([]ServerAddressByClientCIDR, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.ServerAddressByClientCIDRs = nil
		}
		if in.SecretRef != nil {
			in, out := &in.SecretRef, &out.SecretRef
			*out = new(v1.LocalObjectReference)
			**out = **in
		} else {
			out.SecretRef = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ClusterStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterStatus)
		out := out.(*ClusterStatus)
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ClusterCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ClusterCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		if in.Zones != nil {
			in, out := &in.Zones, &out.Zones
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Zones = nil
		}
		out.Region = in.Region
		return nil
	}
}

func DeepCopy_v1beta1_ServerAddressByClientCIDR(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServerAddressByClientCIDR)
		out := out.(*ServerAddressByClientCIDR)
		out.ClientCIDR = in.ClientCIDR
		out.ServerAddress = in.ServerAddress
		return nil
	}
}
