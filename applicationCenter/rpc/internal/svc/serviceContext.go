package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"qbq-open-platform/applicationCenter/rpc/internal/config"
	"qbq-open-platform/common/paas"
	"qbq-open-platform/common/utils"
)

type ServiceContext struct {
	Server       *config.Server
	RedisClient  *redis.Redis
	DbEngin      *gorm.DB
	PaasPlatform *paas.Paas
	Snowflake    *snowflake.Node
}

func NewServiceContext(config *config.Bootstrap) *ServiceContext {
	//启动Gorm支持
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Server.Mysql.DataSource,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 表名前缀，`User`表为`t_users`
			SingularTable: false, // 使用单数表名，启用该选项后，`User` 表将是`application`
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	paasPlatform := paas.New(func(paas *paas.Paas) {
		paas.ClientId = config.Server.Paas.ClientId
		paas.ClientSecret = config.Server.Paas.ClientSecret
		paas.PaasUrl = config.Server.Paas.PaasUrl
		paas.Retry = config.Server.Paas.Retry
	})

	redisClient := redis.New(config.Server.Redis.Host, func(r *redis.Redis) {
		r.Pass = config.Server.Redis.Pass
		r.Type = config.Server.Redis.Type
	})

	node, err := utils.GenerateSnowflakeNode(redisClient, config.Server.Application.ServiceConf.Name)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Server:       config.Server,
		RedisClient:  redisClient,
		DbEngin:      db,
		PaasPlatform: paasPlatform,
		Snowflake:    node,
	}
}
