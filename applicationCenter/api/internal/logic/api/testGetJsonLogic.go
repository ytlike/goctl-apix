package api

import (
	"context"

	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGetJsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestGetJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGetJsonLogic {
	return &TestGetJsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestGetJsonLogic) TestGetJson(req *types.TestGetJsonReq) error {
	// todo: add your logic here and delete this line

	return nil
}
