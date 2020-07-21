package user

import "GoWebScaffold/infras/store/redisStore"

type oauthCache struct {
	commonRedis *redisStore.CommonRedisDao
}

func NewOauthCache() *oauthCache {
	cache := new(oauthCache)
	cache.commonRedis = redisStore.NewCommonRedisDao()
	return cache
}
