// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ProjectionsClient is the client API for Projections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectionsClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Projections_ConnectClient, error)
}

type projectionsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectionsClient(cc grpc.ClientConnInterface) ProjectionsClient {
	return &projectionsClient{cc}
}

func (c *projectionsClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Projections_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Projections_ServiceDesc.Streams[0], "/dolittle.runtime.events.processing.Projections/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectionsConnectClient{stream}
	return x, nil
}

type Projections_ConnectClient interface {
	Send(*ProjectionClientToRuntimeMessage) error
	Recv() (*ProjectionRuntimeToClientMessage, error)
	grpc.ClientStream
}

type projectionsConnectClient struct {
	grpc.ClientStream
}

func (x *projectionsConnectClient) Send(m *ProjectionClientToRuntimeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *projectionsConnectClient) Recv() (*ProjectionRuntimeToClientMessage, error) {
	m := new(ProjectionRuntimeToClientMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProjectionsServer is the server API for Projections service.
// All implementations must embed UnimplementedProjectionsServer
// for forward compatibility
type ProjectionsServer interface {
	Connect(Projections_ConnectServer) error
	mustEmbedUnimplementedProjectionsServer()
}

// UnimplementedProjectionsServer must be embedded to have forward compatible implementations.
type UnimplementedProjectionsServer struct {
}

func (UnimplementedProjectionsServer) Connect(Projections_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
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

func _Projections_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProjectionsServer).Connect(&projectionsConnectServer{stream})
}

type Projections_ConnectServer interface {
	Send(*ProjectionRuntimeToClientMessage) error
	Recv() (*ProjectionClientToRuntimeMessage, error)
	grpc.ServerStream
}

type projectionsConnectServer struct {
	grpc.ServerStream
}

func (x *projectionsConnectServer) Send(m *ProjectionRuntimeToClientMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *projectionsConnectServer) Recv() (*ProjectionClientToRuntimeMessage, error) {
	m := new(ProjectionClientToRuntimeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Projections_ServiceDesc is the grpc.ServiceDesc for Projections service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Projections_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.events.processing.Projections",
	HandlerType: (*ProjectionsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Projections_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Runtime/Events.Processing/Projections.proto",
}
