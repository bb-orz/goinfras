package user

import (
	"GoWebScaffold/infras/store/redisStore"
	"github.com/garyburd/redigo/redis"
)

type userCache struct {
	commonRedis *redisStore.CommonRedisDao
}

func NewUserCache() *userCache {
	cache := new(userCache)
	cache.commonRedis = redisStore.NewCommonRedisDao()
	return cache
}

// 设置鉴权令牌缓存
func (cache *userCache) SetUserToken(userNo, token string) error {
	key := UserCacheTokenPrefix + userNo
	_, err := cache.commonRedis.R("SETEX", key, UserCacheTokenExpire, token)
	if err != nil {
		return err
	}
	return nil
}

// 获取鉴权令牌缓存
func (cache *userCache) GetUserToken(userNo string) (string, error) {
	key := UserCacheTokenPrefix + userNo
	code, err := redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}
	return code, nil
}
