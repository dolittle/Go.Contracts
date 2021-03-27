// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Fundamentals/Services/CallContext.proto

package services

import (
	execution "go.dolittle.io/contracts/v5/execution"
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

type CallRequestContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionContext *execution.ExecutionContext `protobuf:"bytes,1,opt,name=executionContext,proto3" json:"executionContext,omitempty"`
	HeadId           *protobuf.Uuid              `protobuf:"bytes,2,opt,name=headId,proto3" json:"headId,omitempty"`
}

func (x *CallRequestContext) Reset() {
	*x = CallRequestContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Fundamentals_Services_CallContext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequestContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequestContext) ProtoMessage() {}

func (x *CallRequestContext) ProtoReflect() protoreflect.Message {
	mi := &file_Fundamentals_Services_CallContext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequestContext.ProtoReflect.Descriptor instead.
func (*CallRequestContext) Descriptor() ([]byte, []int) {
	return file_Fundamentals_Services_CallContext_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequestContext) GetExecutionContext() *execution.ExecutionContext {
	if x != nil {
		return x.ExecutionContext
	}
	return nil
}

func (x *CallRequestContext) GetHeadId() *protobuf.Uuid {
	if x != nil {
		return x.HeadId
	}
	return nil
}

var File_Fundamentals_Services_CallContext_proto protoreflect.FileDescriptor

var file_Fundamentals_Services_CallContext_proto_rawDesc = []byte{
	0x0a, 0x27, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x2d, 0x46, 0x75,
	0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e,
	0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01,
	0x0a, 0x12, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x50, 0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64, 0x42, 0x44, 0x5a, 0x24, 0x67, 0x6f, 0x2e, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xaa,
	0x02, 0x1b, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Fundamentals_Services_CallContext_proto_rawDescOnce sync.Once
	file_Fundamentals_Services_CallContext_proto_rawDescData = file_Fundamentals_Services_CallContext_proto_rawDesc
)

func file_Fundamentals_Services_CallContext_proto_rawDescGZIP() []byte {
	file_Fundamentals_Services_CallContext_proto_rawDescOnce.Do(func() {
		file_Fundamentals_Services_CallContext_proto_rawDescData = protoimpl.X.CompressGZIP(file_Fundamentals_Services_CallContext_proto_rawDescData)
	})
	return file_Fundamentals_Services_CallContext_proto_rawDescData
}

var file_Fundamentals_Services_CallContext_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Fundamentals_Services_CallContext_proto_goTypes = []interface{}{
	(*CallRequestContext)(nil),         // 0: dolittle.services.CallRequestContext
	(*execution.ExecutionContext)(nil), // 1: dolittle.execution.ExecutionContext
	(*protobuf.Uuid)(nil),              // 2: dolittle.protobuf.Uuid
}
var file_Fundamentals_Services_CallContext_proto_depIdxs = []int32{
	1, // 0: dolittle.services.CallRequestContext.executionContext:type_name -> dolittle.execution.ExecutionContext
	2, // 1: dolittle.services.CallRequestContext.headId:type_name -> dolittle.protobuf.Uuid
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Fundamentals_Services_CallContext_proto_init() }
func file_Fundamentals_Services_CallContext_proto_init() {
	if File_Fundamentals_Services_CallContext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Fundamentals_Services_CallContext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequestContext); i {
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
			RawDescriptor: file_Fundamentals_Services_CallContext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Fundamentals_Services_CallContext_proto_goTypes,
		DependencyIndexes: file_Fundamentals_Services_CallContext_proto_depIdxs,
		MessageInfos:      file_Fundamentals_Services_CallContext_proto_msgTypes,
	}.Build()
	File_Fundamentals_Services_CallContext_proto = out.File
	file_Fundamentals_Services_CallContext_proto_rawDesc = nil
	file_Fundamentals_Services_CallContext_proto_goTypes = nil
	file_Fundamentals_Services_CallContext_proto_depIdxs = nil
}
