package nacos

import "github.com/nacos-group/nacos-sdk-go/common/constant"

type Option func(n *Options)

type Options struct {
	Host        string
	Port        uint64
	NamespaceId string
	Timeout     uint64
	Group       string
	DataId      string

	ApplicationName     string
	ApplicationListenOn string
}

func WithHost(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

func WithPort(port uint64) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func WithNamespaceId(namespaceId string) Option {
	return func(o *Options) {
		o.NamespaceId = namespaceId
	}
}

func WithTimeout(timeout uint64) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func WithApplicationName(applicationName string) Option {
	return func(o *Options) {
		o.ApplicationName = applicationName
	}
}

func WithApplicationListenOn(applicationListenOn string) Option {
	return func(o *Options) {
		o.ApplicationListenOn = applicationListenOn
	}
}

func BuildServerConfigs(options *Options) []constant.ServerConfig {
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			options.Host,
			options.Port,
		),
	}
	return serverConfigs
}

func BuildClientConfig(options *Options) *constant.ClientConfig {
	clientConfig := constant.NewClientConfig(
		//当namespace是public时，此处填空字符串。
		constant.WithNamespaceId(options.NamespaceId),
		constant.WithTimeoutMs(options.Timeout),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("error"),
	)
	return clientConfig
}
