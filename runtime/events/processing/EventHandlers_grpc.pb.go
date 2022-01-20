// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: Runtime/Events.Processing/EventHandlers.proto

package processing

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

// EventHandlersClient is the client API for EventHandlers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventHandlersClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (EventHandlers_ConnectClient, error)
}

type eventHandlersClient struct {
	cc grpc.ClientConnInterface
}

func NewEventHandlersClient(cc grpc.ClientConnInterface) EventHandlersClient {
	return &eventHandlersClient{cc}
}

func (c *eventHandlersClient) Connect(ctx context.Context, opts ...grpc.CallOption) (EventHandlers_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventHandlers_ServiceDesc.Streams[0], "/dolittle.runtime.events.processing.EventHandlers/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventHandlersConnectClient{stream}
	return x, nil
}

type EventHandlers_ConnectClient interface {
	Send(*EventHandlerClientToRuntimeMessage) error
	Recv() (*EventHandlerRuntimeToClientMessage, error)
	grpc.ClientStream
}

type eventHandlersConnectClient struct {
	grpc.ClientStream
}

func (x *eventHandlersConnectClient) Send(m *EventHandlerClientToRuntimeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventHandlersConnectClient) Recv() (*EventHandlerRuntimeToClientMessage, error) {
	m := new(EventHandlerRuntimeToClientMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventHandlersServer is the server API for EventHandlers service.
// All implementations must embed UnimplementedEventHandlersServer
// for forward compatibility
type EventHandlersServer interface {
	Connect(EventHandlers_ConnectServer) error
	mustEmbedUnimplementedEventHandlersServer()
}

// UnimplementedEventHandlersServer must be embedded to have forward compatible implementations.
type UnimplementedEventHandlersServer struct {
}

func (UnimplementedEventHandlersServer) Connect(EventHandlers_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedEventHandlersServer) mustEmbedUnimplementedEventHandlersServer() {}

// UnsafeEventHandlersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventHandlersServer will
// result in compilation errors.
type UnsafeEventHandlersServer interface {
	mustEmbedUnimplementedEventHandlersServer()
}

func RegisterEventHandlersServer(s grpc.ServiceRegistrar, srv EventHandlersServer) {
	s.RegisterService(&EventHandlers_ServiceDesc, srv)
}

func _EventHandlers_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventHandlersServer).Connect(&eventHandlersConnectServer{stream})
}

type EventHandlers_ConnectServer interface {
	Send(*EventHandlerRuntimeToClientMessage) error
	Recv() (*EventHandlerClientToRuntimeMessage, error)
	grpc.ServerStream
}

type eventHandlersConnectServer struct {
	grpc.ServerStream
}

func (x *eventHandlersConnectServer) Send(m *EventHandlerRuntimeToClientMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventHandlersConnectServer) Recv() (*EventHandlerClientToRuntimeMessage, error) {
	m := new(EventHandlerClientToRuntimeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventHandlers_ServiceDesc is the grpc.ServiceDesc for EventHandlers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventHandlers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.events.processing.EventHandlers",
	HandlerType: (*EventHandlersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _EventHandlers_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Runtime/Events.Processing/EventHandlers.proto",
}
