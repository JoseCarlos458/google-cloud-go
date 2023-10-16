// Copyright 2020 Google LLC
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
// source: google/appengine/v1/certificate.proto

package appenginepb

import (
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

// State of certificate management. Refers to the most recent certificate
// acquisition or renewal attempt.
type ManagementStatus int32

const (
	ManagementStatus_MANAGEMENT_STATUS_UNSPECIFIED ManagementStatus = 0
	// Certificate was successfully obtained and inserted into the serving
	// system.
	ManagementStatus_OK ManagementStatus = 1
	// Certificate is under active attempts to acquire or renew.
	ManagementStatus_PENDING ManagementStatus = 2
	// Most recent renewal failed due to an invalid DNS setup and will be
	// retried. Renewal attempts will continue to fail until the certificate
	// domain's DNS configuration is fixed. The last successfully provisioned
	// certificate may still be serving.
	ManagementStatus_FAILED_RETRYING_NOT_VISIBLE ManagementStatus = 4
	// All renewal attempts have been exhausted, likely due to an invalid DNS
	// setup.
	ManagementStatus_FAILED_PERMANENT ManagementStatus = 6
	// Most recent renewal failed due to an explicit CAA record that does not
	// include one of the in-use CAs (Google CA and Let's Encrypt). Renewals will
	// continue to fail until the CAA is reconfigured. The last successfully
	// provisioned certificate may still be serving.
	ManagementStatus_FAILED_RETRYING_CAA_FORBIDDEN ManagementStatus = 7
	// Most recent renewal failed due to a CAA retrieval failure. This means that
	// the domain's DNS provider does not properly handle CAA records, failing
	// requests for CAA records when no CAA records are defined. Renewals will
	// continue to fail until the DNS provider is changed or a CAA record is
	// added for the given domain. The last successfully provisioned certificate
	// may still be serving.
	ManagementStatus_FAILED_RETRYING_CAA_CHECKING ManagementStatus = 8
)

// Enum value maps for ManagementStatus.
var (
	ManagementStatus_name = map[int32]string{
		0: "MANAGEMENT_STATUS_UNSPECIFIED",
		1: "OK",
		2: "PENDING",
		4: "FAILED_RETRYING_NOT_VISIBLE",
		6: "FAILED_PERMANENT",
		7: "FAILED_RETRYING_CAA_FORBIDDEN",
		8: "FAILED_RETRYING_CAA_CHECKING",
	}
	ManagementStatus_value = map[string]int32{
		"MANAGEMENT_STATUS_UNSPECIFIED": 0,
		"OK":                            1,
		"PENDING":                       2,
		"FAILED_RETRYING_NOT_VISIBLE":   4,
		"FAILED_PERMANENT":              6,
		"FAILED_RETRYING_CAA_FORBIDDEN": 7,
		"FAILED_RETRYING_CAA_CHECKING":  8,
	}
)

func (x ManagementStatus) Enum() *ManagementStatus {
	p := new(ManagementStatus)
	*p = x
	return p
}

func (x ManagementStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManagementStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1_certificate_proto_enumTypes[0].Descriptor()
}

func (ManagementStatus) Type() protoreflect.EnumType {
	return &file_google_appengine_v1_certificate_proto_enumTypes[0]
}

func (x ManagementStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManagementStatus.Descriptor instead.
func (ManagementStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1_certificate_proto_rawDescGZIP(), []int{0}
}

// An SSL certificate that a user has been authorized to administer. A user
// is authorized to administer any certificate that applies to one of their
// authorized domains.
type AuthorizedCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full path to the `AuthorizedCertificate` resource in the API. Example:
	// `apps/myapp/authorizedCertificates/12345`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative name of the certificate. This is a unique value autogenerated
	// on `AuthorizedCertificate` resource creation. Example: `12345`.
	//
	// @OutputOnly
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The user-specified display name of the certificate. This is not
	// guaranteed to be unique. Example: `My Certificate`.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Topmost applicable domains of this certificate. This certificate
	// applies to these domains and their subdomains. Example: `example.com`.
	//
	// @OutputOnly
	DomainNames []string `protobuf:"bytes,4,rep,name=domain_names,json=domainNames,proto3" json:"domain_names,omitempty"`
	// The time when this certificate expires. To update the renewal time on this
	// certificate, upload an SSL certificate with a different expiration time
	// using [`AuthorizedCertificates.UpdateAuthorizedCertificate`]().
	//
	// @OutputOnly
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	// The SSL certificate serving the `AuthorizedCertificate` resource. This
	// must be obtained independently from a certificate authority.
	CertificateRawData *CertificateRawData `protobuf:"bytes,6,opt,name=certificate_raw_data,json=certificateRawData,proto3" json:"certificate_raw_data,omitempty"`
	// Only applicable if this certificate is managed by App Engine. Managed
	// certificates are tied to the lifecycle of a `DomainMapping` and cannot be
	// updated or deleted via the `AuthorizedCertificates` API. If this
	// certificate is manually administered by the user, this field will be empty.
	//
	// @OutputOnly
	ManagedCertificate *ManagedCertificate `protobuf:"bytes,7,opt,name=managed_certificate,json=managedCertificate,proto3" json:"managed_certificate,omitempty"`
	// The full paths to user visible Domain Mapping resources that have this
	// certificate mapped. Example: `apps/myapp/domainMappings/example.com`.
	//
	// This may not represent the full list of mapped domain mappings if the user
	// does not have `VIEWER` permissions on all of the applications that have
	// this certificate mapped. See `domain_mappings_count` for a complete count.
	//
	// Only returned by `GET` or `LIST` requests when specifically requested by
	// the `view=FULL_CERTIFICATE` option.
	//
	// @OutputOnly
	VisibleDomainMappings []string `protobuf:"bytes,8,rep,name=visible_domain_mappings,json=visibleDomainMappings,proto3" json:"visible_domain_mappings,omitempty"`
	// Aggregate count of the domain mappings with this certificate mapped. This
	// count includes domain mappings on applications for which the user does not
	// have `VIEWER` permissions.
	//
	// Only returned by `GET` or `LIST` requests when specifically requested by
	// the `view=FULL_CERTIFICATE` option.
	//
	// @OutputOnly
	DomainMappingsCount int32 `protobuf:"varint,9,opt,name=domain_mappings_count,json=domainMappingsCount,proto3" json:"domain_mappings_count,omitempty"`
}

func (x *AuthorizedCertificate) Reset() {
	*x = AuthorizedCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizedCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizedCertificate) ProtoMessage() {}

func (x *AuthorizedCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizedCertificate.ProtoReflect.Descriptor instead.
func (*AuthorizedCertificate) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizedCertificate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorizedCertificate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthorizedCertificate) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AuthorizedCertificate) GetDomainNames() []string {
	if x != nil {
		return x.DomainNames
	}
	return nil
}

func (x *AuthorizedCertificate) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

func (x *AuthorizedCertificate) GetCertificateRawData() *CertificateRawData {
	if x != nil {
		return x.CertificateRawData
	}
	return nil
}

func (x *AuthorizedCertificate) GetManagedCertificate() *ManagedCertificate {
	if x != nil {
		return x.ManagedCertificate
	}
	return nil
}

func (x *AuthorizedCertificate) GetVisibleDomainMappings() []string {
	if x != nil {
		return x.VisibleDomainMappings
	}
	return nil
}

func (x *AuthorizedCertificate) GetDomainMappingsCount() int32 {
	if x != nil {
		return x.DomainMappingsCount
	}
	return 0
}

// An SSL certificate obtained from a certificate authority.
type CertificateRawData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PEM encoded x.509 public key certificate. This field is set once on
	// certificate creation. Must include the header and footer. Example:
	// <pre>
	// -----BEGIN CERTIFICATE-----
	// <certificate_value>
	// -----END CERTIFICATE-----
	// </pre>
	PublicCertificate string `protobuf:"bytes,1,opt,name=public_certificate,json=publicCertificate,proto3" json:"public_certificate,omitempty"`
	// Unencrypted PEM encoded RSA private key. This field is set once on
	// certificate creation and then encrypted. The key size must be 2048
	// bits or fewer. Must include the header and footer. Example:
	// <pre>
	// -----BEGIN RSA PRIVATE KEY-----
	// <unencrypted_key_value>
	// -----END RSA PRIVATE KEY-----
	// </pre>
	// @InputOnly
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (x *CertificateRawData) Reset() {
	*x = CertificateRawData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRawData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRawData) ProtoMessage() {}

func (x *CertificateRawData) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRawData.ProtoReflect.Descriptor instead.
func (*CertificateRawData) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateRawData) GetPublicCertificate() string {
	if x != nil {
		return x.PublicCertificate
	}
	return ""
}

func (x *CertificateRawData) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

// A certificate managed by App Engine.
type ManagedCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time at which the certificate was last renewed. The renewal process is
	// fully managed. Certificate renewal will automatically occur before the
	// certificate expires. Renewal errors can be tracked via `ManagementStatus`.
	//
	// @OutputOnly
	LastRenewalTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_renewal_time,json=lastRenewalTime,proto3" json:"last_renewal_time,omitempty"`
	// Status of certificate management. Refers to the most recent certificate
	// acquisition or renewal attempt.
	//
	// @OutputOnly
	Status ManagementStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.appengine.v1.ManagementStatus" json:"status,omitempty"`
}

func (x *ManagedCertificate) Reset() {
	*x = ManagedCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_certificate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagedCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagedCertificate) ProtoMessage() {}

func (x *ManagedCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_certificate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagedCertificate.ProtoReflect.Descriptor instead.
func (*ManagedCertificate) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *ManagedCertificate) GetLastRenewalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRenewalTime
	}
	return nil
}

func (x *ManagedCertificate) GetStatus() ManagementStatus {
	if x != nil {
		return x.Status
	}
	return ManagementStatus_MANAGEMENT_STATUS_UNSPECIFIED
}

var File_google_appengine_v1_certificate_proto protoreflect.FileDescriptor

var file_google_appengine_v1_certificate_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x03,
	0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x59,
	0x0a, 0x14, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61,
	0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x12, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x58, 0x0a, 0x13, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x12, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x64, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x9b, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x46, 0x0a, 0x11,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2a, 0xc6, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x41, 0x4e, 0x41,
	0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59,
	0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10,
	0x04, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x50, 0x45, 0x52, 0x4d,
	0x41, 0x4e, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x21, 0x0a, 0x1d, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x41, 0x5f, 0x46,
	0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41,
	0x41, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x42, 0xc1, 0x01, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1_certificate_proto_rawDescOnce sync.Once
	file_google_appengine_v1_certificate_proto_rawDescData = file_google_appengine_v1_certificate_proto_rawDesc
)

func file_google_appengine_v1_certificate_proto_rawDescGZIP() []byte {
	file_google_appengine_v1_certificate_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1_certificate_proto_rawDescData)
	})
	return file_google_appengine_v1_certificate_proto_rawDescData
}

var file_google_appengine_v1_certificate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_appengine_v1_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_appengine_v1_certificate_proto_goTypes = []interface{}{
	(ManagementStatus)(0),         // 0: google.appengine.v1.ManagementStatus
	(*AuthorizedCertificate)(nil), // 1: google.appengine.v1.AuthorizedCertificate
	(*CertificateRawData)(nil),    // 2: google.appengine.v1.CertificateRawData
	(*ManagedCertificate)(nil),    // 3: google.appengine.v1.ManagedCertificate
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_google_appengine_v1_certificate_proto_depIdxs = []int32{
	4, // 0: google.appengine.v1.AuthorizedCertificate.expire_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.appengine.v1.AuthorizedCertificate.certificate_raw_data:type_name -> google.appengine.v1.CertificateRawData
	3, // 2: google.appengine.v1.AuthorizedCertificate.managed_certificate:type_name -> google.appengine.v1.ManagedCertificate
	4, // 3: google.appengine.v1.ManagedCertificate.last_renewal_time:type_name -> google.protobuf.Timestamp
	0, // 4: google.appengine.v1.ManagedCertificate.status:type_name -> google.appengine.v1.ManagementStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_appengine_v1_certificate_proto_init() }
func file_google_appengine_v1_certificate_proto_init() {
	if File_google_appengine_v1_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizedCertificate); i {
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
		file_google_appengine_v1_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRawData); i {
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
		file_google_appengine_v1_certificate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagedCertificate); i {
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
			RawDescriptor: file_google_appengine_v1_certificate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1_certificate_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1_certificate_proto_depIdxs,
		EnumInfos:         file_google_appengine_v1_certificate_proto_enumTypes,
		MessageInfos:      file_google_appengine_v1_certificate_proto_msgTypes,
	}.Build()
	File_google_appengine_v1_certificate_proto = out.File
	file_google_appengine_v1_certificate_proto_rawDesc = nil
	file_google_appengine_v1_certificate_proto_goTypes = nil
	file_google_appengine_v1_certificate_proto_depIdxs = nil
}
