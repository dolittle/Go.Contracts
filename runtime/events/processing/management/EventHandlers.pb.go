// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: Runtime/Management/Events.Processing/EventHandlers.proto

package management

import (
	_ "go.dolittle.io/contracts/v6/artifacts"
	protobuf "go.dolittle.io/contracts/v6/protobuf"
	_ "go.dolittle.io/contracts/v6/runtime/events/processing"
	_ "go.dolittle.io/contracts/v6/services"
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

type ReprocessEventsFromRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: Do we want another kind of execution context here?
	TenantId       *protobuf.Uuid `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	ScopeId        *protobuf.Uuid `protobuf:"bytes,2,opt,name=scopeId,proto3" json:"scopeId,omitempty"`
	EventHandlerId *protobuf.Uuid `protobuf:"bytes,3,opt,name=eventHandlerId,proto3" json:"eventHandlerId,omitempty"`
	StreamPosition uint64         `protobuf:"varint,4,opt,name=streamPosition,proto3" json:"streamPosition,omitempty"`
}

func (x *ReprocessEventsFromRequest) Reset() {
	*x = ReprocessEventsFromRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReprocessEventsFromRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReprocessEventsFromRequest) ProtoMessage() {}

func (x *ReprocessEventsFromRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReprocessEventsFromRequest.ProtoReflect.Descriptor instead.
func (*ReprocessEventsFromRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{0}
}

func (x *ReprocessEventsFromRequest) GetTenantId() *protobuf.Uuid {
	if x != nil {
		return x.TenantId
	}
	return nil
}

func (x *ReprocessEventsFromRequest) GetScopeId() *protobuf.Uuid {
	if x != nil {
		return x.ScopeId
	}
	return nil
}

func (x *ReprocessEventsFromRequest) GetEventHandlerId() *protobuf.Uuid {
	if x != nil {
		return x.EventHandlerId
	}
	return nil
}

func (x *ReprocessEventsFromRequest) GetStreamPosition() uint64 {
	if x != nil {
		return x.StreamPosition
	}
	return 0
}

type ReprocessEventsFromResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *protobuf.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *ReprocessEventsFromResponse) Reset() {
	*x = ReprocessEventsFromResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReprocessEventsFromResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReprocessEventsFromResponse) ProtoMessage() {}

func (x *ReprocessEventsFromResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReprocessEventsFromResponse.ProtoReflect.Descriptor instead.
func (*ReprocessEventsFromResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{1}
}

func (x *ReprocessEventsFromResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

type ReprocessAllEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: Do we want another kind of execution context here?
	ScopeId        *protobuf.Uuid `protobuf:"bytes,1,opt,name=scopeId,proto3" json:"scopeId,omitempty"`
	EventHandlerId *protobuf.Uuid `protobuf:"bytes,2,opt,name=eventHandlerId,proto3" json:"eventHandlerId,omitempty"`
}

func (x *ReprocessAllEventsRequest) Reset() {
	*x = ReprocessAllEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReprocessAllEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReprocessAllEventsRequest) ProtoMessage() {}

func (x *ReprocessAllEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReprocessAllEventsRequest.ProtoReflect.Descriptor instead.
func (*ReprocessAllEventsRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{2}
}

func (x *ReprocessAllEventsRequest) GetScopeId() *protobuf.Uuid {
	if x != nil {
		return x.ScopeId
	}
	return nil
}

func (x *ReprocessAllEventsRequest) GetEventHandlerId() *protobuf.Uuid {
	if x != nil {
		return x.EventHandlerId
	}
	return nil
}

type ReprocessAllEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *protobuf.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *ReprocessAllEventsResponse) Reset() {
	*x = ReprocessAllEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReprocessAllEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReprocessAllEventsResponse) ProtoMessage() {}

func (x *ReprocessAllEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReprocessAllEventsResponse.ProtoReflect.Descriptor instead.
func (*ReprocessAllEventsResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{3}
}

func (x *ReprocessAllEventsResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

var File_Runtime_Management_Events_Processing_EventHandlers_proto protoreflect.FileDescriptor

var file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDesc = []byte{
	0x0a, 0x38, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x25, 0x46, 0x75, 0x6e, 0x64, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73,
	0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x50,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x2f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64,
	0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x1b, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52,
	0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x1a, 0x52, 0x65, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x32, 0xea, 0x02,
	0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12,
	0xac, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x49, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa9,
	0x01, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x6c, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x48, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41,
	0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x49, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x7c, 0x5a, 0x40, 0x67, 0x6f,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x36, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xaa, 0x02,
	0x37, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescOnce sync.Once
	file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescData = file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDesc
)

func file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescGZIP() []byte {
	file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescOnce.Do(func() {
		file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescData)
	})
	return file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDescData
}

var file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Runtime_Management_Events_Processing_EventHandlers_proto_goTypes = []interface{}{
	(*ReprocessEventsFromRequest)(nil),  // 0: dolittle.runtime.events.processing.managament.ReprocessEventsFromRequest
	(*ReprocessEventsFromResponse)(nil), // 1: dolittle.runtime.events.processing.managament.ReprocessEventsFromResponse
	(*ReprocessAllEventsRequest)(nil),   // 2: dolittle.runtime.events.processing.managament.ReprocessAllEventsRequest
	(*ReprocessAllEventsResponse)(nil),  // 3: dolittle.runtime.events.processing.managament.ReprocessAllEventsResponse
	(*protobuf.Uuid)(nil),               // 4: dolittle.protobuf.Uuid
	(*protobuf.Failure)(nil),            // 5: dolittle.protobuf.Failure
}
var file_Runtime_Management_Events_Processing_EventHandlers_proto_depIdxs = []int32{
	4, // 0: dolittle.runtime.events.processing.managament.ReprocessEventsFromRequest.tenantId:type_name -> dolittle.protobuf.Uuid
	4, // 1: dolittle.runtime.events.processing.managament.ReprocessEventsFromRequest.scopeId:type_name -> dolittle.protobuf.Uuid
	4, // 2: dolittle.runtime.events.processing.managament.ReprocessEventsFromRequest.eventHandlerId:type_name -> dolittle.protobuf.Uuid
	5, // 3: dolittle.runtime.events.processing.managament.ReprocessEventsFromResponse.failure:type_name -> dolittle.protobuf.Failure
	4, // 4: dolittle.runtime.events.processing.managament.ReprocessAllEventsRequest.scopeId:type_name -> dolittle.protobuf.Uuid
	4, // 5: dolittle.runtime.events.processing.managament.ReprocessAllEventsRequest.eventHandlerId:type_name -> dolittle.protobuf.Uuid
	5, // 6: dolittle.runtime.events.processing.managament.ReprocessAllEventsResponse.failure:type_name -> dolittle.protobuf.Failure
	0, // 7: dolittle.runtime.events.processing.managament.EventHandlers.ReprocessEventsFrom:input_type -> dolittle.runtime.events.processing.managament.ReprocessEventsFromRequest
	2, // 8: dolittle.runtime.events.processing.managament.EventHandlers.ReprocessAllEvents:input_type -> dolittle.runtime.events.processing.managament.ReprocessAllEventsRequest
	1, // 9: dolittle.runtime.events.processing.managament.EventHandlers.ReprocessEventsFrom:output_type -> dolittle.runtime.events.processing.managament.ReprocessEventsFromResponse
	3, // 10: dolittle.runtime.events.processing.managament.EventHandlers.ReprocessAllEvents:output_type -> dolittle.runtime.events.processing.managament.ReprocessAllEventsResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Runtime_Management_Events_Processing_EventHandlers_proto_init() }
func file_Runtime_Management_Events_Processing_EventHandlers_proto_init() {
	if File_Runtime_Management_Events_Processing_EventHandlers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReprocessEventsFromRequest); i {
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
		file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReprocessEventsFromResponse); i {
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
		file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReprocessAllEventsRequest); i {
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
		file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReprocessAllEventsResponse); i {
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
			RawDescriptor: file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_Management_Events_Processing_EventHandlers_proto_goTypes,
		DependencyIndexes: file_Runtime_Management_Events_Processing_EventHandlers_proto_depIdxs,
		MessageInfos:      file_Runtime_Management_Events_Processing_EventHandlers_proto_msgTypes,
	}.Build()
	File_Runtime_Management_Events_Processing_EventHandlers_proto = out.File
	file_Runtime_Management_Events_Processing_EventHandlers_proto_rawDesc = nil
	file_Runtime_Management_Events_Processing_EventHandlers_proto_goTypes = nil
	file_Runtime_Management_Events_Processing_EventHandlers_proto_depIdxs = nil
}
