// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package userservice

import (
	"context"

	"mxshop/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SignOutReq  = pb.SignOutReq
	SignOutResp = pb.SignOutResp

	UserService interface {
		SignOut(ctx context.Context, in *SignOutReq, opts ...grpc.CallOption) (*SignOutResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) SignOut(ctx context.Context, in *SignOutReq, opts ...grpc.CallOption) (*SignOutResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.SignOut(ctx, in, opts...)
}
