// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: Runtime/Embeddings/Store.proto

package embeddings

import (
	protobuf "go.dolittle.io/contracts/v6/protobuf"
	projections "go.dolittle.io/contracts/v6/runtime/projections"
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

type GetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	EmbeddingId *protobuf.Uuid               `protobuf:"bytes,2,opt,name=embeddingId,proto3" json:"embeddingId,omitempty"`
	Key         string                       `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetOneRequest) Reset() {
	*x = GetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Embeddings_Store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRequest) ProtoMessage() {}

func (x *GetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Embeddings_Store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneRequest.ProtoReflect.Descriptor instead.
func (*GetOneRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Embeddings_Store_proto_rawDescGZIP(), []int{0}
}

func (x *GetOneRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *GetOneRequest) GetEmbeddingId() *protobuf.Uuid {
	if x != nil {
		return x.EmbeddingId
	}
	return nil
}

func (x *GetOneRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	EmbeddingId *protobuf.Uuid               `protobuf:"bytes,2,opt,name=embeddingId,proto3" json:"embeddingId,omitempty"`
}

func (x *GetKeysRequest) Reset() {
	*x = GetKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Embeddings_Store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeysRequest) ProtoMessage() {}

func (x *GetKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Embeddings_Store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeysRequest.ProtoReflect.Descriptor instead.
func (*GetKeysRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Embeddings_Store_proto_rawDescGZIP(), []int{1}
}

func (x *GetKeysRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *GetKeysRequest) GetEmbeddingId() *protobuf.Uuid {
	if x != nil {
		return x.EmbeddingId
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	EmbeddingId *protobuf.Uuid               `protobuf:"bytes,2,opt,name=embeddingId,proto3" json:"embeddingId,omitempty"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Embeddings_Store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Embeddings_Store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Embeddings_Store_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllRequest) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *GetAllRequest) GetEmbeddingId() *protobuf.Uuid {
	if x != nil {
		return x.EmbeddingId
	}
	return nil
}

type GetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   *projections.ProjectionCurrentState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Failure *protobuf.Failure                   `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
}

func (x *GetOneResponse) Reset() {
	*x = GetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Embeddings_Store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneResponse) ProtoMessage() {}

func (x *GetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Embeddings_Store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneResponse.ProtoReflect.Descriptor instead.
func (*GetOneResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Embeddings_Store_proto_rawDescGZIP(), []int{3}
}

func (x *GetOneResponse) GetState() *projections.ProjectionCurrentState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *GetOneResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

type GetKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys    []string          `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Failure *protobuf.Failure `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
}

func (x *GetKeysResponse) Reset() {
	*x = GetKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Embeddings_Store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeysResponse) ProtoMessage() {}

func (x *GetKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Embeddings_Store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeysResponse.ProtoReflect.Descriptor instead.
func (*GetKeysResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Embeddings_Store_proto_rawDescGZIP(), []int{4}
}

func (x *GetKeysResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *GetKeysResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States  []*projections.ProjectionCurrentState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
	Failure *protobuf.Failure                     `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"` // not set if not failed
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Embeddings_Store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Embeddings_Store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Embeddings_Store_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllResponse) GetStates() []*projections.ProjectionCurrentState {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *GetAllResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

var File_Runtime_Embeddings_Store_proto protoreflect.FileDescriptor

var file_Runtime_Embeddings_Store_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x2f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x23, 0x46,
	0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73,
	0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x73, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x43, 0x61, 0x6c, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5,
	0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x65, 0x6d, 0x62,
	0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0b, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64,
	0x52, 0x0b, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x93, 0x01,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0b, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52,
	0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0x5b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12,
	0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x32, 0xbc, 0x02, 0x0a,
	0x0e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x61, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x2a, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6d, 0x62,
	0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x64, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2b, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x64, 0x6f, 0x6c,
	0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x2a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x58, 0x5a, 0x2e, 0x67,
	0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x36, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0xaa, 0x02, 0x25,
	0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Embeddings_Store_proto_rawDescOnce sync.Once
	file_Runtime_Embeddings_Store_proto_rawDescData = file_Runtime_Embeddings_Store_proto_rawDesc
)

func file_Runtime_Embeddings_Store_proto_rawDescGZIP() []byte {
	file_Runtime_Embeddings_Store_proto_rawDescOnce.Do(func() {
		file_Runtime_Embeddings_Store_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Embeddings_Store_proto_rawDescData)
	})
	return file_Runtime_Embeddings_Store_proto_rawDescData
}

var file_Runtime_Embeddings_Store_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Runtime_Embeddings_Store_proto_goTypes = []interface{}{
	(*GetOneRequest)(nil),                      // 0: dolittle.runtime.embeddings.GetOneRequest
	(*GetKeysRequest)(nil),                     // 1: dolittle.runtime.embeddings.GetKeysRequest
	(*GetAllRequest)(nil),                      // 2: dolittle.runtime.embeddings.GetAllRequest
	(*GetOneResponse)(nil),                     // 3: dolittle.runtime.embeddings.GetOneResponse
	(*GetKeysResponse)(nil),                    // 4: dolittle.runtime.embeddings.GetKeysResponse
	(*GetAllResponse)(nil),                     // 5: dolittle.runtime.embeddings.GetAllResponse
	(*services.CallRequestContext)(nil),        // 6: dolittle.services.CallRequestContext
	(*protobuf.Uuid)(nil),                      // 7: dolittle.protobuf.Uuid
	(*projections.ProjectionCurrentState)(nil), // 8: dolittle.runtime.projections.ProjectionCurrentState
	(*protobuf.Failure)(nil),                   // 9: dolittle.protobuf.Failure
}
var file_Runtime_Embeddings_Store_proto_depIdxs = []int32{
	6,  // 0: dolittle.runtime.embeddings.GetOneRequest.callContext:type_name -> dolittle.services.CallRequestContext
	7,  // 1: dolittle.runtime.embeddings.GetOneRequest.embeddingId:type_name -> dolittle.protobuf.Uuid
	6,  // 2: dolittle.runtime.embeddings.GetKeysRequest.callContext:type_name -> dolittle.services.CallRequestContext
	7,  // 3: dolittle.runtime.embeddings.GetKeysRequest.embeddingId:type_name -> dolittle.protobuf.Uuid
	6,  // 4: dolittle.runtime.embeddings.GetAllRequest.callContext:type_name -> dolittle.services.CallRequestContext
	7,  // 5: dolittle.runtime.embeddings.GetAllRequest.embeddingId:type_name -> dolittle.protobuf.Uuid
	8,  // 6: dolittle.runtime.embeddings.GetOneResponse.state:type_name -> dolittle.runtime.projections.ProjectionCurrentState
	9,  // 7: dolittle.runtime.embeddings.GetOneResponse.failure:type_name -> dolittle.protobuf.Failure
	9,  // 8: dolittle.runtime.embeddings.GetKeysResponse.failure:type_name -> dolittle.protobuf.Failure
	8,  // 9: dolittle.runtime.embeddings.GetAllResponse.states:type_name -> dolittle.runtime.projections.ProjectionCurrentState
	9,  // 10: dolittle.runtime.embeddings.GetAllResponse.failure:type_name -> dolittle.protobuf.Failure
	0,  // 11: dolittle.runtime.embeddings.EmbeddingStore.GetOne:input_type -> dolittle.runtime.embeddings.GetOneRequest
	1,  // 12: dolittle.runtime.embeddings.EmbeddingStore.GetKeys:input_type -> dolittle.runtime.embeddings.GetKeysRequest
	2,  // 13: dolittle.runtime.embeddings.EmbeddingStore.GetAll:input_type -> dolittle.runtime.embeddings.GetAllRequest
	3,  // 14: dolittle.runtime.embeddings.EmbeddingStore.GetOne:output_type -> dolittle.runtime.embeddings.GetOneResponse
	4,  // 15: dolittle.runtime.embeddings.EmbeddingStore.GetKeys:output_type -> dolittle.runtime.embeddings.GetKeysResponse
	5,  // 16: dolittle.runtime.embeddings.EmbeddingStore.GetAll:output_type -> dolittle.runtime.embeddings.GetAllResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_Runtime_Embeddings_Store_proto_init() }
func file_Runtime_Embeddings_Store_proto_init() {
	if File_Runtime_Embeddings_Store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Embeddings_Store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneRequest); i {
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
		file_Runtime_Embeddings_Store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeysRequest); i {
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
		file_Runtime_Embeddings_Store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_Runtime_Embeddings_Store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneResponse); i {
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
		file_Runtime_Embeddings_Store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeysResponse); i {
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
		file_Runtime_Embeddings_Store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
			RawDescriptor: file_Runtime_Embeddings_Store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_Embeddings_Store_proto_goTypes,
		DependencyIndexes: file_Runtime_Embeddings_Store_proto_depIdxs,
		MessageInfos:      file_Runtime_Embeddings_Store_proto_msgTypes,
	}.Build()
	File_Runtime_Embeddings_Store_proto = out.File
	file_Runtime_Embeddings_Store_proto_rawDesc = nil
	file_Runtime_Embeddings_Store_proto_goTypes = nil
	file_Runtime_Embeddings_Store_proto_depIdxs = nil
}
