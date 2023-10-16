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
// source: google/maps/routing/v2/route_modifiers.proto

package routingpb

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

// Encapsulates a set of optional conditions to satisfy when calculating the
// routes.
type RouteModifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When set to true, avoids toll roads where reasonable, giving preference to
	// routes not containing toll roads. Applies only to the `DRIVE` and
	// `TWO_WHEELER` [RouteTravelMode][google.maps.routing.v2.RouteTravelMode].
	AvoidTolls bool `protobuf:"varint,1,opt,name=avoid_tolls,json=avoidTolls,proto3" json:"avoid_tolls,omitempty"`
	// When set to true, avoids highways where reasonable, giving preference to
	// routes not containing highways. Applies only to the `DRIVE` and
	// `TWO_WHEELER` [RouteTravelMode][google.maps.routing.v2.RouteTravelMode].
	AvoidHighways bool `protobuf:"varint,2,opt,name=avoid_highways,json=avoidHighways,proto3" json:"avoid_highways,omitempty"`
	// When set to true, avoids ferries where reasonable, giving preference to
	// routes not containing ferries. Applies only to the `DRIVE` and`TWO_WHEELER`
	// [RouteTravelMode][google.maps.routing.v2.RouteTravelMode].
	AvoidFerries bool `protobuf:"varint,3,opt,name=avoid_ferries,json=avoidFerries,proto3" json:"avoid_ferries,omitempty"`
	// When set to true, avoids navigating indoors where reasonable, giving
	// preference to routes not containing indoor navigation. Applies only to the
	// `WALK` [RouteTravelMode][google.maps.routing.v2.RouteTravelMode].
	AvoidIndoor bool `protobuf:"varint,4,opt,name=avoid_indoor,json=avoidIndoor,proto3" json:"avoid_indoor,omitempty"`
	// Specifies the vehicle information.
	VehicleInfo *VehicleInfo `protobuf:"bytes,5,opt,name=vehicle_info,json=vehicleInfo,proto3" json:"vehicle_info,omitempty"`
	// Encapsulates information about toll passes.
	// If toll passes are provided, the API tries to return the pass price. If
	// toll passes are not provided, the API treats the toll pass as unknown and
	// tries to return the cash price.
	// Applies only to the `DRIVE` and `TWO_WHEELER`
	// [RouteTravelMode][google.maps.routing.v2.RouteTravelMode].
	TollPasses []TollPass `protobuf:"varint,6,rep,packed,name=toll_passes,json=tollPasses,proto3,enum=google.maps.routing.v2.TollPass" json:"toll_passes,omitempty"`
}

func (x *RouteModifiers) Reset() {
	*x = RouteModifiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routing_v2_route_modifiers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteModifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteModifiers) ProtoMessage() {}

func (x *RouteModifiers) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_route_modifiers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteModifiers.ProtoReflect.Descriptor instead.
func (*RouteModifiers) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_route_modifiers_proto_rawDescGZIP(), []int{0}
}

func (x *RouteModifiers) GetAvoidTolls() bool {
	if x != nil {
		return x.AvoidTolls
	}
	return false
}

func (x *RouteModifiers) GetAvoidHighways() bool {
	if x != nil {
		return x.AvoidHighways
	}
	return false
}

func (x *RouteModifiers) GetAvoidFerries() bool {
	if x != nil {
		return x.AvoidFerries
	}
	return false
}

func (x *RouteModifiers) GetAvoidIndoor() bool {
	if x != nil {
		return x.AvoidIndoor
	}
	return false
}

func (x *RouteModifiers) GetVehicleInfo() *VehicleInfo {
	if x != nil {
		return x.VehicleInfo
	}
	return nil
}

func (x *RouteModifiers) GetTollPasses() []TollPass {
	if x != nil {
		return x.TollPasses
	}
	return nil
}

var File_google_maps_routing_v2_route_modifiers_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_route_modifiers_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d,
	0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x74,
	0x6f, 0x6c, 0x6c, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x54, 0x6f, 0x6c, 0x6c, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x48, 0x69,
	0x67, 0x68, 0x77, 0x61, 0x79, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f,
	0x66, 0x65, 0x72, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61,
	0x76, 0x6f, 0x69, 0x64, 0x46, 0x65, 0x72, 0x72, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x76, 0x6f, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x6f, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x49, 0x6e, 0x64, 0x6f, 0x6f, 0x72, 0x12, 0x46,
	0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x6c, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x6f, 0x6c, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x52, 0x0a, 0x74,
	0x6f, 0x6c, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x65, 0x73, 0x42, 0xc8, 0x01, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x13, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70,
	0x62, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x05, 0x47, 0x4d, 0x52, 0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca,
	0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_route_modifiers_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_route_modifiers_proto_rawDescData = file_google_maps_routing_v2_route_modifiers_proto_rawDesc
)

func file_google_maps_routing_v2_route_modifiers_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_route_modifiers_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_route_modifiers_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_route_modifiers_proto_rawDescData)
	})
	return file_google_maps_routing_v2_route_modifiers_proto_rawDescData
}

var file_google_maps_routing_v2_route_modifiers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_routing_v2_route_modifiers_proto_goTypes = []interface{}{
	(*RouteModifiers)(nil), // 0: google.maps.routing.v2.RouteModifiers
	(*VehicleInfo)(nil),    // 1: google.maps.routing.v2.VehicleInfo
	(TollPass)(0),          // 2: google.maps.routing.v2.TollPass
}
var file_google_maps_routing_v2_route_modifiers_proto_depIdxs = []int32{
	1, // 0: google.maps.routing.v2.RouteModifiers.vehicle_info:type_name -> google.maps.routing.v2.VehicleInfo
	2, // 1: google.maps.routing.v2.RouteModifiers.toll_passes:type_name -> google.maps.routing.v2.TollPass
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_route_modifiers_proto_init() }
func file_google_maps_routing_v2_route_modifiers_proto_init() {
	if File_google_maps_routing_v2_route_modifiers_proto != nil {
		return
	}
	file_google_maps_routing_v2_toll_passes_proto_init()
	file_google_maps_routing_v2_vehicle_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routing_v2_route_modifiers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteModifiers); i {
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
			RawDescriptor: file_google_maps_routing_v2_route_modifiers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_route_modifiers_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_route_modifiers_proto_depIdxs,
		MessageInfos:      file_google_maps_routing_v2_route_modifiers_proto_msgTypes,
	}.Build()
	File_google_maps_routing_v2_route_modifiers_proto = out.File
	file_google_maps_routing_v2_route_modifiers_proto_rawDesc = nil
	file_google_maps_routing_v2_route_modifiers_proto_goTypes = nil
	file_google_maps_routing_v2_route_modifiers_proto_depIdxs = nil
}
