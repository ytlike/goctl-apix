package loader

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/rest"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"qbq-open-platform/applicationCenter/api/internal/config"
	"qbq-open-platform/applicationCenter/api/internal/handler"
	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/authCenter/api/provider/middleware"
	"qbq-open-platform/common/global"
	commonMiddleware "qbq-open-platform/common/middleware"
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
	restConfig := rest.RestConf{
		Middlewares: rest.MiddlewaresConf{
			Trace:      true,
			Log:        true,
			Prometheus: false,
			MaxConns:   true,
			Breaker:    false,
			Shedding:   false,
			Timeout:    true,
			Recover:    false,
			Metrics:    true,
			MaxBytes:   true,
			Gunzip:     false,
		},
	}
	err = copier.CopyWithOption(&restConfig, &serverConfig.Application, copier.Option{DeepCopy: true})
	if err != nil {
		log.Panic(err)
	}
	server := rest.MustNewServer(restConfig)
	ctx := svc.NewServiceContext(bootstrapConfig)
	initGlobal(ctx)

	//originalWriter := logx.Reset()
	//writer := gLog.NewSensitiveLogger(originalWriter)
	//logx.SetWriter(writer)

	// 注册到nacos
	err = nacos.NacosRegisteray(func(opt *nacos.Options) {
		opt.Host = bootstrapConfig.Configure.Host
		opt.Port = bootstrapConfig.Configure.Port
		opt.NamespaceId = bootstrapConfig.Configure.NamespaceId
		opt.Timeout = bootstrapConfig.Configure.Timeout

		opt.ApplicationName = bootstrapConfig.Server.Application.Name
		opt.ApplicationListenOn = fmt.Sprintf("%s:%d", bootstrapConfig.Server.Application.Host, bootstrapConfig.Server.Application.Port)

	})
	if err != nil {
		log.Panic(err)
	}

	// 注册全局中间件
	server.Use(commonMiddleware.LogMiddleware())
	server.Use(middleware.NewRemoteAuthMiddleware(bootstrapConfig.Server.Auth.CheckUrl).Handle)
	handler.RegisterHandlers(server, ctx)

	services := make([]service.Service, 0)
	services = append(services, server)
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

func initGlobal(svc *svc.ServiceContext) {
	global.Config().ApplicationName = svc.Server.Application.Name
	global.Config().IgnoreUrls = svc.Server.Auth.IgnoreUrls
	global.Config().RedisClient = svc.RedisClient
}
