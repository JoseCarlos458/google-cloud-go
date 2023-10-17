// Copyright 2022 Google LLC
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
// source: google/monitoring/dashboard/v1/scorecard.proto

package dashboardpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A widget showing the latest value of a metric, and how this value relates to
// one or more thresholds.
type Scorecard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Fields for querying time series data from the
	// Stackdriver metrics API.
	TimeSeriesQuery *TimeSeriesQuery `protobuf:"bytes,1,opt,name=time_series_query,json=timeSeriesQuery,proto3" json:"time_series_query,omitempty"`
	// Defines the optional additional chart shown on the scorecard. If
	// neither is included - then a default scorecard is shown.
	//
	// Types that are assignable to DataView:
	//	*Scorecard_GaugeView_
	//	*Scorecard_SparkChartView_
	DataView isScorecard_DataView `protobuf_oneof:"data_view"`
	// The thresholds used to determine the state of the scorecard given the
	// time series' current value. For an actual value x, the scorecard is in a
	// danger state if x is less than or equal to a danger threshold that triggers
	// below, or greater than or equal to a danger threshold that triggers above.
	// Similarly, if x is above/below a warning threshold that triggers
	// above/below, then the scorecard is in a warning state - unless x also puts
	// it in a danger state. (Danger trumps warning.)
	//
	// As an example, consider a scorecard with the following four thresholds:
	//
	// ```
	// {
	//   value: 90,
	//   category: 'DANGER',
	//   trigger: 'ABOVE',
	// },
	// {
	//   value: 70,
	//   category: 'WARNING',
	//   trigger: 'ABOVE',
	// },
	// {
	//   value: 10,
	//   category: 'DANGER',
	//   trigger: 'BELOW',
	// },
	// {
	//   value: 20,
	//   category: 'WARNING',
	//   trigger: 'BELOW',
	// }
	// ```
	//
	// Then: values less than or equal to 10 would put the scorecard in a DANGER
	// state, values greater than 10 but less than or equal to 20 a WARNING state,
	// values strictly between 20 and 70 an OK state, values greater than or equal
	// to 70 but less than 90 a WARNING state, and values greater than or equal to
	// 90 a DANGER state.
	Thresholds []*Threshold `protobuf:"bytes,6,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
}

func (x *Scorecard) Reset() {
	*x = Scorecard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scorecard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scorecard) ProtoMessage() {}

func (x *Scorecard) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scorecard.ProtoReflect.Descriptor instead.
func (*Scorecard) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_scorecard_proto_rawDescGZIP(), []int{0}
}

func (x *Scorecard) GetTimeSeriesQuery() *TimeSeriesQuery {
	if x != nil {
		return x.TimeSeriesQuery
	}
	return nil
}

func (m *Scorecard) GetDataView() isScorecard_DataView {
	if m != nil {
		return m.DataView
	}
	return nil
}

func (x *Scorecard) GetGaugeView() *Scorecard_GaugeView {
	if x, ok := x.GetDataView().(*Scorecard_GaugeView_); ok {
		return x.GaugeView
	}
	return nil
}

func (x *Scorecard) GetSparkChartView() *Scorecard_SparkChartView {
	if x, ok := x.GetDataView().(*Scorecard_SparkChartView_); ok {
		return x.SparkChartView
	}
	return nil
}

func (x *Scorecard) GetThresholds() []*Threshold {
	if x != nil {
		return x.Thresholds
	}
	return nil
}

type isScorecard_DataView interface {
	isScorecard_DataView()
}

type Scorecard_GaugeView_ struct {
	// Will cause the scorecard to show a gauge chart.
	GaugeView *Scorecard_GaugeView `protobuf:"bytes,4,opt,name=gauge_view,json=gaugeView,proto3,oneof"`
}

type Scorecard_SparkChartView_ struct {
	// Will cause the scorecard to show a spark chart.
	SparkChartView *Scorecard_SparkChartView `protobuf:"bytes,5,opt,name=spark_chart_view,json=sparkChartView,proto3,oneof"`
}

func (*Scorecard_GaugeView_) isScorecard_DataView() {}

func (*Scorecard_SparkChartView_) isScorecard_DataView() {}

// A gauge chart shows where the current value sits within a pre-defined
// range. The upper and lower bounds should define the possible range of
// values for the scorecard's query (inclusive).
type Scorecard_GaugeView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The lower bound for this gauge chart. The value of the chart should
	// always be greater than or equal to this.
	LowerBound float64 `protobuf:"fixed64,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	// The upper bound for this gauge chart. The value of the chart should
	// always be less than or equal to this.
	UpperBound float64 `protobuf:"fixed64,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
}

func (x *Scorecard_GaugeView) Reset() {
	*x = Scorecard_GaugeView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scorecard_GaugeView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scorecard_GaugeView) ProtoMessage() {}

func (x *Scorecard_GaugeView) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scorecard_GaugeView.ProtoReflect.Descriptor instead.
func (*Scorecard_GaugeView) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_scorecard_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Scorecard_GaugeView) GetLowerBound() float64 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

func (x *Scorecard_GaugeView) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

// A sparkChart is a small chart suitable for inclusion in a table-cell or
// inline in text. This message contains the configuration for a sparkChart
// to show up on a Scorecard, showing recent trends of the scorecard's
// timeseries.
type Scorecard_SparkChartView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The type of sparkchart to show in this chartView.
	SparkChartType SparkChartType `protobuf:"varint,1,opt,name=spark_chart_type,json=sparkChartType,proto3,enum=google.monitoring.dashboard.v1.SparkChartType" json:"spark_chart_type,omitempty"`
	// The lower bound on data point frequency in the chart implemented by
	// specifying the minimum alignment period to use in a time series query.
	// For example, if the data is published once every 10 minutes it would not
	// make sense to fetch and align data at one minute intervals. This field is
	// optional and exists only as a hint.
	MinAlignmentPeriod *durationpb.Duration `protobuf:"bytes,2,opt,name=min_alignment_period,json=minAlignmentPeriod,proto3" json:"min_alignment_period,omitempty"`
}

func (x *Scorecard_SparkChartView) Reset() {
	*x = Scorecard_SparkChartView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scorecard_SparkChartView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scorecard_SparkChartView) ProtoMessage() {}

func (x *Scorecard_SparkChartView) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scorecard_SparkChartView.ProtoReflect.Descriptor instead.
func (*Scorecard_SparkChartView) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_scorecard_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Scorecard_SparkChartView) GetSparkChartType() SparkChartType {
	if x != nil {
		return x.SparkChartType
	}
	return SparkChartType_SPARK_CHART_TYPE_UNSPECIFIED
}

func (x *Scorecard_SparkChartView) GetMinAlignmentPeriod() *durationpb.Duration {
	if x != nil {
		return x.MinAlignmentPeriod
	}
	return nil
}

var File_google_monitoring_dashboard_v1_scorecard_proto protoreflect.FileDescriptor

var file_google_monitoring_dashboard_v1_scorecard_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8f, 0x05, 0x0a, 0x09, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x12, 0x60, 0x0a,
	0x11, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f,
	0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x54, 0x0a, 0x0a, 0x67, 0x61, 0x75, 0x67, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x47,
	0x61, 0x75, 0x67, 0x65, 0x56, 0x69, 0x65, 0x77, 0x48, 0x00, 0x52, 0x09, 0x67, 0x61, 0x75, 0x67,
	0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x64, 0x0a, 0x10, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x56, 0x69, 0x65, 0x77, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x43, 0x68, 0x61, 0x72, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x49, 0x0a, 0x0a, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x1a, 0x4d, 0x0a, 0x09, 0x47, 0x61, 0x75, 0x67, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x1a, 0xbc, 0x01, 0x0a, 0x0e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x5d, 0x0a, 0x10, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x68, 0x61, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x5f, 0x61,
	0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x12, 0x6d, 0x69, 0x6e, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x42, 0xf7, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x70, 0x62, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x5c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_dashboard_v1_scorecard_proto_rawDescOnce sync.Once
	file_google_monitoring_dashboard_v1_scorecard_proto_rawDescData = file_google_monitoring_dashboard_v1_scorecard_proto_rawDesc
)

func file_google_monitoring_dashboard_v1_scorecard_proto_rawDescGZIP() []byte {
	file_google_monitoring_dashboard_v1_scorecard_proto_rawDescOnce.Do(func() {
		file_google_monitoring_dashboard_v1_scorecard_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_dashboard_v1_scorecard_proto_rawDescData)
	})
	return file_google_monitoring_dashboard_v1_scorecard_proto_rawDescData
}

var file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_monitoring_dashboard_v1_scorecard_proto_goTypes = []interface{}{
	(*Scorecard)(nil),                // 0: google.monitoring.dashboard.v1.Scorecard
	(*Scorecard_GaugeView)(nil),      // 1: google.monitoring.dashboard.v1.Scorecard.GaugeView
	(*Scorecard_SparkChartView)(nil), // 2: google.monitoring.dashboard.v1.Scorecard.SparkChartView
	(*TimeSeriesQuery)(nil),          // 3: google.monitoring.dashboard.v1.TimeSeriesQuery
	(*Threshold)(nil),                // 4: google.monitoring.dashboard.v1.Threshold
	(SparkChartType)(0),              // 5: google.monitoring.dashboard.v1.SparkChartType
	(*durationpb.Duration)(nil),      // 6: google.protobuf.Duration
}
var file_google_monitoring_dashboard_v1_scorecard_proto_depIdxs = []int32{
	3, // 0: google.monitoring.dashboard.v1.Scorecard.time_series_query:type_name -> google.monitoring.dashboard.v1.TimeSeriesQuery
	1, // 1: google.monitoring.dashboard.v1.Scorecard.gauge_view:type_name -> google.monitoring.dashboard.v1.Scorecard.GaugeView
	2, // 2: google.monitoring.dashboard.v1.Scorecard.spark_chart_view:type_name -> google.monitoring.dashboard.v1.Scorecard.SparkChartView
	4, // 3: google.monitoring.dashboard.v1.Scorecard.thresholds:type_name -> google.monitoring.dashboard.v1.Threshold
	5, // 4: google.monitoring.dashboard.v1.Scorecard.SparkChartView.spark_chart_type:type_name -> google.monitoring.dashboard.v1.SparkChartType
	6, // 5: google.monitoring.dashboard.v1.Scorecard.SparkChartView.min_alignment_period:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_monitoring_dashboard_v1_scorecard_proto_init() }
func file_google_monitoring_dashboard_v1_scorecard_proto_init() {
	if File_google_monitoring_dashboard_v1_scorecard_proto != nil {
		return
	}
	file_google_monitoring_dashboard_v1_metrics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scorecard); i {
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
		file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scorecard_GaugeView); i {
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
		file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scorecard_SparkChartView); i {
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
	file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Scorecard_GaugeView_)(nil),
		(*Scorecard_SparkChartView_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_monitoring_dashboard_v1_scorecard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_dashboard_v1_scorecard_proto_goTypes,
		DependencyIndexes: file_google_monitoring_dashboard_v1_scorecard_proto_depIdxs,
		MessageInfos:      file_google_monitoring_dashboard_v1_scorecard_proto_msgTypes,
	}.Build()
	File_google_monitoring_dashboard_v1_scorecard_proto = out.File
	file_google_monitoring_dashboard_v1_scorecard_proto_rawDesc = nil
	file_google_monitoring_dashboard_v1_scorecard_proto_goTypes = nil
	file_google_monitoring_dashboard_v1_scorecard_proto_depIdxs = nil
}
