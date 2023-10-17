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
// source: google/api/servicecontrol/v1/operation.proto

package servicecontrolpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Defines the importance of the data contained in the operation.
type Operation_Importance int32

const (
	// Allows data caching, batching, and aggregation. It provides
	// higher performance with higher data loss risk.
	Operation_LOW Operation_Importance = 0
	// Disables data aggregation to minimize data loss. It is for operations
	// that contains significant monetary value or audit trail. This feature
	// only applies to the client libraries.
	Operation_HIGH Operation_Importance = 1
)

// Enum value maps for Operation_Importance.
var (
	Operation_Importance_name = map[int32]string{
		0: "LOW",
		1: "HIGH",
	}
	Operation_Importance_value = map[string]int32{
		"LOW":  0,
		"HIGH": 1,
	}
)

func (x Operation_Importance) Enum() *Operation_Importance {
	p := new(Operation_Importance)
	*p = x
	return p
}

func (x Operation_Importance) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation_Importance) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_servicecontrol_v1_operation_proto_enumTypes[0].Descriptor()
}

func (Operation_Importance) Type() protoreflect.EnumType {
	return &file_google_api_servicecontrol_v1_operation_proto_enumTypes[0]
}

func (x Operation_Importance) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation_Importance.Descriptor instead.
func (Operation_Importance) EnumDescriptor() ([]byte, []int) {
	return file_google_api_servicecontrol_v1_operation_proto_rawDescGZIP(), []int{0, 0}
}

// Represents information regarding an operation.
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity of the operation. This must be unique within the scope of the
	// service that generated the operation. If the service calls
	// Check() and Report() on the same operation, the two calls should carry
	// the same id.
	//
	// UUID version 4 is recommended, though not required.
	// In scenarios where an operation is computed from existing information
	// and an idempotent id is desirable for deduplication purpose, UUID version 5
	// is recommended. See RFC 4122 for details.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Fully qualified name of the operation. Reserved for future use.
	OperationName string `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	// Identity of the consumer who is using the service.
	// This field should be filled in for the operations initiated by a
	// consumer, but not for service-initiated operations that are
	// not related to a specific consumer.
	//
	// - This can be in one of the following formats:
	//     - project:PROJECT_ID,
	//     - project`_`number:PROJECT_NUMBER,
	//     - projects/PROJECT_ID or PROJECT_NUMBER,
	//     - folders/FOLDER_NUMBER,
	//     - organizations/ORGANIZATION_NUMBER,
	//     - api`_`key:API_KEY.
	ConsumerId string `protobuf:"bytes,3,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	// Required. Start time of the operation.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of the operation.
	// Required when the operation is used in
	// [ServiceController.Report][google.api.servicecontrol.v1.ServiceController.Report],
	// but optional when the operation is used in
	// [ServiceController.Check][google.api.servicecontrol.v1.ServiceController.Check].
	EndTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Labels describing the operation. Only the following labels are allowed:
	//
	// - Labels describing monitored resources as defined in
	//   the service configuration.
	// - Default labels of metric values. When specified, labels defined in the
	//   metric value override these default.
	// - The following labels defined by Google Cloud Platform:
	//     - `cloud.googleapis.com/location` describing the location where the
	//        operation happened,
	//     - `servicecontrol.googleapis.com/user_agent` describing the user agent
	//        of the API request,
	//     - `servicecontrol.googleapis.com/service_agent` describing the service
	//        used to handle the API request (e.g. ESP),
	//     - `servicecontrol.googleapis.com/platform` describing the platform
	//        where the API is served, such as App Engine, Compute Engine, or
	//        Kubernetes Engine.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Represents information about this operation. Each MetricValueSet
	// corresponds to a metric defined in the service configuration.
	// The data type used in the MetricValueSet must agree with
	// the data type specified in the metric definition.
	//
	// Within a single operation, it is not allowed to have more than one
	// MetricValue instances that have the same metric names and identical
	// label value combinations. If a request has such duplicated MetricValue
	// instances, the entire request is rejected with
	// an invalid argument error.
	MetricValueSets []*MetricValueSet `protobuf:"bytes,7,rep,name=metric_value_sets,json=metricValueSets,proto3" json:"metric_value_sets,omitempty"`
	// Represents information to be logged.
	LogEntries []*LogEntry `protobuf:"bytes,8,rep,name=log_entries,json=logEntries,proto3" json:"log_entries,omitempty"`
	// DO NOT USE. This is an experimental field.
	Importance Operation_Importance `protobuf:"varint,11,opt,name=importance,proto3,enum=google.api.servicecontrol.v1.Operation_Importance" json:"importance,omitempty"`
	// Unimplemented.
	Extensions []*anypb.Any `protobuf:"bytes,16,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_servicecontrol_v1_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_servicecontrol_v1_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_google_api_servicecontrol_v1_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *Operation) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

func (x *Operation) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *Operation) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Operation) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Operation) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Operation) GetMetricValueSets() []*MetricValueSet {
	if x != nil {
		return x.MetricValueSets
	}
	return nil
}

func (x *Operation) GetLogEntries() []*LogEntry {
	if x != nil {
		return x.LogEntries
	}
	return nil
}

func (x *Operation) GetImportance() Operation_Importance {
	if x != nil {
		return x.Importance
	}
	return Operation_LOW
}

func (x *Operation) GetExtensions() []*anypb.Any {
	if x != nil {
		return x.Extensions
	}
	return nil
}

var File_google_api_servicecontrol_v1_operation_proto protoreflect.FileDescriptor

var file_google_api_servicecontrol_v1_operation_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x05, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x58, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x0b, 0x6c, 0x6f, 0x67,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x52, 0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x42, 0xe9, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_servicecontrol_v1_operation_proto_rawDescOnce sync.Once
	file_google_api_servicecontrol_v1_operation_proto_rawDescData = file_google_api_servicecontrol_v1_operation_proto_rawDesc
)

func file_google_api_servicecontrol_v1_operation_proto_rawDescGZIP() []byte {
	file_google_api_servicecontrol_v1_operation_proto_rawDescOnce.Do(func() {
		file_google_api_servicecontrol_v1_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_servicecontrol_v1_operation_proto_rawDescData)
	})
	return file_google_api_servicecontrol_v1_operation_proto_rawDescData
}

var file_google_api_servicecontrol_v1_operation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_api_servicecontrol_v1_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_api_servicecontrol_v1_operation_proto_goTypes = []interface{}{
	(Operation_Importance)(0),     // 0: google.api.servicecontrol.v1.Operation.Importance
	(*Operation)(nil),             // 1: google.api.servicecontrol.v1.Operation
	nil,                           // 2: google.api.servicecontrol.v1.Operation.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*MetricValueSet)(nil),        // 4: google.api.servicecontrol.v1.MetricValueSet
	(*LogEntry)(nil),              // 5: google.api.servicecontrol.v1.LogEntry
	(*anypb.Any)(nil),             // 6: google.protobuf.Any
}
var file_google_api_servicecontrol_v1_operation_proto_depIdxs = []int32{
	3, // 0: google.api.servicecontrol.v1.Operation.start_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.api.servicecontrol.v1.Operation.end_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.api.servicecontrol.v1.Operation.labels:type_name -> google.api.servicecontrol.v1.Operation.LabelsEntry
	4, // 3: google.api.servicecontrol.v1.Operation.metric_value_sets:type_name -> google.api.servicecontrol.v1.MetricValueSet
	5, // 4: google.api.servicecontrol.v1.Operation.log_entries:type_name -> google.api.servicecontrol.v1.LogEntry
	0, // 5: google.api.servicecontrol.v1.Operation.importance:type_name -> google.api.servicecontrol.v1.Operation.Importance
	6, // 6: google.api.servicecontrol.v1.Operation.extensions:type_name -> google.protobuf.Any
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_api_servicecontrol_v1_operation_proto_init() }
func file_google_api_servicecontrol_v1_operation_proto_init() {
	if File_google_api_servicecontrol_v1_operation_proto != nil {
		return
	}
	file_google_api_servicecontrol_v1_log_entry_proto_init()
	file_google_api_servicecontrol_v1_metric_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_api_servicecontrol_v1_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_servicecontrol_v1_operation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_servicecontrol_v1_operation_proto_goTypes,
		DependencyIndexes: file_google_api_servicecontrol_v1_operation_proto_depIdxs,
		EnumInfos:         file_google_api_servicecontrol_v1_operation_proto_enumTypes,
		MessageInfos:      file_google_api_servicecontrol_v1_operation_proto_msgTypes,
	}.Build()
	File_google_api_servicecontrol_v1_operation_proto = out.File
	file_google_api_servicecontrol_v1_operation_proto_rawDesc = nil
	file_google_api_servicecontrol_v1_operation_proto_goTypes = nil
	file_google_api_servicecontrol_v1_operation_proto_depIdxs = nil
}
