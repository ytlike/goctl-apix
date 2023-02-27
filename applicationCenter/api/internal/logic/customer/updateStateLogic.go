package customer

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"qbq-open-platform/applicationCenter/rpc/pb"
	"qbq-open-platform/common/errorsEnums"

	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStateLogic {
	return &UpdateStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStateLogic) UpdateState(req *types.CustomerUpdateStateReq) error {
	rpcReq := &pb.CustomerUpdateStateReq{}
	err := copier.Copy(rpcReq, req)
	if err != nil {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
	rpcResp, err := l.svcCtx.CustomerApplicationRpcClient.CustomerUpdateState(l.ctx, rpcReq)
	if err != nil {
		return errors.Wrapf(err, errorsEnums.GetMsgByCode(errorsEnums.RPC_CALL_ERROR))
	}
	resp := &types.CustomerUpdateStateReq{}
	err = copier.Copy(resp, rpcResp)
	if err != nil {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
	return err
}
