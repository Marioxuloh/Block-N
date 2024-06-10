// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.0
// source: discovery.proto

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
	Discoverer_Discovery_FullMethodName = "/discovery.Discoverer/Discovery"
)

// DiscovererClient is the client API for Discoverer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscovererClient interface {
	Discovery(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type discovererClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscovererClient(cc grpc.ClientConnInterface) DiscovererClient {
	return &discovererClient{cc}
}

func (c *discovererClient) Discovery(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, Discoverer_Discovery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscovererServer is the server API for Discoverer service.
// All implementations must embed UnimplementedDiscovererServer
// for forward compatibility
type DiscovererServer interface {
	Discovery(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
	mustEmbedUnimplementedDiscovererServer()
}

// UnimplementedDiscovererServer must be embedded to have forward compatible implementations.
type UnimplementedDiscovererServer struct {
}

func (UnimplementedDiscovererServer) Discovery(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discovery not implemented")
}
func (UnimplementedDiscovererServer) mustEmbedUnimplementedDiscovererServer() {}

// UnsafeDiscovererServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscovererServer will
// result in compilation errors.
type UnsafeDiscovererServer interface {
	mustEmbedUnimplementedDiscovererServer()
}

func RegisterDiscovererServer(s grpc.ServiceRegistrar, srv DiscovererServer) {
	s.RegisterService(&Discoverer_ServiceDesc, srv)
}

func _Discoverer_Discovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscovererServer).Discovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Discoverer_Discovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscovererServer).Discovery(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Discoverer_ServiceDesc is the grpc.ServiceDesc for Discoverer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Discoverer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discoverer",
	HandlerType: (*DiscovererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Discovery",
			Handler:    _Discoverer_Discovery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}
