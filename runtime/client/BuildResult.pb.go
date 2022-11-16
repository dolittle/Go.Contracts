// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: Runtime/Client/BuildResult.proto

package client

import (
	artifacts "go.dolittle.io/contracts/v7/artifacts"
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

type BuildResult_Type int32

const (
	BuildResult_INFORMATION BuildResult_Type = 0
	BuildResult_FAILURE     BuildResult_Type = 1
	BuildResult_ERROR       BuildResult_Type = 2
)

// Enum value maps for BuildResult_Type.
var (
	BuildResult_Type_name = map[int32]string{
		0: "INFORMATION",
		1: "FAILURE",
		2: "ERROR",
	}
	BuildResult_Type_value = map[string]int32{
		"INFORMATION": 0,
		"FAILURE":     1,
		"ERROR":       2,
	}
)

func (x BuildResult_Type) Enum() *BuildResult_Type {
	p := new(BuildResult_Type)
	*p = x
	return p
}

func (x BuildResult_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildResult_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_Runtime_Client_BuildResult_proto_enumTypes[0].Descriptor()
}

func (BuildResult_Type) Type() protoreflect.EnumType {
	return &file_Runtime_Client_BuildResult_proto_enumTypes[0]
}

func (x BuildResult_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildResult_Type.Descriptor instead.
func (BuildResult_Type) EnumDescriptor() ([]byte, []int) {
	return file_Runtime_Client_BuildResult_proto_rawDescGZIP(), []int{0, 0}
}

type BuildResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string           `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Type    BuildResult_Type `protobuf:"varint,2,opt,name=type,proto3,enum=dolittle.runtime.client.BuildResult_Type" json:"type,omitempty"`
}

func (x *BuildResult) Reset() {
	*x = BuildResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Client_BuildResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResult) ProtoMessage() {}

func (x *BuildResult) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Client_BuildResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResult.ProtoReflect.Descriptor instead.
func (*BuildResult) Descriptor() ([]byte, []int) {
	return file_Runtime_Client_BuildResult_proto_rawDescGZIP(), []int{0}
}

func (x *BuildResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BuildResult) GetType() BuildResult_Type {
	if x != nil {
		return x.Type
	}
	return BuildResult_INFORMATION
}

type ArtifactBuildResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aritfact    *artifacts.Artifact `protobuf:"bytes,1,opt,name=aritfact,proto3" json:"aritfact,omitempty"`
	Alias       string              `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	BuildResult *BuildResult        `protobuf:"bytes,3,opt,name=buildResult,proto3" json:"buildResult,omitempty"`
}

func (x *ArtifactBuildResult) Reset() {
	*x = ArtifactBuildResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Client_BuildResult_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactBuildResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactBuildResult) ProtoMessage() {}

func (x *ArtifactBuildResult) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Client_BuildResult_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactBuildResult.ProtoReflect.Descriptor instead.
func (*ArtifactBuildResult) Descriptor() ([]byte, []int) {
	return file_Runtime_Client_BuildResult_proto_rawDescGZIP(), []int{1}
}

func (x *ArtifactBuildResult) GetAritfact() *artifacts.Artifact {
	if x != nil {
		return x.Aritfact
	}
	return nil
}

func (x *ArtifactBuildResult) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *ArtifactBuildResult) GetBuildResult() *BuildResult {
	if x != nil {
		return x.BuildResult
	}
	return nil
}

type BuildResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventTypes     []*ArtifactBuildResult `protobuf:"bytes,1,rep,name=eventTypes,proto3" json:"eventTypes,omitempty"`
	AggregateRoots []*ArtifactBuildResult `protobuf:"bytes,2,rep,name=aggregateRoots,proto3" json:"aggregateRoots,omitempty"`
	EventHandlers  []*ArtifactBuildResult `protobuf:"bytes,3,rep,name=eventHandlers,proto3" json:"eventHandlers,omitempty"`
	Projections    []*ArtifactBuildResult `protobuf:"bytes,4,rep,name=projections,proto3" json:"projections,omitempty"`
	Embeddings     []*ArtifactBuildResult `protobuf:"bytes,5,rep,name=embeddings,proto3" json:"embeddings,omitempty"`
	Filters        []*ArtifactBuildResult `protobuf:"bytes,6,rep,name=filters,proto3" json:"filters,omitempty"`
	Other          []*BuildResult         `protobuf:"bytes,7,rep,name=other,proto3" json:"other,omitempty"`
}

func (x *BuildResults) Reset() {
	*x = BuildResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Client_BuildResult_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResults) ProtoMessage() {}

func (x *BuildResults) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Client_BuildResult_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResults.ProtoReflect.Descriptor instead.
func (*BuildResults) Descriptor() ([]byte, []int) {
	return file_Runtime_Client_BuildResult_proto_rawDescGZIP(), []int{2}
}

func (x *BuildResults) GetEventTypes() []*ArtifactBuildResult {
	if x != nil {
		return x.EventTypes
	}
	return nil
}

func (x *BuildResults) GetAggregateRoots() []*ArtifactBuildResult {
	if x != nil {
		return x.AggregateRoots
	}
	return nil
}

func (x *BuildResults) GetEventHandlers() []*ArtifactBuildResult {
	if x != nil {
		return x.EventHandlers
	}
	return nil
}

func (x *BuildResults) GetProjections() []*ArtifactBuildResult {
	if x != nil {
		return x.Projections
	}
	return nil
}

func (x *BuildResults) GetEmbeddings() []*ArtifactBuildResult {
	if x != nil {
		return x.Embeddings
	}
	return nil
}

func (x *BuildResults) GetFilters() []*ArtifactBuildResult {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *BuildResults) GetOther() []*BuildResult {
	if x != nil {
		return x.Other
	}
	return nil
}

var File_Runtime_Client_BuildResult_proto protoreflect.FileDescriptor

var file_Runtime_Client_BuildResult_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2f,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22,
	0xad, 0x01, 0x0a, 0x13, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x72, 0x69, 0x74, 0x66,
	0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x08, 0x61, 0x72, 0x69, 0x74, 0x66, 0x61, 0x63,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0xa8, 0x04, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x4c, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x54,
	0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x0a, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x46, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3a,
	0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x50, 0x5a, 0x2a, 0x67, 0x6f,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x21, 0x44, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Client_BuildResult_proto_rawDescOnce sync.Once
	file_Runtime_Client_BuildResult_proto_rawDescData = file_Runtime_Client_BuildResult_proto_rawDesc
)

func file_Runtime_Client_BuildResult_proto_rawDescGZIP() []byte {
	file_Runtime_Client_BuildResult_proto_rawDescOnce.Do(func() {
		file_Runtime_Client_BuildResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Client_BuildResult_proto_rawDescData)
	})
	return file_Runtime_Client_BuildResult_proto_rawDescData
}

var file_Runtime_Client_BuildResult_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Runtime_Client_BuildResult_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Runtime_Client_BuildResult_proto_goTypes = []interface{}{
	(BuildResult_Type)(0),       // 0: dolittle.runtime.client.BuildResult.Type
	(*BuildResult)(nil),         // 1: dolittle.runtime.client.BuildResult
	(*ArtifactBuildResult)(nil), // 2: dolittle.runtime.client.ArtifactBuildResult
	(*BuildResults)(nil),        // 3: dolittle.runtime.client.BuildResults
	(*artifacts.Artifact)(nil),  // 4: dolittle.artifacts.Artifact
}
var file_Runtime_Client_BuildResult_proto_depIdxs = []int32{
	0,  // 0: dolittle.runtime.client.BuildResult.type:type_name -> dolittle.runtime.client.BuildResult.Type
	4,  // 1: dolittle.runtime.client.ArtifactBuildResult.aritfact:type_name -> dolittle.artifacts.Artifact
	1,  // 2: dolittle.runtime.client.ArtifactBuildResult.buildResult:type_name -> dolittle.runtime.client.BuildResult
	2,  // 3: dolittle.runtime.client.BuildResults.eventTypes:type_name -> dolittle.runtime.client.ArtifactBuildResult
	2,  // 4: dolittle.runtime.client.BuildResults.aggregateRoots:type_name -> dolittle.runtime.client.ArtifactBuildResult
	2,  // 5: dolittle.runtime.client.BuildResults.eventHandlers:type_name -> dolittle.runtime.client.ArtifactBuildResult
	2,  // 6: dolittle.runtime.client.BuildResults.projections:type_name -> dolittle.runtime.client.ArtifactBuildResult
	2,  // 7: dolittle.runtime.client.BuildResults.embeddings:type_name -> dolittle.runtime.client.ArtifactBuildResult
	2,  // 8: dolittle.runtime.client.BuildResults.filters:type_name -> dolittle.runtime.client.ArtifactBuildResult
	1,  // 9: dolittle.runtime.client.BuildResults.other:type_name -> dolittle.runtime.client.BuildResult
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_Runtime_Client_BuildResult_proto_init() }
func file_Runtime_Client_BuildResult_proto_init() {
	if File_Runtime_Client_BuildResult_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Client_BuildResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResult); i {
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
		file_Runtime_Client_BuildResult_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactBuildResult); i {
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
		file_Runtime_Client_BuildResult_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResults); i {
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
			RawDescriptor: file_Runtime_Client_BuildResult_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Runtime_Client_BuildResult_proto_goTypes,
		DependencyIndexes: file_Runtime_Client_BuildResult_proto_depIdxs,
		EnumInfos:         file_Runtime_Client_BuildResult_proto_enumTypes,
		MessageInfos:      file_Runtime_Client_BuildResult_proto_msgTypes,
	}.Build()
	File_Runtime_Client_BuildResult_proto = out.File
	file_Runtime_Client_BuildResult_proto_rawDesc = nil
	file_Runtime_Client_BuildResult_proto_goTypes = nil
	file_Runtime_Client_BuildResult_proto_depIdxs = nil
}