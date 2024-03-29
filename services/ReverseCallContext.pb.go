// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: Services/ReverseCallContext.proto

package services

import (
	execution "go.dolittle.io/contracts/v7/execution"
	protobuf "go.dolittle.io/contracts/v7/protobuf"
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

type ReverseCallArgumentsContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionContext *execution.ExecutionContext `protobuf:"bytes,1,opt,name=executionContext,proto3" json:"executionContext,omitempty"`
	HeadId           *protobuf.Uuid              `protobuf:"bytes,2,opt,name=headId,proto3" json:"headId,omitempty"`
	PingInterval     *durationpb.Duration        `protobuf:"bytes,3,opt,name=pingInterval,proto3" json:"pingInterval,omitempty"`
}

func (x *ReverseCallArgumentsContext) Reset() {
	*x = ReverseCallArgumentsContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Services_ReverseCallContext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseCallArgumentsContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseCallArgumentsContext) ProtoMessage() {}

func (x *ReverseCallArgumentsContext) ProtoReflect() protoreflect.Message {
	mi := &file_Services_ReverseCallContext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseCallArgumentsContext.ProtoReflect.Descriptor instead.
func (*ReverseCallArgumentsContext) Descriptor() ([]byte, []int) {
	return file_Services_ReverseCallContext_proto_rawDescGZIP(), []int{0}
}

func (x *ReverseCallArgumentsContext) GetExecutionContext() *execution.ExecutionContext {
	if x != nil {
		return x.ExecutionContext
	}
	return nil
}

func (x *ReverseCallArgumentsContext) GetHeadId() *protobuf.Uuid {
	if x != nil {
		return x.HeadId
	}
	return nil
}

func (x *ReverseCallArgumentsContext) GetPingInterval() *durationpb.Duration {
	if x != nil {
		return x.PingInterval
	}
	return nil
}

type ReverseCallRequestContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId           *protobuf.Uuid              `protobuf:"bytes,1,opt,name=callId,proto3" json:"callId,omitempty"`
	ExecutionContext *execution.ExecutionContext `protobuf:"bytes,2,opt,name=executionContext,proto3" json:"executionContext,omitempty"`
}

func (x *ReverseCallRequestContext) Reset() {
	*x = ReverseCallRequestContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Services_ReverseCallContext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseCallRequestContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseCallRequestContext) ProtoMessage() {}

func (x *ReverseCallRequestContext) ProtoReflect() protoreflect.Message {
	mi := &file_Services_ReverseCallContext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseCallRequestContext.ProtoReflect.Descriptor instead.
func (*ReverseCallRequestContext) Descriptor() ([]byte, []int) {
	return file_Services_ReverseCallContext_proto_rawDescGZIP(), []int{1}
}

func (x *ReverseCallRequestContext) GetCallId() *protobuf.Uuid {
	if x != nil {
		return x.CallId
	}
	return nil
}

func (x *ReverseCallRequestContext) GetExecutionContext() *execution.ExecutionContext {
	if x != nil {
		return x.ExecutionContext
	}
	return nil
}

type ReverseCallResponseContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId *protobuf.Uuid `protobuf:"bytes,1,opt,name=callId,proto3" json:"callId,omitempty"`
}

func (x *ReverseCallResponseContext) Reset() {
	*x = ReverseCallResponseContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Services_ReverseCallContext_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseCallResponseContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseCallResponseContext) ProtoMessage() {}

func (x *ReverseCallResponseContext) ProtoReflect() protoreflect.Message {
	mi := &file_Services_ReverseCallContext_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseCallResponseContext.ProtoReflect.Descriptor instead.
func (*ReverseCallResponseContext) Descriptor() ([]byte, []int) {
	return file_Services_ReverseCallContext_proto_rawDescGZIP(), []int{2}
}

func (x *ReverseCallResponseContext) GetCallId() *protobuf.Uuid {
	if x != nil {
		return x.CallId
	}
	return nil
}

var File_Services_ReverseCallContext_proto protoreflect.FileDescriptor

var file_Services_ReverseCallContext_proto_rawDesc = []byte{
	0x0a, 0x21, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x20, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01,
	0x0a, 0x1b, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x50, 0x0a,
	0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x10, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x2f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22,
	0x9e, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a,
	0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x50,
	0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x10,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x4d, 0x0a, 0x1a, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2f,
	0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x42,
	0x44, 0x5a, 0x24, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xaa, 0x02, 0x1b, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Services_ReverseCallContext_proto_rawDescOnce sync.Once
	file_Services_ReverseCallContext_proto_rawDescData = file_Services_ReverseCallContext_proto_rawDesc
)

func file_Services_ReverseCallContext_proto_rawDescGZIP() []byte {
	file_Services_ReverseCallContext_proto_rawDescOnce.Do(func() {
		file_Services_ReverseCallContext_proto_rawDescData = protoimpl.X.CompressGZIP(file_Services_ReverseCallContext_proto_rawDescData)
	})
	return file_Services_ReverseCallContext_proto_rawDescData
}

var file_Services_ReverseCallContext_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Services_ReverseCallContext_proto_goTypes = []interface{}{
	(*ReverseCallArgumentsContext)(nil), // 0: dolittle.services.ReverseCallArgumentsContext
	(*ReverseCallRequestContext)(nil),   // 1: dolittle.services.ReverseCallRequestContext
	(*ReverseCallResponseContext)(nil),  // 2: dolittle.services.ReverseCallResponseContext
	(*execution.ExecutionContext)(nil),  // 3: dolittle.execution.ExecutionContext
	(*protobuf.Uuid)(nil),               // 4: dolittle.protobuf.Uuid
	(*durationpb.Duration)(nil),         // 5: google.protobuf.Duration
}
var file_Services_ReverseCallContext_proto_depIdxs = []int32{
	3, // 0: dolittle.services.ReverseCallArgumentsContext.executionContext:type_name -> dolittle.execution.ExecutionContext
	4, // 1: dolittle.services.ReverseCallArgumentsContext.headId:type_name -> dolittle.protobuf.Uuid
	5, // 2: dolittle.services.ReverseCallArgumentsContext.pingInterval:type_name -> google.protobuf.Duration
	4, // 3: dolittle.services.ReverseCallRequestContext.callId:type_name -> dolittle.protobuf.Uuid
	3, // 4: dolittle.services.ReverseCallRequestContext.executionContext:type_name -> dolittle.execution.ExecutionContext
	4, // 5: dolittle.services.ReverseCallResponseContext.callId:type_name -> dolittle.protobuf.Uuid
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Services_ReverseCallContext_proto_init() }
func file_Services_ReverseCallContext_proto_init() {
	if File_Services_ReverseCallContext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Services_ReverseCallContext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseCallArgumentsContext); i {
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
		file_Services_ReverseCallContext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseCallRequestContext); i {
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
		file_Services_ReverseCallContext_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseCallResponseContext); i {
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
			RawDescriptor: file_Services_ReverseCallContext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Services_ReverseCallContext_proto_goTypes,
		DependencyIndexes: file_Services_ReverseCallContext_proto_depIdxs,
		MessageInfos:      file_Services_ReverseCallContext_proto_msgTypes,
	}.Build()
	File_Services_ReverseCallContext_proto = out.File
	file_Services_ReverseCallContext_proto_rawDesc = nil
	file_Services_ReverseCallContext_proto_goTypes = nil
	file_Services_ReverseCallContext_proto_depIdxs = nil
}
