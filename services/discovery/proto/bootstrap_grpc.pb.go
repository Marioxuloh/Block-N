// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.0
// source: bootstrap.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Bootstraper_Bootstrap_FullMethodName = "/discovery.Bootstraper/Bootstrap"
)

// BootstraperClient is the client API for Bootstraper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BootstraperClient interface {
	Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error)
}

type bootstraperClient struct {
	cc grpc.ClientConnInterface
}

func NewBootstraperClient(cc grpc.ClientConnInterface) BootstraperClient {
	return &bootstraperClient{cc}
}

func (c *bootstraperClient) Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BootstrapResponse)
	err := c.cc.Invoke(ctx, Bootstraper_Bootstrap_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstraperServer is the server API for Bootstraper service.
// All implementations must embed UnimplementedBootstraperServer
// for forward compatibility
type BootstraperServer interface {
	Bootstrap(context.Context, *BootstrapRequest) (*BootstrapResponse, error)
	mustEmbedUnimplementedBootstraperServer()
}

// UnimplementedBootstraperServer must be embedded to have forward compatible implementations.
type UnimplementedBootstraperServer struct {
}

func (UnimplementedBootstraperServer) Bootstrap(context.Context, *BootstrapRequest) (*BootstrapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bootstrap not implemented")
}
func (UnimplementedBootstraperServer) mustEmbedUnimplementedBootstraperServer() {}

// UnsafeBootstraperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BootstraperServer will
// result in compilation errors.
type UnsafeBootstraperServer interface {
	mustEmbedUnimplementedBootstraperServer()
}

func RegisterBootstraperServer(s grpc.ServiceRegistrar, srv BootstraperServer) {
	s.RegisterService(&Bootstraper_ServiceDesc, srv)
}

func _Bootstraper_Bootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstraperServer).Bootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bootstraper_Bootstrap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstraperServer).Bootstrap(ctx, req.(*BootstrapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bootstraper_ServiceDesc is the grpc.ServiceDesc for Bootstraper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bootstraper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Bootstraper",
	HandlerType: (*BootstraperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bootstrap",
			Handler:    _Bootstraper_Bootstrap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bootstrap.proto",
}
