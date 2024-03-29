// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: Runtime/Management/Aggregates/AggregateRoots.proto

package management

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

// AggregateRootsClient is the client API for AggregateRoots service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregateRootsClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
}

type aggregateRootsClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregateRootsClient(cc grpc.ClientConnInterface) AggregateRootsClient {
	return &aggregateRootsClient{cc}
}

func (c *aggregateRootsClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.aggregates.management.AggregateRoots/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregateRootsClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.aggregates.management.AggregateRoots/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregateRootsClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.aggregates.management.AggregateRoots/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregateRootsServer is the server API for AggregateRoots service.
// All implementations must embed UnimplementedAggregateRootsServer
// for forward compatibility
type AggregateRootsServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	mustEmbedUnimplementedAggregateRootsServer()
}

// UnimplementedAggregateRootsServer must be embedded to have forward compatible implementations.
type UnimplementedAggregateRootsServer struct {
}

func (UnimplementedAggregateRootsServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAggregateRootsServer) GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedAggregateRootsServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedAggregateRootsServer) mustEmbedUnimplementedAggregateRootsServer() {}

// UnsafeAggregateRootsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregateRootsServer will
// result in compilation errors.
type UnsafeAggregateRootsServer interface {
	mustEmbedUnimplementedAggregateRootsServer()
}

func RegisterAggregateRootsServer(s grpc.ServiceRegistrar, srv AggregateRootsServer) {
	s.RegisterService(&AggregateRoots_ServiceDesc, srv)
}

func _AggregateRoots_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregateRootsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.aggregates.management.AggregateRoots/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregateRootsServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregateRoots_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregateRootsServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.aggregates.management.AggregateRoots/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregateRootsServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AggregateRoots_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregateRootsServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.aggregates.management.AggregateRoots/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregateRootsServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregateRoots_ServiceDesc is the grpc.ServiceDesc for AggregateRoots service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregateRoots_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.aggregates.management.AggregateRoots",
	HandlerType: (*AggregateRootsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _AggregateRoots_GetAll_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _AggregateRoots_GetOne_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _AggregateRoots_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Runtime/Management/Aggregates/AggregateRoots.proto",
}
