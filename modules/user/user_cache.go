package user

import (
	"GoWebScaffold/infras/store/redisStore"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type userCache struct {
	commonRedis *redisStore.CommonRedisDao
}

func NewUserCache() *userCache {
	cache := new(userCache)
	cache.commonRedis = redisStore.NewCommonRedisDao()
	return cache
}

// 保存邮箱验证码缓存
func (cache *userCache) SetUserVerifiedEmailCode(uid int, code string) error {
	key := UserCacheVerifiedEmailCodePrefix + strconv.Itoa(uid)
	_, err := cache.commonRedis.R("SETEX", key, UserCacheVerifiedEmailCodeExpire, code)
	if err != nil {
		return err
	}

	return nil
}

// 获取邮箱验证码缓存
func (cache *userCache) GetUserVerifiedEmailCode(uid int) (string, error) {
	key := UserCacheVerifiedEmailCodePrefix + strconv.Itoa(uid)
	code, err := redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}

	return code, nil
}

// 保存手机验证码缓存
func (cache *userCache) SetUserVerifiedPhoneCode(uid int, code string) error {
	key := UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(uid)
	_, err := cache.commonRedis.R("SETEX", key, UserCacheVerifiedPhoneCodeExpire, code)
	if err != nil {
		return err
	}

	return nil
}

// 获取手机验证码缓存
func (cache *userCache) GetUserVerifiedPhoneCode(uid int) (string, error) {
	key := UserCacheVerifiedPhoneCodePrefix + strconv.Itoa(uid)
	code, err := redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}

	return code, nil
}
