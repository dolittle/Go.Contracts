// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Runtime/Events.Processing/EventHandlers.proto

package processing

import (
	artifacts "go.dolittle.io/contracts/v5/artifacts"
	protobuf "go.dolittle.io/contracts/v5/protobuf"
	services "go.dolittle.io/contracts/v5/services"
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

type EventHandlerRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext    *services.ReverseCallArgumentsContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	ScopeId        *protobuf.Uuid                        `protobuf:"bytes,2,opt,name=scopeId,proto3" json:"scopeId,omitempty"`
	EventHandlerId *protobuf.Uuid                        `protobuf:"bytes,3,opt,name=eventHandlerId,proto3" json:"eventHandlerId,omitempty"`
	Types          []*artifacts.Artifact                 `protobuf:"bytes,4,rep,name=types,proto3" json:"types,omitempty"`
	Partitioned    bool                                  `protobuf:"varint,5,opt,name=partitioned,proto3" json:"partitioned,omitempty"`
}

func (x *EventHandlerRegistrationRequest) Reset() {
	*x = EventHandlerRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHandlerRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHandlerRegistrationRequest) ProtoMessage() {}

func (x *EventHandlerRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHandlerRegistrationRequest.ProtoReflect.Descriptor instead.
func (*EventHandlerRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{0}
}

func (x *EventHandlerRegistrationRequest) GetCallContext() *services.ReverseCallArgumentsContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *EventHandlerRegistrationRequest) GetScopeId() *protobuf.Uuid {
	if x != nil {
		return x.ScopeId
	}
	return nil
}

func (x *EventHandlerRegistrationRequest) GetEventHandlerId() *protobuf.Uuid {
	if x != nil {
		return x.EventHandlerId
	}
	return nil
}

func (x *EventHandlerRegistrationRequest) GetTypes() []*artifacts.Artifact {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *EventHandlerRegistrationRequest) GetPartitioned() bool {
	if x != nil {
		return x.Partitioned
	}
	return false
}

type EventHandlerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext *services.ReverseCallResponseContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Failure     *ProcessorFailure                    `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"` // If not set/empty - no failure
}

func (x *EventHandlerResponse) Reset() {
	*x = EventHandlerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHandlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHandlerResponse) ProtoMessage() {}

func (x *EventHandlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHandlerResponse.ProtoReflect.Descriptor instead.
func (*EventHandlerResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{1}
}

func (x *EventHandlerResponse) GetCallContext() *services.ReverseCallResponseContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *EventHandlerResponse) GetFailure() *ProcessorFailure {
	if x != nil {
		return x.Failure
	}
	return nil
}

type EventHandlerClientToRuntimeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*EventHandlerClientToRuntimeMessage_RegistrationRequest
	//	*EventHandlerClientToRuntimeMessage_HandleResult
	//	*EventHandlerClientToRuntimeMessage_Pong
	Message isEventHandlerClientToRuntimeMessage_Message `protobuf_oneof:"Message"`
}

func (x *EventHandlerClientToRuntimeMessage) Reset() {
	*x = EventHandlerClientToRuntimeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHandlerClientToRuntimeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHandlerClientToRuntimeMessage) ProtoMessage() {}

func (x *EventHandlerClientToRuntimeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHandlerClientToRuntimeMessage.ProtoReflect.Descriptor instead.
func (*EventHandlerClientToRuntimeMessage) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{2}
}

func (m *EventHandlerClientToRuntimeMessage) GetMessage() isEventHandlerClientToRuntimeMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *EventHandlerClientToRuntimeMessage) GetRegistrationRequest() *EventHandlerRegistrationRequest {
	if x, ok := x.GetMessage().(*EventHandlerClientToRuntimeMessage_RegistrationRequest); ok {
		return x.RegistrationRequest
	}
	return nil
}

func (x *EventHandlerClientToRuntimeMessage) GetHandleResult() *EventHandlerResponse {
	if x, ok := x.GetMessage().(*EventHandlerClientToRuntimeMessage_HandleResult); ok {
		return x.HandleResult
	}
	return nil
}

func (x *EventHandlerClientToRuntimeMessage) GetPong() *services.Pong {
	if x, ok := x.GetMessage().(*EventHandlerClientToRuntimeMessage_Pong); ok {
		return x.Pong
	}
	return nil
}

type isEventHandlerClientToRuntimeMessage_Message interface {
	isEventHandlerClientToRuntimeMessage_Message()
}

type EventHandlerClientToRuntimeMessage_RegistrationRequest struct {
	RegistrationRequest *EventHandlerRegistrationRequest `protobuf:"bytes,1,opt,name=registrationRequest,proto3,oneof"`
}

type EventHandlerClientToRuntimeMessage_HandleResult struct {
	HandleResult *EventHandlerResponse `protobuf:"bytes,2,opt,name=handleResult,proto3,oneof"`
}

type EventHandlerClientToRuntimeMessage_Pong struct {
	Pong *services.Pong `protobuf:"bytes,3,opt,name=pong,proto3,oneof"`
}

func (*EventHandlerClientToRuntimeMessage_RegistrationRequest) isEventHandlerClientToRuntimeMessage_Message() {
}

func (*EventHandlerClientToRuntimeMessage_HandleResult) isEventHandlerClientToRuntimeMessage_Message() {
}

func (*EventHandlerClientToRuntimeMessage_Pong) isEventHandlerClientToRuntimeMessage_Message() {}

type EventHandlerRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure *protobuf.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *EventHandlerRegistrationResponse) Reset() {
	*x = EventHandlerRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHandlerRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHandlerRegistrationResponse) ProtoMessage() {}

func (x *EventHandlerRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHandlerRegistrationResponse.ProtoReflect.Descriptor instead.
func (*EventHandlerRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{3}
}

func (x *EventHandlerRegistrationResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

type HandleEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext          *services.ReverseCallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Event                *StreamEvent                        `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	RetryProcessingState *RetryProcessingState               `protobuf:"bytes,3,opt,name=retryProcessingState,proto3" json:"retryProcessingState,omitempty"`
}

func (x *HandleEventRequest) Reset() {
	*x = HandleEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleEventRequest) ProtoMessage() {}

func (x *HandleEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleEventRequest.ProtoReflect.Descriptor instead.
func (*HandleEventRequest) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{4}
}

func (x *HandleEventRequest) GetCallContext() *services.ReverseCallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *HandleEventRequest) GetEvent() *StreamEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *HandleEventRequest) GetRetryProcessingState() *RetryProcessingState {
	if x != nil {
		return x.RetryProcessingState
	}
	return nil
}

type EventHandlerRuntimeToClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*EventHandlerRuntimeToClientMessage_RegistrationResponse
	//	*EventHandlerRuntimeToClientMessage_HandleRequest
	//	*EventHandlerRuntimeToClientMessage_Ping
	Message isEventHandlerRuntimeToClientMessage_Message `protobuf_oneof:"Message"`
}

func (x *EventHandlerRuntimeToClientMessage) Reset() {
	*x = EventHandlerRuntimeToClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventHandlerRuntimeToClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventHandlerRuntimeToClientMessage) ProtoMessage() {}

func (x *EventHandlerRuntimeToClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventHandlerRuntimeToClientMessage.ProtoReflect.Descriptor instead.
func (*EventHandlerRuntimeToClientMessage) Descriptor() ([]byte, []int) {
	return file_Runtime_Events_Processing_EventHandlers_proto_rawDescGZIP(), []int{5}
}

func (m *EventHandlerRuntimeToClientMessage) GetMessage() isEventHandlerRuntimeToClientMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *EventHandlerRuntimeToClientMessage) GetRegistrationResponse() *EventHandlerRegistrationResponse {
	if x, ok := x.GetMessage().(*EventHandlerRuntimeToClientMessage_RegistrationResponse); ok {
		return x.RegistrationResponse
	}
	return nil
}

func (x *EventHandlerRuntimeToClientMessage) GetHandleRequest() *HandleEventRequest {
	if x, ok := x.GetMessage().(*EventHandlerRuntimeToClientMessage_HandleRequest); ok {
		return x.HandleRequest
	}
	return nil
}

func (x *EventHandlerRuntimeToClientMessage) GetPing() *services.Ping {
	if x, ok := x.GetMessage().(*EventHandlerRuntimeToClientMessage_Ping); ok {
		return x.Ping
	}
	return nil
}

type isEventHandlerRuntimeToClientMessage_Message interface {
	isEventHandlerRuntimeToClientMessage_Message()
}

type EventHandlerRuntimeToClientMessage_RegistrationResponse struct {
	RegistrationResponse *EventHandlerRegistrationResponse `protobuf:"bytes,1,opt,name=registrationResponse,proto3,oneof"`
}

type EventHandlerRuntimeToClientMessage_HandleRequest struct {
	HandleRequest *HandleEventRequest `protobuf:"bytes,2,opt,name=handleRequest,proto3,oneof"`
}

type EventHandlerRuntimeToClientMessage_Ping struct {
	Ping *services.Ping `protobuf:"bytes,3,opt,name=ping,proto3,oneof"`
}

func (*EventHandlerRuntimeToClientMessage_RegistrationResponse) isEventHandlerRuntimeToClientMessage_Message() {
}

func (*EventHandlerRuntimeToClientMessage_HandleRequest) isEventHandlerRuntimeToClientMessage_Message() {
}

func (*EventHandlerRuntimeToClientMessage_Ping) isEventHandlerRuntimeToClientMessage_Message() {}

var File_Runtime_Events_Processing_EventHandlers_proto protoreflect.FileDescriptor

var file_Runtime_Events_Processing_EventHandlers_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x1a, 0x25, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x73, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e, 0x64,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x46, 0x75,
	0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x46, 0x75, 0x6e, 0x64, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x2f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x50, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x02, 0x0a,
	0x1f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x50, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x07, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x22, 0xb7, 0x01, 0x0a,
	0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x6f, 0x6c,
	0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x4e, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0xb7, 0x02, 0x0a, 0x22, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x77, 0x0a,
	0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x64, 0x6f, 0x6c,
	0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x48, 0x00, 0x52,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x58, 0x0a, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0x99, 0x02, 0x0a, 0x12, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4e, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x45, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x6c, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x14, 0x72, 0x65, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xba, 0x02, 0x0a, 0x22, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x7a, 0x0a,
	0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x64, 0x6f,
	0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x00, 0x52, 0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xaf, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x9d, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x12, 0x46, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x46, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x66, 0x5a, 0x35, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x76, 0x35, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0xaa, 0x02,
	0x2c, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_Events_Processing_EventHandlers_proto_rawDescOnce sync.Once
	file_Runtime_Events_Processing_EventHandlers_proto_rawDescData = file_Runtime_Events_Processing_EventHandlers_proto_rawDesc
)

func file_Runtime_Events_Processing_EventHandlers_proto_rawDescGZIP() []byte {
	file_Runtime_Events_Processing_EventHandlers_proto_rawDescOnce.Do(func() {
		file_Runtime_Events_Processing_EventHandlers_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_Events_Processing_EventHandlers_proto_rawDescData)
	})
	return file_Runtime_Events_Processing_EventHandlers_proto_rawDescData
}

var file_Runtime_Events_Processing_EventHandlers_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Runtime_Events_Processing_EventHandlers_proto_goTypes = []interface{}{
	(*EventHandlerRegistrationRequest)(nil),      // 0: dolittle.runtime.events.processing.EventHandlerRegistrationRequest
	(*EventHandlerResponse)(nil),                 // 1: dolittle.runtime.events.processing.EventHandlerResponse
	(*EventHandlerClientToRuntimeMessage)(nil),   // 2: dolittle.runtime.events.processing.EventHandlerClientToRuntimeMessage
	(*EventHandlerRegistrationResponse)(nil),     // 3: dolittle.runtime.events.processing.EventHandlerRegistrationResponse
	(*HandleEventRequest)(nil),                   // 4: dolittle.runtime.events.processing.HandleEventRequest
	(*EventHandlerRuntimeToClientMessage)(nil),   // 5: dolittle.runtime.events.processing.EventHandlerRuntimeToClientMessage
	(*services.ReverseCallArgumentsContext)(nil), // 6: dolittle.services.ReverseCallArgumentsContext
	(*protobuf.Uuid)(nil),                        // 7: dolittle.protobuf.Uuid
	(*artifacts.Artifact)(nil),                   // 8: dolittle.artifacts.Artifact
	(*services.ReverseCallResponseContext)(nil),  // 9: dolittle.services.ReverseCallResponseContext
	(*ProcessorFailure)(nil),                     // 10: dolittle.runtime.events.processing.ProcessorFailure
	(*services.Pong)(nil),                        // 11: dolittle.services.Pong
	(*protobuf.Failure)(nil),                     // 12: dolittle.protobuf.Failure
	(*services.ReverseCallRequestContext)(nil),   // 13: dolittle.services.ReverseCallRequestContext
	(*StreamEvent)(nil),                          // 14: dolittle.runtime.events.processing.StreamEvent
	(*RetryProcessingState)(nil),                 // 15: dolittle.runtime.events.processing.RetryProcessingState
	(*services.Ping)(nil),                        // 16: dolittle.services.Ping
}
var file_Runtime_Events_Processing_EventHandlers_proto_depIdxs = []int32{
	6,  // 0: dolittle.runtime.events.processing.EventHandlerRegistrationRequest.callContext:type_name -> dolittle.services.ReverseCallArgumentsContext
	7,  // 1: dolittle.runtime.events.processing.EventHandlerRegistrationRequest.scopeId:type_name -> dolittle.protobuf.Uuid
	7,  // 2: dolittle.runtime.events.processing.EventHandlerRegistrationRequest.eventHandlerId:type_name -> dolittle.protobuf.Uuid
	8,  // 3: dolittle.runtime.events.processing.EventHandlerRegistrationRequest.types:type_name -> dolittle.artifacts.Artifact
	9,  // 4: dolittle.runtime.events.processing.EventHandlerResponse.callContext:type_name -> dolittle.services.ReverseCallResponseContext
	10, // 5: dolittle.runtime.events.processing.EventHandlerResponse.failure:type_name -> dolittle.runtime.events.processing.ProcessorFailure
	0,  // 6: dolittle.runtime.events.processing.EventHandlerClientToRuntimeMessage.registrationRequest:type_name -> dolittle.runtime.events.processing.EventHandlerRegistrationRequest
	1,  // 7: dolittle.runtime.events.processing.EventHandlerClientToRuntimeMessage.handleResult:type_name -> dolittle.runtime.events.processing.EventHandlerResponse
	11, // 8: dolittle.runtime.events.processing.EventHandlerClientToRuntimeMessage.pong:type_name -> dolittle.services.Pong
	12, // 9: dolittle.runtime.events.processing.EventHandlerRegistrationResponse.failure:type_name -> dolittle.protobuf.Failure
	13, // 10: dolittle.runtime.events.processing.HandleEventRequest.callContext:type_name -> dolittle.services.ReverseCallRequestContext
	14, // 11: dolittle.runtime.events.processing.HandleEventRequest.event:type_name -> dolittle.runtime.events.processing.StreamEvent
	15, // 12: dolittle.runtime.events.processing.HandleEventRequest.retryProcessingState:type_name -> dolittle.runtime.events.processing.RetryProcessingState
	3,  // 13: dolittle.runtime.events.processing.EventHandlerRuntimeToClientMessage.registrationResponse:type_name -> dolittle.runtime.events.processing.EventHandlerRegistrationResponse
	4,  // 14: dolittle.runtime.events.processing.EventHandlerRuntimeToClientMessage.handleRequest:type_name -> dolittle.runtime.events.processing.HandleEventRequest
	16, // 15: dolittle.runtime.events.processing.EventHandlerRuntimeToClientMessage.ping:type_name -> dolittle.services.Ping
	2,  // 16: dolittle.runtime.events.processing.EventHandlers.Connect:input_type -> dolittle.runtime.events.processing.EventHandlerClientToRuntimeMessage
	5,  // 17: dolittle.runtime.events.processing.EventHandlers.Connect:output_type -> dolittle.runtime.events.processing.EventHandlerRuntimeToClientMessage
	17, // [17:18] is the sub-list for method output_type
	16, // [16:17] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_Runtime_Events_Processing_EventHandlers_proto_init() }
func file_Runtime_Events_Processing_EventHandlers_proto_init() {
	if File_Runtime_Events_Processing_EventHandlers_proto != nil {
		return
	}
	file_Runtime_Events_Processing_StreamEvent_proto_init()
	file_Runtime_Events_Processing_Processors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHandlerRegistrationRequest); i {
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
		file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHandlerResponse); i {
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
		file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHandlerClientToRuntimeMessage); i {
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
		file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHandlerRegistrationResponse); i {
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
		file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleEventRequest); i {
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
		file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventHandlerRuntimeToClientMessage); i {
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
	file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*EventHandlerClientToRuntimeMessage_RegistrationRequest)(nil),
		(*EventHandlerClientToRuntimeMessage_HandleResult)(nil),
		(*EventHandlerClientToRuntimeMessage_Pong)(nil),
	}
	file_Runtime_Events_Processing_EventHandlers_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*EventHandlerRuntimeToClientMessage_RegistrationResponse)(nil),
		(*EventHandlerRuntimeToClientMessage_HandleRequest)(nil),
		(*EventHandlerRuntimeToClientMessage_Ping)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Runtime_Events_Processing_EventHandlers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_Events_Processing_EventHandlers_proto_goTypes,
		DependencyIndexes: file_Runtime_Events_Processing_EventHandlers_proto_depIdxs,
		MessageInfos:      file_Runtime_Events_Processing_EventHandlers_proto_msgTypes,
	}.Build()
	File_Runtime_Events_Processing_EventHandlers_proto = out.File
	file_Runtime_Events_Processing_EventHandlers_proto_rawDesc = nil
	file_Runtime_Events_Processing_EventHandlers_proto_goTypes = nil
	file_Runtime_Events_Processing_EventHandlers_proto_depIdxs = nil
}
