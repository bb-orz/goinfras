package XRedis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

type CommonRedisCache struct{}

func NewCommonRedisCache() *CommonRedisCache {
	return new(CommonRedisCache)
}

// 添加一个不存在或已超时的键值
func (*CommonRedisCache) Add(k string, v interface{}) error {
	isOK, err := redis.String(XCommand().R("SET", k, v, "NX"))
	if err != nil {
		return err
	}

	if isOK != "OK" {
		return errors.New("ADD FAIL! ")
	}
	return nil
}

// 获取一个键值
func (*CommonRedisCache) Get(k string) (interface{}, bool) {
	reply, err := XCommand().R("GET", k)
	if err != nil {
		return nil, false
	}
	return reply, true
}

// 获取一个字符键值
func (*CommonRedisCache) GetString(k string) (string, error) {
	return redis.String(XCommand().R("GET", k))
}

// 获取一个整型键值
func (*CommonRedisCache) GetInt(k string) (int, error) {
	return redis.Int(XCommand().R("GET", k))
}

// 获取一个布尔键值
func (*CommonRedisCache) GetBool(k string) (bool, error) {
	return redis.Bool(XCommand().R("GET", k))
}

// 获取一个浮点数键值
func (*CommonRedisCache) GetFloat64(k string) (float64, error) {
	return redis.Float64(XCommand().R("GET", k))
}

// 更新或添加一个键值，无论是否已存在
func (*CommonRedisCache) Set(k string, v interface{}) error {
	isOK, err := redis.String(XCommand().R("SET", k, v))
	if err != nil {
		return err
	}

	if isOK != "OK" {
		return errors.New("SET FAIL! ")
	}
	return nil
}

// 更新一个已存在且未过期的键值，不满足条件则报错
func (*CommonRedisCache) Replace(k string, v interface{}) error {
	isOK, err := redis.String(XCommand().R("SET", k, v, "XX"))
	if err != nil {
		return err
	}

	if isOK != "OK" {
		return errors.New("Replace FAIL! ")
	}
	return nil
}

// 添加一个不存在或已超时的键值，带超时
func (*CommonRedisCache) AddWithExp(k string, v interface{}, exp time.Duration) error {
	isOK, err := redis.String(XCommand().R("SET", k, v, "EX", int64(exp), "NX"))
	if err != nil {
		return err
	}

	if isOK != "OK" {
		return errors.New("ADD FAIL! ")
	}
	return nil
}

// 获取一个带过期时间的键值
func (*CommonRedisCache) GetWithExp(k string) (interface{}, time.Time, bool) {
	reply, err := XCommand().R("GET", k)
	if err != nil {
		return nil, time.Time{}, false
	}
	expSecond, err := redis.Int64(XCommand().R("TTL", k))
	if err != nil || expSecond < 0 {
		return reply, time.Time{}, false
	}
	expTime := time.Now().Add(time.Duration(expSecond))
	return reply, expTime, true
}

// 获取一个字符键值
func (*CommonRedisCache) GetStringWithExp(k string) (string, int64, error) {
	v, err := redis.String(XCommand().R("GET", k))
	if err != nil {
		return v, -1, err
	}
	expSecond, err := redis.Int64(XCommand().R("TTL", k))
	if err != nil || expSecond < 0 {
		return v, expSecond, err
	}
	return v, expSecond, nil
}

// 获取一个整型键值
func (*CommonRedisCache) GetIntWithExp(k string) (int, int64, error) {
	v, err := redis.Int(XCommand().R("GET", k))
	if err != nil {
		return v, -1, err
	}
	expSecond, err := redis.Int64(XCommand().R("TTL", k))
	if err != nil || expSecond < 0 {
		return v, expSecond, err
	}
	return v, expSecond, nil
}

// 获取一个布尔键值
func (*CommonRedisCache) GetBoolWithExp(k string) (bool, int64, error) {
	v, err := redis.Bool(XCommand().R("GET", k))
	if err != nil {
		return v, -1, err
	}
	expSecond, err := redis.Int64(XCommand().R("TTL", k))
	if err != nil || expSecond < 0 {
		return v, expSecond, err
	}
	return v, expSecond, nil
}

// 获取一个浮点数键值
func (*CommonRedisCache) GetFloat64WithExp(k string) (float64, int64, error) {
	v, err := redis.Float64(XCommand().R("GET", k))
	if err != nil {
		return v, -1, err
	}
	expSecond, err := redis.Int64(XCommand().R("TTL", k))
	if err != nil || expSecond < 0 {
		return v, expSecond, err
	}
	return v, expSecond, nil
}

// 更新或添加一个键值，无论是否已存在，带超时
func (*CommonRedisCache) SetWithExp(k string, v interface{}, exp time.Duration) error {
	isOK, err := redis.String(XCommand().R("SET", k, v, "EX", int64(exp)))
	if err != nil {
		return err
	}

	if isOK != "OK" {
		return errors.New("SetWithExp FAIL! ")
	}
	return nil
}

// 更新一个已存在且未过期的键值，不满足条件则报错，带超时
func (*CommonRedisCache) ReplaceWithExp(k string, v interface{}, exp time.Duration) error {
	isOK, err := redis.String(XCommand().R("SET", k, v, "EX", int64(exp), "XX"))
	if err != nil {
		return err
	}

	if isOK != "OK" {
		return errors.New("Replace FAIL! ")
	}
	return nil
}

// 自增int64
func (*CommonRedisCache) Increment(k string, v int64) error {
	_, err := redis.Int64(XCommand().R("INCRBY", k, v))
	return err
}

// 自减int64
func (*CommonRedisCache) Decrement(k string, v int64) error {
	_, err := redis.Int64(XCommand().R("DECRBY", k, v))
	return err
}

// 删除键值
func (*CommonRedisCache) Delete(k string) bool {
	delCount, err := redis.Int64(XCommand().R("DEL", k))
	if err != nil {
		return false
	}
	if delCount > 0 {
		return true
	}
	return false
}
