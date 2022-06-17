// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: Runtime/Handshake/Handshake.proto

package handshake

import (
	protobuf "go.dolittle.io/contracts/v7/protobuf"
	versioning "go.dolittle.io/contracts/v7/versioning"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HandshakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SdkIdentifier    string               `protobuf:"bytes,1,opt,name=sdkIdentifier,proto3" json:"sdkIdentifier,omitempty"`
	HeadVersion      *versioning.Version  `protobuf:"bytes,2,opt,name=headVersion,proto3" json:"headVersion,omitempty"`
	SdkVersion       *versioning.Version  `protobuf:"bytes,3,opt,name=sdkVersion,proto3" json:"sdkVersion,omitempty"`
	ContractsVersion *versioning.Version  `protobuf:"bytes,4,opt,name=contractsVersion,proto3" json:"contractsVersion,omitempty"`
	Attempt          uint32               `protobuf:"varint,5,opt,name=attempt,proto3" json:"attempt,omitempty"`
	TimeSpent        *durationpb.Duration `protobuf:"bytes,6,opt,name=timeSpent,proto3" json:"timeSpent,omitempty"`
}

func (x *HandshakeRequest) Reset() {
	*x = HandshakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Handshake_Handshake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeRequest) ProtoMessage() {}

func (x *HandshakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Handshake_Handshake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeRequest.ProtoReflect.Descriptor instead.
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Handshake_Handshake_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakeRequest) GetSdkIdentifier() string {
	if x != nil {
		return x.SdkIdentifier
	}
	return ""
}

func (x *HandshakeRequest) GetHeadVersion() *versioning.Version {
	if x != nil {
		return x.HeadVersion
	}
	return nil
}

func (x *HandshakeRequest) GetSdkVersion() *versioning.Version {
	if x != nil {
		return x.SdkVersion
	}
	return nil
}

func (x *HandshakeRequest) GetContractsVersion() *versioning.Version {
	if x != nil {
		return x.ContractsVersion
	}
	return nil
}

func (x *HandshakeRequest) GetAttempt() uint32 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

func (x *HandshakeRequest) GetTimeSpent() *durationpb.Duration {
	if x != nil {
		return x.TimeSpent
	}
	return nil
}

type HandshakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure          *protobuf.Failure   `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
	RuntimeVersion   *versioning.Version `protobuf:"bytes,2,opt,name=runtimeVersion,proto3" json:"runtimeVersion,omitempty"`
	ContractsVersion *versioning.Version `protobuf:"bytes,3,opt,name=contractsVersion,proto3" json:"contractsVersion,omitempty"`
	CustomerId       *protobuf.Uuid      `protobuf:"bytes,4,opt,name=customerId,proto3" json:"customerId,omitempty"`
	CustomerName     string              `protobuf:"bytes,5,opt,name=customerName,proto3" json:"customerName,omitempty"`
	ApplicationId    *protobuf.Uuid      `protobuf:"bytes,6,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	ApplicationName  string              `protobuf:"bytes,7,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	MicroserviceId   *protobuf.Uuid      `protobuf:"bytes,8,opt,name=microserviceId,proto3" json:"microserviceId,omitempty"`
	MicroserviceName string              `protobuf:"bytes,9,opt,name=microserviceName,proto3" json:"microserviceName,omitempty"`
	EnvironmentName  string              `protobuf:"bytes,10,opt,name=environmentName,proto3" json:"environmentName,omitempty"`
	OtlpEndpoint     *string             `protobuf:"bytes,11,opt,name=otlpEndpoint,proto3,oneof" json:"otlpEndpoint,omitempty"`
}

func (x *HandshakeResponse) Reset() {
	*x = HandshakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Handshake_Handshake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeResponse) ProtoMessage() {}

func (x *HandshakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Handshake_Handshake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeResponse.ProtoReflect.Descriptor instead.
func (*HandshakeResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Handshake_Handshake_proto_rawDescGZIP(), []int{1}
}

func (x *HandshakeResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *HandshakeResponse) GetRuntimeVersion() *versioning.Version {
	if x != nil {
		return x.RuntimeVersion
	}
	return nil
}

func (x *HandshakeResponse) GetContractsVersion() *versioning.Version {
	if x != nil {
		return x.ContractsVersion
	}
	return nil
}

func (x *HandshakeResponse) GetCustomerId() *protobuf.Uuid {
	if x != nil {
		return x.CustomerId
	}
	return nil
}

func (x *HandshakeResponse) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *HandshakeResponse) GetApplicationId() *protobuf.Uuid {
	if x != nil {
		return x.ApplicationId
	}
	return nil
}

func (x *HandshakeResponse) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *HandshakeResponse) GetMicroserviceId() *protobuf.Uuid {
	if x != nil {
		return x.MicroserviceId
	}
	return nil
}

func (x *HandshakeResponse) GetMicroserviceName() string {
	if x != nil {
		return x.MicroserviceName
	}
	return ""
}

func (x *HandshakeResponse) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

func (x *HandshakeResponse) GetOtlpEndpoint() string {
	if x != nil && x.OtlpEndpoint != nil {
		return *x.OtlpEndpoint
	}
	return ""
}

var File_Runtime_Handshake_Handshake_proto protoreflect.FileDescriptor

var file_Runtime_Handshake_Handshake_proto_rawDesc = []byte{
	0x0a, 0x21, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68,
	0x61, 0x6b, 0x65, 0x2f, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x1a,
	0x16, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x64, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x64, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x3e, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x48, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x22, 0xf0, 0x04, 0x0a,
	0x11, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52,
	0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x48,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52,
	0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x10, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0c, 0x6f, 0x74, 0x6c, 0x70, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x6f,
	0x74, 0x6c, 0x70, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x6f, 0x74, 0x6c, 0x70, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x32,
	0x75, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x68, 0x0a, 0x09,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x2c, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x68, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x56, 0x5a, 0x2d, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c,
	0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0xaa, 0x02, 0x24, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Handshake_Handshake_proto_rawDescOnce sync.Once
	file_Runtime_Handshake_Handshake_proto_rawDescData = file_Runtime_Handshake_Handshake_proto_rawDesc
)

func file_Runtime_Handshake_Handshake_proto_rawDescGZIP() []byte {
	file_Runtime_Handshake_Handshake_proto_rawDescOnce.Do(func() {
		file_Runtime_Handshake_Handshake_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Handshake_Handshake_proto_rawDescData)
	})
	return file_Runtime_Handshake_Handshake_proto_rawDescData
}

var file_Runtime_Handshake_Handshake_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Runtime_Handshake_Handshake_proto_goTypes = []interface{}{
	(*HandshakeRequest)(nil),    // 0: dolittle.runtime.handshake.HandshakeRequest
	(*HandshakeResponse)(nil),   // 1: dolittle.runtime.handshake.HandshakeResponse
	(*versioning.Version)(nil),  // 2: dolittle.versioning.Version
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(*protobuf.Failure)(nil),    // 4: dolittle.protobuf.Failure
	(*protobuf.Uuid)(nil),       // 5: dolittle.protobuf.Uuid
}
var file_Runtime_Handshake_Handshake_proto_depIdxs = []int32{
	2,  // 0: dolittle.runtime.handshake.HandshakeRequest.headVersion:type_name -> dolittle.versioning.Version
	2,  // 1: dolittle.runtime.handshake.HandshakeRequest.sdkVersion:type_name -> dolittle.versioning.Version
	2,  // 2: dolittle.runtime.handshake.HandshakeRequest.contractsVersion:type_name -> dolittle.versioning.Version
	3,  // 3: dolittle.runtime.handshake.HandshakeRequest.timeSpent:type_name -> google.protobuf.Duration
	4,  // 4: dolittle.runtime.handshake.HandshakeResponse.failure:type_name -> dolittle.protobuf.Failure
	2,  // 5: dolittle.runtime.handshake.HandshakeResponse.runtimeVersion:type_name -> dolittle.versioning.Version
	2,  // 6: dolittle.runtime.handshake.HandshakeResponse.contractsVersion:type_name -> dolittle.versioning.Version
	5,  // 7: dolittle.runtime.handshake.HandshakeResponse.customerId:type_name -> dolittle.protobuf.Uuid
	5,  // 8: dolittle.runtime.handshake.HandshakeResponse.applicationId:type_name -> dolittle.protobuf.Uuid
	5,  // 9: dolittle.runtime.handshake.HandshakeResponse.microserviceId:type_name -> dolittle.protobuf.Uuid
	0,  // 10: dolittle.runtime.handshake.Handshake.Handshake:input_type -> dolittle.runtime.handshake.HandshakeRequest
	1,  // 11: dolittle.runtime.handshake.Handshake.Handshake:output_type -> dolittle.runtime.handshake.HandshakeResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_Runtime_Handshake_Handshake_proto_init() }
func file_Runtime_Handshake_Handshake_proto_init() {
	if File_Runtime_Handshake_Handshake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Handshake_Handshake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeRequest); i {
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
		file_Runtime_Handshake_Handshake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeResponse); i {
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
	file_Runtime_Handshake_Handshake_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Runtime_Handshake_Handshake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_Handshake_Handshake_proto_goTypes,
		DependencyIndexes: file_Runtime_Handshake_Handshake_proto_depIdxs,
		MessageInfos:      file_Runtime_Handshake_Handshake_proto_msgTypes,
	}.Build()
	File_Runtime_Handshake_Handshake_proto = out.File
	file_Runtime_Handshake_Handshake_proto_rawDesc = nil
	file_Runtime_Handshake_Handshake_proto_goTypes = nil
	file_Runtime_Handshake_Handshake_proto_depIdxs = nil
}
