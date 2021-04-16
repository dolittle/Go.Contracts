// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package projections

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

// ProjectionsClient is the client API for Projections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectionsClient interface {
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type projectionsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectionsClient(cc grpc.ClientConnInterface) ProjectionsClient {
	return &projectionsClient{cc}
}

func (c *projectionsClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.projections.Projections/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectionsClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.projections.Projections/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectionsServer is the server API for Projections service.
// All implementations must embed UnimplementedProjectionsServer
// for forward compatibility
type ProjectionsServer interface {
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedProjectionsServer()
}

// UnimplementedProjectionsServer must be embedded to have forward compatible implementations.
type UnimplementedProjectionsServer struct {
}

func (UnimplementedProjectionsServer) GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedProjectionsServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedProjectionsServer) mustEmbedUnimplementedProjectionsServer() {}

// UnsafeProjectionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectionsServer will
// result in compilation errors.
type UnsafeProjectionsServer interface {
	mustEmbedUnimplementedProjectionsServer()
}

func RegisterProjectionsServer(s grpc.ServiceRegistrar, srv ProjectionsServer) {
	s.RegisterService(&Projections_ServiceDesc, srv)
}

func _Projections_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectionsServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.projections.Projections/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectionsServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projections_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectionsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.projections.Projections/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectionsServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Projections_ServiceDesc is the grpc.ServiceDesc for Projections service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Projections_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.projections.Projections",
	HandlerType: (*ProjectionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOne",
			Handler:    _Projections_GetOne_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Projections_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Runtime/Projections/Store.proto",
}