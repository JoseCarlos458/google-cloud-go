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
// source: google/cloud/discoveryengine/v1beta/common.proto

package discoveryenginepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A floating point interval.
type Interval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The lower bound of the interval. If neither of the min fields are
	// set, then the lower bound is negative infinity.
	//
	// This field must be not larger than max.
	// Otherwise, an `INVALID_ARGUMENT` error is returned.
	//
	// Types that are assignable to Min:
	//	*Interval_Minimum
	//	*Interval_ExclusiveMinimum
	Min isInterval_Min `protobuf_oneof:"min"`
	// The upper bound of the interval. If neither of the max fields are
	// set, then the upper bound is positive infinity.
	//
	// This field must be not smaller than min.
	// Otherwise, an `INVALID_ARGUMENT` error is returned.
	//
	// Types that are assignable to Max:
	//	*Interval_Maximum
	//	*Interval_ExclusiveMaximum
	Max isInterval_Max `protobuf_oneof:"max"`
}

func (x *Interval) Reset() {
	*x = Interval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interval) ProtoMessage() {}

func (x *Interval) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interval.ProtoReflect.Descriptor instead.
func (*Interval) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_common_proto_rawDescGZIP(), []int{0}
}

func (m *Interval) GetMin() isInterval_Min {
	if m != nil {
		return m.Min
	}
	return nil
}

func (x *Interval) GetMinimum() float64 {
	if x, ok := x.GetMin().(*Interval_Minimum); ok {
		return x.Minimum
	}
	return 0
}

func (x *Interval) GetExclusiveMinimum() float64 {
	if x, ok := x.GetMin().(*Interval_ExclusiveMinimum); ok {
		return x.ExclusiveMinimum
	}
	return 0
}

func (m *Interval) GetMax() isInterval_Max {
	if m != nil {
		return m.Max
	}
	return nil
}

func (x *Interval) GetMaximum() float64 {
	if x, ok := x.GetMax().(*Interval_Maximum); ok {
		return x.Maximum
	}
	return 0
}

func (x *Interval) GetExclusiveMaximum() float64 {
	if x, ok := x.GetMax().(*Interval_ExclusiveMaximum); ok {
		return x.ExclusiveMaximum
	}
	return 0
}

type isInterval_Min interface {
	isInterval_Min()
}

type Interval_Minimum struct {
	// Inclusive lower bound.
	Minimum float64 `protobuf:"fixed64,1,opt,name=minimum,proto3,oneof"`
}

type Interval_ExclusiveMinimum struct {
	// Exclusive lower bound.
	ExclusiveMinimum float64 `protobuf:"fixed64,2,opt,name=exclusive_minimum,json=exclusiveMinimum,proto3,oneof"`
}

func (*Interval_Minimum) isInterval_Min() {}

func (*Interval_ExclusiveMinimum) isInterval_Min() {}

type isInterval_Max interface {
	isInterval_Max()
}

type Interval_Maximum struct {
	// Inclusive upper bound.
	Maximum float64 `protobuf:"fixed64,3,opt,name=maximum,proto3,oneof"`
}

type Interval_ExclusiveMaximum struct {
	// Exclusive upper bound.
	ExclusiveMaximum float64 `protobuf:"fixed64,4,opt,name=exclusive_maximum,json=exclusiveMaximum,proto3,oneof"`
}

func (*Interval_Maximum) isInterval_Max() {}

func (*Interval_ExclusiveMaximum) isInterval_Max() {}

// A custom attribute that is not explicitly modeled in a resource, e.g.
// [UserEvent][google.cloud.discoveryengine.v1beta.UserEvent].
type CustomAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The textual values of this custom attribute. For example, `["yellow",
	// "green"]` when the key is "color".
	//
	// Empty string is not allowed. Otherwise, an `INVALID_ARGUMENT` error is
	// returned.
	//
	// Exactly one of
	// [CustomAttribute.text][google.cloud.discoveryengine.v1beta.CustomAttribute.text]
	// or
	// [CustomAttribute.numbers][google.cloud.discoveryengine.v1beta.CustomAttribute.numbers]
	// should be set. Otherwise, an `INVALID_ARGUMENT` error is returned.
	Text []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	// The numerical values of this custom attribute. For example, `[2.3, 15.4]`
	// when the key is "lengths_cm".
	//
	// Exactly one of
	// [CustomAttribute.text][google.cloud.discoveryengine.v1beta.CustomAttribute.text]
	// or
	// [CustomAttribute.numbers][google.cloud.discoveryengine.v1beta.CustomAttribute.numbers]
	// should be set. Otherwise, an `INVALID_ARGUMENT` error is returned.
	Numbers []float64 `protobuf:"fixed64,2,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *CustomAttribute) Reset() {
	*x = CustomAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomAttribute) ProtoMessage() {}

func (x *CustomAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomAttribute.ProtoReflect.Descriptor instead.
func (*CustomAttribute) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_common_proto_rawDescGZIP(), []int{1}
}

func (x *CustomAttribute) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *CustomAttribute) GetNumbers() []float64 {
	if x != nil {
		return x.Numbers
	}
	return nil
}

// Information of an end user.
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Highly recommended for logged-in users. Unique identifier for logged-in
	// user, such as a user name. Don't set for anonymous users.
	//
	// Always use a hashed value for this ID.
	//
	// Don't set the field to the same fixed ID for different users. This mixes
	// the event history of those users together, which results in degraded
	// model quality.
	//
	// The field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an `INVALID_ARGUMENT` error is returned.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// User agent as included in the HTTP header.
	//
	// The field must be a UTF-8 encoded string with a length limit of 1,000
	// characters. Otherwise, an `INVALID_ARGUMENT` error is returned.
	//
	// This should not be set when using the client side event reporting with
	// GTM or JavaScript tag in
	// [UserEventService.CollectUserEvent][google.cloud.discoveryengine.v1beta.UserEventService.CollectUserEvent]
	// or if
	// [UserEvent.direct_user_request][google.cloud.discoveryengine.v1beta.UserEvent.direct_user_request]
	// is set.
	UserAgent string `protobuf:"bytes,2,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_common_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfo) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

// Double list.
type DoubleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Double values.
	Values []float64 `protobuf:"fixed64,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *DoubleList) Reset() {
	*x = DoubleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleList) ProtoMessage() {}

func (x *DoubleList) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleList.ProtoReflect.Descriptor instead.
func (*DoubleList) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_common_proto_rawDescGZIP(), []int{3}
}

func (x *DoubleList) GetValues() []float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_google_cloud_discoveryengine_v1beta_common_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1beta_common_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x1a, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x2d, 0x0a, 0x11, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x07, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x2d, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x01, 0x52, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x4d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x42, 0x05, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x22, 0x3f, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x22, 0x42, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0xd2,
	0x07, 0xea, 0x41, 0xe6, 0x01, 0x0a, 0x25, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x51, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d, 0x12,
	0x6a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d, 0xea, 0x41, 0xc5, 0x01, 0x0a,
	0x28, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x12, 0x58, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x7d, 0xea, 0x41, 0x89, 0x02, 0x0a, 0x2c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12, 0x78, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x7d, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44,
	0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02,
	0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31,
	0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1beta_common_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1beta_common_proto_rawDescData = file_google_cloud_discoveryengine_v1beta_common_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1beta_common_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1beta_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1beta_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1beta_common_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1beta_common_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_discoveryengine_v1beta_common_proto_goTypes = []interface{}{
	(*Interval)(nil),        // 0: google.cloud.discoveryengine.v1beta.Interval
	(*CustomAttribute)(nil), // 1: google.cloud.discoveryengine.v1beta.CustomAttribute
	(*UserInfo)(nil),        // 2: google.cloud.discoveryengine.v1beta.UserInfo
	(*DoubleList)(nil),      // 3: google.cloud.discoveryengine.v1beta.DoubleList
}
var file_google_cloud_discoveryengine_v1beta_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1beta_common_proto_init() }
func file_google_cloud_discoveryengine_v1beta_common_proto_init() {
	if File_google_cloud_discoveryengine_v1beta_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interval); i {
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
		file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomAttribute); i {
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
		file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleList); i {
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
	file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Interval_Minimum)(nil),
		(*Interval_ExclusiveMinimum)(nil),
		(*Interval_Maximum)(nil),
		(*Interval_ExclusiveMaximum)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_discoveryengine_v1beta_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1beta_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1beta_common_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1beta_common_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1beta_common_proto = out.File
	file_google_cloud_discoveryengine_v1beta_common_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1beta_common_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1beta_common_proto_depIdxs = nil
}
