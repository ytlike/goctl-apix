package utils

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

// TryAcquireCtx 获取redisLock，waitTime最大等待获取锁的时间，单位毫秒，最小1000毫秒。
func TryAcquireCtx(ctx context.Context, redisLock *redis.RedisLock, waitTime int64) (bool, error) {
	if waitTime < 1000 {
		waitTime = 1000
	}
	for {
		current := time.Now().UnixMilli()
		ok, redisLockErr := redisLock.AcquireCtx(ctx)
		if ok {
			return true, nil
		}
		time.Sleep(time.Millisecond * 100)
		currentTime := time.Now().UnixMilli()
		waitTime = waitTime - (currentTime - current)
		if waitTime <= 0 {
			return false, redisLockErr
		}
	}
}
