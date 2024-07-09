// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.0
// source: retrieve.proto

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
	Retriever_Retrieve_FullMethodName = "/discovery.Retriever/Retrieve"
)

// RetrieverClient is the client API for Retriever service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetrieverClient interface {
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error)
}

type retrieverClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrieverClient(cc grpc.ClientConnInterface) RetrieverClient {
	return &retrieverClient{cc}
}

func (c *retrieverClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RetrieveResponse)
	err := c.cc.Invoke(ctx, Retriever_Retrieve_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrieverServer is the server API for Retriever service.
// All implementations must embed UnimplementedRetrieverServer
// for forward compatibility
type RetrieverServer interface {
	Retrieve(context.Context, *RetrieveRequest) (*RetrieveResponse, error)
	mustEmbedUnimplementedRetrieverServer()
}

// UnimplementedRetrieverServer must be embedded to have forward compatible implementations.
type UnimplementedRetrieverServer struct {
}

func (UnimplementedRetrieverServer) Retrieve(context.Context, *RetrieveRequest) (*RetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedRetrieverServer) mustEmbedUnimplementedRetrieverServer() {}

// UnsafeRetrieverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetrieverServer will
// result in compilation errors.
type UnsafeRetrieverServer interface {
	mustEmbedUnimplementedRetrieverServer()
}

func RegisterRetrieverServer(s grpc.ServiceRegistrar, srv RetrieverServer) {
	s.RegisterService(&Retriever_ServiceDesc, srv)
}

func _Retriever_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieverServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Retriever_Retrieve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieverServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Retriever_ServiceDesc is the grpc.ServiceDesc for Retriever service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Retriever_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Retriever",
	HandlerType: (*RetrieverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Retrieve",
			Handler:    _Retriever_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "retrieve.proto",
}