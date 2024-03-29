// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: Runtime/Events/Committed.proto

package events

import (
	artifacts "go.dolittle.io/contracts/v7/artifacts"
	execution "go.dolittle.io/contracts/v7/execution"
	protobuf "go.dolittle.io/contracts/v7/protobuf"
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

type CommittedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventLogSequenceNumber         uint64                      `protobuf:"varint,1,opt,name=eventLogSequenceNumber,proto3" json:"eventLogSequenceNumber,omitempty"`
	Occurred                       *timestamppb.Timestamp      `protobuf:"bytes,2,opt,name=occurred,proto3" json:"occurred,omitempty"`
	EventSourceId                  string                      `protobuf:"bytes,3,opt,name=eventSourceId,proto3" json:"eventSourceId,omitempty"`
	ExecutionContext               *execution.ExecutionContext `protobuf:"bytes,4,opt,name=executionContext,proto3" json:"executionContext,omitempty"`
	EventType                      *artifacts.Artifact         `protobuf:"bytes,5,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Public                         bool                        `protobuf:"varint,6,opt,name=public,proto3" json:"public,omitempty"`
	Content                        string                      `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	External                       bool                        `protobuf:"varint,8,opt,name=external,proto3" json:"external,omitempty"`
	ExternalEventLogSequenceNumber uint64                      `protobuf:"varint,9,opt,name=externalEventLogSequenceNumber,proto3" json:"externalEventLogSequenceNumber,omitempty"`
	ExternalEventReceived          *timestamppb.Timestamp      `protobuf:"bytes,10,opt,name=externalEventReceived,proto3" json:"externalEventReceived,omitempty"`
}

func (x *CommittedEvent) Reset() {
	*x = CommittedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Committed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommittedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommittedEvent) ProtoMessage() {}

func (x *CommittedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Committed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommittedEvent.ProtoReflect.Descriptor instead.
func (*CommittedEvent) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Committed_proto_rawDescGZIP(), []int{0}
}

func (x *CommittedEvent) GetEventLogSequenceNumber() uint64 {
	if x != nil {
		return x.EventLogSequenceNumber
	}
	return 0
}

func (x *CommittedEvent) GetOccurred() *timestamppb.Timestamp {
	if x != nil {
		return x.Occurred
	}
	return nil
}

func (x *CommittedEvent) GetEventSourceId() string {
	if x != nil {
		return x.EventSourceId
	}
	return ""
}

func (x *CommittedEvent) GetExecutionContext() *execution.ExecutionContext {
	if x != nil {
		return x.ExecutionContext
	}
	return nil
}

func (x *CommittedEvent) GetEventType() *artifacts.Artifact {
	if x != nil {
		return x.EventType
	}
	return nil
}

func (x *CommittedEvent) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *CommittedEvent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommittedEvent) GetExternal() bool {
	if x != nil {
		return x.External
	}
	return false
}

func (x *CommittedEvent) GetExternalEventLogSequenceNumber() uint64 {
	if x != nil {
		return x.ExternalEventLogSequenceNumber
	}
	return 0
}

func (x *CommittedEvent) GetExternalEventReceived() *timestamppb.Timestamp {
	if x != nil {
		return x.ExternalEventReceived
	}
	return nil
}

type CommittedAggregateEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventSourceId   string         `protobuf:"bytes,1,opt,name=eventSourceId,proto3" json:"eventSourceId,omitempty"`
	AggregateRootId *protobuf.Uuid `protobuf:"bytes,2,opt,name=aggregateRootId,proto3" json:"aggregateRootId,omitempty"`
	// Deprecated: Do not use.
	AggregateRootVersion        uint64                                              `protobuf:"varint,3,opt,name=aggregateRootVersion,proto3" json:"aggregateRootVersion,omitempty"` // DEPRECATED Replaced by currentAggregateRootVersion
	Events                      []*CommittedAggregateEvents_CommittedAggregateEvent `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
	CurrentAggregateRootVersion uint64                                              `protobuf:"varint,5,opt,name=currentAggregateRootVersion,proto3" json:"currentAggregateRootVersion,omitempty"` // Represents the current version of the aggregate root
}

func (x *CommittedAggregateEvents) Reset() {
	*x = CommittedAggregateEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Committed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommittedAggregateEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommittedAggregateEvents) ProtoMessage() {}

func (x *CommittedAggregateEvents) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Committed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommittedAggregateEvents.ProtoReflect.Descriptor instead.
func (*CommittedAggregateEvents) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Committed_proto_rawDescGZIP(), []int{1}
}

func (x *CommittedAggregateEvents) GetEventSourceId() string {
	if x != nil {
		return x.EventSourceId
	}
	return ""
}

func (x *CommittedAggregateEvents) GetAggregateRootId() *protobuf.Uuid {
	if x != nil {
		return x.AggregateRootId
	}
	return nil
}

// Deprecated: Do not use.
func (x *CommittedAggregateEvents) GetAggregateRootVersion() uint64 {
	if x != nil {
		return x.AggregateRootVersion
	}
	return 0
}

func (x *CommittedAggregateEvents) GetEvents() []*CommittedAggregateEvents_CommittedAggregateEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *CommittedAggregateEvents) GetCurrentAggregateRootVersion() uint64 {
	if x != nil {
		return x.CurrentAggregateRootVersion
	}
	return 0
}

type CommittedAggregateEvents_CommittedAggregateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventLogSequenceNumber uint64                      `protobuf:"varint,1,opt,name=eventLogSequenceNumber,proto3" json:"eventLogSequenceNumber,omitempty"`
	Occurred               *timestamppb.Timestamp      `protobuf:"bytes,2,opt,name=occurred,proto3" json:"occurred,omitempty"`
	ExecutionContext       *execution.ExecutionContext `protobuf:"bytes,3,opt,name=executionContext,proto3" json:"executionContext,omitempty"`
	EventType              *artifacts.Artifact         `protobuf:"bytes,4,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Public                 bool                        `protobuf:"varint,5,opt,name=public,proto3" json:"public,omitempty"`
	Content                string                      `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	AggregateRootVersion   uint64                      `protobuf:"varint,7,opt,name=aggregateRootVersion,proto3" json:"aggregateRootVersion,omitempty"` // The aggregate root version the event was applied to
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) Reset() {
	*x = CommittedAggregateEvents_CommittedAggregateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Committed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommittedAggregateEvents_CommittedAggregateEvent) ProtoMessage() {}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Committed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommittedAggregateEvents_CommittedAggregateEvent.ProtoReflect.Descriptor instead.
func (*CommittedAggregateEvents_CommittedAggregateEvent) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Committed_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) GetEventLogSequenceNumber() uint64 {
	if x != nil {
		return x.EventLogSequenceNumber
	}
	return 0
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) GetOccurred() *timestamppb.Timestamp {
	if x != nil {
		return x.Occurred
	}
	return nil
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) GetExecutionContext() *execution.ExecutionContext {
	if x != nil {
		return x.ExecutionContext
	}
	return nil
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) GetEventType() *artifacts.Artifact {
	if x != nil {
		return x.EventType
	}
	return nil
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommittedAggregateEvents_CommittedAggregateEvent) GetAggregateRootVersion() uint64 {
	if x != nil {
		return x.AggregateRootVersion
	}
	return 0
}

var File_Runtime_Events_Committed_proto protoreflect.FileDescriptor

var file_Runtime_Events_Committed_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x18, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36,
	0x0a, 0x16, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x12, 0x46, 0x0a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x22, 0xe0, 0x05, 0x0a, 0x18, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a,
	0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52,
	0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x36, 0x0a, 0x14, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x14, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x1b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x1b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xfd, 0x02,
	0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x36, 0x0a, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x10, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x50, 0x5a,
	0x2a, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0xaa, 0x02, 0x21, 0x44, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Events_Committed_proto_rawDescOnce sync.Once
	file_Runtime_Events_Committed_proto_rawDescData = file_Runtime_Events_Committed_proto_rawDesc
)

func file_Runtime_Events_Committed_proto_rawDescGZIP() []byte {
	file_Runtime_Events_Committed_proto_rawDescOnce.Do(func() {
		file_Runtime_Events_Committed_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Events_Committed_proto_rawDescData)
	})
	return file_Runtime_Events_Committed_proto_rawDescData
}

var file_Runtime_Events_Committed_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Runtime_Events_Committed_proto_goTypes = []interface{}{
	(*CommittedEvent)(nil),                                   // 0: dolittle.runtime.events.CommittedEvent
	(*CommittedAggregateEvents)(nil),                         // 1: dolittle.runtime.events.CommittedAggregateEvents
	(*CommittedAggregateEvents_CommittedAggregateEvent)(nil), // 2: dolittle.runtime.events.CommittedAggregateEvents.CommittedAggregateEvent
	(*timestamppb.Timestamp)(nil),                            // 3: google.protobuf.Timestamp
	(*execution.ExecutionContext)(nil),                       // 4: dolittle.execution.ExecutionContext
	(*artifacts.Artifact)(nil),                               // 5: dolittle.artifacts.Artifact
	(*protobuf.Uuid)(nil),                                    // 6: dolittle.protobuf.Uuid
}
var file_Runtime_Events_Committed_proto_depIdxs = []int32{
	3, // 0: dolittle.runtime.events.CommittedEvent.occurred:type_name -> google.protobuf.Timestamp
	4, // 1: dolittle.runtime.events.CommittedEvent.executionContext:type_name -> dolittle.execution.ExecutionContext
	5, // 2: dolittle.runtime.events.CommittedEvent.eventType:type_name -> dolittle.artifacts.Artifact
	3, // 3: dolittle.runtime.events.CommittedEvent.externalEventReceived:type_name -> google.protobuf.Timestamp
	6, // 4: dolittle.runtime.events.CommittedAggregateEvents.aggregateRootId:type_name -> dolittle.protobuf.Uuid
	2, // 5: dolittle.runtime.events.CommittedAggregateEvents.events:type_name -> dolittle.runtime.events.CommittedAggregateEvents.CommittedAggregateEvent
	3, // 6: dolittle.runtime.events.CommittedAggregateEvents.CommittedAggregateEvent.occurred:type_name -> google.protobuf.Timestamp
	4, // 7: dolittle.runtime.events.CommittedAggregateEvents.CommittedAggregateEvent.executionContext:type_name -> dolittle.execution.ExecutionContext
	5, // 8: dolittle.runtime.events.CommittedAggregateEvents.CommittedAggregateEvent.eventType:type_name -> dolittle.artifacts.Artifact
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_Runtime_Events_Committed_proto_init() }
func file_Runtime_Events_Committed_proto_init() {
	if File_Runtime_Events_Committed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Events_Committed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommittedEvent); i {
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
		file_Runtime_Events_Committed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommittedAggregateEvents); i {
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
		file_Runtime_Events_Committed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommittedAggregateEvents_CommittedAggregateEvent); i {
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
			RawDescriptor: file_Runtime_Events_Committed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Runtime_Events_Committed_proto_goTypes,
		DependencyIndexes: file_Runtime_Events_Committed_proto_depIdxs,
		MessageInfos:      file_Runtime_Events_Committed_proto_msgTypes,
	}.Build()
	File_Runtime_Events_Committed_proto = out.File
	file_Runtime_Events_Committed_proto_rawDesc = nil
	file_Runtime_Events_Committed_proto_goTypes = nil
	file_Runtime_Events_Committed_proto_depIdxs = nil
}
