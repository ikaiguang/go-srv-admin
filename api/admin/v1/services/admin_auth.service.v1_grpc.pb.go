// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/admin/v1/services/admin_auth.service.v1.proto

package adminservicev1

import (
	context "context"
	resources "github.com/ikaiguang/go-srv-admin/api/admin/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SrvAdminAuthClient is the client API for SrvAdminAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvAdminAuthClient interface {
	// LoginByEmail Email登录
	LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// Ping ping pong
	Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error)
}

type srvAdminAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvAdminAuthClient(cc grpc.ClientConnInterface) SrvAdminAuthClient {
	return &srvAdminAuthClient{cc}
}

func (c *srvAdminAuthClient) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, "/admin.api.admin.adminservicev1.SrvAdminAuth/LoginByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvAdminAuthClient) Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error) {
	out := new(resources.PingResp)
	err := c.cc.Invoke(ctx, "/admin.api.admin.adminservicev1.SrvAdminAuth/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvAdminAuthServer is the server API for SrvAdminAuth service.
// All implementations must embed UnimplementedSrvAdminAuthServer
// for forward compatibility
type SrvAdminAuthServer interface {
	// LoginByEmail Email登录
	LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error)
	// Ping ping pong
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	mustEmbedUnimplementedSrvAdminAuthServer()
}

// UnimplementedSrvAdminAuthServer must be embedded to have forward compatible implementations.
type UnimplementedSrvAdminAuthServer struct {
}

func (UnimplementedSrvAdminAuthServer) LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByEmail not implemented")
}
func (UnimplementedSrvAdminAuthServer) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSrvAdminAuthServer) mustEmbedUnimplementedSrvAdminAuthServer() {}

// UnsafeSrvAdminAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvAdminAuthServer will
// result in compilation errors.
type UnsafeSrvAdminAuthServer interface {
	mustEmbedUnimplementedSrvAdminAuthServer()
}

func RegisterSrvAdminAuthServer(s grpc.ServiceRegistrar, srv SrvAdminAuthServer) {
	s.RegisterService(&SrvAdminAuth_ServiceDesc, srv)
}

func _SrvAdminAuth_LoginByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvAdminAuthServer).LoginByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.api.admin.adminservicev1.SrvAdminAuth/LoginByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvAdminAuthServer).LoginByEmail(ctx, req.(*resources.LoginByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvAdminAuth_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvAdminAuthServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.api.admin.adminservicev1.SrvAdminAuth/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvAdminAuthServer).Ping(ctx, req.(*resources.PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvAdminAuth_ServiceDesc is the grpc.ServiceDesc for SrvAdminAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvAdminAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.api.admin.adminservicev1.SrvAdminAuth",
	HandlerType: (*SrvAdminAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByEmail",
			Handler:    _SrvAdminAuth_LoginByEmail_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _SrvAdminAuth_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/admin/v1/services/admin_auth.service.v1.proto",
}
