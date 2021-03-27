// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Runtime/Events/Uncommitted.proto

package events

import (
	artifacts "go.dolittle.io/contracts/v5/artifacts"
	protobuf "go.dolittle.io/contracts/v5/protobuf"
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

type UncommittedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artifact      *artifacts.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	EventSourceId *protobuf.Uuid      `protobuf:"bytes,2,opt,name=eventSourceId,proto3" json:"eventSourceId,omitempty"`
	Public        bool                `protobuf:"varint,3,opt,name=public,proto3" json:"public,omitempty"`
	Content       string              `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UncommittedEvent) Reset() {
	*x = UncommittedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Uncommitted_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UncommittedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncommittedEvent) ProtoMessage() {}

func (x *UncommittedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Uncommitted_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncommittedEvent.ProtoReflect.Descriptor instead.
func (*UncommittedEvent) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Uncommitted_proto_rawDescGZIP(), []int{0}
}

func (x *UncommittedEvent) GetArtifact() *artifacts.Artifact {
	if x != nil {
		return x.Artifact
	}
	return nil
}

func (x *UncommittedEvent) GetEventSourceId() *protobuf.Uuid {
	if x != nil {
		return x.EventSourceId
	}
	return nil
}

func (x *UncommittedEvent) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *UncommittedEvent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UncommittedAggregateEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregateRootId              *protobuf.Uuid                                          `protobuf:"bytes,1,opt,name=aggregateRootId,proto3" json:"aggregateRootId,omitempty"`
	EventSourceId                *protobuf.Uuid                                          `protobuf:"bytes,2,opt,name=eventSourceId,proto3" json:"eventSourceId,omitempty"`
	ExpectedAggregateRootVersion uint64                                                  `protobuf:"varint,3,opt,name=expectedAggregateRootVersion,proto3" json:"expectedAggregateRootVersion,omitempty"`
	Events                       []*UncommittedAggregateEvents_UncommittedAggregateEvent `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *UncommittedAggregateEvents) Reset() {
	*x = UncommittedAggregateEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Uncommitted_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UncommittedAggregateEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncommittedAggregateEvents) ProtoMessage() {}

func (x *UncommittedAggregateEvents) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Uncommitted_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncommittedAggregateEvents.ProtoReflect.Descriptor instead.
func (*UncommittedAggregateEvents) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Uncommitted_proto_rawDescGZIP(), []int{1}
}

func (x *UncommittedAggregateEvents) GetAggregateRootId() *protobuf.Uuid {
	if x != nil {
		return x.AggregateRootId
	}
	return nil
}

func (x *UncommittedAggregateEvents) GetEventSourceId() *protobuf.Uuid {
	if x != nil {
		return x.EventSourceId
	}
	return nil
}

func (x *UncommittedAggregateEvents) GetExpectedAggregateRootVersion() uint64 {
	if x != nil {
		return x.ExpectedAggregateRootVersion
	}
	return 0
}

func (x *UncommittedAggregateEvents) GetEvents() []*UncommittedAggregateEvents_UncommittedAggregateEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type UncommittedAggregateEvents_UncommittedAggregateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artifact *artifacts.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	Public   bool                `protobuf:"varint,2,opt,name=public,proto3" json:"public,omitempty"`
	Content  string              `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UncommittedAggregateEvents_UncommittedAggregateEvent) Reset() {
	*x = UncommittedAggregateEvents_UncommittedAggregateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Uncommitted_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UncommittedAggregateEvents_UncommittedAggregateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncommittedAggregateEvents_UncommittedAggregateEvent) ProtoMessage() {}

func (x *UncommittedAggregateEvents_UncommittedAggregateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Uncommitted_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncommittedAggregateEvents_UncommittedAggregateEvent.ProtoReflect.Descriptor instead.
func (*UncommittedAggregateEvents_UncommittedAggregateEvent) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Uncommitted_proto_rawDescGZIP(), []int{1, 0}
}

func (x *UncommittedAggregateEvents_UncommittedAggregateEvent) GetArtifact() *artifacts.Artifact {
	if x != nil {
		return x.Artifact
	}
	return nil
}

func (x *UncommittedAggregateEvents_UncommittedAggregateEvent) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *UncommittedAggregateEvents_UncommittedAggregateEvent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_Runtime_Events_Uncommitted_proto protoreflect.FileDescriptor

var file_Runtime_Events_Uncommitted_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x25, 0x46, 0x75, 0x6e,
	0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73,
	0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x10, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c,
	0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0xd3, 0x03, 0x0a, 0x1a, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x1c, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1c, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55,
	0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0x87, 0x01, 0x0a, 0x19, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38,
	0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x08,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x50, 0x5a, 0x2a, 0x67, 0x6f,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0xaa, 0x02, 0x21, 0x44, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Events_Uncommitted_proto_rawDescOnce sync.Once
	file_Runtime_Events_Uncommitted_proto_rawDescData = file_Runtime_Events_Uncommitted_proto_rawDesc
)

func file_Runtime_Events_Uncommitted_proto_rawDescGZIP() []byte {
	file_Runtime_Events_Uncommitted_proto_rawDescOnce.Do(func() {
		file_Runtime_Events_Uncommitted_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Events_Uncommitted_proto_rawDescData)
	})
	return file_Runtime_Events_Uncommitted_proto_rawDescData
}

var file_Runtime_Events_Uncommitted_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Runtime_Events_Uncommitted_proto_goTypes = []interface{}{
	(*UncommittedEvent)(nil),                                     // 0: dolittle.runtime.events.UncommittedEvent
	(*UncommittedAggregateEvents)(nil),                           // 1: dolittle.runtime.events.UncommittedAggregateEvents
	(*UncommittedAggregateEvents_UncommittedAggregateEvent)(nil), // 2: dolittle.runtime.events.UncommittedAggregateEvents.UncommittedAggregateEvent
	(*artifacts.Artifact)(nil),                                   // 3: dolittle.artifacts.Artifact
	(*protobuf.Uuid)(nil),                                        // 4: dolittle.protobuf.Uuid
}
var file_Runtime_Events_Uncommitted_proto_depIdxs = []int32{
	3, // 0: dolittle.runtime.events.UncommittedEvent.artifact:type_name -> dolittle.artifacts.Artifact
	4, // 1: dolittle.runtime.events.UncommittedEvent.eventSourceId:type_name -> dolittle.protobuf.Uuid
	4, // 2: dolittle.runtime.events.UncommittedAggregateEvents.aggregateRootId:type_name -> dolittle.protobuf.Uuid
	4, // 3: dolittle.runtime.events.UncommittedAggregateEvents.eventSourceId:type_name -> dolittle.protobuf.Uuid
	2, // 4: dolittle.runtime.events.UncommittedAggregateEvents.events:type_name -> dolittle.runtime.events.UncommittedAggregateEvents.UncommittedAggregateEvent
	3, // 5: dolittle.runtime.events.UncommittedAggregateEvents.UncommittedAggregateEvent.artifact:type_name -> dolittle.artifacts.Artifact
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Runtime_Events_Uncommitted_proto_init() }
func file_Runtime_Events_Uncommitted_proto_init() {
	if File_Runtime_Events_Uncommitted_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Events_Uncommitted_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UncommittedEvent); i {
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
		file_Runtime_Events_Uncommitted_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UncommittedAggregateEvents); i {
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
		file_Runtime_Events_Uncommitted_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UncommittedAggregateEvents_UncommittedAggregateEvent); i {
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
			RawDescriptor: file_Runtime_Events_Uncommitted_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Runtime_Events_Uncommitted_proto_goTypes,
		DependencyIndexes: file_Runtime_Events_Uncommitted_proto_depIdxs,
		MessageInfos:      file_Runtime_Events_Uncommitted_proto_msgTypes,
	}.Build()
	File_Runtime_Events_Uncommitted_proto = out.File
	file_Runtime_Events_Uncommitted_proto_rawDesc = nil
	file_Runtime_Events_Uncommitted_proto_goTypes = nil
	file_Runtime_Events_Uncommitted_proto_depIdxs = nil
}
