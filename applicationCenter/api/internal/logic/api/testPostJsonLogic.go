package api

import (
	"context"

	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestPostJsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestPostJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestPostJsonLogic {
	return &TestPostJsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestPostJsonLogic) TestPostJson(req *types.TestPostJsonReq) error {
	// todo: add your logic here and delete this line

	return nil
}
