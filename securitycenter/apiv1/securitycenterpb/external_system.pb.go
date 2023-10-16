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
// source: google/cloud/securitycenter/v1/external_system.proto

package securitycenterpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Representation of third party SIEM/SOAR fields within SCC.
type ExternalSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full resource name of the external system, for example:
	// "organizations/1234/sources/5678/findings/123456/externalSystems/jira",
	// "folders/1234/sources/5678/findings/123456/externalSystems/jira",
	// "projects/1234/sources/5678/findings/123456/externalSystems/jira"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// References primary/secondary etc assignees in the external system.
	Assignees []string `protobuf:"bytes,2,rep,name=assignees,proto3" json:"assignees,omitempty"`
	// Identifier that's used to track the given finding in the external system.
	ExternalUid string `protobuf:"bytes,3,opt,name=external_uid,json=externalUid,proto3" json:"external_uid,omitempty"`
	// Most recent status of the corresponding finding's ticket/tracker in the
	// external system.
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// The most recent time when the corresponding finding's ticket/tracker was
	// updated in the external system.
	ExternalSystemUpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=external_system_update_time,json=externalSystemUpdateTime,proto3" json:"external_system_update_time,omitempty"`
}

func (x *ExternalSystem) Reset() {
	*x = ExternalSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_external_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalSystem) ProtoMessage() {}

func (x *ExternalSystem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_external_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalSystem.ProtoReflect.Descriptor instead.
func (*ExternalSystem) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_external_system_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalSystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExternalSystem) GetAssignees() []string {
	if x != nil {
		return x.Assignees
	}
	return nil
}

func (x *ExternalSystem) GetExternalUid() string {
	if x != nil {
		return x.ExternalUid
	}
	return ""
}

func (x *ExternalSystem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ExternalSystem) GetExternalSystemUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExternalSystemUpdateTime
	}
	return nil
}

var File_google_cloud_securitycenter_v1_external_system_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_external_system_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa0, 0x04, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x59, 0x0a, 0x1b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x18, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0xc5, 0x02,
	0xea, 0x41, 0xc1, 0x02, 0x0a, 0x2c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x61, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x7d, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x7d, 0x12, 0x55, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x7b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x7d, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x7d, 0x12, 0x57, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x7d, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x7d, 0x42, 0xed, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_external_system_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_external_system_proto_rawDescData = file_google_cloud_securitycenter_v1_external_system_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_external_system_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_external_system_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_external_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_external_system_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_external_system_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_external_system_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_external_system_proto_goTypes = []interface{}{
	(*ExternalSystem)(nil),        // 0: google.cloud.securitycenter.v1.ExternalSystem
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_google_cloud_securitycenter_v1_external_system_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.ExternalSystem.external_system_update_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_external_system_proto_init() }
func file_google_cloud_securitycenter_v1_external_system_proto_init() {
	if File_google_cloud_securitycenter_v1_external_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_external_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalSystem); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_external_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_external_system_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_external_system_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_external_system_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_external_system_proto = out.File
	file_google_cloud_securitycenter_v1_external_system_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_external_system_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_external_system_proto_depIdxs = nil
}
