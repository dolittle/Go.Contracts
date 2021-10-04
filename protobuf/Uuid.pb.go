// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: Fundamentals/Protobuf/Uuid.proto

package protobuf

import (
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

type Uuid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Uuid) Reset() {
	*x = Uuid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Fundamentals_Protobuf_Uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uuid) ProtoMessage() {}

func (x *Uuid) ProtoReflect() protoreflect.Message {
	mi := &file_Fundamentals_Protobuf_Uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uuid.ProtoReflect.Descriptor instead.
func (*Uuid) Descriptor() ([]byte, []int) {
	return file_Fundamentals_Protobuf_Uuid_proto_rawDescGZIP(), []int{0}
}

func (x *Uuid) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_Fundamentals_Protobuf_Uuid_proto protoreflect.FileDescriptor

var file_Fundamentals_Protobuf_Uuid_proto_rawDesc = []byte{
	0x0a, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x44, 0x5a, 0x24, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x76, 0x36, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xaa, 0x02, 0x1b, 0x44, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Fundamentals_Protobuf_Uuid_proto_rawDescOnce sync.Once
	file_Fundamentals_Protobuf_Uuid_proto_rawDescData = file_Fundamentals_Protobuf_Uuid_proto_rawDesc
)

func file_Fundamentals_Protobuf_Uuid_proto_rawDescGZIP() []byte {
	file_Fundamentals_Protobuf_Uuid_proto_rawDescOnce.Do(func() {
		file_Fundamentals_Protobuf_Uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_Fundamentals_Protobuf_Uuid_proto_rawDescData)
	})
	return file_Fundamentals_Protobuf_Uuid_proto_rawDescData
}

var file_Fundamentals_Protobuf_Uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Fundamentals_Protobuf_Uuid_proto_goTypes = []interface{}{
	(*Uuid)(nil), // 0: dolittle.protobuf.Uuid
}
var file_Fundamentals_Protobuf_Uuid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Fundamentals_Protobuf_Uuid_proto_init() }
func file_Fundamentals_Protobuf_Uuid_proto_init() {
	if File_Fundamentals_Protobuf_Uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Fundamentals_Protobuf_Uuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uuid); i {
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
			RawDescriptor: file_Fundamentals_Protobuf_Uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Fundamentals_Protobuf_Uuid_proto_goTypes,
		DependencyIndexes: file_Fundamentals_Protobuf_Uuid_proto_depIdxs,
		MessageInfos:      file_Fundamentals_Protobuf_Uuid_proto_msgTypes,
	}.Build()
	File_Fundamentals_Protobuf_Uuid_proto = out.File
	file_Fundamentals_Protobuf_Uuid_proto_rawDesc = nil
	file_Fundamentals_Protobuf_Uuid_proto_goTypes = nil
	file_Fundamentals_Protobuf_Uuid_proto_depIdxs = nil
}
