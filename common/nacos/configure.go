package nacos

import (
	"context"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"log"
	"qbq-open-platform/common/utils"
)

var (
	ChangeContent    chan string
	ChangeContentCtx context.Context
	cancel           context.CancelFunc
)

// 连接配置中心（nacos）
func NacosConfigure(opts ...Option) (string, error) {
	ChangeContentCtx, cancel = context.WithCancel(context.Background())

	ChangeContent = make(chan string)
	options := &Options{
		Host:        "127.0.0.1",
		Port:        8848,
		NamespaceId: "",
		Timeout:     5000,
	}

	for _, opt := range opts {
		opt(options)
	}

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  BuildClientConfig(options),
			ServerConfigs: BuildServerConfigs(options),
		},
	)
	if err != nil {
		log.Panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: options.DataId,
		Group:  options.Group,
		Type:   vo.YAML,
	})
	content = utils.ParseYamlEnv(content)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: options.DataId,
		Group:  options.Group,
		Type:   vo.YAML,
		OnChange: func(namespace, group, dataId, data string) {
			ChangeContent <- data
		},
	})

	proc.AddShutdownListener(func() {
		cancel()
		err = configClient.CancelListenConfig(vo.ConfigParam{
			DataId: options.DataId,
			Group:  options.Group,
			Type:   vo.YAML,
		})
		if err != nil {
			logx.Info("cancel listenConfig error: ", err.Error())
		} else {
			logx.Info("cancel listenConfig from nacos server.")
		}
	})

	return content, nil
}
