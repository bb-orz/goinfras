package XJwt

import (
	"errors"
	"fmt"
	"github.com/bb-orz/goinfras/XCache"
	"github.com/bb-orz/goinfras/XCache/XGocache"
	"strconv"
	"time"
)

// 创建一个默认配置的带redis缓存的TokenUtils
func CreateDefaultTkuWithCache(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}

	// 检查redis连接池组件或创建默认池
	if !XGocache.CheckGocache() {
		XGocache.CreateDefaultCache(nil)
	}

	tku = NewTokenUtilsWithCache(config)
}

type tokenUtilsCache struct {
	tokenUtils
	CommonCache XCache.ICommonCache // 使用CommonCache 缓存
	keyPrefix   string              // 键名前缀
}

func NewTokenUtilsWithCache(config *Config) *tokenUtilsCache {
	ts := new(tokenUtilsCache)
	ts.privateKey = []byte(config.PrivateKey)
	ts.expTime = time.Second * time.Duration(config.ExpSeconds)
	ts.CommonCache = XCache.XCommon()
	ts.keyPrefix = config.TokenCacheKeyPrefix
	return ts
}

func (tks *tokenUtilsCache) Encode(user UserClaim) (string, error) {
	var err error
	// 编码
	token, err := tks.encode(user)
	if err != nil {
		return "", err
	}

	// 将token缓存到redis
	if user.Id <= 0 {
		return "", errors.New("Empty UserId is not allowed! ")
	}

	key := tks.keyPrefix + strconv.Itoa(int(user.Id))
	err = tks.CommonCache.SetWithExp(key, token, tks.expTime)

	return token, nil
}

func (tks *tokenUtilsCache) Decode(tokenString string) (*CustomerClaim, error) {
	// 如不能解码，直接返回err
	claim, err := tks.decode(tokenString)
	if err != nil {
		return claim, err
	}
	key := tks.keyPrefix + strconv.Itoa(int(claim.UserClaim.Id))
	var val interface{}
	var b bool
	if val, b = tks.CommonCache.Get(key); b {
		// 鉴定缓存数据
		cacheToken := fmt.Sprintf("%s", val)
		if cacheToken != tokenString {
			return nil, errors.New("Token string is invalid with cache data ")
		}
	} else {
		return nil, errors.New("Token key cache is not exist! ")
	}

	return claim, nil
}

func (tks *tokenUtilsCache) Remove(tokenString string) error {
	// 如不能解码，直接返回err
	claim, err := tks.decode(tokenString)
	if err != nil {
		return err
	}
	key := tks.keyPrefix + strconv.Itoa(int(claim.UserClaim.Id))

	b := tks.CommonCache.Delete(key)
	if !b {
		return errors.New("Remove Token Cache Fail ")
	}
	return nil
}
