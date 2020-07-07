package mail

import (
	"GoWebScaffold/infras/store/redisStore"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type mailCache struct {
	commonRedis *redisStore.CommonRedisDao
}

func NewMailCache() *mailCache {
	cache := new(mailCache)
	cache.commonRedis = redisStore.NewCommonRedisDao()
	return cache
}

// 保存邮箱验证码缓存
func (cache *mailCache) SetUserVerifiedEmailCode(uid int, code string) error {
	key := UserCacheVerifiedEmailCodePrefix + strconv.Itoa(uid)
	_, err := cache.commonRedis.R("SETEX", key, UserCacheVerifiedEmailCodeExpire, code)
	if err != nil {
		return err
	}

	return nil
}

// 获取邮箱验证码缓存
func (cache *mailCache) GetUserVerifiedEmailCode(uid int) (string, error) {
	key := UserCacheVerifiedEmailCodePrefix + strconv.Itoa(uid)
	code, err := redis.String(cache.commonRedis.R("GET", key))
	if err != nil {
		return "", err
	}

	return code, nil
}
