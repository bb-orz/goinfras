package verified

import (
	"GoWebScaffold/infras/store/redisStore"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type verifiedCache struct {
	commonRedis *redisStore.CommonRedisDao
}

func NewMailCache() *verifiedCache {
	cache := new(verifiedCache)
	cache.commonRedis = redisStore.NewCommonRedisDao()
	return cache
}

// 保存邮箱验证码缓存
func (cache *verifiedCache) SetForgetPasswordVerifiedCode(uid uint, code string) error {
	var err error
	var key string
	key = UserCacheForgetPasswordVerifiedCodePrefix + strconv.Itoa(int(uid))
	_, err = cache.commonRedis.R("SETEX", key, UserCacheForgetPasswordVerifiedCodeExpire, code)
	if err != nil {
		return err
	}

	return nil
}

// 获取邮箱验证码缓存
func (cache *verifiedCache) GetForgetPasswordVerifiedCode(uid uint) (string, error) {
	var err error
	var key string
	var code string
	key = UserCacheForgetPasswordVerifiedCodePrefix + strconv.Itoa(int(uid))
	code, err = redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}

	return code, nil
}

// 保存邮箱验证码缓存
func (cache *verifiedCache) SetUserVerifiedEmailCode(uid uint, code string) error {
	var err error
	var key string
	key = UserCacheVerifiedEmailCodePrefix + strconv.Itoa(int(uid))
	_, err = cache.commonRedis.R("SETEX", key, UserCacheVerifiedEmailCodeExpire, code)
	if err != nil {
		return err
	}

	return nil
}

// 获取邮箱验证码缓存
func (cache *verifiedCache) GetUserVerifiedEmailCode(uid uint) (string, error) {
	var err error
	var key string
	var code string

	key = UserCacheVerifiedEmailCodePrefix + strconv.Itoa(int(uid))
	code, err = redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}

	return code, nil
}

// 保存手机验证码缓存
func (cache *verifiedCache) SetUserVerifiedPhoneCode(uid uint, code string) error {
	var err error
	var key string

	key = UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(int(uid))
	_, err = cache.commonRedis.R("SETEX", key, UserCacheVerifiedPhoneCodeExpire, code)
	if err != nil {
		return err
	}

	return nil
}

// 获取手机验证码缓存
func (cache *verifiedCache) GetUserVerifiedPhoneCode(uid uint) (string, error) {
	var err error
	var key string
	var code string

	key = UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(int(uid))
	code, err = redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}

	return code, nil
}
