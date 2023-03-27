// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/lock_service.proto

package lockv1

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

// LockServiceClient is the client API for LockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LockServiceClient interface {
	GetLockState(ctx context.Context, in *GetLockStateRequest, opts ...grpc.CallOption) (*GetLockStateResponse, error)
	SetLockState(ctx context.Context, in *SetLockStateRequest, opts ...grpc.CallOption) (*SetLockStateResponse, error)
}

type lockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLockServiceClient(cc grpc.ClientConnInterface) LockServiceClient {
	return &lockServiceClient{cc}
}

func (c *lockServiceClient) GetLockState(ctx context.Context, in *GetLockStateRequest, opts ...grpc.CallOption) (*GetLockStateResponse, error) {
	out := new(GetLockStateResponse)
	err := c.cc.Invoke(ctx, "/lock.v1.LockService/GetLockState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) SetLockState(ctx context.Context, in *SetLockStateRequest, opts ...grpc.CallOption) (*SetLockStateResponse, error) {
	out := new(SetLockStateResponse)
	err := c.cc.Invoke(ctx, "/lock.v1.LockService/SetLockState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockServiceServer is the server API for LockService service.
// All implementations should embed UnimplementedLockServiceServer
// for forward compatibility
type LockServiceServer interface {
	GetLockState(context.Context, *GetLockStateRequest) (*GetLockStateResponse, error)
	SetLockState(context.Context, *SetLockStateRequest) (*SetLockStateResponse, error)
}

// UnimplementedLockServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLockServiceServer struct {
}

func (UnimplementedLockServiceServer) GetLockState(context.Context, *GetLockStateRequest) (*GetLockStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLockState not implemented")
}
func (UnimplementedLockServiceServer) SetLockState(context.Context, *SetLockStateRequest) (*SetLockStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLockState not implemented")
}

// UnsafeLockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LockServiceServer will
// result in compilation errors.
type UnsafeLockServiceServer interface {
	mustEmbedUnimplementedLockServiceServer()
}

func RegisterLockServiceServer(s grpc.ServiceRegistrar, srv LockServiceServer) {
	s.RegisterService(&LockService_ServiceDesc, srv)
}

func _LockService_GetLockState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLockStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).GetLockState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lock.v1.LockService/GetLockState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).GetLockState(ctx, req.(*GetLockStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_SetLockState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLockStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).SetLockState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lock.v1.LockService/SetLockState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).SetLockState(ctx, req.(*SetLockStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LockService_ServiceDesc is the grpc.ServiceDesc for LockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lock.v1.LockService",
	HandlerType: (*LockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLockState",
			Handler:    _LockService_GetLockState_Handler,
		},
		{
			MethodName: "SetLockState",
			Handler:    _LockService_SetLockState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/lock_service.proto",
}
