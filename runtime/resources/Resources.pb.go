// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: Runtime/Resources/Resources.proto

package resources

import (
	protobuf "go.dolittle.io/contracts/v6/protobuf"
	services "go.dolittle.io/contracts/v6/services"
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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Resources_Resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Resources_Resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Resources_Resources_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

type GetMongoDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionString string            `protobuf:"bytes,1,opt,name=connectionString,proto3" json:"connectionString,omitempty"`
	Failure          *protobuf.Failure `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
}

func (x *GetMongoDBResponse) Reset() {
	*x = GetMongoDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Resources_Resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMongoDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMongoDBResponse) ProtoMessage() {}

func (x *GetMongoDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Resources_Resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMongoDBResponse.ProtoReflect.Descriptor instead.
func (*GetMongoDBResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Resources_Resources_proto_rawDescGZIP(), []int{1}
}

func (x *GetMongoDBResponse) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

func (x *GetMongoDBResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

var File_Runtime_Resources_Resources_proto protoreflect.FileDescriptor

var file_Runtime_Resources_Resources_proto_rawDesc = []byte{
	0x0a, 0x21, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x23, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x73, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x43, 0x61, 0x6c, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x63,
	0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x76, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x67, 0x6f,
	0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x32, 0x71, 0x0a, 0x09,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x12, 0x26, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x56, 0x5a, 0x2d, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x36, 0x2f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xaa, 0x02, 0x24, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Resources_Resources_proto_rawDescOnce sync.Once
	file_Runtime_Resources_Resources_proto_rawDescData = file_Runtime_Resources_Resources_proto_rawDesc
)

func file_Runtime_Resources_Resources_proto_rawDescGZIP() []byte {
	file_Runtime_Resources_Resources_proto_rawDescOnce.Do(func() {
		file_Runtime_Resources_Resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Resources_Resources_proto_rawDescData)
	})
	return file_Runtime_Resources_Resources_proto_rawDescData
}

var file_Runtime_Resources_Resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Runtime_Resources_Resources_proto_goTypes = []interface{}{
	(*GetRequest)(nil),                  // 0: dolittle.runtime.resources.GetRequest
	(*GetMongoDBResponse)(nil),          // 1: dolittle.runtime.resources.GetMongoDBResponse
	(*services.CallRequestContext)(nil), // 2: dolittle.services.CallRequestContext
	(*protobuf.Failure)(nil),            // 3: dolittle.protobuf.Failure
}
var file_Runtime_Resources_Resources_proto_depIdxs = []int32{
	2, // 0: dolittle.runtime.resources.GetRequest.callContext:type_name -> dolittle.services.CallRequestContext
	3, // 1: dolittle.runtime.resources.GetMongoDBResponse.failure:type_name -> dolittle.protobuf.Failure
	0, // 2: dolittle.runtime.resources.Resources.GetMongoDB:input_type -> dolittle.runtime.resources.GetRequest
	1, // 3: dolittle.runtime.resources.Resources.GetMongoDB:output_type -> dolittle.runtime.resources.GetMongoDBResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Runtime_Resources_Resources_proto_init() }
func file_Runtime_Resources_Resources_proto_init() {
	if File_Runtime_Resources_Resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Resources_Resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_Runtime_Resources_Resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMongoDBResponse); i {
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
			RawDescriptor: file_Runtime_Resources_Resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_Resources_Resources_proto_goTypes,
		DependencyIndexes: file_Runtime_Resources_Resources_proto_depIdxs,
		MessageInfos:      file_Runtime_Resources_Resources_proto_msgTypes,
	}.Build()
	File_Runtime_Resources_Resources_proto = out.File
	file_Runtime_Resources_Resources_proto_rawDesc = nil
	file_Runtime_Resources_Resources_proto_goTypes = nil
	file_Runtime_Resources_Resources_proto_depIdxs = nil
}