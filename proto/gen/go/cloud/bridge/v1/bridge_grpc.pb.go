// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bridge

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

// BridgeServiceClient is the client API for BridgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	StreamActions(ctx context.Context, in *StreamActionsRequest, opts ...grpc.CallOption) (BridgeService_StreamActionsClient, error)
}

type bridgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeServiceClient(cc grpc.ClientConnInterface) BridgeServiceClient {
	return &bridgeServiceClient{cc}
}

func (c *bridgeServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/dk.nohns.cloud.bridge.BridgeService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) StreamActions(ctx context.Context, in *StreamActionsRequest, opts ...grpc.CallOption) (BridgeService_StreamActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BridgeService_ServiceDesc.Streams[0], "/dk.nohns.cloud.bridge.BridgeService/StreamActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &bridgeServiceStreamActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BridgeService_StreamActionsClient interface {
	Recv() (*StreamAction, error)
	grpc.ClientStream
}

type bridgeServiceStreamActionsClient struct {
	grpc.ClientStream
}

func (x *bridgeServiceStreamActionsClient) Recv() (*StreamAction, error) {
	m := new(StreamAction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BridgeServiceServer is the server API for BridgeService service.
// All implementations should embed UnimplementedBridgeServiceServer
// for forward compatibility
type BridgeServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	StreamActions(*StreamActionsRequest, BridgeService_StreamActionsServer) error
}

// UnimplementedBridgeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBridgeServiceServer struct {
}

func (UnimplementedBridgeServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedBridgeServiceServer) StreamActions(*StreamActionsRequest, BridgeService_StreamActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamActions not implemented")
}

// UnsafeBridgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeServiceServer will
// result in compilation errors.
type UnsafeBridgeServiceServer interface {
	mustEmbedUnimplementedBridgeServiceServer()
}

func RegisterBridgeServiceServer(s grpc.ServiceRegistrar, srv BridgeServiceServer) {
	s.RegisterService(&BridgeService_ServiceDesc, srv)
}

func _BridgeService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dk.nohns.cloud.bridge.BridgeService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_StreamActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamActionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BridgeServiceServer).StreamActions(m, &bridgeServiceStreamActionsServer{stream})
}

type BridgeService_StreamActionsServer interface {
	Send(*StreamAction) error
	grpc.ServerStream
}

type bridgeServiceStreamActionsServer struct {
	grpc.ServerStream
}

func (x *bridgeServiceStreamActionsServer) Send(m *StreamAction) error {
	return x.ServerStream.SendMsg(m)
}

// BridgeService_ServiceDesc is the grpc.ServiceDesc for BridgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BridgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dk.nohns.cloud.bridge.BridgeService",
	HandlerType: (*BridgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _BridgeService_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamActions",
			Handler:       _BridgeService_StreamActions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cloud/bridge/v1/bridge.proto",
}
