package api

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"
	"qbq-open-platform/applicationCenter/rpc/pb"
	"qbq-open-platform/common/errorsEnums"
)

type TargetAppListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTargetAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TargetAppListLogic {
	return &TargetAppListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TargetAppListLogic) TargetAppList(req *types.ApiTargetAppListReq) (resp *types.ApiTargetAppListResp, err error) {
	rpcReq := &pb.ApiTargetAppListReq{}
	err = copier.Copy(rpcReq, req)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
	rpcResp, err := l.svcCtx.ApiApplicationRpcClient.ApiTargetAppList(l.ctx, rpcReq)
	if err != nil {
		return nil, errors.Wrapf(err, errorsEnums.GetMsgByCode(errorsEnums.RPC_CALL_ERROR))
	}
	resp = &types.ApiTargetAppListResp{}
	err = copier.Copy(resp, rpcResp)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
	return resp, err
}
