package XJwt

import (
	"GoWebScaffold/infras/XStore/XRedis"
	"github.com/garyburd/redigo/redis"
)

type redisCache struct {
	dao *XRedis.CommonRedisDao // 服务端缓存保存用于校验
}

func NewRedisCache() *redisCache {
	cache := new(redisCache)
	cache.dao = XRedis.XCommon()
	return cache
}

const TokenCacheKeyPrefix = "user.jwt.token."

// 设置Token到redis
func (cache *redisCache) SetToken(id, token string) error {
	key := TokenCacheKeyPrefix + id
	_, err := cache.dao.R("SET", key, token, "EX", 3600)
	if err != nil {
		return err
	}

	return nil

}

// 从redis获取token
func (cache *redisCache) GetToken(id string) (string, error) {
	key := TokenCacheKeyPrefix + id

	rs, err := redis.String(cache.dao.R("GET", key))
	if err != nil {
		return "", err
	}

	return rs, nil
}

// 从redis删除token
func (cache *redisCache) DeleteToken(id string) error {
	key := TokenCacheKeyPrefix + id

	_, err := cache.dao.R("DEL", key)
	if err != nil {
		return err
	}
	return nil
}
