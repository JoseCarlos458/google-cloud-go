// Copyright 2021 Google LLC
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
// source: google/cloud/retail/v2alpha/export_config.proto

package retailpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration of destination for Export related errors.
type ExportErrorsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Errors destination.
	//
	// Types that are assignable to Destination:
	//	*ExportErrorsConfig_GcsPrefix
	Destination isExportErrorsConfig_Destination `protobuf_oneof:"destination"`
}

func (x *ExportErrorsConfig) Reset() {
	*x = ExportErrorsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportErrorsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportErrorsConfig) ProtoMessage() {}

func (x *ExportErrorsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportErrorsConfig.ProtoReflect.Descriptor instead.
func (*ExportErrorsConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP(), []int{0}
}

func (m *ExportErrorsConfig) GetDestination() isExportErrorsConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *ExportErrorsConfig) GetGcsPrefix() string {
	if x, ok := x.GetDestination().(*ExportErrorsConfig_GcsPrefix); ok {
		return x.GcsPrefix
	}
	return ""
}

type isExportErrorsConfig_Destination interface {
	isExportErrorsConfig_Destination()
}

type ExportErrorsConfig_GcsPrefix struct {
	// Google Cloud Storage path for import errors. This must be an empty,
	// existing Cloud Storage bucket. Export errors will be written to a file in
	// this bucket, one per line, as a JSON-encoded
	// `google.rpc.Status` message.
	GcsPrefix string `protobuf:"bytes,1,opt,name=gcs_prefix,json=gcsPrefix,proto3,oneof"`
}

func (*ExportErrorsConfig_GcsPrefix) isExportErrorsConfig_Destination() {}

// Metadata related to the progress of the Export operation. This is
// returned by the google.longrunning.Operation.metadata field.
type ExportMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Operation last update time. If the operation is done, this is also the
	// finish time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *ExportMetadata) Reset() {
	*x = ExportMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetadata) ProtoMessage() {}

func (x *ExportMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetadata.ProtoReflect.Descriptor instead.
func (*ExportMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP(), []int{1}
}

func (x *ExportMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ExportMetadata) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Response of the ExportProductsRequest. If the long running
// operation is done, then this message is returned by the
// google.longrunning.Operations.response field if the operation was successful.
type ExportProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A sample of errors encountered while processing the request.
	ErrorSamples []*status.Status `protobuf:"bytes,1,rep,name=error_samples,json=errorSamples,proto3" json:"error_samples,omitempty"`
	// This field is never set.
	ErrorsConfig *ExportErrorsConfig `protobuf:"bytes,2,opt,name=errors_config,json=errorsConfig,proto3" json:"errors_config,omitempty"`
	// Output result indicating where the data were exported to.
	OutputResult *OutputResult `protobuf:"bytes,3,opt,name=output_result,json=outputResult,proto3" json:"output_result,omitempty"`
}

func (x *ExportProductsResponse) Reset() {
	*x = ExportProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportProductsResponse) ProtoMessage() {}

func (x *ExportProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportProductsResponse.ProtoReflect.Descriptor instead.
func (*ExportProductsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP(), []int{2}
}

func (x *ExportProductsResponse) GetErrorSamples() []*status.Status {
	if x != nil {
		return x.ErrorSamples
	}
	return nil
}

func (x *ExportProductsResponse) GetErrorsConfig() *ExportErrorsConfig {
	if x != nil {
		return x.ErrorsConfig
	}
	return nil
}

func (x *ExportProductsResponse) GetOutputResult() *OutputResult {
	if x != nil {
		return x.OutputResult
	}
	return nil
}

// Response of the ExportUserEventsRequest. If the long running
// operation was successful, then this message is returned by the
// google.longrunning.Operations.response field if the operation was successful.
type ExportUserEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A sample of errors encountered while processing the request.
	ErrorSamples []*status.Status `protobuf:"bytes,1,rep,name=error_samples,json=errorSamples,proto3" json:"error_samples,omitempty"`
	// This field is never set.
	ErrorsConfig *ExportErrorsConfig `protobuf:"bytes,2,opt,name=errors_config,json=errorsConfig,proto3" json:"errors_config,omitempty"`
	// Output result indicating where the data were exported to.
	OutputResult *OutputResult `protobuf:"bytes,3,opt,name=output_result,json=outputResult,proto3" json:"output_result,omitempty"`
}

func (x *ExportUserEventsResponse) Reset() {
	*x = ExportUserEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportUserEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportUserEventsResponse) ProtoMessage() {}

func (x *ExportUserEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportUserEventsResponse.ProtoReflect.Descriptor instead.
func (*ExportUserEventsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP(), []int{3}
}

func (x *ExportUserEventsResponse) GetErrorSamples() []*status.Status {
	if x != nil {
		return x.ErrorSamples
	}
	return nil
}

func (x *ExportUserEventsResponse) GetErrorsConfig() *ExportErrorsConfig {
	if x != nil {
		return x.ErrorsConfig
	}
	return nil
}

func (x *ExportUserEventsResponse) GetOutputResult() *OutputResult {
	if x != nil {
		return x.OutputResult
	}
	return nil
}

// Output result that stores the information about where the exported data is
// stored.
type OutputResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The BigQuery location where the result is stored.
	BigqueryResult []*BigQueryOutputResult `protobuf:"bytes,1,rep,name=bigquery_result,json=bigqueryResult,proto3" json:"bigquery_result,omitempty"`
	// The Google Cloud Storage location where the result is stored.
	GcsResult []*GcsOutputResult `protobuf:"bytes,2,rep,name=gcs_result,json=gcsResult,proto3" json:"gcs_result,omitempty"`
}

func (x *OutputResult) Reset() {
	*x = OutputResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputResult) ProtoMessage() {}

func (x *OutputResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputResult.ProtoReflect.Descriptor instead.
func (*OutputResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP(), []int{4}
}

func (x *OutputResult) GetBigqueryResult() []*BigQueryOutputResult {
	if x != nil {
		return x.BigqueryResult
	}
	return nil
}

func (x *OutputResult) GetGcsResult() []*GcsOutputResult {
	if x != nil {
		return x.GcsResult
	}
	return nil
}

// A BigQuery output result.
type BigQueryOutputResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of a BigQuery Dataset.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The ID of a BigQuery Table.
	TableId string `protobuf:"bytes,2,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
}

func (x *BigQueryOutputResult) Reset() {
	*x = BigQueryOutputResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryOutputResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryOutputResult) ProtoMessage() {}

func (x *BigQueryOutputResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryOutputResult.ProtoReflect.Descriptor instead.
func (*BigQueryOutputResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP(), []int{5}
}

func (x *BigQueryOutputResult) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *BigQueryOutputResult) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

// A Gcs output result.
type GcsOutputResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uri of Gcs output
	OutputUri string `protobuf:"bytes,1,opt,name=output_uri,json=outputUri,proto3" json:"output_uri,omitempty"`
}

func (x *GcsOutputResult) Reset() {
	*x = GcsOutputResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsOutputResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsOutputResult) ProtoMessage() {}

func (x *GcsOutputResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsOutputResult.ProtoReflect.Descriptor instead.
func (*GcsOutputResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP(), []int{6}
}

func (x *GcsOutputResult) GetOutputUri() string {
	if x != nil {
		return x.OutputUri
	}
	return ""
}

var File_google_cloud_retail_v2alpha_export_config_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_export_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0a, 0x67, 0x63,
	0x73, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x67, 0x63, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x0d, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf7, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x0d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x4e, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0xf9, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a,
	0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xb7, 0x01,
	0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5a,
	0x0a, 0x0f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x67, 0x63,
	0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x63, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x09, 0x67, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x50, 0x0a, 0x14, 0x42, 0x69, 0x67, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x63, 0x73,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x55, 0x72, 0x69, 0x42, 0xd5, 0x01, 0x0a, 0x1f,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x11, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02, 0x06,
	0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_export_config_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_export_config_proto_rawDescData = file_google_cloud_retail_v2alpha_export_config_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_export_config_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_export_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_export_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_export_config_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_export_config_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_export_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_retail_v2alpha_export_config_proto_goTypes = []interface{}{
	(*ExportErrorsConfig)(nil),       // 0: google.cloud.retail.v2alpha.ExportErrorsConfig
	(*ExportMetadata)(nil),           // 1: google.cloud.retail.v2alpha.ExportMetadata
	(*ExportProductsResponse)(nil),   // 2: google.cloud.retail.v2alpha.ExportProductsResponse
	(*ExportUserEventsResponse)(nil), // 3: google.cloud.retail.v2alpha.ExportUserEventsResponse
	(*OutputResult)(nil),             // 4: google.cloud.retail.v2alpha.OutputResult
	(*BigQueryOutputResult)(nil),     // 5: google.cloud.retail.v2alpha.BigQueryOutputResult
	(*GcsOutputResult)(nil),          // 6: google.cloud.retail.v2alpha.GcsOutputResult
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*status.Status)(nil),            // 8: google.rpc.Status
}
var file_google_cloud_retail_v2alpha_export_config_proto_depIdxs = []int32{
	7,  // 0: google.cloud.retail.v2alpha.ExportMetadata.create_time:type_name -> google.protobuf.Timestamp
	7,  // 1: google.cloud.retail.v2alpha.ExportMetadata.update_time:type_name -> google.protobuf.Timestamp
	8,  // 2: google.cloud.retail.v2alpha.ExportProductsResponse.error_samples:type_name -> google.rpc.Status
	0,  // 3: google.cloud.retail.v2alpha.ExportProductsResponse.errors_config:type_name -> google.cloud.retail.v2alpha.ExportErrorsConfig
	4,  // 4: google.cloud.retail.v2alpha.ExportProductsResponse.output_result:type_name -> google.cloud.retail.v2alpha.OutputResult
	8,  // 5: google.cloud.retail.v2alpha.ExportUserEventsResponse.error_samples:type_name -> google.rpc.Status
	0,  // 6: google.cloud.retail.v2alpha.ExportUserEventsResponse.errors_config:type_name -> google.cloud.retail.v2alpha.ExportErrorsConfig
	4,  // 7: google.cloud.retail.v2alpha.ExportUserEventsResponse.output_result:type_name -> google.cloud.retail.v2alpha.OutputResult
	5,  // 8: google.cloud.retail.v2alpha.OutputResult.bigquery_result:type_name -> google.cloud.retail.v2alpha.BigQueryOutputResult
	6,  // 9: google.cloud.retail.v2alpha.OutputResult.gcs_result:type_name -> google.cloud.retail.v2alpha.GcsOutputResult
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_export_config_proto_init() }
func file_google_cloud_retail_v2alpha_export_config_proto_init() {
	if File_google_cloud_retail_v2alpha_export_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportErrorsConfig); i {
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
		file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetadata); i {
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
		file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportProductsResponse); i {
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
		file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportUserEventsResponse); i {
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
		file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputResult); i {
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
		file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryOutputResult); i {
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
		file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsOutputResult); i {
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
	file_google_cloud_retail_v2alpha_export_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExportErrorsConfig_GcsPrefix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_retail_v2alpha_export_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_export_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_export_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2alpha_export_config_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_export_config_proto = out.File
	file_google_cloud_retail_v2alpha_export_config_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_export_config_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_export_config_proto_depIdxs = nil
}
