package api

import (
	"context"

	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGetQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestGetQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGetQueryLogic {
	return &TestGetQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestGetQueryLogic) TestGetQuery(req *types.TestGetQueryReq) error {
	// todo: add your logic here and delete this line

	return nil
}
