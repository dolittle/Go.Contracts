// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Fundamentals/Versioning/Version.proto

package versioning

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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major            int32  `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor            int32  `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch            int32  `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	Build            int32  `protobuf:"varint,4,opt,name=build,proto3" json:"build,omitempty"`
	PreReleaseString string `protobuf:"bytes,5,opt,name=preReleaseString,proto3" json:"preReleaseString,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Fundamentals_Versioning_Version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_Fundamentals_Versioning_Version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_Fundamentals_Versioning_Version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() int32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

func (x *Version) GetBuild() int32 {
	if x != nil {
		return x.Build
	}
	return 0
}

func (x *Version) GetPreReleaseString() string {
	if x != nil {
		return x.PreReleaseString
	}
	return ""
}

var File_Fundamentals_Versioning_Version_proto protoreflect.FileDescriptor

var file_Fundamentals_Versioning_Version_proto_rawDesc = []byte{
	0x0a, 0x25, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x8d, 0x01, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d,
	0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x65, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x48, 0x5a, 0x26,
	0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0xaa, 0x02, 0x1d, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Fundamentals_Versioning_Version_proto_rawDescOnce sync.Once
	file_Fundamentals_Versioning_Version_proto_rawDescData = file_Fundamentals_Versioning_Version_proto_rawDesc
)

func file_Fundamentals_Versioning_Version_proto_rawDescGZIP() []byte {
	file_Fundamentals_Versioning_Version_proto_rawDescOnce.Do(func() {
		file_Fundamentals_Versioning_Version_proto_rawDescData = protoimpl.X.CompressGZIP(file_Fundamentals_Versioning_Version_proto_rawDescData)
	})
	return file_Fundamentals_Versioning_Version_proto_rawDescData
}

var file_Fundamentals_Versioning_Version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Fundamentals_Versioning_Version_proto_goTypes = []interface{}{
	(*Version)(nil), // 0: dolittle.versioning.Version
}
var file_Fundamentals_Versioning_Version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Fundamentals_Versioning_Version_proto_init() }
func file_Fundamentals_Versioning_Version_proto_init() {
	if File_Fundamentals_Versioning_Version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Fundamentals_Versioning_Version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_Fundamentals_Versioning_Version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Fundamentals_Versioning_Version_proto_goTypes,
		DependencyIndexes: file_Fundamentals_Versioning_Version_proto_depIdxs,
		MessageInfos:      file_Fundamentals_Versioning_Version_proto_msgTypes,
	}.Build()
	File_Fundamentals_Versioning_Version_proto = out.File
	file_Fundamentals_Versioning_Version_proto_rawDesc = nil
	file_Fundamentals_Versioning_Version_proto_goTypes = nil
	file_Fundamentals_Versioning_Version_proto_depIdxs = nil
}
