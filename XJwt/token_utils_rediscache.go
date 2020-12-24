package XJwt

import (
	"errors"
	"github.com/bb-orz/goinfras/XCache/XRedis"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 创建一个默认配置的带redis缓存的TokenUtils
func CreateDefaultTkuWithRedisCache(config *Config) error {
	if config == nil {
		config = DefaultConfig()
	}

	// 检查redis连接池组件或创建默认池
	if !XRedis.CheckPool() {
		var err error
		err = XRedis.CreateDefaultPool(nil)
		if err != nil {
			return err
		}
	}

	tku = NewTokenUtilsWithRedisCache(config)

	return nil
}

type tokenUtilsRedisCache struct {
	tokenUtils
	redisDao  *XRedis.CommonRedisDao // 使用redis缓存
	keyPrefix string                 // 键名前缀
}

func NewTokenUtilsWithRedisCache(config *Config) *tokenUtilsRedisCache {
	ts := new(tokenUtilsRedisCache)
	ts.privateKey = []byte(config.PrivateKey)
	ts.expTime = time.Now().Add(time.Second * time.Duration(config.ExpSeconds))
	ts.redisDao = XRedis.XCommon()
	ts.keyPrefix = config.TokenCacheKeyPrefix
	return ts
}

func (tks *tokenUtilsRedisCache) Encode(user UserClaim) (string, error) {
	var err error
	// 编码
	token, err := tks.encode(user)
	if err != nil {
		return "", err
	}

	// 将token缓存到redis
	if user.Id == "" {
		return "", errors.New("Empty UserId is not allowed! ")
	}
	exp := tks.expTime.Sub(time.Now())
	key := tks.keyPrefix + user.Id
	_, err = tks.redisDao.R("SET", key, token, "EX", exp)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (tks *tokenUtilsRedisCache) Decode(tokenString string) (*CustomerClaim, error) {
	// 如不能解码，直接返回err
	claim, err := tks.decode(tokenString)
	if err != nil {
		return claim, err
	}
	key := tks.keyPrefix + claim.UserClaim.Id
	cacheToken, err := redis.String(tks.redisDao.R("GET", key))

	// redis 鉴定缓存数据
	if cacheToken != tokenString {
		return nil, errors.New("Token string is invalid with cache data ")
	}
	return claim, nil
}
