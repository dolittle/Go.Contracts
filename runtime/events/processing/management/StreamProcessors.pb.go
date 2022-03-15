// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: Runtime/Management/Events.Processing/StreamProcessors.proto

package management

import (
	protobuf "go.dolittle.io/contracts/v7/protobuf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnpartitionedTenantScopedStreamProcessorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFailing     bool                   `protobuf:"varint,1,opt,name=isFailing,proto3" json:"isFailing,omitempty"`
	FailureReason string                 `protobuf:"bytes,2,opt,name=failureReason,proto3" json:"failureReason,omitempty"`
	RetryCount    uint32                 `protobuf:"varint,3,opt,name=retryCount,proto3" json:"retryCount,omitempty"`
	RetryTime     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=retryTime,proto3" json:"retryTime,omitempty"`
}

func (x *UnpartitionedTenantScopedStreamProcessorStatus) Reset() {
	*x = UnpartitionedTenantScopedStreamProcessorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnpartitionedTenantScopedStreamProcessorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnpartitionedTenantScopedStreamProcessorStatus) ProtoMessage() {}

func (x *UnpartitionedTenantScopedStreamProcessorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnpartitionedTenantScopedStreamProcessorStatus.ProtoReflect.Descriptor instead.
func (*UnpartitionedTenantScopedStreamProcessorStatus) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescGZIP(), []int{0}
}

func (x *UnpartitionedTenantScopedStreamProcessorStatus) GetIsFailing() bool {
	if x != nil {
		return x.IsFailing
	}
	return false
}

func (x *UnpartitionedTenantScopedStreamProcessorStatus) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

func (x *UnpartitionedTenantScopedStreamProcessorStatus) GetRetryCount() uint32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

func (x *UnpartitionedTenantScopedStreamProcessorStatus) GetRetryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RetryTime
	}
	return nil
}

type PartitionedTenantScopedStreamProcessorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailingPartitions []*FailingPartition `protobuf:"bytes,1,rep,name=failingPartitions,proto3" json:"failingPartitions,omitempty"`
}

func (x *PartitionedTenantScopedStreamProcessorStatus) Reset() {
	*x = PartitionedTenantScopedStreamProcessorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartitionedTenantScopedStreamProcessorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionedTenantScopedStreamProcessorStatus) ProtoMessage() {}

func (x *PartitionedTenantScopedStreamProcessorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionedTenantScopedStreamProcessorStatus.ProtoReflect.Descriptor instead.
func (*PartitionedTenantScopedStreamProcessorStatus) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescGZIP(), []int{1}
}

func (x *PartitionedTenantScopedStreamProcessorStatus) GetFailingPartitions() []*FailingPartition {
	if x != nil {
		return x.FailingPartitions
	}
	return nil
}

type FailingPartition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartitionId    string                 `protobuf:"bytes,1,opt,name=partitionId,proto3" json:"partitionId,omitempty"`
	StreamPosition uint64                 `protobuf:"varint,2,opt,name=streamPosition,proto3" json:"streamPosition,omitempty"`
	FailureReason  string                 `protobuf:"bytes,3,opt,name=failureReason,proto3" json:"failureReason,omitempty"`
	RetryCount     uint32                 `protobuf:"varint,4,opt,name=retryCount,proto3" json:"retryCount,omitempty"`
	RetryTime      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=retryTime,proto3" json:"retryTime,omitempty"`
	LastFailed     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=lastFailed,proto3" json:"lastFailed,omitempty"`
}

func (x *FailingPartition) Reset() {
	*x = FailingPartition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailingPartition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailingPartition) ProtoMessage() {}

func (x *FailingPartition) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailingPartition.ProtoReflect.Descriptor instead.
func (*FailingPartition) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescGZIP(), []int{2}
}

func (x *FailingPartition) GetPartitionId() string {
	if x != nil {
		return x.PartitionId
	}
	return ""
}

func (x *FailingPartition) GetStreamPosition() uint64 {
	if x != nil {
		return x.StreamPosition
	}
	return 0
}

func (x *FailingPartition) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

func (x *FailingPartition) GetRetryCount() uint32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

func (x *FailingPartition) GetRetryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RetryTime
	}
	return nil
}

func (x *FailingPartition) GetLastFailed() *timestamppb.Timestamp {
	if x != nil {
		return x.LastFailed
	}
	return nil
}

type TenantScopedStreamProcessorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId                  *protobuf.Uuid         `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	StreamPosition            uint64                 `protobuf:"varint,2,opt,name=streamPosition,proto3" json:"streamPosition,omitempty"`
	LastSuccessfullyProcessed *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=lastSuccessfullyProcessed,proto3" json:"lastSuccessfullyProcessed,omitempty"`
	// Types that are assignable to Status:
	//	*TenantScopedStreamProcessorStatus_Unpartitioned
	//	*TenantScopedStreamProcessorStatus_Partitioned
	Status isTenantScopedStreamProcessorStatus_Status `protobuf_oneof:"Status"`
}

func (x *TenantScopedStreamProcessorStatus) Reset() {
	*x = TenantScopedStreamProcessorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantScopedStreamProcessorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantScopedStreamProcessorStatus) ProtoMessage() {}

func (x *TenantScopedStreamProcessorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantScopedStreamProcessorStatus.ProtoReflect.Descriptor instead.
func (*TenantScopedStreamProcessorStatus) Descriptor() ([]byte, []int) {
	return file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescGZIP(), []int{3}
}

func (x *TenantScopedStreamProcessorStatus) GetTenantId() *protobuf.Uuid {
	if x != nil {
		return x.TenantId
	}
	return nil
}

func (x *TenantScopedStreamProcessorStatus) GetStreamPosition() uint64 {
	if x != nil {
		return x.StreamPosition
	}
	return 0
}

func (x *TenantScopedStreamProcessorStatus) GetLastSuccessfullyProcessed() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSuccessfullyProcessed
	}
	return nil
}

func (m *TenantScopedStreamProcessorStatus) GetStatus() isTenantScopedStreamProcessorStatus_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (x *TenantScopedStreamProcessorStatus) GetUnpartitioned() *UnpartitionedTenantScopedStreamProcessorStatus {
	if x, ok := x.GetStatus().(*TenantScopedStreamProcessorStatus_Unpartitioned); ok {
		return x.Unpartitioned
	}
	return nil
}

func (x *TenantScopedStreamProcessorStatus) GetPartitioned() *PartitionedTenantScopedStreamProcessorStatus {
	if x, ok := x.GetStatus().(*TenantScopedStreamProcessorStatus_Partitioned); ok {
		return x.Partitioned
	}
	return nil
}

type isTenantScopedStreamProcessorStatus_Status interface {
	isTenantScopedStreamProcessorStatus_Status()
}

type TenantScopedStreamProcessorStatus_Unpartitioned struct {
	Unpartitioned *UnpartitionedTenantScopedStreamProcessorStatus `protobuf:"bytes,4,opt,name=unpartitioned,proto3,oneof"`
}

type TenantScopedStreamProcessorStatus_Partitioned struct {
	Partitioned *PartitionedTenantScopedStreamProcessorStatus `protobuf:"bytes,5,opt,name=partitioned,proto3,oneof"`
}

func (*TenantScopedStreamProcessorStatus_Unpartitioned) isTenantScopedStreamProcessorStatus_Status() {
}

func (*TenantScopedStreamProcessorStatus_Partitioned) isTenantScopedStreamProcessorStatus_Status() {}

var File_Runtime_Management_Events_Processing_StreamProcessors_proto protoreflect.FileDescriptor

var file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x20, 0x46, 0x75,
	0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xce, 0x01, 0x0a, 0x2e, 0x55, 0x6e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x9d, 0x01, 0x0a, 0x2c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x6d, 0x0a, 0x11, 0x66, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x66,
	0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x98, 0x02, 0x0a, 0x10, 0x46, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0xed, 0x03, 0x0a, 0x21,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58,
	0x0a, 0x19, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x6c, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x19, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x85, 0x01, 0x0a, 0x0d, 0x75, 0x6e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x5d, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x6e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x0d, 0x75, 0x6e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64,
	0x12, 0x7f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5b, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x64, 0x42, 0x08, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x7c, 0x5a, 0x40, 0x67,
	0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xaa,
	0x02, 0x37, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescOnce sync.Once
	file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescData = file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDesc
)

func file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescGZIP() []byte {
	file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescOnce.Do(func() {
		file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescData)
	})
	return file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDescData
}

var file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Runtime_Management_Events_Processing_StreamProcessors_proto_goTypes = []interface{}{
	(*UnpartitionedTenantScopedStreamProcessorStatus)(nil), // 0: dolittle.runtime.events.processing.management.UnpartitionedTenantScopedStreamProcessorStatus
	(*PartitionedTenantScopedStreamProcessorStatus)(nil),   // 1: dolittle.runtime.events.processing.management.PartitionedTenantScopedStreamProcessorStatus
	(*FailingPartition)(nil),                               // 2: dolittle.runtime.events.processing.management.FailingPartition
	(*TenantScopedStreamProcessorStatus)(nil),              // 3: dolittle.runtime.events.processing.management.TenantScopedStreamProcessorStatus
	(*timestamppb.Timestamp)(nil),                          // 4: google.protobuf.Timestamp
	(*protobuf.Uuid)(nil),                                  // 5: dolittle.protobuf.Uuid
}
var file_Runtime_Management_Events_Processing_StreamProcessors_proto_depIdxs = []int32{
	4, // 0: dolittle.runtime.events.processing.management.UnpartitionedTenantScopedStreamProcessorStatus.retryTime:type_name -> google.protobuf.Timestamp
	2, // 1: dolittle.runtime.events.processing.management.PartitionedTenantScopedStreamProcessorStatus.failingPartitions:type_name -> dolittle.runtime.events.processing.management.FailingPartition
	4, // 2: dolittle.runtime.events.processing.management.FailingPartition.retryTime:type_name -> google.protobuf.Timestamp
	4, // 3: dolittle.runtime.events.processing.management.FailingPartition.lastFailed:type_name -> google.protobuf.Timestamp
	5, // 4: dolittle.runtime.events.processing.management.TenantScopedStreamProcessorStatus.tenantId:type_name -> dolittle.protobuf.Uuid
	4, // 5: dolittle.runtime.events.processing.management.TenantScopedStreamProcessorStatus.lastSuccessfullyProcessed:type_name -> google.protobuf.Timestamp
	0, // 6: dolittle.runtime.events.processing.management.TenantScopedStreamProcessorStatus.unpartitioned:type_name -> dolittle.runtime.events.processing.management.UnpartitionedTenantScopedStreamProcessorStatus
	1, // 7: dolittle.runtime.events.processing.management.TenantScopedStreamProcessorStatus.partitioned:type_name -> dolittle.runtime.events.processing.management.PartitionedTenantScopedStreamProcessorStatus
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_Runtime_Management_Events_Processing_StreamProcessors_proto_init() }
func file_Runtime_Management_Events_Processing_StreamProcessors_proto_init() {
	if File_Runtime_Management_Events_Processing_StreamProcessors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnpartitionedTenantScopedStreamProcessorStatus); i {
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
		file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartitionedTenantScopedStreamProcessorStatus); i {
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
		file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailingPartition); i {
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
		file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantScopedStreamProcessorStatus); i {
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
	file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TenantScopedStreamProcessorStatus_Unpartitioned)(nil),
		(*TenantScopedStreamProcessorStatus_Partitioned)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Runtime_Management_Events_Processing_StreamProcessors_proto_goTypes,
		DependencyIndexes: file_Runtime_Management_Events_Processing_StreamProcessors_proto_depIdxs,
		MessageInfos:      file_Runtime_Management_Events_Processing_StreamProcessors_proto_msgTypes,
	}.Build()
	File_Runtime_Management_Events_Processing_StreamProcessors_proto = out.File
	file_Runtime_Management_Events_Processing_StreamProcessors_proto_rawDesc = nil
	file_Runtime_Management_Events_Processing_StreamProcessors_proto_goTypes = nil
	file_Runtime_Management_Events_Processing_StreamProcessors_proto_depIdxs = nil
}
