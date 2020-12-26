package XGocache

import "time"

type CommonGocache struct{}

func NewCommonGocache() *CommonGocache {
	return new(CommonGocache)
}

// 添加一个不存在或已超时的键值
func (*CommonGocache) Add(k string, v interface{}) error {
	return X().Add(k, v, 0)
}

// 获取一个键值
func (*CommonGocache) Get(k string) (interface{}, bool) {
	return X().Get(k)
}

// 更新或添加一个键值，无论是否已存在
func (*CommonGocache) Set(k string, v interface{}) error {
	X().Set(k, v, 0)
	return nil
}

// 更新一个已存在且未过期的键值，不满足条件则报错
func (*CommonGocache) Replace(k string, v interface{}) error {
	return X().Replace(k, v, 0)
}

// 添加一个不存在或已超时的键值，带超时,单位s
func (*CommonGocache) AddWithExp(k string, v interface{}, exp time.Duration) error {
	return X().Add(k, v, time.Duration(exp)*time.Second)
}

// 获取一个带过期时间的键值
func (*CommonGocache) GetWithExp(k string) (interface{}, time.Time, bool) {
	return X().GetWithExpiration(k)
}

// 更新或添加一个键值，无论是否已存在，带超时,单位s
func (*CommonGocache) SetWithExp(k string, v interface{}, exp time.Duration) error {
	X().Set(k, v, time.Duration(exp)*time.Second)
	return nil
}

// 更新一个已存在且未过期的键值，不满足条件则报错，带超时,单位s
func (*CommonGocache) ReplaceWithExp(k string, v interface{}, exp time.Duration) error {
	return X().Replace(k, v, time.Duration(exp)*time.Second)
}

// 自增int64
func (*CommonGocache) Increment(k string, v int64) error {
	return X().Increment(k, v)
}

// 自减int64
func (*CommonGocache) Decrement(k string, v int64) error {
	return X().Decrement(k, v)
}

// 删除键值
func (*CommonGocache) Delete(k string) bool {
	X().Delete(k)
	return true
}
