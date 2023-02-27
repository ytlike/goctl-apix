// Code generated by goctl. DO NOT EDIT!
// Source: application.proto

package client

import (
	"context"

	"qbq-open-platform/applicationCenter/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiTargetApp                      = pb.ApiTargetApp
	ApiTargetAppListReq               = pb.ApiTargetAppListReq
	ApiTargetAppListResp              = pb.ApiTargetAppListResp
	CreateApplicationCallbackOrderReq = pb.CreateApplicationCallbackOrderReq
	CreateApplicationCallbackReq      = pb.CreateApplicationCallbackReq
	CreateApplicationCallbackResp     = pb.CreateApplicationCallbackResp
	CreateApplicationReq              = pb.CreateApplicationReq
	CreateApplicationResp             = pb.CreateApplicationResp
	CustomerUpdateStateReq            = pb.CustomerUpdateStateReq
	CustomerUpdateStateResp           = pb.CustomerUpdateStateResp
	GetApplicationCallbackReq         = pb.GetApplicationCallbackReq
	GetApplicationCallbackResp        = pb.GetApplicationCallbackResp
	OpenApplicationReq                = pb.OpenApplicationReq
	OpenApplicationResp               = pb.OpenApplicationResp
	UpdateApplicationCallbackReq      = pb.UpdateApplicationCallbackReq
	UpdateApplicationCallbackResp     = pb.UpdateApplicationCallbackResp
	UpdateApplicationInfoReq          = pb.UpdateApplicationInfoReq
	UpdateApplicationInfoResp         = pb.UpdateApplicationInfoResp

	ApplicationCustomer interface {
		//  应用启用/禁用
		CustomerUpdateState(ctx context.Context, in *CustomerUpdateStateReq, opts ...grpc.CallOption) (*CustomerUpdateStateResp, error)
		//  创建应用
		CreateApplication(ctx context.Context, in *CreateApplicationReq, opts ...grpc.CallOption) (*CreateApplicationResp, error)
		//  开通应用
		OpenApplication(ctx context.Context, in *OpenApplicationReq, opts ...grpc.CallOption) (*OpenApplicationResp, error)
		//  修改应用信息
		UpdateApplicationInfo(ctx context.Context, in *UpdateApplicationInfoReq, opts ...grpc.CallOption) (*UpdateApplicationInfoResp, error)
	}

	defaultApplicationCustomer struct {
		cli zrpc.Client
	}
)

func NewApplicationCustomer(cli zrpc.Client) ApplicationCustomer {
	return &defaultApplicationCustomer{
		cli: cli,
	}
}

//  应用启用/禁用
func (m *defaultApplicationCustomer) CustomerUpdateState(ctx context.Context, in *CustomerUpdateStateReq, opts ...grpc.CallOption) (*CustomerUpdateStateResp, error) {
	client := pb.NewApplicationCustomerClient(m.cli.Conn())
	return client.CustomerUpdateState(ctx, in, opts...)
}

//  创建应用
func (m *defaultApplicationCustomer) CreateApplication(ctx context.Context, in *CreateApplicationReq, opts ...grpc.CallOption) (*CreateApplicationResp, error) {
	client := pb.NewApplicationCustomerClient(m.cli.Conn())
	return client.CreateApplication(ctx, in, opts...)
}

//  开通应用
func (m *defaultApplicationCustomer) OpenApplication(ctx context.Context, in *OpenApplicationReq, opts ...grpc.CallOption) (*OpenApplicationResp, error) {
	client := pb.NewApplicationCustomerClient(m.cli.Conn())
	return client.OpenApplication(ctx, in, opts...)
}

//  修改应用信息
func (m *defaultApplicationCustomer) UpdateApplicationInfo(ctx context.Context, in *UpdateApplicationInfoReq, opts ...grpc.CallOption) (*UpdateApplicationInfoResp, error) {
	client := pb.NewApplicationCustomerClient(m.cli.Conn())
	return client.UpdateApplicationInfo(ctx, in, opts...)
}
