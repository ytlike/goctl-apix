package global

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

// Config 全局配置对象
var globalConfig = &config{}

const (
	PAAS_CLIENT_TOKEN          = "qbq-open-platform:paas:client:token"        // paas平台client模式的key
	PAAS_EXCHANGE_TOKEN        = "qbq-open-platform:paas:exchange:token"      // paas平台exchange模式的key
	EMAIL_CODE_PREFIX          = "qbq-open-platform:user-center:email:code:"  // 用户中心邮箱验证码 key
	EMAIL_CODE_EXIST_PREFIX    = "qbq-open-platform:user-center:email:exist:" // 用户中心邮箱已经存在的验证码 key
	EMAIL_CODE_EXPIRATION_TIME = 600                                          // 验证码过期时间
	EMAIL_CODE_RESEND_TIME     = 60                                           // 验证码重复发送间隔时间
	INIT_APPLICATION_MAP_KEY   = "qbq-open-platform:init:applicationMap"      // 初始化应用map集合的key
	CACHE_CODE_UTIL_VALUE      = "qbq-open-platform:code-util"                //
	APPLICATION_OPEN_KEY       = "qbq-open-platform:application-center:open:" // 应用中心开通应用key
	AUTH_REDIS_PREFIX          = "qbq-open-platform:auth:"                    // redis中的认证的前缀
	CACHE_AUTHENTICATION_KEY   = "authentication"                             // 本地缓存的authorization的key
	SNOWFLAKE_NODE_PREFIX      = "qbq-open-platform:snowflake-node:"          // 雪花算法node的前缀
)

type config struct {
	ApplicationName string
	IgnoreUrls      []string
	RedisClient     *redis.Redis
	DbEngin         *gorm.DB
}

func Config() *config {
	return globalConfig
}
