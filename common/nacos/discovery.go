package nacos

import (
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"qbq-open-platform/common"
)

// 注册到nacos
func NacosRegisteray(opts ...Option) error {
	options := &Options{
		Host:        "127.0.0.1",
		Port:        8848,
		NamespaceId: "",
		Timeout:     5000,
	}

	for _, opt := range opts {
		opt(options)
	}

	optss := nacos.NewNacosConfig(options.ApplicationName, options.ApplicationListenOn,
		BuildServerConfigs(options), BuildClientConfig(options))

	info := common.GetVersion()
	optss.Metadata["AppName"] = options.ApplicationName
	optss.Metadata["Version"] = info.Version
	optss.Metadata["Branch"] = info.GitBranch
	optss.Metadata["BuildDate"] = info.BuildDate
	optss.Metadata["GoVersion"] = info.GoVersion
	optss.Metadata["OS/Arch"] = info.Platform

	err := nacos.RegisterService(optss)
	return err
}
