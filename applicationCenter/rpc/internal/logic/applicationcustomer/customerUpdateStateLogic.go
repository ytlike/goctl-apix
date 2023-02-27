package applicationcustomerlogic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"qbq-open-platform/applicationCenter/gorm/model"
	"qbq-open-platform/applicationCenter/rpc/mq"
	"qbq-open-platform/authCenter/api/provider/template"
	"qbq-open-platform/common/errorsEnums"
	vo "qbq-open-platform/common/vo/mq"
	"time"

	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	"qbq-open-platform/applicationCenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerUpdateStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerUpdateStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerUpdateStateLogic {
	return &CustomerUpdateStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用启用/禁用
func (l *CustomerUpdateStateLogic) CustomerUpdateState(in *pb.CustomerUpdateStateReq) (*pb.CustomerUpdateStateResp, error) {
	customerUpdateStateResp := &pb.CustomerUpdateStateResp{}
	appModel := &model.ApplicationBaseInfoModel{}

	userDetails, err := template.GetUserDetails(l.ctx)
	if err != nil {
		return nil, err
	}

	result := l.svcCtx.DbEngin.WithContext(l.ctx).
		Where("app_code = ? and deleted = 0 and app_developer_id = ?", in.AppCode, userDetails.GetUserId()).
		First(appModel)
	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.APP_NOT_EXIST_ERROR), errorsEnums.GetMsgByCode(errorsEnums.APP_NOT_EXIST_ERROR))
	}
	if result.Error != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
	}

	now := time.Now()
	appModel.UpdateTime = &now
	result = l.svcCtx.DbEngin.WithContext(l.ctx).Model(&appModel).
		Where("app_code = ? and deleted = 0", in.AppCode).
		Update("app_status", in.State)
	if result.Error != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
	}

	applicationMQVo := &vo.ApplicationMQVo{}
	err = copier.Copy(applicationMQVo, &appModel)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
	appJson, err := json.Marshal(applicationMQVo)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "json序列化失败, %v", err)
	}

	err = mq.Producer.SendSync("taskCenter-application-topic", appJson)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.ROCKETMQ_SEND_ERROR), "%v", err)
	}

	return customerUpdateStateResp, nil
}
