package mq

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"qbq-open-platform/applicationCenter/gorm/model"
	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/global"
	"qbq-open-platform/common/utils"
	vo "qbq-open-platform/common/vo/mq"
	"strings"
)

type ApplicationCenterConsumer struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	consumer rocketmq.PushConsumer
}

func NewApplicationCenterConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *ApplicationCenterConsumer {
	ns, err := net.LookupHost(svcCtx.Server.Rocketmq.NameServer[:strings.Index(svcCtx.Server.Rocketmq.NameServer, ":")])
	if err != nil {
		logx.WithContext(ctx).Errorf("创建mq消费者失败，原因：%v", err)
	}
	// 设置rocketmq组件的日志打印级别
	rlog.SetLogLevel(logrus.ErrorLevel.String())
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("applicationCenter-group"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{ns[0] + ":9876"})),
	)
	if err != nil {
		logx.WithContext(ctx).Errorf("创建mq消费者失败，原因：%v", err)
	}

	return &ApplicationCenterConsumer{
		ctx:      ctx,
		svcCtx:   svcCtx,
		consumer: c,
	}
}

func (c *ApplicationCenterConsumer) Start() {
	err := c.consumer.Subscribe("applicationCenter-application-topic", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			vo := &vo.ApplicationMQVo{}
			err := json.Unmarshal(msgs[i].Body, vo)
			if err != nil {
				logx.WithContext(ctx).Errorf("mq消费应用更新消息失败，原因：%v", err)
				return consumer.ConsumeRetryLater, err
			} else {
				logx.Info("消费创建应用消息： %v", vo)
				err = c.saveOrUpdateApplication(vo)
				if err != nil {
					logx.WithContext(ctx).Errorf("mq消费应用更新消息失败，原因：%+v", err)
					return consumer.ConsumeRetryLater, err
				}

				//发送消息到任务中心
				appJson, err := json.Marshal(vo)
				if err != nil {
					return consumer.ConsumeRetryLater, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "json序列化失败, %v", err)
				}
				sendMqErr := Producer.SendSync("taskCenter-application-topic", appJson)
				if sendMqErr != nil {
					return consumer.ConsumeRetryLater, errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.ROCKETMQ_SEND_ERROR), "%v", sendMqErr)
				}
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		logx.WithContext(c.ctx).Errorf("mq订阅topic失败，原因：%v", err)
	}
	err = c.consumer.Start()
	if err != nil {
		logx.WithContext(c.ctx).Errorf("mq启动失败，原因：%v", err)
	}
}

func (c *ApplicationCenterConsumer) Stop() {
	c.consumer.Shutdown()
}

func (c *ApplicationCenterConsumer) saveOrUpdateApplication(vo *vo.ApplicationMQVo) error {
	jsonByte, err := json.Marshal(vo)
	logx.WithContext(c.ctx).Infof("mq消费到应用更新消息，参数：%s", string(jsonByte))

	application := &model.ApplicationBaseInfoModel{}
	err = copier.Copy(application, vo)
	if err != nil {
		logx.WithContext(c.ctx).WithContext(c.ctx).Errorf("mq消费应用更新消息失败，原因：%s", err.Error())
		return err
	}

	var count int64
	result := c.svcCtx.DbEngin.Model(application).Where("app_code = ?", vo.AppCode).Count(&count)
	if result.Error != nil {
		logx.WithContext(c.ctx).Errorf("%s，原因：%s", errorsEnums.GetMsgByCode(errorsEnums.DB_ERROR), result.Error)
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
	}
	if count > 0 {
		// 更新
		result = c.svcCtx.DbEngin.Save(application)
		if result.Error != nil {
			logx.WithContext(c.ctx).Errorf("%s，原因：%s", errorsEnums.GetMsgByCode(errorsEnums.DB_ERROR), result.Error)
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
		}
	} else {
		// 新增
		result = c.svcCtx.DbEngin.Create(application)
		if result.Error != nil {
			logx.WithContext(c.ctx).Errorf("%s，原因：%s", errorsEnums.GetMsgByCode(errorsEnums.DB_ERROR), result.Error)
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
		}
	}
	appJson, err := utils.Struct2json(application)
	if err != nil {
		return err
	}
	err = global.Config().RedisClient.HsetCtx(c.ctx, global.INIT_APPLICATION_MAP_KEY, application.AppCode, appJson)
	if err != nil {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
	}
	return nil
}
