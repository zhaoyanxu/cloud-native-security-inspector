//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cache) DeepCopyInto(out *Cache) {
	*out = *in
	if in.CredentialRef != nil {
		in, out := &in.CredentialRef, &out.CredentialRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(int)
		**out = **in
	}
	in.Settings.DeepCopyInto(&out.Settings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cache.
func (in *Cache) DeepCopy() *Cache {
	if in == nil {
		return nil
	}
	out := new(Cache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSettings) DeepCopyInto(out *CacheSettings) {
	*out = *in
	if in.SkipTLSVerify != nil {
		in, out := &in.SkipTLSVerify, &out.SkipTLSVerify
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSettings.
func (in *CacheSettings) DeepCopy() *CacheSettings {
	if in == nil {
		return nil
	}
	out := new(CacheSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceBaseline) DeepCopyInto(out *ComplianceBaseline) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceBaseline.
func (in *ComplianceBaseline) DeepCopy() *ComplianceBaseline {
	if in == nil {
		return nil
	}
	out := new(ComplianceBaseline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProvider) DeepCopyInto(out *DataProvider) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(Credential)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(Cache)
		(*in).DeepCopyInto(*out)
	}
	out.Connection = in.Connection
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProvider.
func (in *DataProvider) DeepCopy() *DataProvider {
	if in == nil {
		return nil
	}
	out := new(DataProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSource) DeepCopyInto(out *DataSource) {
	*out = *in
	in.Registry.DeepCopyInto(&out.Registry)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSource.
func (in *DataSource) DeepCopy() *DataSource {
	if in == nil {
		return nil
	}
	out := new(DataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportConfig) DeepCopyInto(out *ExportConfig) {
	*out = *in
	out.OpenSearch = in.OpenSearch
	out.Governor = in.Governor
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportConfig.
func (in *ExportConfig) DeepCopy() *ExportConfig {
	if in == nil {
		return nil
	}
	out := new(ExportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FollowupAction) DeepCopyInto(out *FollowupAction) {
	*out = *in
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FollowupAction.
func (in *FollowupAction) DeepCopy() *FollowupAction {
	if in == nil {
		return nil
	}
	out := new(FollowupAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GovernorOutputConfig) DeepCopyInto(out *GovernorOutputConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GovernorOutputConfig.
func (in *GovernorOutputConfig) DeepCopy() *GovernorOutputConfig {
	if in == nil {
		return nil
	}
	out := new(GovernorOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionConfiguration) DeepCopyInto(out *InspectionConfiguration) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*FollowupAction, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FollowupAction)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Baselines != nil {
		in, out := &in.Baselines, &out.Baselines
		*out = make([]*ComplianceBaseline, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ComplianceBaseline)
				**out = **in
			}
		}
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadSelector != nil {
		in, out := &in.WorkloadSelector, &out.WorkloadSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionConfiguration.
func (in *InspectionConfiguration) DeepCopy() *InspectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(InspectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicy) DeepCopyInto(out *InspectionPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicy.
func (in *InspectionPolicy) DeepCopy() *InspectionPolicy {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectionPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicyList) DeepCopyInto(out *InspectionPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InspectionPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicyList.
func (in *InspectionPolicyList) DeepCopy() *InspectionPolicyList {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectionPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicySpec) DeepCopyInto(out *InspectionPolicySpec) {
	*out = *in
	if in.WorkNamespace != nil {
		in, out := &in.WorkNamespace, &out.WorkNamespace
		*out = new(string)
		**out = **in
	}
	in.Inspection.DeepCopyInto(&out.Inspection)
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.Inspector != nil {
		in, out := &in.Inspector, &out.Inspector
		*out = new(Inspector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicySpec.
func (in *InspectionPolicySpec) DeepCopy() *InspectionPolicySpec {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicyStatus) DeepCopyInto(out *InspectionPolicyStatus) {
	*out = *in
	if in.InspectionExecutor != nil {
		in, out := &in.InspectionExecutor, &out.InspectionExecutor
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.KubebenchExecutor != nil {
		in, out := &in.KubebenchExecutor, &out.KubebenchExecutor
		*out = make([]*corev1.ObjectReference, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.ObjectReference)
				**out = **in
			}
		}
	}
	if in.RiskExecutor != nil {
		in, out := &in.RiskExecutor, &out.RiskExecutor
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.WorkloadScannerExecutor != nil {
		in, out := &in.WorkloadScannerExecutor, &out.WorkloadScannerExecutor
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicyStatus.
func (in *InspectionPolicyStatus) DeepCopy() *InspectionPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inspector) DeepCopyInto(out *Inspector) {
	*out = *in
	out.ExportConfig = in.ExportConfig
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inspector.
func (in *Inspector) DeepCopy() *Inspector {
	if in == nil {
		return nil
	}
	out := new(Inspector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnownRegistry) DeepCopyInto(out *KnownRegistry) {
	*out = *in
	in.Registry.DeepCopyInto(&out.Registry)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnownRegistry.
func (in *KnownRegistry) DeepCopy() *KnownRegistry {
	if in == nil {
		return nil
	}
	out := new(KnownRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpensearchOutputConfig) DeepCopyInto(out *OpensearchOutputConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpensearchOutputConfig.
func (in *OpensearchOutputConfig) DeepCopy() *OpensearchOutputConfig {
	if in == nil {
		return nil
	}
	out := new(OpensearchOutputConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.CredentialRef != nil {
		in, out := &in.CredentialRef, &out.CredentialRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportData) DeepCopyInto(out *ReportData) {
	*out = *in
	out.ExportConfig = in.ExportConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportData.
func (in *ReportData) DeepCopy() *ReportData {
	if in == nil {
		return nil
	}
	out := new(ReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Setting) DeepCopyInto(out *Setting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Setting.
func (in *Setting) DeepCopy() *Setting {
	if in == nil {
		return nil
	}
	out := new(Setting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Setting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingList) DeepCopyInto(out *SettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Setting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingList.
func (in *SettingList) DeepCopy() *SettingList {
	if in == nil {
		return nil
	}
	out := new(SettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingSpec) DeepCopyInto(out *SettingSpec) {
	*out = *in
	if in.KnownRegistries != nil {
		in, out := &in.KnownRegistries, &out.KnownRegistries
		*out = make([]KnownRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.DataSource.DeepCopyInto(&out.DataSource)
	in.VacDataSource.DeepCopyInto(&out.VacDataSource)
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(Cache)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingSpec.
func (in *SettingSpec) DeepCopy() *SettingSpec {
	if in == nil {
		return nil
	}
	out := new(SettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingStatus) DeepCopyInto(out *SettingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingStatus.
func (in *SettingStatus) DeepCopy() *SettingStatus {
	if in == nil {
		return nil
	}
	out := new(SettingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Strategy) DeepCopyInto(out *Strategy) {
	*out = *in
	if in.HistoryLimit != nil {
		in, out := &in.HistoryLimit, &out.HistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Strategy.
func (in *Strategy) DeepCopy() *Strategy {
	if in == nil {
		return nil
	}
	out := new(Strategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VacDataSource) DeepCopyInto(out *VacDataSource) {
	*out = *in
	if in.CredentialRef != nil {
		in, out := &in.CredentialRef, &out.CredentialRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VacDataSource.
func (in *VacDataSource) DeepCopy() *VacDataSource {
	if in == nil {
		return nil
	}
	out := new(VacDataSource)
	in.DeepCopyInto(out)
	return out
}
