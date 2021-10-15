// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: Runtime/EventHorizon/Consumer.proto

package eventhorizon

import (
	protobuf "go.dolittle.io/contracts/v6/protobuf"
	events "go.dolittle.io/contracts/v6/runtime/events"
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

type EventHorizonEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamSequenceNumber uint64                 `protobuf:"varint,1,opt,name=streamSequenceNumber,proto3" json:"streamSequenceNumber,omitempty"`
	Event                *events.CommittedEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventHorizonEvent) Reset() {
	*x = EventHorizonEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHorizonEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHorizonEvent) ProtoMessage() {}

func (x *EventHorizonEvent) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHorizonEvent.ProtoReflect.Descriptor instead.
func (*EventHorizonEvent) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Consumer_proto_rawDescGZIP(), []int{0}
}

func (x *EventHorizonEvent) GetStreamSequenceNumber() uint64 {
	if x != nil {
		return x.StreamSequenceNumber
	}
	return 0
}

func (x *EventHorizonEvent) GetEvent() *events.CommittedEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type ConsumerSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext    *services.ReverseCallArgumentsContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	TenantId       *protobuf.Uuid                        `protobuf:"bytes,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	StreamId       *protobuf.Uuid                        `protobuf:"bytes,3,opt,name=streamId,proto3" json:"streamId,omitempty"`
	PartitionId    string                                `protobuf:"bytes,4,opt,name=partitionId,proto3" json:"partitionId,omitempty"`
	StreamPosition uint64                                `protobuf:"varint,5,opt,name=streamPosition,proto3" json:"streamPosition,omitempty"`
}

func (x *ConsumerSubscriptionRequest) Reset() {
	*x = ConsumerSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerSubscriptionRequest) ProtoMessage() {}

func (x *ConsumerSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*ConsumerSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Consumer_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumerSubscriptionRequest) GetCallContext() *services.ReverseCallArgumentsContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *ConsumerSubscriptionRequest) GetTenantId() *protobuf.Uuid {
	if x != nil {
		return x.TenantId
	}
	return nil
}

func (x *ConsumerSubscriptionRequest) GetStreamId() *protobuf.Uuid {
	if x != nil {
		return x.StreamId
	}
	return nil
}

func (x *ConsumerSubscriptionRequest) GetPartitionId() string {
	if x != nil {
		return x.PartitionId
	}
	return ""
}

func (x *ConsumerSubscriptionRequest) GetStreamPosition() uint64 {
	if x != nil {
		return x.StreamPosition
	}
	return 0
}

type ConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.ReverseCallResponseContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Failure     *protobuf.Failure                    `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"` // If not set/empty - no failure
}

func (x *ConsumerResponse) Reset() {
	*x = ConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerResponse) ProtoMessage() {}

func (x *ConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerResponse.ProtoReflect.Descriptor instead.
func (*ConsumerResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Consumer_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumerResponse) GetCallContext() *services.ReverseCallResponseContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *ConsumerResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

type EventHorizonConsumerToProducerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*EventHorizonConsumerToProducerMessage_SubscriptionRequest
	//	*EventHorizonConsumerToProducerMessage_Response
	//	*EventHorizonConsumerToProducerMessage_Pong
	Message isEventHorizonConsumerToProducerMessage_Message `protobuf_oneof:"Message"`
}

func (x *EventHorizonConsumerToProducerMessage) Reset() {
	*x = EventHorizonConsumerToProducerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHorizonConsumerToProducerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHorizonConsumerToProducerMessage) ProtoMessage() {}

func (x *EventHorizonConsumerToProducerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHorizonConsumerToProducerMessage.ProtoReflect.Descriptor instead.
func (*EventHorizonConsumerToProducerMessage) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Consumer_proto_rawDescGZIP(), []int{3}
}

func (m *EventHorizonConsumerToProducerMessage) GetMessage() isEventHorizonConsumerToProducerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *EventHorizonConsumerToProducerMessage) GetSubscriptionRequest() *ConsumerSubscriptionRequest {
	if x, ok := x.GetMessage().(*EventHorizonConsumerToProducerMessage_SubscriptionRequest); ok {
		return x.SubscriptionRequest
	}
	return nil
}

func (x *EventHorizonConsumerToProducerMessage) GetResponse() *ConsumerResponse {
	if x, ok := x.GetMessage().(*EventHorizonConsumerToProducerMessage_Response); ok {
		return x.Response
	}
	return nil
}

func (x *EventHorizonConsumerToProducerMessage) GetPong() *services.Pong {
	if x, ok := x.GetMessage().(*EventHorizonConsumerToProducerMessage_Pong); ok {
		return x.Pong
	}
	return nil
}

type isEventHorizonConsumerToProducerMessage_Message interface {
	isEventHorizonConsumerToProducerMessage_Message()
}

type EventHorizonConsumerToProducerMessage_SubscriptionRequest struct {
	SubscriptionRequest *ConsumerSubscriptionRequest `protobuf:"bytes,1,opt,name=subscriptionRequest,proto3,oneof"`
}

type EventHorizonConsumerToProducerMessage_Response struct {
	Response *ConsumerResponse `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

type EventHorizonConsumerToProducerMessage_Pong struct {
	Pong *services.Pong `protobuf:"bytes,3,opt,name=pong,proto3,oneof"`
}

func (*EventHorizonConsumerToProducerMessage_SubscriptionRequest) isEventHorizonConsumerToProducerMessage_Message() {
}

func (*EventHorizonConsumerToProducerMessage_Response) isEventHorizonConsumerToProducerMessage_Message() {
}

func (*EventHorizonConsumerToProducerMessage_Pong) isEventHorizonConsumerToProducerMessage_Message() {
}

type ConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.ReverseCallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Event       *EventHorizonEvent                  `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *ConsumerRequest) Reset() {
	*x = ConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerRequest) ProtoMessage() {}

func (x *ConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerRequest.ProtoReflect.Descriptor instead.
func (*ConsumerRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Consumer_proto_rawDescGZIP(), []int{4}
}

func (x *ConsumerRequest) GetCallContext() *services.ReverseCallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *ConsumerRequest) GetEvent() *EventHorizonEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type EventHorizonProducerToConsumerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*EventHorizonProducerToConsumerMessage_SubscriptionResponse
	//	*EventHorizonProducerToConsumerMessage_Request
	//	*EventHorizonProducerToConsumerMessage_Ping
	Message isEventHorizonProducerToConsumerMessage_Message `protobuf_oneof:"Message"`
}

func (x *EventHorizonProducerToConsumerMessage) Reset() {
	*x = EventHorizonProducerToConsumerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHorizonProducerToConsumerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHorizonProducerToConsumerMessage) ProtoMessage() {}

func (x *EventHorizonProducerToConsumerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Consumer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHorizonProducerToConsumerMessage.ProtoReflect.Descriptor instead.
func (*EventHorizonProducerToConsumerMessage) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Consumer_proto_rawDescGZIP(), []int{5}
}

func (m *EventHorizonProducerToConsumerMessage) GetMessage() isEventHorizonProducerToConsumerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *EventHorizonProducerToConsumerMessage) GetSubscriptionResponse() *SubscriptionResponse {
	if x, ok := x.GetMessage().(*EventHorizonProducerToConsumerMessage_SubscriptionResponse); ok {
		return x.SubscriptionResponse
	}
	return nil
}

func (x *EventHorizonProducerToConsumerMessage) GetRequest() *ConsumerRequest {
	if x, ok := x.GetMessage().(*EventHorizonProducerToConsumerMessage_Request); ok {
		return x.Request
	}
	return nil
}

func (x *EventHorizonProducerToConsumerMessage) GetPing() *services.Ping {
	if x, ok := x.GetMessage().(*EventHorizonProducerToConsumerMessage_Ping); ok {
		return x.Ping
	}
	return nil
}

type isEventHorizonProducerToConsumerMessage_Message interface {
	isEventHorizonProducerToConsumerMessage_Message()
}

type EventHorizonProducerToConsumerMessage_SubscriptionResponse struct {
	SubscriptionResponse *SubscriptionResponse `protobuf:"bytes,1,opt,name=subscriptionResponse,proto3,oneof"`
}

type EventHorizonProducerToConsumerMessage_Request struct {
	Request *ConsumerRequest `protobuf:"bytes,2,opt,name=request,proto3,oneof"`
}

type EventHorizonProducerToConsumerMessage_Ping struct {
	Ping *services.Ping `protobuf:"bytes,3,opt,name=ping,proto3,oneof"`
}

func (*EventHorizonProducerToConsumerMessage_SubscriptionResponse) isEventHorizonProducerToConsumerMessage_Message() {
}

func (*EventHorizonProducerToConsumerMessage_Request) isEventHorizonProducerToConsumerMessage_Message() {
}

func (*EventHorizonProducerToConsumerMessage_Ping) isEventHorizonProducerToConsumerMessage_Message() {
}

var File_Runtime_EventHorizon_Consumer_proto protoreflect.FileDescriptor

var file_Runtime_EventHorizon_Consumer_proto_rawDesc = []byte{
	0x0a, 0x23, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48,
	0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x6f, 0x6e, 0x1a, 0x23, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x46, 0x75, 0x6e,
	0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e,
	0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x50, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x72, 0x69,
	0x7a, 0x6f, 0x6e, 0x2f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a,
	0x14, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x3d, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0xa3, 0x02, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x50, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75,
	0x69, 0x64, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x63,
	0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x34, 0x0a, 0x07,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x22, 0xa0, 0x02, 0x0a, 0x25, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x72, 0x69,
	0x7a, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x6e, 0x0a, 0x13,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x70,
	0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f,
	0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x46, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74,
	0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f,
	0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x98, 0x02, 0x0a, 0x25, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x72, 0x69, 0x7a,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x69, 0x0a, 0x14, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x14, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa8, 0x01, 0x0a,
	0x08, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x9b, 0x01, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x44, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x72,
	0x69, 0x7a, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x44, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x5c, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x76, 0x36, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0xaa, 0x02, 0x27, 0x44, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_EventHorizon_Consumer_proto_rawDescOnce sync.Once
	file_Runtime_EventHorizon_Consumer_proto_rawDescData = file_Runtime_EventHorizon_Consumer_proto_rawDesc
)

func file_Runtime_EventHorizon_Consumer_proto_rawDescGZIP() []byte {
	file_Runtime_EventHorizon_Consumer_proto_rawDescOnce.Do(func() {
		file_Runtime_EventHorizon_Consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_EventHorizon_Consumer_proto_rawDescData)
	})
	return file_Runtime_EventHorizon_Consumer_proto_rawDescData
}

var file_Runtime_EventHorizon_Consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Runtime_EventHorizon_Consumer_proto_goTypes = []interface{}{
	(*EventHorizonEvent)(nil),                     // 0: dolittle.runtime.eventhorizon.EventHorizonEvent
	(*ConsumerSubscriptionRequest)(nil),           // 1: dolittle.runtime.eventhorizon.ConsumerSubscriptionRequest
	(*ConsumerResponse)(nil),                      // 2: dolittle.runtime.eventhorizon.ConsumerResponse
	(*EventHorizonConsumerToProducerMessage)(nil), // 3: dolittle.runtime.eventhorizon.EventHorizonConsumerToProducerMessage
	(*ConsumerRequest)(nil),                       // 4: dolittle.runtime.eventhorizon.ConsumerRequest
	(*EventHorizonProducerToConsumerMessage)(nil), // 5: dolittle.runtime.eventhorizon.EventHorizonProducerToConsumerMessage
	(*events.CommittedEvent)(nil),                 // 6: dolittle.runtime.events.CommittedEvent
	(*services.ReverseCallArgumentsContext)(nil),  // 7: dolittle.services.ReverseCallArgumentsContext
	(*protobuf.Uuid)(nil),                         // 8: dolittle.protobuf.Uuid
	(*services.ReverseCallResponseContext)(nil),   // 9: dolittle.services.ReverseCallResponseContext
	(*protobuf.Failure)(nil),                      // 10: dolittle.protobuf.Failure
	(*services.Pong)(nil),                         // 11: dolittle.services.Pong
	(*services.ReverseCallRequestContext)(nil),    // 12: dolittle.services.ReverseCallRequestContext
	(*SubscriptionResponse)(nil),                  // 13: dolittle.runtime.eventhorizon.SubscriptionResponse
	(*services.Ping)(nil),                         // 14: dolittle.services.Ping
}
var file_Runtime_EventHorizon_Consumer_proto_depIdxs = []int32{
	6,  // 0: dolittle.runtime.eventhorizon.EventHorizonEvent.event:type_name -> dolittle.runtime.events.CommittedEvent
	7,  // 1: dolittle.runtime.eventhorizon.ConsumerSubscriptionRequest.callContext:type_name -> dolittle.services.ReverseCallArgumentsContext
	8,  // 2: dolittle.runtime.eventhorizon.ConsumerSubscriptionRequest.tenantId:type_name -> dolittle.protobuf.Uuid
	8,  // 3: dolittle.runtime.eventhorizon.ConsumerSubscriptionRequest.streamId:type_name -> dolittle.protobuf.Uuid
	9,  // 4: dolittle.runtime.eventhorizon.ConsumerResponse.callContext:type_name -> dolittle.services.ReverseCallResponseContext
	10, // 5: dolittle.runtime.eventhorizon.ConsumerResponse.failure:type_name -> dolittle.protobuf.Failure
	1,  // 6: dolittle.runtime.eventhorizon.EventHorizonConsumerToProducerMessage.subscriptionRequest:type_name -> dolittle.runtime.eventhorizon.ConsumerSubscriptionRequest
	2,  // 7: dolittle.runtime.eventhorizon.EventHorizonConsumerToProducerMessage.response:type_name -> dolittle.runtime.eventhorizon.ConsumerResponse
	11, // 8: dolittle.runtime.eventhorizon.EventHorizonConsumerToProducerMessage.pong:type_name -> dolittle.services.Pong
	12, // 9: dolittle.runtime.eventhorizon.ConsumerRequest.callContext:type_name -> dolittle.services.ReverseCallRequestContext
	0,  // 10: dolittle.runtime.eventhorizon.ConsumerRequest.event:type_name -> dolittle.runtime.eventhorizon.EventHorizonEvent
	13, // 11: dolittle.runtime.eventhorizon.EventHorizonProducerToConsumerMessage.subscriptionResponse:type_name -> dolittle.runtime.eventhorizon.SubscriptionResponse
	4,  // 12: dolittle.runtime.eventhorizon.EventHorizonProducerToConsumerMessage.request:type_name -> dolittle.runtime.eventhorizon.ConsumerRequest
	14, // 13: dolittle.runtime.eventhorizon.EventHorizonProducerToConsumerMessage.ping:type_name -> dolittle.services.Ping
	3,  // 14: dolittle.runtime.eventhorizon.Consumer.Subscribe:input_type -> dolittle.runtime.eventhorizon.EventHorizonConsumerToProducerMessage
	5,  // 15: dolittle.runtime.eventhorizon.Consumer.Subscribe:output_type -> dolittle.runtime.eventhorizon.EventHorizonProducerToConsumerMessage
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_Runtime_EventHorizon_Consumer_proto_init() }
func file_Runtime_EventHorizon_Consumer_proto_init() {
	if File_Runtime_EventHorizon_Consumer_proto != nil {
		return
	}
	file_Runtime_EventHorizon_Subscriptions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Runtime_EventHorizon_Consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHorizonEvent); i {
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
		file_Runtime_EventHorizon_Consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerSubscriptionRequest); i {
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
		file_Runtime_EventHorizon_Consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerResponse); i {
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
		file_Runtime_EventHorizon_Consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHorizonConsumerToProducerMessage); i {
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
		file_Runtime_EventHorizon_Consumer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerRequest); i {
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
		file_Runtime_EventHorizon_Consumer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHorizonProducerToConsumerMessage); i {
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
	file_Runtime_EventHorizon_Consumer_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*EventHorizonConsumerToProducerMessage_SubscriptionRequest)(nil),
		(*EventHorizonConsumerToProducerMessage_Response)(nil),
		(*EventHorizonConsumerToProducerMessage_Pong)(nil),
	}
	file_Runtime_EventHorizon_Consumer_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*EventHorizonProducerToConsumerMessage_SubscriptionResponse)(nil),
		(*EventHorizonProducerToConsumerMessage_Request)(nil),
		(*EventHorizonProducerToConsumerMessage_Ping)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Runtime_EventHorizon_Consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_EventHorizon_Consumer_proto_goTypes,
		DependencyIndexes: file_Runtime_EventHorizon_Consumer_proto_depIdxs,
		MessageInfos:      file_Runtime_EventHorizon_Consumer_proto_msgTypes,
	}.Build()
	File_Runtime_EventHorizon_Consumer_proto = out.File
	file_Runtime_EventHorizon_Consumer_proto_rawDesc = nil
	file_Runtime_EventHorizon_Consumer_proto_goTypes = nil
	file_Runtime_EventHorizon_Consumer_proto_depIdxs = nil
}
