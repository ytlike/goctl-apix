package token

import (
	"context"
)

// TokenHandle token处理接口
type TokenHandle interface {
	// ApiCheckToken API层token校验
	ApiCheckToken(ctx *context.Context, accessToken string) error
	// RpcCheckToken RPC层token校验
	RpcCheckToken(ctx context.Context) error
	// RenewToken token续期
	RenewToken(ctx context.Context, value interface{}, accessToken string) (bool, error)
	// DeleteToken 删除token
	DeleteToken(ctx context.Context, accessToken string) error
	// RpcGetToken RPC层获取token
	RpcGetToken(ctx context.Context) (string, error)
	// RpcGetValue RPC层获取token对应的值
	RpcGetValue(ctx context.Context) (string, error)
}
