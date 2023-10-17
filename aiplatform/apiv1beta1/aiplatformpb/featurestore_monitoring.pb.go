// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/aiplatform/v1beta1/featurestore_monitoring.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The state defines whether to enable ImportFeature analysis.
type FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State int32

const (
	// Should not be used.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_STATE_UNSPECIFIED FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State = 0
	// The default behavior of whether to enable the monitoring.
	// EntityType-level config: disabled.
	// Feature-level config: inherited from the configuration of EntityType
	// this Feature belongs to.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_DEFAULT FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State = 1
	// Explicitly enables import features analysis.
	// EntityType-level config: by default enables import features analysis
	// for all Features under it. Feature-level config: enables import
	// features analysis regardless of the EntityType-level config.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_ENABLED FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State = 2
	// Explicitly disables import features analysis.
	// EntityType-level config: by default disables import features analysis
	// for all Features under it. Feature-level config: disables import
	// features analysis regardless of the EntityType-level config.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_DISABLED FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State = 3
)

// Enum value maps for FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State.
var (
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "DEFAULT",
		2: "ENABLED",
		3: "DISABLED",
	}
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"DEFAULT":           1,
		"ENABLED":           2,
		"DISABLED":          3,
	}
)

func (x FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State) Enum() *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State {
	p := new(FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State)
	*p = x
	return p
}

func (x FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_enumTypes[0].Descriptor()
}

func (FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_enumTypes[0]
}

func (x FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State.Descriptor instead.
func (FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Defines the baseline to do anomaly detection for feature values imported
// by each
// [ImportFeatureValues][google.cloud.aiplatform.v1beta1.FeaturestoreService.ImportFeatureValues]
// operation.
type FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline int32

const (
	// Should not be used.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_BASELINE_UNSPECIFIED FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline = 0
	// Choose the later one statistics generated by either most recent
	// snapshot analysis or previous import features analysis. If non of them
	// exists, skip anomaly detection and only generate a statistics.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_LATEST_STATS FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline = 1
	// Use the statistics generated by the most recent snapshot analysis if
	// exists.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_MOST_RECENT_SNAPSHOT_STATS FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline = 2
	// Use the statistics generated by the previous import features analysis
	// if exists.
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_PREVIOUS_IMPORT_FEATURES_STATS FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline = 3
)

// Enum value maps for FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline.
var (
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline_name = map[int32]string{
		0: "BASELINE_UNSPECIFIED",
		1: "LATEST_STATS",
		2: "MOST_RECENT_SNAPSHOT_STATS",
		3: "PREVIOUS_IMPORT_FEATURES_STATS",
	}
	FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline_value = map[string]int32{
		"BASELINE_UNSPECIFIED":           0,
		"LATEST_STATS":                   1,
		"MOST_RECENT_SNAPSHOT_STATS":     2,
		"PREVIOUS_IMPORT_FEATURES_STATS": 3,
	}
)

func (x FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline) Enum() *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline {
	p := new(FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline)
	*p = x
	return p
}

func (x FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_enumTypes[1].Descriptor()
}

func (FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_enumTypes[1]
}

func (x FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline.Descriptor instead.
func (FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescGZIP(), []int{0, 1, 1}
}

// Configuration of how features in Featurestore are monitored.
type FeaturestoreMonitoringConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The config for Snapshot Analysis Based Feature Monitoring.
	SnapshotAnalysis *FeaturestoreMonitoringConfig_SnapshotAnalysis `protobuf:"bytes,1,opt,name=snapshot_analysis,json=snapshotAnalysis,proto3" json:"snapshot_analysis,omitempty"`
	// The config for ImportFeatures Analysis Based Feature Monitoring.
	ImportFeaturesAnalysis *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis `protobuf:"bytes,2,opt,name=import_features_analysis,json=importFeaturesAnalysis,proto3" json:"import_features_analysis,omitempty"`
	// Threshold for numerical features of anomaly detection.
	// This is shared by all objectives of Featurestore Monitoring for numerical
	// features (i.e. Features with type
	// ([Feature.ValueType][google.cloud.aiplatform.v1beta1.Feature.ValueType])
	// DOUBLE or INT64).
	NumericalThresholdConfig *FeaturestoreMonitoringConfig_ThresholdConfig `protobuf:"bytes,3,opt,name=numerical_threshold_config,json=numericalThresholdConfig,proto3" json:"numerical_threshold_config,omitempty"`
	// Threshold for categorical features of anomaly detection.
	// This is shared by all types of Featurestore Monitoring for categorical
	// features (i.e. Features with type
	// ([Feature.ValueType][google.cloud.aiplatform.v1beta1.Feature.ValueType])
	// BOOL or STRING).
	CategoricalThresholdConfig *FeaturestoreMonitoringConfig_ThresholdConfig `protobuf:"bytes,4,opt,name=categorical_threshold_config,json=categoricalThresholdConfig,proto3" json:"categorical_threshold_config,omitempty"`
}

func (x *FeaturestoreMonitoringConfig) Reset() {
	*x = FeaturestoreMonitoringConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeaturestoreMonitoringConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeaturestoreMonitoringConfig) ProtoMessage() {}

func (x *FeaturestoreMonitoringConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeaturestoreMonitoringConfig.ProtoReflect.Descriptor instead.
func (*FeaturestoreMonitoringConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *FeaturestoreMonitoringConfig) GetSnapshotAnalysis() *FeaturestoreMonitoringConfig_SnapshotAnalysis {
	if x != nil {
		return x.SnapshotAnalysis
	}
	return nil
}

func (x *FeaturestoreMonitoringConfig) GetImportFeaturesAnalysis() *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis {
	if x != nil {
		return x.ImportFeaturesAnalysis
	}
	return nil
}

func (x *FeaturestoreMonitoringConfig) GetNumericalThresholdConfig() *FeaturestoreMonitoringConfig_ThresholdConfig {
	if x != nil {
		return x.NumericalThresholdConfig
	}
	return nil
}

func (x *FeaturestoreMonitoringConfig) GetCategoricalThresholdConfig() *FeaturestoreMonitoringConfig_ThresholdConfig {
	if x != nil {
		return x.CategoricalThresholdConfig
	}
	return nil
}

// Configuration of the Featurestore's Snapshot Analysis Based Monitoring.
// This type of analysis generates statistics for each Feature based on a
// snapshot of the latest feature value of each entities every
// monitoring_interval.
type FeaturestoreMonitoringConfig_SnapshotAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The monitoring schedule for snapshot analysis.
	// For EntityType-level config:
	//   unset / disabled = true indicates disabled by
	//   default for Features under it; otherwise by default enable snapshot
	//   analysis monitoring with monitoring_interval for Features under it.
	// Feature-level config:
	//   disabled = true indicates disabled regardless of the EntityType-level
	//   config; unset monitoring_interval indicates going with EntityType-level
	//   config; otherwise run snapshot analysis monitoring with
	//   monitoring_interval regardless of the EntityType-level config.
	// Explicitly Disable the snapshot analysis based monitoring.
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Configuration of the snapshot analysis based monitoring pipeline running
	// interval. The value is rolled up to full day.
	// If both
	// [monitoring_interval_days][google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days]
	// and the deprecated `monitoring_interval` field
	// are set when creating/updating EntityTypes/Features,
	// [monitoring_interval_days][google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval_days]
	// will be used.
	//
	// Deprecated: Marked as deprecated in google/cloud/aiplatform/v1beta1/featurestore_monitoring.proto.
	MonitoringInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=monitoring_interval,json=monitoringInterval,proto3" json:"monitoring_interval,omitempty"`
	// Configuration of the snapshot analysis based monitoring pipeline
	// running interval. The value indicates number of days.
	MonitoringIntervalDays int32 `protobuf:"varint,3,opt,name=monitoring_interval_days,json=monitoringIntervalDays,proto3" json:"monitoring_interval_days,omitempty"`
	// Customized export features time window for snapshot analysis. Unit is one
	// day. Default value is 3 weeks. Minimum value is 1 day. Maximum value is
	// 4000 days.
	StalenessDays int32 `protobuf:"varint,4,opt,name=staleness_days,json=stalenessDays,proto3" json:"staleness_days,omitempty"`
}

func (x *FeaturestoreMonitoringConfig_SnapshotAnalysis) Reset() {
	*x = FeaturestoreMonitoringConfig_SnapshotAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeaturestoreMonitoringConfig_SnapshotAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeaturestoreMonitoringConfig_SnapshotAnalysis) ProtoMessage() {}

func (x *FeaturestoreMonitoringConfig_SnapshotAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeaturestoreMonitoringConfig_SnapshotAnalysis.ProtoReflect.Descriptor instead.
func (*FeaturestoreMonitoringConfig_SnapshotAnalysis) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FeaturestoreMonitoringConfig_SnapshotAnalysis) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

// Deprecated: Marked as deprecated in google/cloud/aiplatform/v1beta1/featurestore_monitoring.proto.
func (x *FeaturestoreMonitoringConfig_SnapshotAnalysis) GetMonitoringInterval() *durationpb.Duration {
	if x != nil {
		return x.MonitoringInterval
	}
	return nil
}

func (x *FeaturestoreMonitoringConfig_SnapshotAnalysis) GetMonitoringIntervalDays() int32 {
	if x != nil {
		return x.MonitoringIntervalDays
	}
	return 0
}

func (x *FeaturestoreMonitoringConfig_SnapshotAnalysis) GetStalenessDays() int32 {
	if x != nil {
		return x.StalenessDays
	}
	return 0
}

// Configuration of the Featurestore's ImportFeature Analysis Based
// Monitoring. This type of analysis generates statistics for values of each
// Feature imported by every
// [ImportFeatureValues][google.cloud.aiplatform.v1beta1.FeaturestoreService.ImportFeatureValues]
// operation.
type FeaturestoreMonitoringConfig_ImportFeaturesAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether to enable / disable / inherite default hebavior for import
	// features analysis.
	State FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State" json:"state,omitempty"`
	// The baseline used to do anomaly detection for the statistics generated by
	// import features analysis.
	AnomalyDetectionBaseline FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline `protobuf:"varint,2,opt,name=anomaly_detection_baseline,json=anomalyDetectionBaseline,proto3,enum=google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline" json:"anomaly_detection_baseline,omitempty"`
}

func (x *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis) Reset() {
	*x = FeaturestoreMonitoringConfig_ImportFeaturesAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeaturestoreMonitoringConfig_ImportFeaturesAnalysis) ProtoMessage() {}

func (x *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeaturestoreMonitoringConfig_ImportFeaturesAnalysis.ProtoReflect.Descriptor instead.
func (*FeaturestoreMonitoringConfig_ImportFeaturesAnalysis) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescGZIP(), []int{0, 1}
}

func (x *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis) GetState() FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State {
	if x != nil {
		return x.State
	}
	return FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_STATE_UNSPECIFIED
}

func (x *FeaturestoreMonitoringConfig_ImportFeaturesAnalysis) GetAnomalyDetectionBaseline() FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline {
	if x != nil {
		return x.AnomalyDetectionBaseline
	}
	return FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_BASELINE_UNSPECIFIED
}

// The config for Featurestore Monitoring threshold.
type FeaturestoreMonitoringConfig_ThresholdConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Threshold:
	//	*FeaturestoreMonitoringConfig_ThresholdConfig_Value
	Threshold isFeaturestoreMonitoringConfig_ThresholdConfig_Threshold `protobuf_oneof:"threshold"`
}

func (x *FeaturestoreMonitoringConfig_ThresholdConfig) Reset() {
	*x = FeaturestoreMonitoringConfig_ThresholdConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeaturestoreMonitoringConfig_ThresholdConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeaturestoreMonitoringConfig_ThresholdConfig) ProtoMessage() {}

func (x *FeaturestoreMonitoringConfig_ThresholdConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeaturestoreMonitoringConfig_ThresholdConfig.ProtoReflect.Descriptor instead.
func (*FeaturestoreMonitoringConfig_ThresholdConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescGZIP(), []int{0, 2}
}

func (m *FeaturestoreMonitoringConfig_ThresholdConfig) GetThreshold() isFeaturestoreMonitoringConfig_ThresholdConfig_Threshold {
	if m != nil {
		return m.Threshold
	}
	return nil
}

func (x *FeaturestoreMonitoringConfig_ThresholdConfig) GetValue() float64 {
	if x, ok := x.GetThreshold().(*FeaturestoreMonitoringConfig_ThresholdConfig_Value); ok {
		return x.Value
	}
	return 0
}

type isFeaturestoreMonitoringConfig_ThresholdConfig_Threshold interface {
	isFeaturestoreMonitoringConfig_ThresholdConfig_Threshold()
}

type FeaturestoreMonitoringConfig_ThresholdConfig_Value struct {
	// Specify a threshold value that can trigger the alert.
	// 1. For categorical feature, the distribution distance is calculated by
	// L-inifinity norm.
	// 2. For numerical feature, the distribution distance is calculated by
	// Jensen–Shannon divergence. Each feature must have a non-zero threshold
	// if they need to be monitored. Otherwise no alert will be triggered for
	// that feature.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3,oneof"`
}

func (*FeaturestoreMonitoringConfig_ThresholdConfig_Value) isFeaturestoreMonitoringConfig_ThresholdConfig_Threshold() {
}

var File_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd5, 0x0a, 0x0a, 0x1c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x7b, 0x0a, 0x11, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x10, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x8e,
	0x01, 0x0a, 0x18, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x16, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12,
	0x8b, 0x01, 0x0a, 0x1a, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x18, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x8f, 0x01,
	0x0a, 0x1c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x1a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0xdf, 0x01, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x4e, 0x0a, 0x13, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x18, 0x01, 0x52, 0x12, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x38, 0x0a, 0x18, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x16, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x79, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x79,
	0x73, 0x1a, 0xec, 0x03, 0x0a, 0x16, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x70, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x9b,
	0x01, 0x0a, 0x1a, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x18, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x46, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x22, 0x7a, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x18, 0x0a, 0x14, 0x42, 0x41, 0x53, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a,
	0x4d, 0x4f, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x4e, 0x41, 0x50,
	0x53, 0x48, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e,
	0x50, 0x52, 0x45, 0x56, 0x49, 0x4f, 0x55, 0x53, 0x5f, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x03,
	0x1a, 0x36, 0x0a, 0x0f, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0xf2, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x1b, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56,
	0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_goTypes = []interface{}{
	(FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_State)(0),    // 0: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.State
	(FeaturestoreMonitoringConfig_ImportFeaturesAnalysis_Baseline)(0), // 1: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.Baseline
	(*FeaturestoreMonitoringConfig)(nil),                              // 2: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig
	(*FeaturestoreMonitoringConfig_SnapshotAnalysis)(nil),             // 3: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis
	(*FeaturestoreMonitoringConfig_ImportFeaturesAnalysis)(nil),       // 4: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis
	(*FeaturestoreMonitoringConfig_ThresholdConfig)(nil),              // 5: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ThresholdConfig
	(*durationpb.Duration)(nil),                                       // 6: google.protobuf.Duration
}
var file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_depIdxs = []int32{
	3, // 0: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.snapshot_analysis:type_name -> google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis
	4, // 1: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.import_features_analysis:type_name -> google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis
	5, // 2: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.numerical_threshold_config:type_name -> google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ThresholdConfig
	5, // 3: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.categorical_threshold_config:type_name -> google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ThresholdConfig
	6, // 4: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.SnapshotAnalysis.monitoring_interval:type_name -> google.protobuf.Duration
	0, // 5: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.state:type_name -> google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.State
	1, // 6: google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.anomaly_detection_baseline:type_name -> google.cloud.aiplatform.v1beta1.FeaturestoreMonitoringConfig.ImportFeaturesAnalysis.Baseline
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_init() }
func file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeaturestoreMonitoringConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeaturestoreMonitoringConfig_SnapshotAnalysis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeaturestoreMonitoringConfig_ImportFeaturesAnalysis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeaturestoreMonitoringConfig_ThresholdConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*FeaturestoreMonitoringConfig_ThresholdConfig_Value)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto = out.File
	file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_featurestore_monitoring_proto_depIdxs = nil
}
