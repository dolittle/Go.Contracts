// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: Runtime/EventHorizon/Subscriptions.proto

package eventhorizon

import (
	protobuf "go.dolittle.io/contracts/v7/protobuf"
	services "go.dolittle.io/contracts/v7/services"
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

type SubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure   *protobuf.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
	ConsentId *protobuf.Uuid    `protobuf:"bytes,2,opt,name=consentId,proto3" json:"consentId,omitempty"`
}

func (x *SubscriptionResponse) Reset() {
	*x = SubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Subscriptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionResponse) ProtoMessage() {}

func (x *SubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Subscriptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Subscriptions_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionResponse) GetFailure() *protobuf.Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *SubscriptionResponse) GetConsentId() *protobuf.Uuid {
	if x != nil {
		return x.ConsentId
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallContext    *services.CallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	MicroserviceId *protobuf.Uuid               `protobuf:"bytes,2,opt,name=microserviceId,proto3" json:"microserviceId,omitempty"`
	TenantId       *protobuf.Uuid               `protobuf:"bytes,3,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	StreamId       *protobuf.Uuid               `protobuf:"bytes,4,opt,name=streamId,proto3" json:"streamId,omitempty"`
	PartitionId    string                       `protobuf:"bytes,5,opt,name=partitionId,proto3" json:"partitionId,omitempty"`
	ScopeId        *protobuf.Uuid               `protobuf:"bytes,6,opt,name=scopeId,proto3" json:"scopeId,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Runtime_EventHorizon_Subscriptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_Runtime_EventHorizon_Subscriptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_Runtime_EventHorizon_Subscriptions_proto_rawDescGZIP(), []int{1}
}

func (x *Subscription) GetCallContext() *services.CallRequestContext {
	if x != nil {
		return x.CallContext
	}
	return nil
}

func (x *Subscription) GetMicroserviceId() *protobuf.Uuid {
	if x != nil {
		return x.MicroserviceId
	}
	return nil
}

func (x *Subscription) GetTenantId() *protobuf.Uuid {
	if x != nil {
		return x.TenantId
	}
	return nil
}

func (x *Subscription) GetStreamId() *protobuf.Uuid {
	if x != nil {
		return x.StreamId
	}
	return nil
}

func (x *Subscription) GetPartitionId() string {
	if x != nil {
		return x.PartitionId
	}
	return ""
}

func (x *Subscription) GetScopeId() *protobuf.Uuid {
	if x != nil {
		return x.ScopeId
	}
	return nil
}

var File_Runtime_EventHorizon_Subscriptions_proto protoreflect.FileDescriptor

var file_Runtime_EventHorizon_Subscriptions_proto_rawDesc = []byte{
	0x0a, 0x28, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48,
	0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x64, 0x6f, 0x6c, 0x69,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x1a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x55, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64,
	0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x12, 0x35, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xd7, 0x02, 0x0a, 0x0c, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c,
	0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c,
	0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x49, 0x64, 0x32, 0x7e, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x6d, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x2b, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x33, 0x2e,
	0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x5c, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x64, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76,
	0x37, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0xaa, 0x02, 0x27, 0x44, 0x6f, 0x6c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48,
	0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Runtime_EventHorizon_Subscriptions_proto_rawDescOnce sync.Once
	file_Runtime_EventHorizon_Subscriptions_proto_rawDescData = file_Runtime_EventHorizon_Subscriptions_proto_rawDesc
)

func file_Runtime_EventHorizon_Subscriptions_proto_rawDescGZIP() []byte {
	file_Runtime_EventHorizon_Subscriptions_proto_rawDescOnce.Do(func() {
		file_Runtime_EventHorizon_Subscriptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_Runtime_EventHorizon_Subscriptions_proto_rawDescData)
	})
	return file_Runtime_EventHorizon_Subscriptions_proto_rawDescData
}

var file_Runtime_EventHorizon_Subscriptions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Runtime_EventHorizon_Subscriptions_proto_goTypes = []interface{}{
	(*SubscriptionResponse)(nil),        // 0: dolittle.runtime.eventhorizon.SubscriptionResponse
	(*Subscription)(nil),                // 1: dolittle.runtime.eventhorizon.Subscription
	(*protobuf.Failure)(nil),            // 2: dolittle.protobuf.Failure
	(*protobuf.Uuid)(nil),               // 3: dolittle.protobuf.Uuid
	(*services.CallRequestContext)(nil), // 4: dolittle.services.CallRequestContext
}
var file_Runtime_EventHorizon_Subscriptions_proto_depIdxs = []int32{
	2, // 0: dolittle.runtime.eventhorizon.SubscriptionResponse.failure:type_name -> dolittle.protobuf.Failure
	3, // 1: dolittle.runtime.eventhorizon.SubscriptionResponse.consentId:type_name -> dolittle.protobuf.Uuid
	4, // 2: dolittle.runtime.eventhorizon.Subscription.callContext:type_name -> dolittle.services.CallRequestContext
	3, // 3: dolittle.runtime.eventhorizon.Subscription.microserviceId:type_name -> dolittle.protobuf.Uuid
	3, // 4: dolittle.runtime.eventhorizon.Subscription.tenantId:type_name -> dolittle.protobuf.Uuid
	3, // 5: dolittle.runtime.eventhorizon.Subscription.streamId:type_name -> dolittle.protobuf.Uuid
	3, // 6: dolittle.runtime.eventhorizon.Subscription.scopeId:type_name -> dolittle.protobuf.Uuid
	1, // 7: dolittle.runtime.eventhorizon.Subscriptions.Subscribe:input_type -> dolittle.runtime.eventhorizon.Subscription
	0, // 8: dolittle.runtime.eventhorizon.Subscriptions.Subscribe:output_type -> dolittle.runtime.eventhorizon.SubscriptionResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Runtime_EventHorizon_Subscriptions_proto_init() }
func file_Runtime_EventHorizon_Subscriptions_proto_init() {
	if File_Runtime_EventHorizon_Subscriptions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Runtime_EventHorizon_Subscriptions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionResponse); i {
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
		file_Runtime_EventHorizon_Subscriptions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
			RawDescriptor: file_Runtime_EventHorizon_Subscriptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Runtime_EventHorizon_Subscriptions_proto_goTypes,
		DependencyIndexes: file_Runtime_EventHorizon_Subscriptions_proto_depIdxs,
		MessageInfos:      file_Runtime_EventHorizon_Subscriptions_proto_msgTypes,
	}.Build()
	File_Runtime_EventHorizon_Subscriptions_proto = out.File
	file_Runtime_EventHorizon_Subscriptions_proto_rawDesc = nil
	file_Runtime_EventHorizon_Subscriptions_proto_goTypes = nil
	file_Runtime_EventHorizon_Subscriptions_proto_depIdxs = nil
}
