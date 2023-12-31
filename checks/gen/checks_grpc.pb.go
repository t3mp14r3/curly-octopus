// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: checks.proto

package gen

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

const (
	Checks_Create_FullMethodName = "/Checks/Create"
)

// ChecksClient is the client API for Checks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChecksClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type checksClient struct {
	cc grpc.ClientConnInterface
}

func NewChecksClient(cc grpc.ClientConnInterface) ChecksClient {
	return &checksClient{cc}
}

func (c *checksClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Checks_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChecksServer is the server API for Checks service.
// All implementations must embed UnimplementedChecksServer
// for forward compatibility
type ChecksServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedChecksServer()
}

// UnimplementedChecksServer must be embedded to have forward compatible implementations.
type UnimplementedChecksServer struct {
}

func (UnimplementedChecksServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChecksServer) mustEmbedUnimplementedChecksServer() {}

// UnsafeChecksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChecksServer will
// result in compilation errors.
type UnsafeChecksServer interface {
	mustEmbedUnimplementedChecksServer()
}

func RegisterChecksServer(s grpc.ServiceRegistrar, srv ChecksServer) {
	s.RegisterService(&Checks_ServiceDesc, srv)
}

func _Checks_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChecksServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Checks_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChecksServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Checks_ServiceDesc is the grpc.ServiceDesc for Checks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Checks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Checks",
	HandlerType: (*ChecksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Checks_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checks.proto",
}
