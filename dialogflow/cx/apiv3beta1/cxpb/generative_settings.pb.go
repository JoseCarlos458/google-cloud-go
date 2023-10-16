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
// source: google/cloud/dialogflow/cx/v3beta1/generative_settings.proto

package cxpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Settings for Generative AI.
type GenerativeSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/generativeSettings`.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Settings for Generative Fallback.
	FallbackSettings *GenerativeSettings_FallbackSettings `protobuf:"bytes,1,opt,name=fallback_settings,json=fallbackSettings,proto3" json:"fallback_settings,omitempty"`
	// Settings for Generative Safety.
	GenerativeSafetySettings *SafetySettings `protobuf:"bytes,3,opt,name=generative_safety_settings,json=generativeSafetySettings,proto3" json:"generative_safety_settings,omitempty"`
	// Settings for knowledge connector.
	KnowledgeConnectorSettings *GenerativeSettings_KnowledgeConnectorSettings `protobuf:"bytes,7,opt,name=knowledge_connector_settings,json=knowledgeConnectorSettings,proto3" json:"knowledge_connector_settings,omitempty"`
	// Language for this settings.
	LanguageCode string `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *GenerativeSettings) Reset() {
	*x = GenerativeSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerativeSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerativeSettings) ProtoMessage() {}

func (x *GenerativeSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerativeSettings.ProtoReflect.Descriptor instead.
func (*GenerativeSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescGZIP(), []int{0}
}

func (x *GenerativeSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenerativeSettings) GetFallbackSettings() *GenerativeSettings_FallbackSettings {
	if x != nil {
		return x.FallbackSettings
	}
	return nil
}

func (x *GenerativeSettings) GetGenerativeSafetySettings() *SafetySettings {
	if x != nil {
		return x.GenerativeSafetySettings
	}
	return nil
}

func (x *GenerativeSettings) GetKnowledgeConnectorSettings() *GenerativeSettings_KnowledgeConnectorSettings {
	if x != nil {
		return x.KnowledgeConnectorSettings
	}
	return nil
}

func (x *GenerativeSettings) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

// Settings for Generative Fallback.
type GenerativeSettings_FallbackSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display name of the selected prompt.
	SelectedPrompt string `protobuf:"bytes,3,opt,name=selected_prompt,json=selectedPrompt,proto3" json:"selected_prompt,omitempty"`
	// Stored prompts that can be selected, for example default templates like
	// "conservative" or "chatty", or user defined ones.
	PromptTemplates []*GenerativeSettings_FallbackSettings_PromptTemplate `protobuf:"bytes,4,rep,name=prompt_templates,json=promptTemplates,proto3" json:"prompt_templates,omitempty"`
}

func (x *GenerativeSettings_FallbackSettings) Reset() {
	*x = GenerativeSettings_FallbackSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerativeSettings_FallbackSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerativeSettings_FallbackSettings) ProtoMessage() {}

func (x *GenerativeSettings_FallbackSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerativeSettings_FallbackSettings.ProtoReflect.Descriptor instead.
func (*GenerativeSettings_FallbackSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GenerativeSettings_FallbackSettings) GetSelectedPrompt() string {
	if x != nil {
		return x.SelectedPrompt
	}
	return ""
}

func (x *GenerativeSettings_FallbackSettings) GetPromptTemplates() []*GenerativeSettings_FallbackSettings_PromptTemplate {
	if x != nil {
		return x.PromptTemplates
	}
	return nil
}

// Settings for knowledge connector. These parameters are used for LLM prompt
// like "You are <agent>. You are a helpful and verbose <agent_identity> at
// <business>, <business_description>. Your task is to help humans on
// <agent_scope>".
type GenerativeSettings_KnowledgeConnectorSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the company, organization or other entity that the agent
	// represents. Used for knowledge connector LLM prompt and for knowledge
	// search.
	Business string `protobuf:"bytes,1,opt,name=business,proto3" json:"business,omitempty"`
	// Name of the virtual agent. Used for LLM prompt. Can be left empty.
	Agent string `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	// Identity of the agent, e.g. "virtual agent", "AI assistant".
	AgentIdentity string `protobuf:"bytes,3,opt,name=agent_identity,json=agentIdentity,proto3" json:"agent_identity,omitempty"`
	// Company description, used for LLM prompt, e.g. "a family company selling
	// freshly roasted coffee beans".
	BusinessDescription string `protobuf:"bytes,4,opt,name=business_description,json=businessDescription,proto3" json:"business_description,omitempty"`
	// Agent scope, e.g. "Example company website", "internal Example
	// company website for employees", "manual of car owner".
	AgentScope string `protobuf:"bytes,5,opt,name=agent_scope,json=agentScope,proto3" json:"agent_scope,omitempty"`
}

func (x *GenerativeSettings_KnowledgeConnectorSettings) Reset() {
	*x = GenerativeSettings_KnowledgeConnectorSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerativeSettings_KnowledgeConnectorSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerativeSettings_KnowledgeConnectorSettings) ProtoMessage() {}

func (x *GenerativeSettings_KnowledgeConnectorSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerativeSettings_KnowledgeConnectorSettings.ProtoReflect.Descriptor instead.
func (*GenerativeSettings_KnowledgeConnectorSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GenerativeSettings_KnowledgeConnectorSettings) GetBusiness() string {
	if x != nil {
		return x.Business
	}
	return ""
}

func (x *GenerativeSettings_KnowledgeConnectorSettings) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *GenerativeSettings_KnowledgeConnectorSettings) GetAgentIdentity() string {
	if x != nil {
		return x.AgentIdentity
	}
	return ""
}

func (x *GenerativeSettings_KnowledgeConnectorSettings) GetBusinessDescription() string {
	if x != nil {
		return x.BusinessDescription
	}
	return ""
}

func (x *GenerativeSettings_KnowledgeConnectorSettings) GetAgentScope() string {
	if x != nil {
		return x.AgentScope
	}
	return ""
}

// Prompt template.
type GenerativeSettings_FallbackSettings_PromptTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt name.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Prompt text that is sent to a LLM on no-match default, placeholders are
	// filled downstream. For example: "Here is a conversation $conversation,
	// a response is: "
	PromptText string `protobuf:"bytes,2,opt,name=prompt_text,json=promptText,proto3" json:"prompt_text,omitempty"`
	// If the flag is true, the prompt is frozen and cannot be modified by
	// users.
	Frozen bool `protobuf:"varint,3,opt,name=frozen,proto3" json:"frozen,omitempty"`
}

func (x *GenerativeSettings_FallbackSettings_PromptTemplate) Reset() {
	*x = GenerativeSettings_FallbackSettings_PromptTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerativeSettings_FallbackSettings_PromptTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerativeSettings_FallbackSettings_PromptTemplate) ProtoMessage() {}

func (x *GenerativeSettings_FallbackSettings_PromptTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerativeSettings_FallbackSettings_PromptTemplate.ProtoReflect.Descriptor instead.
func (*GenerativeSettings_FallbackSettings_PromptTemplate) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *GenerativeSettings_FallbackSettings_PromptTemplate) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GenerativeSettings_FallbackSettings_PromptTemplate) GetPromptText() string {
	if x != nil {
		return x.PromptText
	}
	return ""
}

func (x *GenerativeSettings_FallbackSettings_PromptTemplate) GetFrozen() bool {
	if x != nil {
		return x.Frozen
	}
	return false
}

var File_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x08, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x74, 0x0a, 0x11, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x70, 0x0a, 0x1a, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x18, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x61, 0x66, 0x65,
	0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x93, 0x01, 0x0a, 0x1c, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76,
	0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x1a, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0xad, 0x02, 0x0a, 0x10, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x56,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x6c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66,
	0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x1a, 0xc9, 0x01, 0x0a, 0x1a, 0x4b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a,
	0x14, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x3a, 0x81, 0x01, 0xea, 0x41, 0x7e, 0x0a, 0x31, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x49, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x7d, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0xd1, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x78, 0x70, 0x62, 0x3b, 0x63,
	0x78, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33, 0x42, 0x65, 0x74, 0x61, 0x31,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x58,
	0x3a, 0x3a, 0x56, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_goTypes = []interface{}{
	(*GenerativeSettings)(nil),                                 // 0: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings
	(*GenerativeSettings_FallbackSettings)(nil),                // 1: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.FallbackSettings
	(*GenerativeSettings_KnowledgeConnectorSettings)(nil),      // 2: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.KnowledgeConnectorSettings
	(*GenerativeSettings_FallbackSettings_PromptTemplate)(nil), // 3: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.FallbackSettings.PromptTemplate
	(*SafetySettings)(nil),                                     // 4: google.cloud.dialogflow.cx.v3beta1.SafetySettings
}
var file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_depIdxs = []int32{
	1, // 0: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.fallback_settings:type_name -> google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.FallbackSettings
	4, // 1: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.generative_safety_settings:type_name -> google.cloud.dialogflow.cx.v3beta1.SafetySettings
	2, // 2: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.knowledge_connector_settings:type_name -> google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.KnowledgeConnectorSettings
	3, // 3: google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.FallbackSettings.prompt_templates:type_name -> google.cloud.dialogflow.cx.v3beta1.GenerativeSettings.FallbackSettings.PromptTemplate
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto != nil {
		return
	}
	file_google_cloud_dialogflow_cx_v3beta1_safety_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerativeSettings); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerativeSettings_FallbackSettings); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerativeSettings_KnowledgeConnectorSettings); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerativeSettings_FallbackSettings_PromptTemplate); i {
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
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_depIdxs,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_generative_settings_proto_depIdxs = nil
}
