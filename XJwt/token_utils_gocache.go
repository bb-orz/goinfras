package XJwt

import (
	"errors"
	"github.com/bb-orz/goinfras/XCache/XGocache"
	"github.com/pmylund/go-cache"
	"time"
)

// 创建一个默认配置的带redis缓存的TokenUtils
func CreateDefaultTkuWithGoCache(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}

	// 检查redis连接池组件或创建默认池
	if !XGocache.CheckGocache() {
		XGocache.CreateDefaultCache(nil)
	}

	tku = NewTokenUtilsWithGoCache(config)
}

type tokenUtilsGoCache struct {
	tokenUtils
	goCache   *cache.Cache // 使用go-cache 缓存
	keyPrefix string       // 键名前缀
}

func NewTokenUtilsWithGoCache(config *Config) *tokenUtilsGoCache {
	ts := new(tokenUtilsGoCache)
	ts.privateKey = []byte(config.PrivateKey)
	ts.expTime = time.Now().Add(time.Second * time.Duration(config.ExpSeconds))
	ts.goCache = XGocache.X()
	ts.keyPrefix = config.TokenCacheKeyPrefix
	return ts
}

func (tks *tokenUtilsGoCache) Encode(user UserClaim) (string, error) {
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
	tks.goCache.Set(key, token, exp)

	return token, nil
}

func (tks *tokenUtilsGoCache) Decode(tokenString string) (*CustomerClaim, error) {
	// 如不能解码，直接返回err
	claim, err := tks.decode(tokenString)
	if err != nil {
		return claim, err
	}
	key := tks.keyPrefix + claim.UserClaim.Id
	var val interface{}
	var b bool
	if val, b = tks.goCache.Get(key); b {
		// redis 鉴定缓存数据
		cacheToken := val.(string)
		if cacheToken != tokenString {
			return nil, errors.New("Token string is invalid with cache data ")
		}
	} else {
		return nil, errors.New("Token key cache is not exist! ")
	}

	return claim, nil
}
