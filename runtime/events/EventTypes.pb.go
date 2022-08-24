// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: Runtime/Events/EventTypes.proto

package events

import (
	artifacts "go.dolittle.io/contracts/v7/artifacts"
	protobuf "go.dolittle.io/contracts/v7/protobuf"
	services "go.dolittle.io/contracts/v7/services"
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

type EventTypeRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	EventType   *artifacts.Artifact          `protobuf:"bytes,2,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Alias       *string                      `protobuf:"bytes,3,opt,name=alias,proto3,oneof" json:"alias,omitempty"`
}

func (x *EventTypeRegistrationRequest) Reset() {
	*x = EventTypeRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventTypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeRegistrationRequest) ProtoMessage() {}

func (x *EventTypeRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventTypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeRegistrationRequest.ProtoReflect.Descriptor instead.
func (*EventTypeRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventTypes_proto_rawDescGZIP(), []int{0}
}

func (x *EventTypeRegistrationRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *EventTypeRegistrationRequest) GetEventType() *artifacts.Artifact {
	if x != nil {
		return x.EventType
	}
	return nil
}

func (x *EventTypeRegistrationRequest) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

type EventTypeRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *protobuf.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *EventTypeRegistrationResponse) Reset() {
	*x = EventTypeRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_EventTypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeRegistrationResponse) ProtoMessage() {}

func (x *EventTypeRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_EventTypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeRegistrationResponse.ProtoReflect.Descriptor instead.
func (*EventTypeRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_EventTypes_proto_rawDescGZIP(), []int{1}
}

func (x *EventTypeRegistrationResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

var File_Runtime_Events_EventTypes_proto protoreflect.FileDescriptor

var file_Runtime_Events_EventTypes_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x18, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x1c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x22, 0x55, 0x0a, 0x1d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x0a, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x79, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x35, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x2a, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x2f, 0x76, 0x37, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0xaa, 0x02, 0x21, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Events_EventTypes_proto_rawDescOnce sync.Once
	file_Runtime_Events_EventTypes_proto_rawDescData = file_Runtime_Events_EventTypes_proto_rawDesc
)

func file_Runtime_Events_EventTypes_proto_rawDescGZIP() []byte {
	file_Runtime_Events_EventTypes_proto_rawDescOnce.Do(func() {
		file_Runtime_Events_EventTypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Events_EventTypes_proto_rawDescData)
	})
	return file_Runtime_Events_EventTypes_proto_rawDescData
}

var file_Runtime_Events_EventTypes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Runtime_Events_EventTypes_proto_goTypes = []interface{}{
	(*EventTypeRegistrationRequest)(nil),  // 0: dolittle.runtime.events.EventTypeRegistrationRequest
	(*EventTypeRegistrationResponse)(nil), // 1: dolittle.runtime.events.EventTypeRegistrationResponse
	(*services.CallRequestContext)(nil),   // 2: dolittle.services.CallRequestContext
	(*artifacts.Artifact)(nil),            // 3: dolittle.artifacts.Artifact
	(*protobuf.Failure)(nil),              // 4: dolittle.protobuf.Failure
}
var file_Runtime_Events_EventTypes_proto_depIdxs = []int32{
	2, // 0: dolittle.runtime.events.EventTypeRegistrationRequest.callContext:type_name -> dolittle.services.CallRequestContext
	3, // 1: dolittle.runtime.events.EventTypeRegistrationRequest.eventType:type_name -> dolittle.artifacts.Artifact
	4, // 2: dolittle.runtime.events.EventTypeRegistrationResponse.failure:type_name -> dolittle.protobuf.Failure
	0, // 3: dolittle.runtime.events.EventTypes.Register:input_type -> dolittle.runtime.events.EventTypeRegistrationRequest
	1, // 4: dolittle.runtime.events.EventTypes.Register:output_type -> dolittle.runtime.events.EventTypeRegistrationResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Runtime_Events_EventTypes_proto_init() }
func file_Runtime_Events_EventTypes_proto_init() {
	if File_Runtime_Events_EventTypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Events_EventTypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeRegistrationRequest); i {
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
		file_Runtime_Events_EventTypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeRegistrationResponse); i {
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
	file_Runtime_Events_EventTypes_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Runtime_Events_EventTypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_Events_EventTypes_proto_goTypes,
		DependencyIndexes: file_Runtime_Events_EventTypes_proto_depIdxs,
		MessageInfos:      file_Runtime_Events_EventTypes_proto_msgTypes,
	}.Build()
	File_Runtime_Events_EventTypes_proto = out.File
	file_Runtime_Events_EventTypes_proto_rawDesc = nil
	file_Runtime_Events_EventTypes_proto_goTypes = nil
	file_Runtime_Events_EventTypes_proto_depIdxs = nil
}
