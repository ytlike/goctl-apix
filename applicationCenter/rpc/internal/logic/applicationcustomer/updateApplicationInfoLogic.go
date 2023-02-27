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
	vo "qbq-open-platform/common/vo/mq"

	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	"qbq-open-platform/applicationCenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApplicationInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateApplicationInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApplicationInfoLogic {
	return &UpdateApplicationInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改应用信息
func (l *UpdateApplicationInfoLogic) UpdateApplicationInfo(in *pb.UpdateApplicationInfoReq) (*pb.UpdateApplicationInfoResp, error) {
	err := l.svcCtx.DbEngin.Transaction(func(tx *gorm.DB) error {
		userDetails, err := template.GetUserDetails(l.ctx)
		if err != nil {
			return err
		}

		appId := in.GetAppId()
		var appModel = &model.ApplicationBaseInfoModel{}
		result := tx.WithContext(l.ctx).Where("id = ? and app_developer_id = ?", appId, userDetails.GetUserId()).First(appModel)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.APP_NOT_EXIST_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.APP_NOT_EXIST_ERROR))
		}

		//更新数据库
		res := tx.WithContext(l.ctx).Model(appModel).Where("id = ?", appId).Updates(map[string]interface{}{
			"app_name": in.GetAppName(),
			"app_icon": in.GetAppIconFileId(),
			"app_desc": in.GetAppDesc(),
		})
		if res.Error != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", res.Error)
		}

		//同步应用信息到任务中心
		applicationMQVo := &vo.ApplicationMQVo{}
		err = copier.Copy(applicationMQVo, &appModel)
		if err != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "数据转换失败, %v", err)
		}
		appJson, err := json.Marshal(applicationMQVo)
		if err != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "json序列化失败, %v", err)
		}

		err = mq.Producer.SendSync("taskCenter-application-topic", appJson)
		if err != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.ROCKETMQ_SEND_ERROR), "%v", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateApplicationInfoResp{}, nil
}
