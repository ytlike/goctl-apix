package applicationcustomerlogic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"qbq-open-platform/applicationCenter/gorm/model"
	"qbq-open-platform/applicationCenter/rpc/internal/config"
	"qbq-open-platform/applicationCenter/rpc/mq"
	"qbq-open-platform/authCenter/api/provider/template"
	"qbq-open-platform/common/enums"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/global"
	"qbq-open-platform/common/paas"
	vo "qbq-open-platform/common/vo/mq"
	"strconv"

	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	"qbq-open-platform/applicationCenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenApplicationLogic {
	return &OpenApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 开通应用
func (l *OpenApplicationLogic) OpenApplication(in *pb.OpenApplicationReq) (*pb.OpenApplicationResp, error) {
	err := l.svcCtx.DbEngin.Transaction(func(tx *gorm.DB) error {
		userDetails, err := template.GetUserDetails(l.ctx)
		if err != nil {
			return err
		}

		openAppKey := global.APPLICATION_OPEN_KEY + strconv.FormatInt(in.GetApplicationId(), 10)
		redisLock := redis.NewRedisLock(l.svcCtx.RedisClient, openAppKey)
		redisLock.SetExpire(10)
		defer func() {
			if r := recover(); r != nil {
				err = errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.INTERNAL_SERVER_ERROR), "开通应用异常: %v", r)
				logx.WithContext(l.ctx).Errorf("%+v", err)
			}
			redisLock.Release()
		}()
		if ok, err := redisLock.Acquire(); !ok || err != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_LOCK_ERROR), "当前有其他用户正在进行操作，请稍后重试: %v", err)
		}

		appId := in.GetApplicationId()
		var appModel = &model.ApplicationBaseInfoModel{}
		result := tx.WithContext(l.ctx).Where("id = ? and app_developer_id = ?", appId, userDetails.GetUserId()).First(appModel)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.APP_NOT_EXIST_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.APP_NOT_EXIST_ERROR))
		}
		appKeyExist := appModel.AppKey
		if appKeyExist != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.APP_OPENED_ERROR), "%s", errorsEnums.GetMsgByCode(errorsEnums.APP_OPENED_ERROR))
		}

		//获取PaasAppId
		var abilityUrl string
		var paasAppIdStart int64 = 0
		if in.GetOption() == enums.APP_ABILITY_URL_TYPE_ALIYUN {
			abilityUrl = l.svcCtx.Server.AbilityUrl.AliCloudUrl
			paasAppIdStart = l.svcCtx.Server.OpenApplication.AliCloudAppIdBegin
		} else if in.GetOption() == enums.APP_ABILITY_URL_TYPE_DIGITAL_CHONGQING {
			abilityUrl = l.svcCtx.Server.AbilityUrl.DigitalChongQingUrl
			paasAppIdStart = l.svcCtx.Server.OpenApplication.DigitalChongQingAppIdBegin
		} else {
		}
		paasPlatform := paas.New(func(paas *paas.Paas) {
		})
		paasPlatform.PaasUrl = abilityUrl
		paasPlatform.ClientId = l.svcCtx.Server.Paas.ClientId
		paasPlatform.ClientSecret = l.svcCtx.Server.Paas.ClientSecret
		appDto := &model.ApplicationBaseInfoModel{}
		result = tx.WithContext(l.ctx).Select("paas_app_id").Where("app_ability_url = ?", paasPlatform.PaasUrl).Order("paas_app_id desc").First(appDto)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
		}
		var paasAppId int64
		if result.RowsAffected == 0 {
			paasAppId = paasAppIdStart + 1
		} else {
			maxPaasAppId, convertErr := strconv.ParseInt(*appDto.PaasAppId, 10, 64)
			if convertErr != nil {
				return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DATA_TYPE_CONVERT_ERROR), "%v", convertErr)
			}
			paasAppId = maxPaasAppId + 1
		}

		//调用paas接口新增client
		appKey, appSecret, err := paasAddClient(l.ctx, paasPlatform, strconv.FormatInt(paasAppId, 10), appModel.AppName+":"+strconv.FormatInt(appModel.Id, 10), l.svcCtx.Server)
		if err != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.APP_OPEN_ABILITY_ERROR), "%s：%s", errorsEnums.GetMsgByCode(errorsEnums.APP_OPEN_ABILITY_ERROR), err.Error())
		}

		//更新应用信息表
		res := tx.WithContext(l.ctx).Model(appModel).Where("id = ?", appId).Updates(map[string]interface{}{
			"app_key":         appKey,
			"app_secret":      appSecret,
			"app_ability_url": abilityUrl,
			"paas_app_id":     paasAppId,
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
	return &pb.OpenApplicationResp{}, nil
}

func paasAddClient(ctx context.Context, paasPlatform *paas.Paas, appId string, appName string, openApplication *config.Server) (string, string, error) {
	clientAddReq := &paas.ClientAddReq{
		ClientName:           appName,
		AppId:                appId,
		PId:                  openApplication.PId,
		ResourceIds:          openApplication.ResourceIds,
		AccessTokenValidity:  openApplication.AccessTokenValidity,
		RefreshTokenValidity: openApplication.RefreshTokenValidity,
		GrantTypes:           openApplication.GrantTypes,
	}
	addClientResp, err := paasPlatform.PaasAddClient(ctx, clientAddReq)
	if err != nil {
		return "", "", err
	}
	return addClientResp.ClientId, addClientResp.ClientSecret, nil
}
