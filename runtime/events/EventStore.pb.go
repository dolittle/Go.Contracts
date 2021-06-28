// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.15.6
// source: Runtime/Events/EventStore.proto

package events

import (
	protobuf "go.dolittle.io/contracts/v5/protobuf"
	services "go.dolittle.io/contracts/v5/services"
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

type CommitEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Events      []*UncommittedEvent          `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *CommitEventsRequest) Reset() {
	*x = CommitEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventStore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitEventsRequest) ProtoMessage() {}

func (x *CommitEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventStore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitEventsRequest.ProtoReflect.Descriptor instead.
func (*CommitEventsRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventStore_proto_rawDescGZIP(), []int{0}
}

func (x *CommitEventsRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *CommitEventsRequest) GetEvents() []*UncommittedEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type CommitAggregateEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Events      *UncommittedAggregateEvents  `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *CommitAggregateEventsRequest) Reset() {
	*x = CommitAggregateEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventStore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAggregateEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAggregateEventsRequest) ProtoMessage() {}

func (x *CommitAggregateEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventStore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAggregateEventsRequest.ProtoReflect.Descriptor instead.
func (*CommitAggregateEventsRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventStore_proto_rawDescGZIP(), []int{1}
}

func (x *CommitAggregateEventsRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *CommitAggregateEventsRequest) GetEvents() *UncommittedAggregateEvents {
	if x != nil {
		return x.Events
	}
	return nil
}

type FetchForAggregateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Aggregate   *Aggregate                   `protobuf:"bytes,2,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
}

func (x *FetchForAggregateRequest) Reset() {
	*x = FetchForAggregateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventStore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchForAggregateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchForAggregateRequest) ProtoMessage() {}

func (x *FetchForAggregateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventStore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchForAggregateRequest.ProtoReflect.Descriptor instead.
func (*FetchForAggregateRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventStore_proto_rawDescGZIP(), []int{2}
}

func (x *FetchForAggregateRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *FetchForAggregateRequest) GetAggregate() *Aggregate {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

type CommitEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *protobuf.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
	Events  []*CommittedEvent `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *CommitEventsResponse) Reset() {
	*x = CommitEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventStore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitEventsResponse) ProtoMessage() {}

func (x *CommitEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventStore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitEventsResponse.ProtoReflect.Descriptor instead.
func (*CommitEventsResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventStore_proto_rawDescGZIP(), []int{3}
}

func (x *CommitEventsResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *CommitEventsResponse) GetEvents() []*CommittedEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type CommitAggregateEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *protobuf.Failure         `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
	Events  *CommittedAggregateEvents `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *CommitAggregateEventsResponse) Reset() {
	*x = CommitAggregateEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventStore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitAggregateEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitAggregateEventsResponse) ProtoMessage() {}

func (x *CommitAggregateEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventStore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitAggregateEventsResponse.ProtoReflect.Descriptor instead.
func (*CommitAggregateEventsResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventStore_proto_rawDescGZIP(), []int{4}
}

func (x *CommitAggregateEventsResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *CommitAggregateEventsResponse) GetEvents() *CommittedAggregateEvents {
	if x != nil {
		return x.Events
	}
	return nil
}

type FetchForAggregateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *protobuf.Failure         `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
	Events  *CommittedAggregateEvents `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *FetchForAggregateResponse) Reset() {
	*x = FetchForAggregateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventStore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchForAggregateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchForAggregateResponse) ProtoMessage() {}

func (x *FetchForAggregateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventStore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchForAggregateResponse.ProtoReflect.Descriptor instead.
func (*FetchForAggregateResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventStore_proto_rawDescGZIP(), []int{5}
}

func (x *FetchForAggregateResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *FetchForAggregateResponse) GetEvents() *CommittedAggregateEvents {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_Runtime_Events_EventStore_proto protoreflect.FileDescriptor

var file_Runtime_Events_EventStore_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x23, 0x46, 0x75, 0x6e, 0x64,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b,
	0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xb4,
	0x01, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46,
	0x6f, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b,
	0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x22, 0x8d, 0x01,
	0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x3f, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa0, 0x01,
	0x0a, 0x1d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x9c, 0x01, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32,
	0xf5, 0x02, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x65,
	0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x46, 0x6f, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7a, 0x0a, 0x11, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x12, 0x31, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x46, 0x6f, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x46, 0x6f, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x2a, 0x67, 0x6f, 0x2e, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0xaa, 0x02, 0x21, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Runtime_Events_EventStore_proto_rawDescOnce sync.Once
	file_Runtime_Events_EventStore_proto_rawDescData = file_Runtime_Events_EventStore_proto_rawDesc
)

func file_Runtime_Events_EventStore_proto_rawDescGZIP() []byte {
	file_Runtime_Events_EventStore_proto_rawDescOnce.Do(func() {
		file_Runtime_Events_EventStore_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Events_EventStore_proto_rawDescData)
	})
	return file_Runtime_Events_EventStore_proto_rawDescData
}

var file_Runtime_Events_EventStore_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Runtime_Events_EventStore_proto_goTypes = []interface{}{
	(*CommitEventsRequest)(nil),           // 0: dolittle.runtime.events.CommitEventsRequest
	(*CommitAggregateEventsRequest)(nil),  // 1: dolittle.runtime.events.CommitAggregateEventsRequest
	(*FetchForAggregateRequest)(nil),      // 2: dolittle.runtime.events.FetchForAggregateRequest
	(*CommitEventsResponse)(nil),          // 3: dolittle.runtime.events.CommitEventsResponse
	(*CommitAggregateEventsResponse)(nil), // 4: dolittle.runtime.events.CommitAggregateEventsResponse
	(*FetchForAggregateResponse)(nil),     // 5: dolittle.runtime.events.FetchForAggregateResponse
	(*services.CallRequestContext)(nil),   // 6: dolittle.services.CallRequestContext
	(*UncommittedEvent)(nil),              // 7: dolittle.runtime.events.UncommittedEvent
	(*UncommittedAggregateEvents)(nil),    // 8: dolittle.runtime.events.UncommittedAggregateEvents
	(*Aggregate)(nil),                     // 9: dolittle.runtime.events.Aggregate
	(*protobuf.Failure)(nil),              // 10: dolittle.protobuf.Failure
	(*CommittedEvent)(nil),                // 11: dolittle.runtime.events.CommittedEvent
	(*CommittedAggregateEvents)(nil),      // 12: dolittle.runtime.events.CommittedAggregateEvents
}
var file_Runtime_Events_EventStore_proto_depIdxs = []int32{
	6,  // 0: dolittle.runtime.events.CommitEventsRequest.callContext:type_name -> dolittle.services.CallRequestContext
	7,  // 1: dolittle.runtime.events.CommitEventsRequest.events:type_name -> dolittle.runtime.events.UncommittedEvent
	6,  // 2: dolittle.runtime.events.CommitAggregateEventsRequest.callContext:type_name -> dolittle.services.CallRequestContext
	8,  // 3: dolittle.runtime.events.CommitAggregateEventsRequest.events:type_name -> dolittle.runtime.events.UncommittedAggregateEvents
	6,  // 4: dolittle.runtime.events.FetchForAggregateRequest.callContext:type_name -> dolittle.services.CallRequestContext
	9,  // 5: dolittle.runtime.events.FetchForAggregateRequest.aggregate:type_name -> dolittle.runtime.events.Aggregate
	10, // 6: dolittle.runtime.events.CommitEventsResponse.failure:type_name -> dolittle.protobuf.Failure
	11, // 7: dolittle.runtime.events.CommitEventsResponse.events:type_name -> dolittle.runtime.events.CommittedEvent
	10, // 8: dolittle.runtime.events.CommitAggregateEventsResponse.failure:type_name -> dolittle.protobuf.Failure
	12, // 9: dolittle.runtime.events.CommitAggregateEventsResponse.events:type_name -> dolittle.runtime.events.CommittedAggregateEvents
	10, // 10: dolittle.runtime.events.FetchForAggregateResponse.failure:type_name -> dolittle.protobuf.Failure
	12, // 11: dolittle.runtime.events.FetchForAggregateResponse.events:type_name -> dolittle.runtime.events.CommittedAggregateEvents
	0,  // 12: dolittle.runtime.events.EventStore.Commit:input_type -> dolittle.runtime.events.CommitEventsRequest
	1,  // 13: dolittle.runtime.events.EventStore.CommitForAggregate:input_type -> dolittle.runtime.events.CommitAggregateEventsRequest
	2,  // 14: dolittle.runtime.events.EventStore.FetchForAggregate:input_type -> dolittle.runtime.events.FetchForAggregateRequest
	3,  // 15: dolittle.runtime.events.EventStore.Commit:output_type -> dolittle.runtime.events.CommitEventsResponse
	4,  // 16: dolittle.runtime.events.EventStore.CommitForAggregate:output_type -> dolittle.runtime.events.CommitAggregateEventsResponse
	5,  // 17: dolittle.runtime.events.EventStore.FetchForAggregate:output_type -> dolittle.runtime.events.FetchForAggregateResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_Runtime_Events_EventStore_proto_init() }
func file_Runtime_Events_EventStore_proto_init() {
	if File_Runtime_Events_EventStore_proto != nil {
		return
	}
	file_Runtime_Events_Aggregate_proto_init()
	file_Runtime_Events_Committed_proto_init()
	file_Runtime_Events_Uncommitted_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Events_EventStore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitEventsRequest); i {
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
		file_Runtime_Events_EventStore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAggregateEventsRequest); i {
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
		file_Runtime_Events_EventStore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchForAggregateRequest); i {
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
		file_Runtime_Events_EventStore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitEventsResponse); i {
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
		file_Runtime_Events_EventStore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitAggregateEventsResponse); i {
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
		file_Runtime_Events_EventStore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchForAggregateResponse); i {
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
			RawDescriptor: file_Runtime_Events_EventStore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_Events_EventStore_proto_goTypes,
		DependencyIndexes: file_Runtime_Events_EventStore_proto_depIdxs,
		MessageInfos:      file_Runtime_Events_EventStore_proto_msgTypes,
	}.Build()
	File_Runtime_Events_EventStore_proto = out.File
	file_Runtime_Events_EventStore_proto_rawDesc = nil
	file_Runtime_Events_EventStore_proto_goTypes = nil
	file_Runtime_Events_EventStore_proto_depIdxs = nil
}
