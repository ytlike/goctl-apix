package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"os"
	"qbq-open-platform/applicationCenter/rpc/loader"
	"qbq-open-platform/common"
)

var configFile = flag.String("f", "etc/bootstrap.yaml", "the config file")

func main() {
	if common.ExecArgs(os.Args) {
		//writer := graylog.NewGrayLogWriter(func(logger *logrus.Logger) {
		//	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		//})
		//logx.SetWriter(writer)

		serviceGroup := service.NewServiceGroup()
		defer serviceGroup.Stop()

		serviceLoader := new(loader.ServiceLoader)
		config, serviceList := serviceLoader.LoadService(*configFile)
		for _, sev := range serviceList {
			serviceGroup.Add(sev)
		}
		logx.Infof("starting %s server at %s:%d", config.ServiceConf.Name, config.Application.Host, config.Application.Port)
		serviceGroup.Start()
	}
}
