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
// source: google/datastore/admin/v1/migration.proto

package adminpb

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

// States for a migration.
type MigrationState int32

const (
	// Unspecified.
	MigrationState_MIGRATION_STATE_UNSPECIFIED MigrationState = 0
	// The migration is running.
	MigrationState_RUNNING MigrationState = 1
	// The migration is paused.
	MigrationState_PAUSED MigrationState = 2
	// The migration is complete.
	MigrationState_COMPLETE MigrationState = 3
)

// Enum value maps for MigrationState.
var (
	MigrationState_name = map[int32]string{
		0: "MIGRATION_STATE_UNSPECIFIED",
		1: "RUNNING",
		2: "PAUSED",
		3: "COMPLETE",
	}
	MigrationState_value = map[string]int32{
		"MIGRATION_STATE_UNSPECIFIED": 0,
		"RUNNING":                     1,
		"PAUSED":                      2,
		"COMPLETE":                    3,
	}
)

func (x MigrationState) Enum() *MigrationState {
	p := new(MigrationState)
	*p = x
	return p
}

func (x MigrationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_datastore_admin_v1_migration_proto_enumTypes[0].Descriptor()
}

func (MigrationState) Type() protoreflect.EnumType {
	return &file_google_datastore_admin_v1_migration_proto_enumTypes[0]
}

func (x MigrationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationState.Descriptor instead.
func (MigrationState) EnumDescriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_migration_proto_rawDescGZIP(), []int{0}
}

// Steps in a migration.
type MigrationStep int32

const (
	// Unspecified.
	MigrationStep_MIGRATION_STEP_UNSPECIFIED MigrationStep = 0
	// Pre-migration: the database is prepared for migration.
	MigrationStep_PREPARE MigrationStep = 6
	// Start of migration.
	MigrationStep_START MigrationStep = 1
	// Writes are applied synchronously to at least one replica.
	MigrationStep_APPLY_WRITES_SYNCHRONOUSLY MigrationStep = 7
	// Data is copied to Cloud Firestore and then verified to match the data in
	// Cloud Datastore.
	MigrationStep_COPY_AND_VERIFY MigrationStep = 2
	// Eventually-consistent reads are redirected to Cloud Firestore.
	MigrationStep_REDIRECT_EVENTUALLY_CONSISTENT_READS MigrationStep = 3
	// Strongly-consistent reads are redirected to Cloud Firestore.
	MigrationStep_REDIRECT_STRONGLY_CONSISTENT_READS MigrationStep = 4
	// Writes are redirected to Cloud Firestore.
	MigrationStep_REDIRECT_WRITES MigrationStep = 5
)

// Enum value maps for MigrationStep.
var (
	MigrationStep_name = map[int32]string{
		0: "MIGRATION_STEP_UNSPECIFIED",
		6: "PREPARE",
		1: "START",
		7: "APPLY_WRITES_SYNCHRONOUSLY",
		2: "COPY_AND_VERIFY",
		3: "REDIRECT_EVENTUALLY_CONSISTENT_READS",
		4: "REDIRECT_STRONGLY_CONSISTENT_READS",
		5: "REDIRECT_WRITES",
	}
	MigrationStep_value = map[string]int32{
		"MIGRATION_STEP_UNSPECIFIED":           0,
		"PREPARE":                              6,
		"START":                                1,
		"APPLY_WRITES_SYNCHRONOUSLY":           7,
		"COPY_AND_VERIFY":                      2,
		"REDIRECT_EVENTUALLY_CONSISTENT_READS": 3,
		"REDIRECT_STRONGLY_CONSISTENT_READS":   4,
		"REDIRECT_WRITES":                      5,
	}
)

func (x MigrationStep) Enum() *MigrationStep {
	p := new(MigrationStep)
	*p = x
	return p
}

func (x MigrationStep) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationStep) Descriptor() protoreflect.EnumDescriptor {
	return file_google_datastore_admin_v1_migration_proto_enumTypes[1].Descriptor()
}

func (MigrationStep) Type() protoreflect.EnumType {
	return &file_google_datastore_admin_v1_migration_proto_enumTypes[1]
}

func (x MigrationStep) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationStep.Descriptor instead.
func (MigrationStep) EnumDescriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_migration_proto_rawDescGZIP(), []int{1}
}

// Concurrency modes for transactions in Cloud Firestore.
type MigrationProgressEvent_ConcurrencyMode int32

const (
	// Unspecified.
	MigrationProgressEvent_CONCURRENCY_MODE_UNSPECIFIED MigrationProgressEvent_ConcurrencyMode = 0
	// Pessimistic concurrency.
	MigrationProgressEvent_PESSIMISTIC MigrationProgressEvent_ConcurrencyMode = 1
	// Optimistic concurrency.
	MigrationProgressEvent_OPTIMISTIC MigrationProgressEvent_ConcurrencyMode = 2
	// Optimistic concurrency with entity groups.
	MigrationProgressEvent_OPTIMISTIC_WITH_ENTITY_GROUPS MigrationProgressEvent_ConcurrencyMode = 3
)

// Enum value maps for MigrationProgressEvent_ConcurrencyMode.
var (
	MigrationProgressEvent_ConcurrencyMode_name = map[int32]string{
		0: "CONCURRENCY_MODE_UNSPECIFIED",
		1: "PESSIMISTIC",
		2: "OPTIMISTIC",
		3: "OPTIMISTIC_WITH_ENTITY_GROUPS",
	}
	MigrationProgressEvent_ConcurrencyMode_value = map[string]int32{
		"CONCURRENCY_MODE_UNSPECIFIED":  0,
		"PESSIMISTIC":                   1,
		"OPTIMISTIC":                    2,
		"OPTIMISTIC_WITH_ENTITY_GROUPS": 3,
	}
)

func (x MigrationProgressEvent_ConcurrencyMode) Enum() *MigrationProgressEvent_ConcurrencyMode {
	p := new(MigrationProgressEvent_ConcurrencyMode)
	*p = x
	return p
}

func (x MigrationProgressEvent_ConcurrencyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationProgressEvent_ConcurrencyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_datastore_admin_v1_migration_proto_enumTypes[2].Descriptor()
}

func (MigrationProgressEvent_ConcurrencyMode) Type() protoreflect.EnumType {
	return &file_google_datastore_admin_v1_migration_proto_enumTypes[2]
}

func (x MigrationProgressEvent_ConcurrencyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationProgressEvent_ConcurrencyMode.Descriptor instead.
func (MigrationProgressEvent_ConcurrencyMode) EnumDescriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_migration_proto_rawDescGZIP(), []int{1, 0}
}

// An event signifying a change in state of a [migration from Cloud Datastore to
// Cloud Firestore in Datastore
// mode](https://cloud.google.com/datastore/docs/upgrade-to-firestore).
type MigrationStateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new state of the migration.
	State MigrationState `protobuf:"varint,1,opt,name=state,proto3,enum=google.datastore.admin.v1.MigrationState" json:"state,omitempty"`
}

func (x *MigrationStateEvent) Reset() {
	*x = MigrationStateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationStateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationStateEvent) ProtoMessage() {}

func (x *MigrationStateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationStateEvent.ProtoReflect.Descriptor instead.
func (*MigrationStateEvent) Descriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_migration_proto_rawDescGZIP(), []int{0}
}

func (x *MigrationStateEvent) GetState() MigrationState {
	if x != nil {
		return x.State
	}
	return MigrationState_MIGRATION_STATE_UNSPECIFIED
}

// An event signifying the start of a new step in a [migration from Cloud
// Datastore to Cloud Firestore in Datastore
// mode](https://cloud.google.com/datastore/docs/upgrade-to-firestore).
type MigrationProgressEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The step that is starting.
	//
	// An event with step set to `START` indicates that the migration
	// has been reverted back to the initial pre-migration state.
	Step MigrationStep `protobuf:"varint,1,opt,name=step,proto3,enum=google.datastore.admin.v1.MigrationStep" json:"step,omitempty"`
	// Details about this step.
	//
	// Types that are assignable to StepDetails:
	//	*MigrationProgressEvent_PrepareStepDetails_
	//	*MigrationProgressEvent_RedirectWritesStepDetails_
	StepDetails isMigrationProgressEvent_StepDetails `protobuf_oneof:"step_details"`
}

func (x *MigrationProgressEvent) Reset() {
	*x = MigrationProgressEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationProgressEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationProgressEvent) ProtoMessage() {}

func (x *MigrationProgressEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationProgressEvent.ProtoReflect.Descriptor instead.
func (*MigrationProgressEvent) Descriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_migration_proto_rawDescGZIP(), []int{1}
}

func (x *MigrationProgressEvent) GetStep() MigrationStep {
	if x != nil {
		return x.Step
	}
	return MigrationStep_MIGRATION_STEP_UNSPECIFIED
}

func (m *MigrationProgressEvent) GetStepDetails() isMigrationProgressEvent_StepDetails {
	if m != nil {
		return m.StepDetails
	}
	return nil
}

func (x *MigrationProgressEvent) GetPrepareStepDetails() *MigrationProgressEvent_PrepareStepDetails {
	if x, ok := x.GetStepDetails().(*MigrationProgressEvent_PrepareStepDetails_); ok {
		return x.PrepareStepDetails
	}
	return nil
}

func (x *MigrationProgressEvent) GetRedirectWritesStepDetails() *MigrationProgressEvent_RedirectWritesStepDetails {
	if x, ok := x.GetStepDetails().(*MigrationProgressEvent_RedirectWritesStepDetails_); ok {
		return x.RedirectWritesStepDetails
	}
	return nil
}

type isMigrationProgressEvent_StepDetails interface {
	isMigrationProgressEvent_StepDetails()
}

type MigrationProgressEvent_PrepareStepDetails_ struct {
	// Details for the `PREPARE` step.
	PrepareStepDetails *MigrationProgressEvent_PrepareStepDetails `protobuf:"bytes,2,opt,name=prepare_step_details,json=prepareStepDetails,proto3,oneof"`
}

type MigrationProgressEvent_RedirectWritesStepDetails_ struct {
	// Details for the `REDIRECT_WRITES` step.
	RedirectWritesStepDetails *MigrationProgressEvent_RedirectWritesStepDetails `protobuf:"bytes,3,opt,name=redirect_writes_step_details,json=redirectWritesStepDetails,proto3,oneof"`
}

func (*MigrationProgressEvent_PrepareStepDetails_) isMigrationProgressEvent_StepDetails() {}

func (*MigrationProgressEvent_RedirectWritesStepDetails_) isMigrationProgressEvent_StepDetails() {}

// Details for the `PREPARE` step.
type MigrationProgressEvent_PrepareStepDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The concurrency mode this database will use when it reaches the
	// `REDIRECT_WRITES` step.
	ConcurrencyMode MigrationProgressEvent_ConcurrencyMode `protobuf:"varint,1,opt,name=concurrency_mode,json=concurrencyMode,proto3,enum=google.datastore.admin.v1.MigrationProgressEvent_ConcurrencyMode" json:"concurrency_mode,omitempty"`
}

func (x *MigrationProgressEvent_PrepareStepDetails) Reset() {
	*x = MigrationProgressEvent_PrepareStepDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationProgressEvent_PrepareStepDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationProgressEvent_PrepareStepDetails) ProtoMessage() {}

func (x *MigrationProgressEvent_PrepareStepDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationProgressEvent_PrepareStepDetails.ProtoReflect.Descriptor instead.
func (*MigrationProgressEvent_PrepareStepDetails) Descriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_migration_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MigrationProgressEvent_PrepareStepDetails) GetConcurrencyMode() MigrationProgressEvent_ConcurrencyMode {
	if x != nil {
		return x.ConcurrencyMode
	}
	return MigrationProgressEvent_CONCURRENCY_MODE_UNSPECIFIED
}

// Details for the `REDIRECT_WRITES` step.
type MigrationProgressEvent_RedirectWritesStepDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ths concurrency mode for this database.
	ConcurrencyMode MigrationProgressEvent_ConcurrencyMode `protobuf:"varint,1,opt,name=concurrency_mode,json=concurrencyMode,proto3,enum=google.datastore.admin.v1.MigrationProgressEvent_ConcurrencyMode" json:"concurrency_mode,omitempty"`
}

func (x *MigrationProgressEvent_RedirectWritesStepDetails) Reset() {
	*x = MigrationProgressEvent_RedirectWritesStepDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationProgressEvent_RedirectWritesStepDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationProgressEvent_RedirectWritesStepDetails) ProtoMessage() {}

func (x *MigrationProgressEvent_RedirectWritesStepDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_admin_v1_migration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationProgressEvent_RedirectWritesStepDetails.ProtoReflect.Descriptor instead.
func (*MigrationProgressEvent_RedirectWritesStepDetails) Descriptor() ([]byte, []int) {
	return file_google_datastore_admin_v1_migration_proto_rawDescGZIP(), []int{1, 1}
}

func (x *MigrationProgressEvent_RedirectWritesStepDetails) GetConcurrencyMode() MigrationProgressEvent_ConcurrencyMode {
	if x != nil {
		return x.ConcurrencyMode
	}
	return MigrationProgressEvent_CONCURRENCY_MODE_UNSPECIFIED
}

var File_google_datastore_admin_v1_migration_proto protoreflect.FileDescriptor

var file_google_datastore_admin_v1_migration_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x13, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xfb,
	0x05, 0x0a, 0x16, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65,
	0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x78, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x12, 0x70,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x8e, 0x01, 0x0a, 0x1c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x73, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x19, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x73, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x1a, 0x82, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6c, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x1a, 0x89, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x73, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6c, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x22, 0x77, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x43, 0x55, 0x52,
	0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x45, 0x53, 0x53,
	0x49, 0x4d, 0x49, 0x53, 0x54, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x50, 0x54,
	0x49, 0x4d, 0x49, 0x53, 0x54, 0x49, 0x43, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x4f, 0x50, 0x54,
	0x49, 0x4d, 0x49, 0x53, 0x54, 0x49, 0x43, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x45, 0x4e, 0x54,
	0x49, 0x54, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x03, 0x42, 0x0e, 0x0a, 0x0c,
	0x73, 0x74, 0x65, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0x58, 0x0a, 0x0e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x0a, 0x1b, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x2a, 0xe3, 0x01, 0x0a, 0x0d, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x49, 0x47, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x50,
	0x41, 0x52, 0x45, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01,
	0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x53,
	0x5f, 0x53, 0x59, 0x4e, 0x43, 0x48, 0x52, 0x4f, 0x4e, 0x4f, 0x55, 0x53, 0x4c, 0x59, 0x10, 0x07,
	0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x50, 0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x56, 0x45, 0x52,
	0x49, 0x46, 0x59, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43,
	0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x55, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x43, 0x4f, 0x4e,
	0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x53, 0x10, 0x03, 0x12,
	0x26, 0x0a, 0x22, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x4f,
	0x4e, 0x47, 0x4c, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x5f,
	0x52, 0x45, 0x41, 0x44, 0x53, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x53, 0x10, 0x05, 0x42, 0xd6, 0x01, 0x0a,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_datastore_admin_v1_migration_proto_rawDescOnce sync.Once
	file_google_datastore_admin_v1_migration_proto_rawDescData = file_google_datastore_admin_v1_migration_proto_rawDesc
)

func file_google_datastore_admin_v1_migration_proto_rawDescGZIP() []byte {
	file_google_datastore_admin_v1_migration_proto_rawDescOnce.Do(func() {
		file_google_datastore_admin_v1_migration_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_datastore_admin_v1_migration_proto_rawDescData)
	})
	return file_google_datastore_admin_v1_migration_proto_rawDescData
}

var file_google_datastore_admin_v1_migration_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_datastore_admin_v1_migration_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_datastore_admin_v1_migration_proto_goTypes = []interface{}{
	(MigrationState)(0),                                      // 0: google.datastore.admin.v1.MigrationState
	(MigrationStep)(0),                                       // 1: google.datastore.admin.v1.MigrationStep
	(MigrationProgressEvent_ConcurrencyMode)(0),              // 2: google.datastore.admin.v1.MigrationProgressEvent.ConcurrencyMode
	(*MigrationStateEvent)(nil),                              // 3: google.datastore.admin.v1.MigrationStateEvent
	(*MigrationProgressEvent)(nil),                           // 4: google.datastore.admin.v1.MigrationProgressEvent
	(*MigrationProgressEvent_PrepareStepDetails)(nil),        // 5: google.datastore.admin.v1.MigrationProgressEvent.PrepareStepDetails
	(*MigrationProgressEvent_RedirectWritesStepDetails)(nil), // 6: google.datastore.admin.v1.MigrationProgressEvent.RedirectWritesStepDetails
}
var file_google_datastore_admin_v1_migration_proto_depIdxs = []int32{
	0, // 0: google.datastore.admin.v1.MigrationStateEvent.state:type_name -> google.datastore.admin.v1.MigrationState
	1, // 1: google.datastore.admin.v1.MigrationProgressEvent.step:type_name -> google.datastore.admin.v1.MigrationStep
	5, // 2: google.datastore.admin.v1.MigrationProgressEvent.prepare_step_details:type_name -> google.datastore.admin.v1.MigrationProgressEvent.PrepareStepDetails
	6, // 3: google.datastore.admin.v1.MigrationProgressEvent.redirect_writes_step_details:type_name -> google.datastore.admin.v1.MigrationProgressEvent.RedirectWritesStepDetails
	2, // 4: google.datastore.admin.v1.MigrationProgressEvent.PrepareStepDetails.concurrency_mode:type_name -> google.datastore.admin.v1.MigrationProgressEvent.ConcurrencyMode
	2, // 5: google.datastore.admin.v1.MigrationProgressEvent.RedirectWritesStepDetails.concurrency_mode:type_name -> google.datastore.admin.v1.MigrationProgressEvent.ConcurrencyMode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_datastore_admin_v1_migration_proto_init() }
func file_google_datastore_admin_v1_migration_proto_init() {
	if File_google_datastore_admin_v1_migration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_datastore_admin_v1_migration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationStateEvent); i {
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
		file_google_datastore_admin_v1_migration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationProgressEvent); i {
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
		file_google_datastore_admin_v1_migration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationProgressEvent_PrepareStepDetails); i {
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
		file_google_datastore_admin_v1_migration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationProgressEvent_RedirectWritesStepDetails); i {
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
	file_google_datastore_admin_v1_migration_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MigrationProgressEvent_PrepareStepDetails_)(nil),
		(*MigrationProgressEvent_RedirectWritesStepDetails_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_datastore_admin_v1_migration_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_datastore_admin_v1_migration_proto_goTypes,
		DependencyIndexes: file_google_datastore_admin_v1_migration_proto_depIdxs,
		EnumInfos:         file_google_datastore_admin_v1_migration_proto_enumTypes,
		MessageInfos:      file_google_datastore_admin_v1_migration_proto_msgTypes,
	}.Build()
	File_google_datastore_admin_v1_migration_proto = out.File
	file_google_datastore_admin_v1_migration_proto_rawDesc = nil
	file_google_datastore_admin_v1_migration_proto_goTypes = nil
	file_google_datastore_admin_v1_migration_proto_depIdxs = nil
}
