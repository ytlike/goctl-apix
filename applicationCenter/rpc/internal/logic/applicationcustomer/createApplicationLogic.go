package applicationcustomerlogic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"qbq-open-platform/applicationCenter/gorm/model"
	"qbq-open-platform/applicationCenter/rpc/mq"
	"qbq-open-platform/authCenter/api/provider/template"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/utils"
	vo "qbq-open-platform/common/vo/mq"

	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	"qbq-open-platform/applicationCenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApplicationLogic {
	return &CreateApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建应用
func (l *CreateApplicationLogic) CreateApplication(in *pb.CreateApplicationReq) (*pb.CreateApplicationResp, error) {
	//获取当前用户信息
	userDetails, err := template.GetUserDetails(l.ctx)
	if err != nil {
		return nil, err
	}

	//用户名不能重复
	app := &model.ApplicationBaseInfoModel{}
	result := l.svcCtx.DbEngin.WithContext(l.ctx).Where("app_name = ? and app_developer_id = ?", in.AppName, userDetails.GetUserId()).First(app)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
	}
	if result.RowsAffected != 0 {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.APPLICATION_NAME_EXIST_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.APPLICATION_NAME_EXIST_ERROR))
	}

	//创建应用
	err = copier.Copy(&app, in)
	if err != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
	}
	app.Id = l.svcCtx.Snowflake.Generate().Int64()
	app.AppCode = utils.GetCode("AC")
	app.AppDeveloperId = userDetails.GetUserId()
	res := l.svcCtx.DbEngin.WithContext(l.ctx).Create(&app)
	if res.Error != nil {
		return nil, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", res.Error)
	}

	//同步到任务中心
	applicationMQVo := &vo.ApplicationMQVo{}
	err = copier.Copy(applicationMQVo, &app)
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

	return &pb.CreateApplicationResp{}, nil
}
