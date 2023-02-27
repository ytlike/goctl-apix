package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"qbq-open-platform/common/global"
)

func GenerateSnowflakeNode(redisClient *redis.Redis, appName string) (*snowflake.Node, error) {
	var maxNode int64 = 1023
	key := global.SNOWFLAKE_NODE_PREFIX + appName
	nodeId, err := redisClient.Incr(key)
	if err != nil {
		return nil, err
	}
	nodeId = nodeId % maxNode
	if nodeId == 0 {
		return GenerateSnowflakeNode(redisClient, appName)
	}
	// 由于历史原因，workerId不使用0，以避免和历史数据id重复
	// 1 ~ 1023
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		return nil, err
	}
	return node, nil
}
