// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// EmbeddingsClient is the client API for Embeddings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmbeddingsClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Embeddings_ConnectClient, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type embeddingsClient struct {
	cc grpc.ClientConnInterface
}

func NewEmbeddingsClient(cc grpc.ClientConnInterface) EmbeddingsClient {
	return &embeddingsClient{cc}
}

func (c *embeddingsClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Embeddings_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Embeddings_ServiceDesc.Streams[0], "/dolittle.runtime.embeddings.Embeddings/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &embeddingsConnectClient{stream}
	return x, nil
}

type Embeddings_ConnectClient interface {
	Send(*EmbeddingClientToRuntimeMessage) error
	Recv() (*EmbeddingRuntimeToClientMessage, error)
	grpc.ClientStream
}

type embeddingsConnectClient struct {
	grpc.ClientStream
}

func (x *embeddingsConnectClient) Send(m *EmbeddingClientToRuntimeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *embeddingsConnectClient) Recv() (*EmbeddingRuntimeToClientMessage, error) {
	m := new(EmbeddingRuntimeToClientMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *embeddingsClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.embeddings.Embeddings/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *embeddingsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/dolittle.runtime.embeddings.Embeddings/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmbeddingsServer is the server API for Embeddings service.
// All implementations must embed UnimplementedEmbeddingsServer
// for forward compatibility
type EmbeddingsServer interface {
	Connect(Embeddings_ConnectServer) error
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedEmbeddingsServer()
}

// UnimplementedEmbeddingsServer must be embedded to have forward compatible implementations.
type UnimplementedEmbeddingsServer struct {
}

func (UnimplementedEmbeddingsServer) Connect(Embeddings_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedEmbeddingsServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEmbeddingsServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEmbeddingsServer) mustEmbedUnimplementedEmbeddingsServer() {}

// UnsafeEmbeddingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmbeddingsServer will
// result in compilation errors.
type UnsafeEmbeddingsServer interface {
	mustEmbedUnimplementedEmbeddingsServer()
}

func RegisterEmbeddingsServer(s grpc.ServiceRegistrar, srv EmbeddingsServer) {
	s.RegisterService(&Embeddings_ServiceDesc, srv)
}

func _Embeddings_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmbeddingsServer).Connect(&embeddingsConnectServer{stream})
}

type Embeddings_ConnectServer interface {
	Send(*EmbeddingRuntimeToClientMessage) error
	Recv() (*EmbeddingClientToRuntimeMessage, error)
	grpc.ServerStream
}

type embeddingsConnectServer struct {
	grpc.ServerStream
}

func (x *embeddingsConnectServer) Send(m *EmbeddingRuntimeToClientMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *embeddingsConnectServer) Recv() (*EmbeddingClientToRuntimeMessage, error) {
	m := new(EmbeddingClientToRuntimeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Embeddings_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbeddingsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.embeddings.Embeddings/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbeddingsServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Embeddings_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbeddingsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolittle.runtime.embeddings.Embeddings/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbeddingsServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Embeddings_ServiceDesc is the grpc.ServiceDesc for Embeddings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Embeddings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.embeddings.Embeddings",
	HandlerType: (*EmbeddingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Embeddings_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Embeddings_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Embeddings_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Runtime/Embeddings/Embeddings.proto",
}
