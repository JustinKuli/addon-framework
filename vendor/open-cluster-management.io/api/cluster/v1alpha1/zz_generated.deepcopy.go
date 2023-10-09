//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScore) DeepCopyInto(out *AddOnPlacementScore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScore.
func (in *AddOnPlacementScore) DeepCopy() *AddOnPlacementScore {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnPlacementScore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScoreItem) DeepCopyInto(out *AddOnPlacementScoreItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScoreItem.
func (in *AddOnPlacementScoreItem) DeepCopy() *AddOnPlacementScoreItem {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScoreItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScoreList) DeepCopyInto(out *AddOnPlacementScoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddOnPlacementScore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScoreList.
func (in *AddOnPlacementScoreList) DeepCopy() *AddOnPlacementScoreList {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddOnPlacementScoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnPlacementScoreStatus) DeepCopyInto(out *AddOnPlacementScoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scores != nil {
		in, out := &in.Scores, &out.Scores
		*out = make([]AddOnPlacementScoreItem, len(*in))
		copy(*out, *in)
	}
	if in.ValidUntil != nil {
		in, out := &in.ValidUntil, &out.ValidUntil
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnPlacementScoreStatus.
func (in *AddOnPlacementScoreStatus) DeepCopy() *AddOnPlacementScoreStatus {
	if in == nil {
		return nil
	}
	out := new(AddOnPlacementScoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaim) DeepCopyInto(out *ClusterClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaim.
func (in *ClusterClaim) DeepCopy() *ClusterClaim {
	if in == nil {
		return nil
	}
	out := new(ClusterClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaimList) DeepCopyInto(out *ClusterClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaimList.
func (in *ClusterClaimList) DeepCopy() *ClusterClaimList {
	if in == nil {
		return nil
	}
	out := new(ClusterClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClaimSpec) DeepCopyInto(out *ClusterClaimSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClaimSpec.
func (in *ClusterClaimSpec) DeepCopy() *ClusterClaimSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRolloutStatus) DeepCopyInto(out *ClusterRolloutStatus) {
	*out = *in
	out.GroupKey = in.GroupKey
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.TimeOutTime != nil {
		in, out := &in.TimeOutTime, &out.TimeOutTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRolloutStatus.
func (in *ClusterRolloutStatus) DeepCopy() *ClusterRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MandatoryDecisionGroup) DeepCopyInto(out *MandatoryDecisionGroup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MandatoryDecisionGroup.
func (in *MandatoryDecisionGroup) DeepCopy() *MandatoryDecisionGroup {
	if in == nil {
		return nil
	}
	out := new(MandatoryDecisionGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MandatoryDecisionGroups) DeepCopyInto(out *MandatoryDecisionGroups) {
	*out = *in
	if in.MandatoryDecisionGroups != nil {
		in, out := &in.MandatoryDecisionGroups, &out.MandatoryDecisionGroups
		*out = make([]MandatoryDecisionGroup, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MandatoryDecisionGroups.
func (in *MandatoryDecisionGroups) DeepCopy() *MandatoryDecisionGroups {
	if in == nil {
		return nil
	}
	out := new(MandatoryDecisionGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutAll) DeepCopyInto(out *RolloutAll) {
	*out = *in
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutAll.
func (in *RolloutAll) DeepCopy() *RolloutAll {
	if in == nil {
		return nil
	}
	out := new(RolloutAll)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutProgressive) DeepCopyInto(out *RolloutProgressive) {
	*out = *in
	in.MandatoryDecisionGroups.DeepCopyInto(&out.MandatoryDecisionGroups)
	out.MaxConcurrency = in.MaxConcurrency
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutProgressive.
func (in *RolloutProgressive) DeepCopy() *RolloutProgressive {
	if in == nil {
		return nil
	}
	out := new(RolloutProgressive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutProgressivePerGroup) DeepCopyInto(out *RolloutProgressivePerGroup) {
	*out = *in
	in.MandatoryDecisionGroups.DeepCopyInto(&out.MandatoryDecisionGroups)
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutProgressivePerGroup.
func (in *RolloutProgressivePerGroup) DeepCopy() *RolloutProgressivePerGroup {
	if in == nil {
		return nil
	}
	out := new(RolloutProgressivePerGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutResult) DeepCopyInto(out *RolloutResult) {
	*out = *in
	if in.ClustersToRollout != nil {
		in, out := &in.ClustersToRollout, &out.ClustersToRollout
		*out = make([]ClusterRolloutStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ClustersTimeOut != nil {
		in, out := &in.ClustersTimeOut, &out.ClustersTimeOut
		*out = make([]ClusterRolloutStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutResult.
func (in *RolloutResult) DeepCopy() *RolloutResult {
	if in == nil {
		return nil
	}
	out := new(RolloutResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStrategy) DeepCopyInto(out *RolloutStrategy) {
	*out = *in
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = new(RolloutAll)
		**out = **in
	}
	if in.Progressive != nil {
		in, out := &in.Progressive, &out.Progressive
		*out = new(RolloutProgressive)
		(*in).DeepCopyInto(*out)
	}
	if in.ProgressivePerGroup != nil {
		in, out := &in.ProgressivePerGroup, &out.ProgressivePerGroup
		*out = new(RolloutProgressivePerGroup)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStrategy.
func (in *RolloutStrategy) DeepCopy() *RolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(RolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeout) DeepCopyInto(out *Timeout) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeout.
func (in *Timeout) DeepCopy() *Timeout {
	if in == nil {
		return nil
	}
	out := new(Timeout)
	in.DeepCopyInto(out)
	return out
}
