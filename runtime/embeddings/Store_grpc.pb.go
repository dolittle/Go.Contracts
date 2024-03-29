// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: Runtime/Embeddings/Store.proto

package embeddings

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

// EmbeddingStoreClient is the client API for EmbeddingStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmbeddingStoreClient interface {
	GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (*GetKeysResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type embeddingStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewEmbeddingStoreClient(cc grpc.ClientConnInterface) EmbeddingStoreClient {
	return &embeddingStoreClient{cc}
}

func (c *embeddingStoreClient) GetOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.embeddings.EmbeddingStore/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embeddingStoreClient) GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (*GetKeysResponse, error) {
	out := new(GetKeysResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.embeddings.EmbeddingStore/GetKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embeddingStoreClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.embeddings.EmbeddingStore/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmbeddingStoreServer is the server API for EmbeddingStore service.
// All implementations must embed UnimplementedEmbeddingStoreServer
// for forward compatibility
type EmbeddingStoreServer interface {
	GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	GetKeys(context.Context, *GetKeysRequest) (*GetKeysResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedEmbeddingStoreServer()
}

// UnimplementedEmbeddingStoreServer must be embedded to have forward compatible implementations.
type UnimplementedEmbeddingStoreServer struct {
}

func (UnimplementedEmbeddingStoreServer) GetOne(context.Context, *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedEmbeddingStoreServer) GetKeys(context.Context, *GetKeysRequest) (*GetKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeys not implemented")
}
func (UnimplementedEmbeddingStoreServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedEmbeddingStoreServer) mustEmbedUnimplementedEmbeddingStoreServer() {}

// UnsafeEmbeddingStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmbeddingStoreServer will
// result in compilation errors.
type UnsafeEmbeddingStoreServer interface {
	mustEmbedUnimplementedEmbeddingStoreServer()
}

func RegisterEmbeddingStoreServer(s grpc.ServiceRegistrar, srv EmbeddingStoreServer) {
	s.RegisterService(&EmbeddingStore_ServiceDesc, srv)
}

func _EmbeddingStore_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbeddingStoreServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.embeddings.EmbeddingStore/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbeddingStoreServer).GetOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmbeddingStore_GetKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbeddingStoreServer).GetKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.embeddings.EmbeddingStore/GetKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbeddingStoreServer).GetKeys(ctx, req.(*GetKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmbeddingStore_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbeddingStoreServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.embeddings.EmbeddingStore/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbeddingStoreServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmbeddingStore_ServiceDesc is the grpc.ServiceDesc for EmbeddingStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmbeddingStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.embeddings.EmbeddingStore",
	HandlerType: (*EmbeddingStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOne",
			Handler:    _EmbeddingStore_GetOne_Handler,
		},
		{
			MethodName: "GetKeys",
			Handler:    _EmbeddingStore_GetKeys_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _EmbeddingStore_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Runtime/Embeddings/Store.proto",
}
