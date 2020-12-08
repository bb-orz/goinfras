package XJwt

import (
	"github.com/garyburd/redigo/redis"
	"goinfras/XStore/XRedis"
)

type redisCache struct {
	cacheDao *XRedis.CommonRedisDao // 服务端缓存保存用于校验
}

func NewRedisCache() *redisCache {
	cache := new(redisCache)
	// 使用XRedis组件的通用操作实例
	cache.cacheDao = XRedis.XCommon()
	return cache
}

const TokenCacheKeyPrefix = "user.jwt.token."

// 设置Token到redis
func (cache *redisCache) SetToken(id, token string, exp int) error {
	key := TokenCacheKeyPrefix + id
	_, err := cache.cacheDao.R("SET", key, token, "EX", exp)
	if err != nil {
		return err
	}

	return nil

}

// 从redis获取token
func (cache *redisCache) GetToken(id string) (string, error) {
	key := TokenCacheKeyPrefix + id

	rs, err := redis.String(cache.cacheDao.R("GET", key))
	if err != nil {
		return "", err
	}

	return rs, nil
}

// 从redis删除token
func (cache *redisCache) DeleteToken(id string) error {
	key := TokenCacheKeyPrefix + id

	_, err := cache.cacheDao.R("DEL", key)
	if err != nil {
		return err
	}
	return nil
}