// Copyright 2020 Google LLC
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
// source: google/cloud/automl/v1beta1/text_extraction.proto

package automlpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Annotation for identifying spans of text.
type TextExtractionAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Text extraction annotations can either be a text segment or a
	// text relation.
	//
	// Types that are assignable to Annotation:
	//	*TextExtractionAnnotation_TextSegment
	Annotation isTextExtractionAnnotation_Annotation `protobuf_oneof:"annotation"`
	// Output only. A confidence estimate between 0.0 and 1.0. A higher value
	// means greater confidence in correctness of the annotation.
	Score float32 `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *TextExtractionAnnotation) Reset() {
	*x = TextExtractionAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextExtractionAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextExtractionAnnotation) ProtoMessage() {}

func (x *TextExtractionAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextExtractionAnnotation.ProtoReflect.Descriptor instead.
func (*TextExtractionAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescGZIP(), []int{0}
}

func (m *TextExtractionAnnotation) GetAnnotation() isTextExtractionAnnotation_Annotation {
	if m != nil {
		return m.Annotation
	}
	return nil
}

func (x *TextExtractionAnnotation) GetTextSegment() *TextSegment {
	if x, ok := x.GetAnnotation().(*TextExtractionAnnotation_TextSegment); ok {
		return x.TextSegment
	}
	return nil
}

func (x *TextExtractionAnnotation) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type isTextExtractionAnnotation_Annotation interface {
	isTextExtractionAnnotation_Annotation()
}

type TextExtractionAnnotation_TextSegment struct {
	// An entity annotation will set this, which is the part of the original
	// text to which the annotation pertains.
	TextSegment *TextSegment `protobuf:"bytes,3,opt,name=text_segment,json=textSegment,proto3,oneof"`
}

func (*TextExtractionAnnotation_TextSegment) isTextExtractionAnnotation_Annotation() {}

// Model evaluation metrics for text extraction problems.
type TextExtractionEvaluationMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The Area under precision recall curve metric.
	AuPrc float32 `protobuf:"fixed32,1,opt,name=au_prc,json=auPrc,proto3" json:"au_prc,omitempty"`
	// Output only. Metrics that have confidence thresholds.
	// Precision-recall curve can be derived from it.
	ConfidenceMetricsEntries []*TextExtractionEvaluationMetrics_ConfidenceMetricsEntry `protobuf:"bytes,2,rep,name=confidence_metrics_entries,json=confidenceMetricsEntries,proto3" json:"confidence_metrics_entries,omitempty"`
}

func (x *TextExtractionEvaluationMetrics) Reset() {
	*x = TextExtractionEvaluationMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextExtractionEvaluationMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextExtractionEvaluationMetrics) ProtoMessage() {}

func (x *TextExtractionEvaluationMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextExtractionEvaluationMetrics.ProtoReflect.Descriptor instead.
func (*TextExtractionEvaluationMetrics) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescGZIP(), []int{1}
}

func (x *TextExtractionEvaluationMetrics) GetAuPrc() float32 {
	if x != nil {
		return x.AuPrc
	}
	return 0
}

func (x *TextExtractionEvaluationMetrics) GetConfidenceMetricsEntries() []*TextExtractionEvaluationMetrics_ConfidenceMetricsEntry {
	if x != nil {
		return x.ConfidenceMetricsEntries
	}
	return nil
}

// Metrics for a single confidence threshold.
type TextExtractionEvaluationMetrics_ConfidenceMetricsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The confidence threshold value used to compute the metrics.
	// Only annotations with score of at least this threshold are considered to
	// be ones the model would return.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Output only. Recall under the given confidence threshold.
	Recall float32 `protobuf:"fixed32,3,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. Precision under the given confidence threshold.
	Precision float32 `protobuf:"fixed32,4,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,5,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
}

func (x *TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) Reset() {
	*x = TextExtractionEvaluationMetrics_ConfidenceMetricsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) ProtoMessage() {}

func (x *TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextExtractionEvaluationMetrics_ConfidenceMetricsEntry.ProtoReflect.Descriptor instead.
func (*TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if x != nil {
		return x.ConfidenceThreshold
	}
	return 0
}

func (x *TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) GetRecall() float32 {
	if x != nil {
		return x.Recall
	}
	return 0
}

func (x *TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) GetPrecision() float32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *TextExtractionEvaluationMetrics_ConfidenceMetricsEntry) GetF1Score() float32 {
	if x != nil {
		return x.F1Score
	}
	return 0
}

var File_google_cloud_automl_v1beta1_text_extraction_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_text_extraction_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x01, 0x0a, 0x18, 0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a,
	0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x0b, 0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xeb, 0x02, 0x0a, 0x1f, 0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x75, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x61, 0x75, 0x50, 0x72, 0x63, 0x12, 0x91, 0x01, 0x0a, 0x1a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x53, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a,
	0x9c, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x31, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x31, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x9b,
	0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0xca, 0x02, 0x1b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74,
	0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74,
	0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescData = file_google_cloud_automl_v1beta1_text_extraction_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_text_extraction_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_automl_v1beta1_text_extraction_proto_goTypes = []interface{}{
	(*TextExtractionAnnotation)(nil),                               // 0: google.cloud.automl.v1beta1.TextExtractionAnnotation
	(*TextExtractionEvaluationMetrics)(nil),                        // 1: google.cloud.automl.v1beta1.TextExtractionEvaluationMetrics
	(*TextExtractionEvaluationMetrics_ConfidenceMetricsEntry)(nil), // 2: google.cloud.automl.v1beta1.TextExtractionEvaluationMetrics.ConfidenceMetricsEntry
	(*TextSegment)(nil),                                            // 3: google.cloud.automl.v1beta1.TextSegment
}
var file_google_cloud_automl_v1beta1_text_extraction_proto_depIdxs = []int32{
	3, // 0: google.cloud.automl.v1beta1.TextExtractionAnnotation.text_segment:type_name -> google.cloud.automl.v1beta1.TextSegment
	2, // 1: google.cloud.automl.v1beta1.TextExtractionEvaluationMetrics.confidence_metrics_entries:type_name -> google.cloud.automl.v1beta1.TextExtractionEvaluationMetrics.ConfidenceMetricsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_text_extraction_proto_init() }
func file_google_cloud_automl_v1beta1_text_extraction_proto_init() {
	if File_google_cloud_automl_v1beta1_text_extraction_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_text_segment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextExtractionAnnotation); i {
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
		file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextExtractionEvaluationMetrics); i {
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
		file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextExtractionEvaluationMetrics_ConfidenceMetricsEntry); i {
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
	file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TextExtractionAnnotation_TextSegment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_automl_v1beta1_text_extraction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_text_extraction_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_text_extraction_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_text_extraction_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_text_extraction_proto = out.File
	file_google_cloud_automl_v1beta1_text_extraction_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_text_extraction_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_text_extraction_proto_depIdxs = nil
}
