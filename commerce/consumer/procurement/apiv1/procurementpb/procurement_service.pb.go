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
// source: google/cloud/commerce/consumer/procurement/v1/procurement_service.proto

package procurementpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [ConsumerProcurementService.PlaceOrder][google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.PlaceOrder].
type PlaceOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the parent resource.
	// This field has the form  `billingAccounts/{billing-account-id}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The user-specified name of the order being placed.
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. Places order for offer. Required when an offer-based order is
	// being placed.
	LineItemInfo []*LineItemInfo `protobuf:"bytes,10,rep,name=line_item_info,json=lineItemInfo,proto3" json:"line_item_info,omitempty"`
	// Optional. A unique identifier for this request.
	// The server will ignore subsequent requests that provide a duplicate request
	// ID for at least 120 minutes after the first request.
	//
	// The request ID must be a valid
	// [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier#Format).
	RequestId string `protobuf:"bytes,7,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *PlaceOrderRequest) Reset() {
	*x = PlaceOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderRequest) ProtoMessage() {}

func (x *PlaceOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderRequest.ProtoReflect.Descriptor instead.
func (*PlaceOrderRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescGZIP(), []int{0}
}

func (x *PlaceOrderRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PlaceOrderRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *PlaceOrderRequest) GetLineItemInfo() []*LineItemInfo {
	if x != nil {
		return x.LineItemInfo
	}
	return nil
}

func (x *PlaceOrderRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

// Message stored in the metadata field of the Operation returned by
// [ConsumerProcurementService.PlaceOrder][google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.PlaceOrder].
type PlaceOrderMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlaceOrderMetadata) Reset() {
	*x = PlaceOrderMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderMetadata) ProtoMessage() {}

func (x *PlaceOrderMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderMetadata.ProtoReflect.Descriptor instead.
func (*PlaceOrderMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescGZIP(), []int{1}
}

// Request message for
// [ConsumerProcurementService.GetOrder][google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.GetOrder]
type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the order to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for
// [ConsumerProcurementService.ListOrders][google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.ListOrders].
type ListOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource to query for orders.
	// This field has the form `billingAccounts/{billing-account-id}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of entries requested.
	// The default page size is 25 and the maximum page size is 200.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The token for fetching the next page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Filter that you can use to limit the list request.
	//
	// A query string that can match a selected set of attributes
	// with string values. For example, `display_name=abc`.
	// Supported query attributes are
	//
	// * `display_name`
	//
	// If the query contains special characters other than letters,
	// underscore, or digits, the phrase must be quoted with double quotes. For
	// example, `display_name="foo:bar"`, where the display name needs to be
	// quoted because it contains special character colon.
	//
	// Queries can be combined with `OR`, and `NOT` to form more complex queries.
	// You can also group them to force a desired evaluation order.
	// For example, `display_name=abc OR display_name=def`.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrdersRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListOrdersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListOrdersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListOrdersRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// Response message for
// [ConsumerProcurementService.ListOrders][google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.ListOrders].
type ListOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of orders in this response.
	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	// The token for fetching the next page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *ListOrdersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto protoreflect.FileDescriptor

var file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x63,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x11,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4a, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x32, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x66, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0c, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xcc, 0x05, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc3, 0x01, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0xca, 0x41, 0x1b, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x12, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x22,
	0x2b, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0xb6, 0x01, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0x34, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xc9, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0xda, 0x41, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x63, 0xca, 0x41, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x63, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x9f, 0x02, 0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x53,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0xaa, 0x02, 0x2d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x2d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x5c, 0x50, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x32, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescOnce sync.Once
	file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescData = file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDesc
)

func file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescGZIP() []byte {
	file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescData)
	})
	return file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDescData
}

var file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_goTypes = []interface{}{
	(*PlaceOrderRequest)(nil),       // 0: google.cloud.commerce.consumer.procurement.v1.PlaceOrderRequest
	(*PlaceOrderMetadata)(nil),      // 1: google.cloud.commerce.consumer.procurement.v1.PlaceOrderMetadata
	(*GetOrderRequest)(nil),         // 2: google.cloud.commerce.consumer.procurement.v1.GetOrderRequest
	(*ListOrdersRequest)(nil),       // 3: google.cloud.commerce.consumer.procurement.v1.ListOrdersRequest
	(*ListOrdersResponse)(nil),      // 4: google.cloud.commerce.consumer.procurement.v1.ListOrdersResponse
	(*LineItemInfo)(nil),            // 5: google.cloud.commerce.consumer.procurement.v1.LineItemInfo
	(*Order)(nil),                   // 6: google.cloud.commerce.consumer.procurement.v1.Order
	(*longrunningpb.Operation)(nil), // 7: google.longrunning.Operation
}
var file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_depIdxs = []int32{
	5, // 0: google.cloud.commerce.consumer.procurement.v1.PlaceOrderRequest.line_item_info:type_name -> google.cloud.commerce.consumer.procurement.v1.LineItemInfo
	6, // 1: google.cloud.commerce.consumer.procurement.v1.ListOrdersResponse.orders:type_name -> google.cloud.commerce.consumer.procurement.v1.Order
	0, // 2: google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.PlaceOrder:input_type -> google.cloud.commerce.consumer.procurement.v1.PlaceOrderRequest
	2, // 3: google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.GetOrder:input_type -> google.cloud.commerce.consumer.procurement.v1.GetOrderRequest
	3, // 4: google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.ListOrders:input_type -> google.cloud.commerce.consumer.procurement.v1.ListOrdersRequest
	7, // 5: google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.PlaceOrder:output_type -> google.longrunning.Operation
	6, // 6: google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.GetOrder:output_type -> google.cloud.commerce.consumer.procurement.v1.Order
	4, // 7: google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService.ListOrders:output_type -> google.cloud.commerce.consumer.procurement.v1.ListOrdersResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_init() }
func file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_init() {
	if File_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto != nil {
		return
	}
	file_google_cloud_commerce_consumer_procurement_v1_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderRequest); i {
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
		file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderMetadata); i {
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
		file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderRequest); i {
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
		file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrdersRequest); i {
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
		file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrdersResponse); i {
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
			RawDescriptor: file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_msgTypes,
	}.Build()
	File_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto = out.File
	file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_rawDesc = nil
	file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_goTypes = nil
	file_google_cloud_commerce_consumer_procurement_v1_procurement_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsumerProcurementServiceClient is the client API for ConsumerProcurementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsumerProcurementServiceClient interface {
	// Creates a new [Order][google.cloud.commerce.consumer.procurement.v1.Order].
	//
	// This API only supports GCP spend-based committed use
	// discounts specified by GCP documentation.
	//
	// The returned long-running operation is in-progress until the backend
	// completes the creation of the resource. Once completed, the order is
	// in
	// [OrderState.ORDER_STATE_ACTIVE][google.cloud.commerce.consumer.procurement.v1.OrderState.ORDER_STATE_ACTIVE].
	// In case of failure, the order resource will be removed.
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Returns the requested
	// [Order][google.cloud.commerce.consumer.procurement.v1.Order] resource.
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
	// Lists [Order][google.cloud.commerce.consumer.procurement.v1.Order]
	// resources that the user has access to, within the scope of the parent
	// resource.
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
}

type consumerProcurementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerProcurementServiceClient(cc grpc.ClientConnInterface) ConsumerProcurementServiceClient {
	return &consumerProcurementServiceClient{cc}
}

func (c *consumerProcurementServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerProcurementServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerProcurementServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService/ListOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerProcurementServiceServer is the server API for ConsumerProcurementService service.
type ConsumerProcurementServiceServer interface {
	// Creates a new [Order][google.cloud.commerce.consumer.procurement.v1.Order].
	//
	// This API only supports GCP spend-based committed use
	// discounts specified by GCP documentation.
	//
	// The returned long-running operation is in-progress until the backend
	// completes the creation of the resource. Once completed, the order is
	// in
	// [OrderState.ORDER_STATE_ACTIVE][google.cloud.commerce.consumer.procurement.v1.OrderState.ORDER_STATE_ACTIVE].
	// In case of failure, the order resource will be removed.
	PlaceOrder(context.Context, *PlaceOrderRequest) (*longrunningpb.Operation, error)
	// Returns the requested
	// [Order][google.cloud.commerce.consumer.procurement.v1.Order] resource.
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	// Lists [Order][google.cloud.commerce.consumer.procurement.v1.Order]
	// resources that the user has access to, within the scope of the parent
	// resource.
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
}

// UnimplementedConsumerProcurementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsumerProcurementServiceServer struct {
}

func (*UnimplementedConsumerProcurementServiceServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (*UnimplementedConsumerProcurementServiceServer) GetOrder(context.Context, *GetOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (*UnimplementedConsumerProcurementServiceServer) ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}

func RegisterConsumerProcurementServiceServer(s *grpc.Server, srv ConsumerProcurementServiceServer) {
	s.RegisterService(&_ConsumerProcurementService_serviceDesc, srv)
}

func _ConsumerProcurementService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerProcurementServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerProcurementServiceServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerProcurementService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerProcurementServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerProcurementServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerProcurementService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerProcurementServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerProcurementServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsumerProcurementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.commerce.consumer.procurement.v1.ConsumerProcurementService",
	HandlerType: (*ConsumerProcurementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _ConsumerProcurementService_PlaceOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _ConsumerProcurementService_GetOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _ConsumerProcurementService_ListOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/commerce/consumer/procurement/v1/procurement_service.proto",
}
