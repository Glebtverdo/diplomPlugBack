// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: plugBack/plugBack.proto

package plugBackGrpc

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
	PlugBackGrpc_Login_FullMethodName       = "/plugBackGrpc.PlugBackGrpc/Login"
	PlugBackGrpc_GetRequests_FullMethodName = "/plugBackGrpc.PlugBackGrpc/GetRequests"
)

// PlugBackGrpcClient is the client API for PlugBackGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlugBackGrpcClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetRequests(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Requests, error)
}

type plugBackGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewPlugBackGrpcClient(cc grpc.ClientConnInterface) PlugBackGrpcClient {
	return &plugBackGrpcClient{cc}
}

func (c *plugBackGrpcClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, PlugBackGrpc_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plugBackGrpcClient) GetRequests(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Requests, error) {
	out := new(Requests)
	err := c.cc.Invoke(ctx, PlugBackGrpc_GetRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlugBackGrpcServer is the server API for PlugBackGrpc service.
// All implementations must embed UnimplementedPlugBackGrpcServer
// for forward compatibility
type PlugBackGrpcServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetRequests(context.Context, *EmptyRequest) (*Requests, error)
	mustEmbedUnimplementedPlugBackGrpcServer()
}

// UnimplementedPlugBackGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedPlugBackGrpcServer struct {
}

func (UnimplementedPlugBackGrpcServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedPlugBackGrpcServer) GetRequests(context.Context, *EmptyRequest) (*Requests, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequests not implemented")
}
func (UnimplementedPlugBackGrpcServer) mustEmbedUnimplementedPlugBackGrpcServer() {}

// UnsafePlugBackGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlugBackGrpcServer will
// result in compilation errors.
type UnsafePlugBackGrpcServer interface {
	mustEmbedUnimplementedPlugBackGrpcServer()
}

func RegisterPlugBackGrpcServer(s grpc.ServiceRegistrar, srv PlugBackGrpcServer) {
	s.RegisterService(&PlugBackGrpc_ServiceDesc, srv)
}

func _PlugBackGrpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlugBackGrpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlugBackGrpc_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlugBackGrpcServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlugBackGrpc_GetRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlugBackGrpcServer).GetRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlugBackGrpc_GetRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlugBackGrpcServer).GetRequests(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlugBackGrpc_ServiceDesc is the grpc.ServiceDesc for PlugBackGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlugBackGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugBackGrpc.PlugBackGrpc",
	HandlerType: (*PlugBackGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _PlugBackGrpc_Login_Handler,
		},
		{
			MethodName: "GetRequests",
			Handler:    _PlugBackGrpc_GetRequests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugBack/plugBack.proto",
}