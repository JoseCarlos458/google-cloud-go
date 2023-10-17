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
// source: google/cloud/asset/v1p5beta1/asset_service.proto

package assetpb

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Asset content type.
type ContentType int32

const (
	// Unspecified content type.
	ContentType_CONTENT_TYPE_UNSPECIFIED ContentType = 0
	// Resource metadata.
	ContentType_RESOURCE ContentType = 1
	// The actual IAM policy set on a resource.
	ContentType_IAM_POLICY ContentType = 2
	// The organization policy set on an asset.
	ContentType_ORG_POLICY ContentType = 4
	// The Access Context Manager policy set on an asset.
	ContentType_ACCESS_POLICY ContentType = 5
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0: "CONTENT_TYPE_UNSPECIFIED",
		1: "RESOURCE",
		2: "IAM_POLICY",
		4: "ORG_POLICY",
		5: "ACCESS_POLICY",
	}
	ContentType_value = map[string]int32{
		"CONTENT_TYPE_UNSPECIFIED": 0,
		"RESOURCE":                 1,
		"IAM_POLICY":               2,
		"ORG_POLICY":               4,
		"ACCESS_POLICY":            5,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_asset_v1p5beta1_asset_service_proto_enumTypes[0].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_google_cloud_asset_v1p5beta1_asset_service_proto_enumTypes[0]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescGZIP(), []int{0}
}

// ListAssets request.
type ListAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the organization or project the assets belong to. Format:
	// "organizations/[organization-number]" (such as "organizations/123"),
	// "projects/[project-id]" (such as "projects/my-project-id"), or
	// "projects/[project-number]" (such as "projects/12345").
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Timestamp to take an asset snapshot. This can only be set to a timestamp
	// between the current time and the current time minus 35 days (inclusive).
	// If not specified, the current time will be used. Due to delays in resource
	// data collection and indexing, there is a volatile window during which
	// running the same query may get different results.
	ReadTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// A list of asset types to take a snapshot for. For example:
	// "compute.googleapis.com/Disk".
	//
	// Regular expression is also supported. For example:
	//
	// * "compute.googleapis.com.*" snapshots resources whose asset type starts
	// with "compute.googleapis.com".
	// * ".*Instance" snapshots resources whose asset type ends with "Instance".
	// * ".*Instance.*" snapshots resources whose asset type contains "Instance".
	//
	// See [RE2](https://github.com/google/re2/wiki/Syntax) for all supported
	// regular expression syntax. If the regular expression does not match any
	// supported asset type, an INVALID_ARGUMENT error will be returned.
	//
	// If specified, only matching assets will be returned, otherwise, it will
	// snapshot all asset types. See [Introduction to Cloud Asset
	// Inventory](https://cloud.google.com/asset-inventory/docs/overview)
	// for all supported asset types.
	AssetTypes []string `protobuf:"bytes,3,rep,name=asset_types,json=assetTypes,proto3" json:"asset_types,omitempty"`
	// Asset content type. If not specified, no content but the asset name will
	// be returned.
	ContentType ContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=google.cloud.asset.v1p5beta1.ContentType" json:"content_type,omitempty"`
	// The maximum number of assets to be returned in a single response. Default
	// is 100, minimum is 1, and maximum is 1000.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The `next_page_token` returned from the previous `ListAssetsResponse`, or
	// unspecified for the first `ListAssetsRequest`. It is a continuation of a
	// prior `ListAssets` call, and the API should return the next page of assets.
	PageToken string `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListAssetsRequest) Reset() {
	*x = ListAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsRequest) ProtoMessage() {}

func (x *ListAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsRequest.ProtoReflect.Descriptor instead.
func (*ListAssetsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListAssetsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListAssetsRequest) GetReadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadTime
	}
	return nil
}

func (x *ListAssetsRequest) GetAssetTypes() []string {
	if x != nil {
		return x.AssetTypes
	}
	return nil
}

func (x *ListAssetsRequest) GetContentType() ContentType {
	if x != nil {
		return x.ContentType
	}
	return ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (x *ListAssetsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAssetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListAssets response.
type ListAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time the snapshot was taken.
	ReadTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
	// Assets.
	Assets []*Asset `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets,omitempty"`
	// Token to retrieve the next page of results. It expires 72 hours after the
	// page token for the first page is generated. Set to empty if there are no
	// remaining results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListAssetsResponse) Reset() {
	*x = ListAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsResponse) ProtoMessage() {}

func (x *ListAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsResponse.ProtoReflect.Descriptor instead.
func (*ListAssetsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListAssetsResponse) GetReadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadTime
	}
	return nil
}

func (x *ListAssetsResponse) GetAssets() []*Asset {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *ListAssetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_asset_v1p5beta1_asset_service_proto protoreflect.FileDescriptor

var file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb8, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x21, 0x12, 0x1f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x4c, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x35,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb2, 0x01, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2a,
	0x6c, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x41,
	0x4d, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x52,
	0x47, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x05, 0x32, 0x80, 0x02,
	0x0a, 0x0c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0,
	0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x2f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2f, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x2a, 0x2f, 0x2a, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x1a, 0x4d, 0xca, 0x41, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x42, 0xad, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x70, 0x35,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x11, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x70, 0x62, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x70, 0x62, 0xaa, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x56, 0x31, 0x50, 0x35, 0x42, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x70, 0x35, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescOnce sync.Once
	file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescData = file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDesc
)

func file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescGZIP() []byte {
	file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescData)
	})
	return file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDescData
}

var file_google_cloud_asset_v1p5beta1_asset_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_asset_v1p5beta1_asset_service_proto_goTypes = []interface{}{
	(ContentType)(0),              // 0: google.cloud.asset.v1p5beta1.ContentType
	(*ListAssetsRequest)(nil),     // 1: google.cloud.asset.v1p5beta1.ListAssetsRequest
	(*ListAssetsResponse)(nil),    // 2: google.cloud.asset.v1p5beta1.ListAssetsResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Asset)(nil),                 // 4: google.cloud.asset.v1p5beta1.Asset
}
var file_google_cloud_asset_v1p5beta1_asset_service_proto_depIdxs = []int32{
	3, // 0: google.cloud.asset.v1p5beta1.ListAssetsRequest.read_time:type_name -> google.protobuf.Timestamp
	0, // 1: google.cloud.asset.v1p5beta1.ListAssetsRequest.content_type:type_name -> google.cloud.asset.v1p5beta1.ContentType
	3, // 2: google.cloud.asset.v1p5beta1.ListAssetsResponse.read_time:type_name -> google.protobuf.Timestamp
	4, // 3: google.cloud.asset.v1p5beta1.ListAssetsResponse.assets:type_name -> google.cloud.asset.v1p5beta1.Asset
	1, // 4: google.cloud.asset.v1p5beta1.AssetService.ListAssets:input_type -> google.cloud.asset.v1p5beta1.ListAssetsRequest
	2, // 5: google.cloud.asset.v1p5beta1.AssetService.ListAssets:output_type -> google.cloud.asset.v1p5beta1.ListAssetsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_asset_v1p5beta1_asset_service_proto_init() }
func file_google_cloud_asset_v1p5beta1_asset_service_proto_init() {
	if File_google_cloud_asset_v1p5beta1_asset_service_proto != nil {
		return
	}
	file_google_cloud_asset_v1p5beta1_assets_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsRequest); i {
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
		file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsResponse); i {
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
			RawDescriptor: file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_asset_v1p5beta1_asset_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_asset_v1p5beta1_asset_service_proto_depIdxs,
		EnumInfos:         file_google_cloud_asset_v1p5beta1_asset_service_proto_enumTypes,
		MessageInfos:      file_google_cloud_asset_v1p5beta1_asset_service_proto_msgTypes,
	}.Build()
	File_google_cloud_asset_v1p5beta1_asset_service_proto = out.File
	file_google_cloud_asset_v1p5beta1_asset_service_proto_rawDesc = nil
	file_google_cloud_asset_v1p5beta1_asset_service_proto_goTypes = nil
	file_google_cloud_asset_v1p5beta1_asset_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Lists assets with time and resource types and returns paged results in
	// response.
	ListAssets(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) ListAssets(ctx context.Context, in *ListAssetsRequest, opts ...grpc.CallOption) (*ListAssetsResponse, error) {
	out := new(ListAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.asset.v1p5beta1.AssetService/ListAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
type AssetServiceServer interface {
	// Lists assets with time and resource types and returns paged results in
	// response.
	ListAssets(context.Context, *ListAssetsRequest) (*ListAssetsResponse, error)
}

// UnimplementedAssetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (*UnimplementedAssetServiceServer) ListAssets(context.Context, *ListAssetsRequest) (*ListAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAssets not implemented")
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_ListAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).ListAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.asset.v1p5beta1.AssetService/ListAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).ListAssets(ctx, req.(*ListAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.asset.v1p5beta1.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAssets",
			Handler:    _AssetService_ListAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/asset/v1p5beta1/asset_service.proto",
}
