package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	"qbq-open-platform/common/errorsEnums"
	"strings"
)

var (
	Producer *ApplicationCenterProducer
)

type ApplicationCenterProducer struct {
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	producer rocketmq.Producer
}

func NewApplicationCenterProducer(ctx context.Context, svcCtx *svc.ServiceContext) *ApplicationCenterProducer {
	ns, err := net.LookupHost(svcCtx.Server.Rocketmq.NameServer[:strings.Index(svcCtx.Server.Rocketmq.NameServer, ":")])
	if err != nil {
		panic(err)
	}
	// 设置rocketmq组件的日志打印级别
	rlog.SetLogLevel(logrus.ErrorLevel.String())
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{ns[0] + ":9876"})),
		producer.WithRetry(3),
		producer.WithGroupName("applicationCenter-group"),
	)
	if err != nil {
		panic(err)
	}

	return &ApplicationCenterProducer{
		ctx:      ctx,
		svcCtx:   svcCtx,
		producer: p,
	}
}

func (p *ApplicationCenterProducer) Start() {
	err := p.producer.Start()
	if err != nil {
		logx.WithContext(p.ctx).Errorf("mq启动失败，原因：%v", err)
	} else {
		Producer = p
	}
}

func (p *ApplicationCenterProducer) Stop() {
	p.producer.Shutdown()
}

func (p *ApplicationCenterProducer) SendSync(topic string, data []byte) error {
	sendResult, err := p.producer.SendSync(p.ctx, &primitive.Message{
		Topic: topic,
		Body:  data,
	})
	if err != nil {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.ROCKETMQ_SEND_ERROR), "%v", err)
	}
	if sendResult.Status != primitive.SendOK {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.ROCKETMQ_SEND_ERROR), "%s", sendResult.String())
	}
	return nil
}
