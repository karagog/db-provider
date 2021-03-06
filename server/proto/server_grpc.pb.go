// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// IntegrationTestClient is the client API for IntegrationTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrationTestClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	// GetDatabaseInstance requests a dedicated test instance.
	//
	// Communication is bi-directional and the caller is expected to
	// keep the RPC open as long as they're using the test instance.
	// As soon as the call finishes, the instance will be given to another
	// requestor.
	//
	// See protobuf messages for protocol details.
	GetDatabaseInstance(ctx context.Context, opts ...grpc.CallOption) (IntegrationTest_GetDatabaseInstanceClient, error)
}

type integrationTestClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationTestClient(cc grpc.ClientConnInterface) IntegrationTestClient {
	return &integrationTestClient{cc}
}

func (c *integrationTestClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/server.IntegrationTest/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationTestClient) GetDatabaseInstance(ctx context.Context, opts ...grpc.CallOption) (IntegrationTest_GetDatabaseInstanceClient, error) {
	stream, err := c.cc.NewStream(ctx, &IntegrationTest_ServiceDesc.Streams[0], "/server.IntegrationTest/GetDatabaseInstance", opts...)
	if err != nil {
		return nil, err
	}
	x := &integrationTestGetDatabaseInstanceClient{stream}
	return x, nil
}

type IntegrationTest_GetDatabaseInstanceClient interface {
	Send(*GetDatabaseInstanceRequest) error
	Recv() (*GetDatabaseInstanceResponse, error)
	grpc.ClientStream
}

type integrationTestGetDatabaseInstanceClient struct {
	grpc.ClientStream
}

func (x *integrationTestGetDatabaseInstanceClient) Send(m *GetDatabaseInstanceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *integrationTestGetDatabaseInstanceClient) Recv() (*GetDatabaseInstanceResponse, error) {
	m := new(GetDatabaseInstanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IntegrationTestServer is the server API for IntegrationTest service.
// All implementations must embed UnimplementedIntegrationTestServer
// for forward compatibility
type IntegrationTestServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	// GetDatabaseInstance requests a dedicated test instance.
	//
	// Communication is bi-directional and the caller is expected to
	// keep the RPC open as long as they're using the test instance.
	// As soon as the call finishes, the instance will be given to another
	// requestor.
	//
	// See protobuf messages for protocol details.
	GetDatabaseInstance(IntegrationTest_GetDatabaseInstanceServer) error
	mustEmbedUnimplementedIntegrationTestServer()
}

// UnimplementedIntegrationTestServer must be embedded to have forward compatible implementations.
type UnimplementedIntegrationTestServer struct {
}

func (UnimplementedIntegrationTestServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedIntegrationTestServer) GetDatabaseInstance(IntegrationTest_GetDatabaseInstanceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDatabaseInstance not implemented")
}
func (UnimplementedIntegrationTestServer) mustEmbedUnimplementedIntegrationTestServer() {}

// UnsafeIntegrationTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrationTestServer will
// result in compilation errors.
type UnsafeIntegrationTestServer interface {
	mustEmbedUnimplementedIntegrationTestServer()
}

func RegisterIntegrationTestServer(s grpc.ServiceRegistrar, srv IntegrationTestServer) {
	s.RegisterService(&IntegrationTest_ServiceDesc, srv)
}

func _IntegrationTest_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationTestServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.IntegrationTest/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationTestServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationTest_GetDatabaseInstance_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IntegrationTestServer).GetDatabaseInstance(&integrationTestGetDatabaseInstanceServer{stream})
}

type IntegrationTest_GetDatabaseInstanceServer interface {
	Send(*GetDatabaseInstanceResponse) error
	Recv() (*GetDatabaseInstanceRequest, error)
	grpc.ServerStream
}

type integrationTestGetDatabaseInstanceServer struct {
	grpc.ServerStream
}

func (x *integrationTestGetDatabaseInstanceServer) Send(m *GetDatabaseInstanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *integrationTestGetDatabaseInstanceServer) Recv() (*GetDatabaseInstanceRequest, error) {
	m := new(GetDatabaseInstanceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IntegrationTest_ServiceDesc is the grpc.ServiceDesc for IntegrationTest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrationTest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.IntegrationTest",
	HandlerType: (*IntegrationTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _IntegrationTest_GetStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDatabaseInstance",
			Handler:       _IntegrationTest_GetDatabaseInstance_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server/proto/server.proto",
}
