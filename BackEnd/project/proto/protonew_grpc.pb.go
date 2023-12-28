// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: protonew.proto

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

const (
	YourService_Login_FullMethodName    = "/ProStoreGolang.YourService/Login"
	YourService_Register_FullMethodName = "/ProStoreGolang.YourService/Register"
)

// YourServiceClient is the client API for YourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YourServiceClient interface {
	Login(ctx context.Context, in *YourLoginRequest, opts ...grpc.CallOption) (*YourLoginResponse, error)
	// New Registration RPC
	Register(ctx context.Context, in *YourRegistrationRequest, opts ...grpc.CallOption) (*YourRegistrationResponse, error)
}

type yourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYourServiceClient(cc grpc.ClientConnInterface) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) Login(ctx context.Context, in *YourLoginRequest, opts ...grpc.CallOption) (*YourLoginResponse, error) {
	out := new(YourLoginResponse)
	err := c.cc.Invoke(ctx, YourService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yourServiceClient) Register(ctx context.Context, in *YourRegistrationRequest, opts ...grpc.CallOption) (*YourRegistrationResponse, error) {
	out := new(YourRegistrationResponse)
	err := c.cc.Invoke(ctx, YourService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YourServiceServer is the server API for YourService service.
// All implementations must embed UnimplementedYourServiceServer
// for forward compatibility
type YourServiceServer interface {
	Login(context.Context, *YourLoginRequest) (*YourLoginResponse, error)
	// New Registration RPC
	Register(context.Context, *YourRegistrationRequest) (*YourRegistrationResponse, error)
	mustEmbedUnimplementedYourServiceServer()
}

// UnimplementedYourServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYourServiceServer struct {
}

func (UnimplementedYourServiceServer) Login(context.Context, *YourLoginRequest) (*YourLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedYourServiceServer) Register(context.Context, *YourRegistrationRequest) (*YourRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedYourServiceServer) mustEmbedUnimplementedYourServiceServer() {}

// UnsafeYourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YourServiceServer will
// result in compilation errors.
type UnsafeYourServiceServer interface {
	mustEmbedUnimplementedYourServiceServer()
}

func RegisterYourServiceServer(s grpc.ServiceRegistrar, srv YourServiceServer) {
	s.RegisterService(&YourService_ServiceDesc, srv)
}

func _YourService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YourLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YourService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Login(ctx, req.(*YourLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YourService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YourRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YourService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Register(ctx, req.(*YourRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YourService_ServiceDesc is the grpc.ServiceDesc for YourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProStoreGolang.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _YourService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _YourService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protonew.proto",
}