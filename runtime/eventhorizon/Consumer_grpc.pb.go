// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: Runtime/EventHorizon/Consumer.proto

package eventhorizon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConsumerClient is the client API for Consumer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Consumer_SubscribeClient, error)
}

type consumerClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerClient(cc grpc.ClientConnInterface) ConsumerClient {
	return &consumerClient{cc}
}

func (c *consumerClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Consumer_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Consumer_ServiceDesc.Streams[0], "/dolittle.runtime.eventhorizon.Consumer/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerSubscribeClient{stream}
	return x, nil
}

type Consumer_SubscribeClient interface {
	Send(*EventHorizonConsumerToProducerMessage) error
	Recv() (*EventHorizonProducerToConsumerMessage, error)
	grpc.ClientStream
}

type consumerSubscribeClient struct {
	grpc.ClientStream
}

func (x *consumerSubscribeClient) Send(m *EventHorizonConsumerToProducerMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *consumerSubscribeClient) Recv() (*EventHorizonProducerToConsumerMessage, error) {
	m := new(EventHorizonProducerToConsumerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConsumerServer is the server API for Consumer service.
// All implementations must embed UnimplementedConsumerServer
// for forward compatibility
type ConsumerServer interface {
	Subscribe(Consumer_SubscribeServer) error
	mustEmbedUnimplementedConsumerServer()
}

// UnimplementedConsumerServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServer struct {
}

func (UnimplementedConsumerServer) Subscribe(Consumer_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedConsumerServer) mustEmbedUnimplementedConsumerServer() {}

// UnsafeConsumerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServer will
// result in compilation errors.
type UnsafeConsumerServer interface {
	mustEmbedUnimplementedConsumerServer()
}

func RegisterConsumerServer(s grpc.ServiceRegistrar, srv ConsumerServer) {
	s.RegisterService(&Consumer_ServiceDesc, srv)
}

func _Consumer_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConsumerServer).Subscribe(&consumerSubscribeServer{stream})
}

type Consumer_SubscribeServer interface {
	Send(*EventHorizonProducerToConsumerMessage) error
	Recv() (*EventHorizonConsumerToProducerMessage, error)
	grpc.ServerStream
}

type consumerSubscribeServer struct {
	grpc.ServerStream
}

func (x *consumerSubscribeServer) Send(m *EventHorizonProducerToConsumerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *consumerSubscribeServer) Recv() (*EventHorizonConsumerToProducerMessage, error) {
	m := new(EventHorizonConsumerToProducerMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Consumer_ServiceDesc is the grpc.ServiceDesc for Consumer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Consumer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.eventhorizon.Consumer",
	HandlerType: (*ConsumerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Consumer_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Runtime/EventHorizon/Consumer.proto",
}
