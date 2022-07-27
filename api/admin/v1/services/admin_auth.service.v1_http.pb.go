// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package adminv1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/ikaiguang/go-srv-admin/api/admin/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvAdminAuthLoginByEmail = "/admin.api.admin.adminservicev1.SrvAdminAuth/LoginByEmail"
const OperationSrvAdminAuthPing = "/admin.api.admin.adminservicev1.SrvAdminAuth/Ping"

type SrvAdminAuthHTTPServer interface {
	LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error)
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
}

func RegisterSrvAdminAuthHTTPServer(s *http.Server, srv SrvAdminAuthHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/login-by-email", _SrvAdminAuth_LoginByEmail0_HTTP_Handler(srv))
	r.GET("/admin/v1/ping", _SrvAdminAuth_Ping0_HTTP_Handler(srv))
}

func _SrvAdminAuth_LoginByEmail0_HTTP_Handler(srv SrvAdminAuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.LoginByEmailReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvAdminAuthLoginByEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByEmail(ctx, req.(*resources.LoginByEmailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.LoginResp)
		return ctx.Result(200, reply)
	}
}

func _SrvAdminAuth_Ping0_HTTP_Handler(srv SrvAdminAuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.PingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvAdminAuthPing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*resources.PingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.PingResp)
		return ctx.Result(200, reply)
	}
}

type SrvAdminAuthHTTPClient interface {
	LoginByEmail(ctx context.Context, req *resources.LoginByEmailReq, opts ...http.CallOption) (rsp *resources.LoginResp, err error)
	Ping(ctx context.Context, req *resources.PingReq, opts ...http.CallOption) (rsp *resources.PingResp, err error)
}

type SrvAdminAuthHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvAdminAuthHTTPClient(client *http.Client) SrvAdminAuthHTTPClient {
	return &SrvAdminAuthHTTPClientImpl{client}
}

func (c *SrvAdminAuthHTTPClientImpl) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...http.CallOption) (*resources.LoginResp, error) {
	var out resources.LoginResp
	pattern := "/admin/v1/login-by-email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvAdminAuthLoginByEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvAdminAuthHTTPClientImpl) Ping(ctx context.Context, in *resources.PingReq, opts ...http.CallOption) (*resources.PingResp, error) {
	var out resources.PingResp
	pattern := "/admin/v1/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvAdminAuthPing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
