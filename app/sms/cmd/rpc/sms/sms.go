// Code generated by goctl. DO NOT EDIT!
// Source: sms.proto

package sms

import (
	"context"

	"github/community-online/app/sms/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetMsgCodeReq  = pb.GetMsgCodeReq
	GetMsgCodeResp = pb.GetMsgCodeResp

	Sms interface {
		GetMsgCode(ctx context.Context, in *GetMsgCodeReq, opts ...grpc.CallOption) (*GetMsgCodeResp, error)
	}

	defaultSms struct {
		cli zrpc.Client
	}
)

func NewSms(cli zrpc.Client) Sms {
	return &defaultSms{
		cli: cli,
	}
}

func (m *defaultSms) GetMsgCode(ctx context.Context, in *GetMsgCodeReq, opts ...grpc.CallOption) (*GetMsgCodeResp, error) {
	client := pb.NewSmsClient(m.cli.Conn())
	return client.GetMsgCode(ctx, in, opts...)
}
