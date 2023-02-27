package loader

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"qbq-open-platform/applicationCenter/gorm/model"
	"qbq-open-platform/applicationCenter/rpc/internal/config"
	apiServer "qbq-open-platform/applicationCenter/rpc/internal/server/applicationapi"
	customerServer "qbq-open-platform/applicationCenter/rpc/internal/server/applicationcustomer"
	"qbq-open-platform/applicationCenter/rpc/internal/svc"
	mq2 "qbq-open-platform/applicationCenter/rpc/mq"
	"qbq-open-platform/applicationCenter/rpc/pb"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/global"
	"qbq-open-platform/common/interceptor/unary"
	"qbq-open-platform/common/nacos"
	"qbq-open-platform/common/utils"
)

type ServiceLoader int

func (s ServiceLoader) LoadService(configPath string) (*config.Server, []service.Service) {
	// 读取bootstrap.yml配置文件
	bootstrapYmlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panic(err)
	}

	content := string(bootstrapYmlFile)
	content = utils.ParseYamlEnv(string(content))
	bootstrapConfig := &config.Bootstrap{}
	err = yaml.Unmarshal([]byte(content), bootstrapConfig)

	if err != nil {
		log.Panic(err)
	}

	// 从nacos读取配置文件
	content, err = nacos.NacosConfigure(func(opt *nacos.Options) {
		opt.Host = bootstrapConfig.Configure.Host
		opt.Port = bootstrapConfig.Configure.Port
		opt.NamespaceId = bootstrapConfig.Configure.NamespaceId
		opt.Timeout = bootstrapConfig.Configure.Timeout
		opt.Group = bootstrapConfig.Configure.Group
		opt.DataId = bootstrapConfig.Configure.DataId
	})
	if err != nil {
		log.Panic(err)
	}

	serverConfig := &config.Server{}
	err = yaml.Unmarshal([]byte(content), serverConfig)
	if err != nil {
		log.Panic(err)
	}
	bootstrapConfig.Server = serverConfig
	threading.GoSafe(func() {
		processContentChange(bootstrapConfig.Server)
	})

	ctx := svc.NewServiceContext(bootstrapConfig)

	rpcServerConf := &zrpc.RpcServerConf{
		Middlewares: zrpc.ServerMiddlewaresConf{
			Trace:      true,
			Recover:    false,
			Stat:       false,
			Prometheus: false,
			Breaker:    false,
		},
	}
	err = copier.CopyWithOption(&rpcServerConf, bootstrapConfig.Server, copier.Option{DeepCopy: true})
	if err != nil {
		log.Panic(err)
	}
	initGlobal(ctx)

	server := zrpc.MustNewServer(*rpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterApplicationApiServer(grpcServer, apiServer.NewApplicationApiServer(ctx))
		pb.RegisterApplicationCustomerServer(grpcServer, customerServer.NewApplicationCustomerServer(ctx))

		if bootstrapConfig.Server.Application.Mode == service.DevMode || bootstrapConfig.Server.Application.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	server.AddUnaryInterceptors(unary.ServerLoggerInterceptor)
	server.AddUnaryInterceptors(unary.ServerTokenInterceptor)

	// 注册到nacos
	err = nacos.NacosRegisteray(func(opt *nacos.Options) {
		opt.Host = bootstrapConfig.Configure.Host
		opt.Port = bootstrapConfig.Configure.Port
		opt.NamespaceId = bootstrapConfig.Configure.NamespaceId
		opt.Timeout = bootstrapConfig.Configure.Timeout

		opt.ApplicationName = bootstrapConfig.Server.Application.Name
		opt.ApplicationListenOn = bootstrapConfig.Server.Application.ListenOn

	})
	if err != nil {
		log.Panic(err)
	}

	err = initApplication(context.Background(), ctx)
	if err != nil {
		log.Panic(err)
	}

	// rocketmq
	rocketMQProducerServices := mq2.NewApplicationCenterProducer(context.Background(), ctx)
	rocketMQConsumerServices := mq2.NewApplicationCenterConsumer(context.Background(), ctx)

	services := make([]service.Service, 0)
	services = append(services, server)
	services = append(services, rocketMQProducerServices)
	services = append(services, rocketMQConsumerServices)
	return serverConfig, services
}

func processContentChange(server *config.Server) {
	defer close(nacos.ChangeContent)
	for {
		select {
		case content := <-nacos.ChangeContent:
			s := &config.Server{}
			err := yaml.Unmarshal([]byte(content), s)
			err = copier.CopyWithOption(server, s, copier.Option{DeepCopy: true})
			if err != nil {
				logx.Error(err)
			}
		case <-nacos.ChangeContentCtx.Done():
			return
		}
	}
}

func initApplication(ctx context.Context, svc *svc.ServiceContext) error {
	var applicationList []*model.ApplicationBaseInfoModel
	result := svc.DbEngin.WithContext(ctx).Where("app_status = 1 and deleted = 0").Find(&applicationList)
	if result.Error != nil {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.DB_ERROR), "%v", result.Error)
	}

	for _, app := range applicationList {
		appJson, err := utils.Struct2json(app)
		if err != nil {
			return err
		}
		err = svc.RedisClient.HsetCtx(ctx, global.INIT_APPLICATION_MAP_KEY, app.AppCode, appJson)
		if err != nil {
			return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.REDIS_ERROR), "%v", err)
		}
	}
	return nil
}

func initGlobal(svc *svc.ServiceContext) {
	global.Config().ApplicationName = svc.Server.Application.Name
	global.Config().RedisClient = svc.RedisClient
}
