package api

import (
	"context"

	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestPostFormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestPostFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestPostFormLogic {
	return &TestPostFormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestPostFormLogic) TestPostForm(req *types.TestPostFormReq) error {
	// todo: add your logic here and delete this line

	return nil
}
