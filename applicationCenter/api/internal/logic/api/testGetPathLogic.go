package api

import (
	"context"

	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestGetPathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestGetPathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestGetPathLogic {
	return &TestGetPathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestGetPathLogic) TestGetPath(req *types.TestGetPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
