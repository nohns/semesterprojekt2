// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/pairing_service.proto

package pairingv1

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

// PairingServiceClient is the client API for PairingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PairingServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type pairingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPairingServiceClient(cc grpc.ClientConnInterface) PairingServiceClient {
	return &pairingServiceClient{cc}
}

func (c *pairingServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/pairing.v1.PairingService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PairingServiceServer is the server API for PairingService service.
// All implementations should embed UnimplementedPairingServiceServer
// for forward compatibility
type PairingServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedPairingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPairingServiceServer struct {
}

func (UnimplementedPairingServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

// UnsafePairingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PairingServiceServer will
// result in compilation errors.
type UnsafePairingServiceServer interface {
	mustEmbedUnimplementedPairingServiceServer()
}

func RegisterPairingServiceServer(s grpc.ServiceRegistrar, srv PairingServiceServer) {
	s.RegisterService(&PairingService_ServiceDesc, srv)
}

func _PairingService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pairing.v1.PairingService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PairingService_ServiceDesc is the grpc.ServiceDesc for PairingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PairingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pairing.v1.PairingService",
	HandlerType: (*PairingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PairingService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/pairing_service.proto",
}
