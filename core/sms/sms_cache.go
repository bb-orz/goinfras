package sms

import (
	"GoWebScaffold/infras/store/redisStore"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type smsCache struct {
	commonRedis *redisStore.CommonRedisDao
}

func NewSmsCache() *smsCache {
	cache := new(smsCache)
	cache.commonRedis = redisStore.NewCommonRedisDao()
	return cache
}

// 保存手机验证码缓存
func (cache *smsCache) SetUserVerifiedPhoneCode(uid int, code string) error {
	key := UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(uid)
	_, err := cache.commonRedis.R("SETEX", key, UserCacheVerifiedPhoneCodeExpire, code)
	if err != nil {
		return err
	}

	return nil
}

// 获取手机验证码缓存
func (cache *smsCache) GetUserVerifiedPhoneCode(uid int) (string, error) {
	key := UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(uid)
	code, err := redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}

	return code, nil
}
