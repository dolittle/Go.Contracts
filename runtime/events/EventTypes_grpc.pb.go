// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package events

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

// EventTypesClient is the client API for EventTypes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventTypesClient interface {
	Register(ctx context.Context, in *EventTypeRegistrationRequest, opts ...grpc.CallOption) (*EventTypeRegistrationResponse, error)
}

type eventTypesClient struct {
	cc grpc.ClientConnInterface
}

func NewEventTypesClient(cc grpc.ClientConnInterface) EventTypesClient {
	return &eventTypesClient{cc}
}

func (c *eventTypesClient) Register(ctx context.Context, in *EventTypeRegistrationRequest, opts ...grpc.CallOption) (*EventTypeRegistrationResponse, error) {
	out := new(EventTypeRegistrationResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.events.EventTypes/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventTypesServer is the server API for EventTypes service.
// All implementations must embed UnimplementedEventTypesServer
// for forward compatibility
type EventTypesServer interface {
	Register(context.Context, *EventTypeRegistrationRequest) (*EventTypeRegistrationResponse, error)
	mustEmbedUnimplementedEventTypesServer()
}

// UnimplementedEventTypesServer must be embedded to have forward compatible implementations.
type UnimplementedEventTypesServer struct {
}

func (UnimplementedEventTypesServer) Register(context.Context, *EventTypeRegistrationRequest) (*EventTypeRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedEventTypesServer) mustEmbedUnimplementedEventTypesServer() {}

// UnsafeEventTypesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventTypesServer will
// result in compilation errors.
type UnsafeEventTypesServer interface {
	mustEmbedUnimplementedEventTypesServer()
}

func RegisterEventTypesServer(s grpc.ServiceRegistrar, srv EventTypesServer) {
	s.RegisterService(&EventTypes_ServiceDesc, srv)
}

func _EventTypes_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventTypeRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTypesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.events.EventTypes/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTypesServer).Register(ctx, req.(*EventTypeRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventTypes_ServiceDesc is the grpc.ServiceDesc for EventTypes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventTypes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.events.EventTypes",
	HandlerType: (*EventTypesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _EventTypes_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Runtime/Events/EventTypes.proto",
}
