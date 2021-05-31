// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package portpb

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

// PortServiceClient is the client API for PortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortServiceClient interface {
	// Client Streaming
	PortsUpdate(ctx context.Context, opts ...grpc.CallOption) (PortService_PortsUpdateClient, error)
}

type portServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortServiceClient(cc grpc.ClientConnInterface) PortServiceClient {
	return &portServiceClient{cc}
}

func (c *portServiceClient) PortsUpdate(ctx context.Context, opts ...grpc.CallOption) (PortService_PortsUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &PortService_ServiceDesc.Streams[0], "/port.PortService/PortsUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &portServicePortsUpdateClient{stream}
	return x, nil
}

type PortService_PortsUpdateClient interface {
	Send(*PortRequest) error
	CloseAndRecv() (*PortResponse, error)
	grpc.ClientStream
}

type portServicePortsUpdateClient struct {
	grpc.ClientStream
}

func (x *portServicePortsUpdateClient) Send(m *PortRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *portServicePortsUpdateClient) CloseAndRecv() (*PortResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PortResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PortServiceServer is the server API for PortService service.
// All implementations must embed UnimplementedPortServiceServer
// for forward compatibility
type PortServiceServer interface {
	// Client Streaming
	PortsUpdate(PortService_PortsUpdateServer) error
	mustEmbedUnimplementedPortServiceServer()
}

// UnimplementedPortServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortServiceServer struct {
}

func (UnimplementedPortServiceServer) PortsUpdate(PortService_PortsUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method PortsUpdate not implemented")
}
func (UnimplementedPortServiceServer) mustEmbedUnimplementedPortServiceServer() {}

// UnsafePortServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortServiceServer will
// result in compilation errors.
type UnsafePortServiceServer interface {
	mustEmbedUnimplementedPortServiceServer()
}

func RegisterPortServiceServer(s grpc.ServiceRegistrar, srv PortServiceServer) {
	s.RegisterService(&PortService_ServiceDesc, srv)
}

func _PortService_PortsUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PortServiceServer).PortsUpdate(&portServicePortsUpdateServer{stream})
}

type PortService_PortsUpdateServer interface {
	SendAndClose(*PortResponse) error
	Recv() (*PortRequest, error)
	grpc.ServerStream
}

type portServicePortsUpdateServer struct {
	grpc.ServerStream
}

func (x *portServicePortsUpdateServer) SendAndClose(m *PortResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *portServicePortsUpdateServer) Recv() (*PortRequest, error) {
	m := new(PortRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PortService_ServiceDesc is the grpc.ServiceDesc for PortService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "port.PortService",
	HandlerType: (*PortServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PortsUpdate",
			Handler:       _PortService_PortsUpdate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "portpb/ports.proto",
}
