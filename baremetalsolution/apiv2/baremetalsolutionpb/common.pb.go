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
// source: google/cloud/baremetalsolution/v2/common.proto

package baremetalsolutionpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Performance tier of the Volume.
type VolumePerformanceTier int32

const (
	// Value is not specified.
	VolumePerformanceTier_VOLUME_PERFORMANCE_TIER_UNSPECIFIED VolumePerformanceTier = 0
	// Regular volumes, shared aggregates.
	VolumePerformanceTier_VOLUME_PERFORMANCE_TIER_SHARED VolumePerformanceTier = 1
	// Assigned aggregates.
	VolumePerformanceTier_VOLUME_PERFORMANCE_TIER_ASSIGNED VolumePerformanceTier = 2
	// High throughput aggregates.
	VolumePerformanceTier_VOLUME_PERFORMANCE_TIER_HT VolumePerformanceTier = 3
)

// Enum value maps for VolumePerformanceTier.
var (
	VolumePerformanceTier_name = map[int32]string{
		0: "VOLUME_PERFORMANCE_TIER_UNSPECIFIED",
		1: "VOLUME_PERFORMANCE_TIER_SHARED",
		2: "VOLUME_PERFORMANCE_TIER_ASSIGNED",
		3: "VOLUME_PERFORMANCE_TIER_HT",
	}
	VolumePerformanceTier_value = map[string]int32{
		"VOLUME_PERFORMANCE_TIER_UNSPECIFIED": 0,
		"VOLUME_PERFORMANCE_TIER_SHARED":      1,
		"VOLUME_PERFORMANCE_TIER_ASSIGNED":    2,
		"VOLUME_PERFORMANCE_TIER_HT":          3,
	}
)

func (x VolumePerformanceTier) Enum() *VolumePerformanceTier {
	p := new(VolumePerformanceTier)
	*p = x
	return p
}

func (x VolumePerformanceTier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VolumePerformanceTier) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_baremetalsolution_v2_common_proto_enumTypes[0].Descriptor()
}

func (VolumePerformanceTier) Type() protoreflect.EnumType {
	return &file_google_cloud_baremetalsolution_v2_common_proto_enumTypes[0]
}

func (x VolumePerformanceTier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VolumePerformanceTier.Descriptor instead.
func (VolumePerformanceTier) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_common_proto_rawDescGZIP(), []int{0}
}

// The possible values for a workload profile.
type WorkloadProfile int32

const (
	// The workload profile is in an unknown state.
	WorkloadProfile_WORKLOAD_PROFILE_UNSPECIFIED WorkloadProfile = 0
	// The workload profile is generic.
	WorkloadProfile_WORKLOAD_PROFILE_GENERIC WorkloadProfile = 1
	// The workload profile is hana.
	WorkloadProfile_WORKLOAD_PROFILE_HANA WorkloadProfile = 2
)

// Enum value maps for WorkloadProfile.
var (
	WorkloadProfile_name = map[int32]string{
		0: "WORKLOAD_PROFILE_UNSPECIFIED",
		1: "WORKLOAD_PROFILE_GENERIC",
		2: "WORKLOAD_PROFILE_HANA",
	}
	WorkloadProfile_value = map[string]int32{
		"WORKLOAD_PROFILE_UNSPECIFIED": 0,
		"WORKLOAD_PROFILE_GENERIC":     1,
		"WORKLOAD_PROFILE_HANA":        2,
	}
)

func (x WorkloadProfile) Enum() *WorkloadProfile {
	p := new(WorkloadProfile)
	*p = x
	return p
}

func (x WorkloadProfile) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkloadProfile) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_baremetalsolution_v2_common_proto_enumTypes[1].Descriptor()
}

func (WorkloadProfile) Type() protoreflect.EnumType {
	return &file_google_cloud_baremetalsolution_v2_common_proto_enumTypes[1]
}

func (x WorkloadProfile) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkloadProfile.Descriptor instead.
func (WorkloadProfile) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_common_proto_rawDescGZIP(), []int{1}
}

var File_google_cloud_baremetalsolution_v2_common_proto protoreflect.FileDescriptor

var file_google_cloud_baremetalsolution_v2_common_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2a, 0xaa, 0x01, 0x0a, 0x15, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x69, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x23, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x49, 0x45,
	0x52, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x56, 0x4f,
	0x4c, 0x55, 0x4d, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45,
	0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x1e, 0x0a, 0x1a, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x48, 0x54, 0x10, 0x03,
	0x2a, 0x6c, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x57, 0x4f, 0x52, 0x4b, 0x4c, 0x4f, 0x41, 0x44, 0x5f,
	0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x57, 0x4f, 0x52, 0x4b, 0x4c, 0x4f, 0x41,
	0x44, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49,
	0x43, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x4f, 0x52, 0x4b, 0x4c, 0x4f, 0x41, 0x44, 0x5f,
	0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x48, 0x41, 0x4e, 0x41, 0x10, 0x02, 0x42, 0xfa,
	0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x72,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x61, 0x72, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32,
	0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_baremetalsolution_v2_common_proto_rawDescOnce sync.Once
	file_google_cloud_baremetalsolution_v2_common_proto_rawDescData = file_google_cloud_baremetalsolution_v2_common_proto_rawDesc
)

func file_google_cloud_baremetalsolution_v2_common_proto_rawDescGZIP() []byte {
	file_google_cloud_baremetalsolution_v2_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_baremetalsolution_v2_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_baremetalsolution_v2_common_proto_rawDescData)
	})
	return file_google_cloud_baremetalsolution_v2_common_proto_rawDescData
}

var file_google_cloud_baremetalsolution_v2_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_baremetalsolution_v2_common_proto_goTypes = []interface{}{
	(VolumePerformanceTier)(0), // 0: google.cloud.baremetalsolution.v2.VolumePerformanceTier
	(WorkloadProfile)(0),       // 1: google.cloud.baremetalsolution.v2.WorkloadProfile
}
var file_google_cloud_baremetalsolution_v2_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_baremetalsolution_v2_common_proto_init() }
func file_google_cloud_baremetalsolution_v2_common_proto_init() {
	if File_google_cloud_baremetalsolution_v2_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_baremetalsolution_v2_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_baremetalsolution_v2_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_baremetalsolution_v2_common_proto_depIdxs,
		EnumInfos:         file_google_cloud_baremetalsolution_v2_common_proto_enumTypes,
	}.Build()
	File_google_cloud_baremetalsolution_v2_common_proto = out.File
	file_google_cloud_baremetalsolution_v2_common_proto_rawDesc = nil
	file_google_cloud_baremetalsolution_v2_common_proto_goTypes = nil
	file_google_cloud_baremetalsolution_v2_common_proto_depIdxs = nil
}
