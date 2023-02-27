package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"qbq-open-platform/applicationCenter/api/internal/config"
	apiClient "qbq-open-platform/applicationCenter/rpc/client/applicationapi"
	customerClient "qbq-open-platform/applicationCenter/rpc/client/applicationcustomer"
	"qbq-open-platform/common/interceptor/unary"
)

type ServiceContext struct {
	Server                       *config.Server
	RedisClient                  *redis.Redis
	ApiApplicationRpcClient      apiClient.ApplicationApi
	CustomerApplicationRpcClient customerClient.ApplicationCustomer
}

func NewServiceContext(config *config.Bootstrap) *ServiceContext {
	appRpcTarget := fmt.Sprintf("nacos://%s:%d/%s?namespaceid=%s&timeout=5000s",
		config.Nacos.Discovery.Host, config.Nacos.Discovery.Port, config.Server.Rpc.App[0].Name, config.Nacos.Discovery.NamespaceId)
	appRpcConf := zrpc.RpcClientConf{
		Target:   appRpcTarget,
		Timeout:  config.Server.Rpc.App[0].Timeout,
		NonBlock: true,
		Middlewares: zrpc.ClientMiddlewaresConf{
			Trace:      true,
			Duration:   true,
			Prometheus: false,
			Breaker:    false,
			Timeout:    true,
		},
	}

	return &ServiceContext{
		Server: config.Server,
		RedisClient: redis.New(config.Server.Redis.Host, func(r *redis.Redis) {
			r.Pass = config.Server.Redis.Pass
			r.Type = config.Server.Redis.Type
		}),
		ApiApplicationRpcClient:      apiClient.NewApplicationApi(zrpc.MustNewClient(appRpcConf, zrpc.WithUnaryClientInterceptor(unary.ClientTokenInterceptor))),
		CustomerApplicationRpcClient: customerClient.NewApplicationCustomer(zrpc.MustNewClient(appRpcConf, zrpc.WithUnaryClientInterceptor(unary.ClientTokenInterceptor))),
	}
}
