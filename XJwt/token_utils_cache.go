package XJwt

import (
	"errors"
	"github.com/bb-orz/goinfras/XCache/XRedis"
	"time"
)

// 创建一个默认配置的带redis缓存的TokenUtils
func CreateDefaultTkuX(config *Config) error {
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

	tku = NewTokenUtilsX(config)

	return nil
}

type tokenUtilsX struct {
	tokenUtils
	cache *redisCache
}

func NewTokenUtilsX(config *Config) *tokenUtilsX {
	ts := new(tokenUtilsX)
	ts.privateKey = []byte(config.PrivateKey)
	ts.expTime = time.Now().Add(time.Second * time.Duration(config.ExpSeconds))
	ts.cache = NewRedisCache()
	return ts
}

func (tks *tokenUtilsX) Encode(user UserClaim) (string, error) {
	token, err := tks.encode(user)
	if err != nil {
		return "", err
	}

	// 将token缓存到redis
	if user.Id == "" {
		return "", errors.New("user id empty not allowed")
	}
	exp := tks.expTime.Sub(time.Now())
	err = tks.cache.SetToken(user.Id, token, int(exp))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (tks *tokenUtilsX) Decode(tokenString string) (*CustomerClaim, error) {
	// 如不能解码，直接返回err
	claim, err := tks.decode(tokenString)
	if err != nil {
		return claim, err
	}

	// redis 鉴定缓存数据
	cacheToken, err := tks.cache.GetToken(claim.UserClaim.Id)
	if cacheToken != tokenString {
		return nil, errors.New("Token string is invalid with cache data ")
	}
	return claim, nil
}
