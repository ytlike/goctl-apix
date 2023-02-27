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

type OpenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenLogic {
	return &OpenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenLogic) Open(req *types.OpenApplicationReq) error {
	rpcReq := &pb.OpenApplicationReq{}
	err := copier.Copy(rpcReq, req)
	if err != nil {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
	_, err = l.svcCtx.CustomerApplicationRpcClient.OpenApplication(l.ctx, rpcReq)
	if err != nil {
		return errors.Wrapf(err, errorsEnums.GetMsgByCode(errorsEnums.RPC_CALL_ERROR))
	}
	return nil
}
