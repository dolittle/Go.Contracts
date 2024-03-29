// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: Execution/ExecutionContext.proto

package execution

import (
	protobuf "go.dolittle.io/contracts/v7/protobuf"
	security "go.dolittle.io/contracts/v7/security"
	versioning "go.dolittle.io/contracts/v7/versioning"
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

type ExecutionContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MicroserviceId *protobuf.Uuid      `protobuf:"bytes,1,opt,name=microserviceId,proto3" json:"microserviceId,omitempty"`
	TenantId       *protobuf.Uuid      `protobuf:"bytes,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Version        *versioning.Version `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	CorrelationId  *protobuf.Uuid      `protobuf:"bytes,4,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	Claims         []*security.Claim   `protobuf:"bytes,5,rep,name=claims,proto3" json:"claims,omitempty"`
	Environment    string              `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	SpanId         []byte              `protobuf:"bytes,7,opt,name=spanId,proto3,oneof" json:"spanId,omitempty"`
}

func (x *ExecutionContext) Reset() {
	*x = ExecutionContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Execution_ExecutionContext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionContext) ProtoMessage() {}

func (x *ExecutionContext) ProtoReflect() protoreflect.Message {
	mi := &file_Execution_ExecutionContext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionContext.ProtoReflect.Descriptor instead.
func (*ExecutionContext) Descriptor() ([]byte, []int) {
	return file_Execution_ExecutionContext_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionContext) GetMicroserviceId() *protobuf.Uuid {
	if x != nil {
		return x.MicroserviceId
	}
	return nil
}

func (x *ExecutionContext) GetTenantId() *protobuf.Uuid {
	if x != nil {
		return x.TenantId
	}
	return nil
}

func (x *ExecutionContext) GetVersion() *versioning.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *ExecutionContext) GetCorrelationId() *protobuf.Uuid {
	if x != nil {
		return x.CorrelationId
	}
	return nil
}

func (x *ExecutionContext) GetClaims() []*security.Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *ExecutionContext) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *ExecutionContext) GetSpanId() []byte {
	if x != nil {
		return x.SpanId
	}
	return nil
}

var File_Execution_ExecutionContext_proto protoreflect.FileDescriptor

var file_Execution_ExecutionContext_proto_rawDesc = []byte{
	0x0a, 0x20, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02, 0x0a, 0x10,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x3f, 0x0a, 0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69,
	0x64, 0x52, 0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0d,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x42, 0x46, 0x5a, 0x25, 0x67, 0x6f, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0xaa, 0x02, 0x1c, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Execution_ExecutionContext_proto_rawDescOnce sync.Once
	file_Execution_ExecutionContext_proto_rawDescData = file_Execution_ExecutionContext_proto_rawDesc
)

func file_Execution_ExecutionContext_proto_rawDescGZIP() []byte {
	file_Execution_ExecutionContext_proto_rawDescOnce.Do(func() {
		file_Execution_ExecutionContext_proto_rawDescData = protoimpl.X.CompressGZIP(file_Execution_ExecutionContext_proto_rawDescData)
	})
	return file_Execution_ExecutionContext_proto_rawDescData
}

var file_Execution_ExecutionContext_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Execution_ExecutionContext_proto_goTypes = []interface{}{
	(*ExecutionContext)(nil),   // 0: dolittle.execution.ExecutionContext
	(*protobuf.Uuid)(nil),      // 1: dolittle.protobuf.Uuid
	(*versioning.Version)(nil), // 2: dolittle.versioning.Version
	(*security.Claim)(nil),     // 3: dolittle.security.Claim
}
var file_Execution_ExecutionContext_proto_depIdxs = []int32{
	1, // 0: dolittle.execution.ExecutionContext.microserviceId:type_name -> dolittle.protobuf.Uuid
	1, // 1: dolittle.execution.ExecutionContext.tenantId:type_name -> dolittle.protobuf.Uuid
	2, // 2: dolittle.execution.ExecutionContext.version:type_name -> dolittle.versioning.Version
	1, // 3: dolittle.execution.ExecutionContext.correlationId:type_name -> dolittle.protobuf.Uuid
	3, // 4: dolittle.execution.ExecutionContext.claims:type_name -> dolittle.security.Claim
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Execution_ExecutionContext_proto_init() }
func file_Execution_ExecutionContext_proto_init() {
	if File_Execution_ExecutionContext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Execution_ExecutionContext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionContext); i {
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
	file_Execution_ExecutionContext_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Execution_ExecutionContext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Execution_ExecutionContext_proto_goTypes,
		DependencyIndexes: file_Execution_ExecutionContext_proto_depIdxs,
		MessageInfos:      file_Execution_ExecutionContext_proto_msgTypes,
	}.Build()
	File_Execution_ExecutionContext_proto = out.File
	file_Execution_ExecutionContext_proto_rawDesc = nil
	file_Execution_ExecutionContext_proto_goTypes = nil
	file_Execution_ExecutionContext_proto_depIdxs = nil
}
