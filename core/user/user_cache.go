package user

import (
	"GoWebScaffold/infras/store/redisStore"
)

type userCache struct {
	commonRedis *redisStore.CommonRedisDao
}

func NewUserCache() *userCache {
	cache := new(userCache)
	cache.commonRedis = redisStore.NewCommonRedisDao()
	return cache
}
