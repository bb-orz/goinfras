package XCache

import "time"

// 一个通用缓存实例，可择机使用redis或go-cache缓存
var commonCache ICommonCache

func SettingCommonCache(cc ICommonCache) {
	commonCache = cc
}

func XCommon() ICommonCache {
	return commonCache
}

func CheckXCommon() bool {
	if commonCache != nil {
		return true
	}
	return false
}

type ICommonCache interface {
	Add(k string, v interface{}) error     // 添加一个不存在或已超时的键值
	Get(k string) (interface{}, bool)      // 获取一个键值
	Set(k string, v interface{}) error     // 更新或添加一个键值，无论是否已存在
	Replace(k string, v interface{}) error // 更新一个已存在且未过期的键值，不满足条件则报错

	AddWithExp(k string, v interface{}, exp time.Duration) error     // 添加一个不存在或已超时的键值，带超时,单位s
	GetWithExp(k string) (interface{}, time.Time, bool)              // 获取一个带过期时间的键值
	SetWithExp(k string, v interface{}, exp time.Duration) error     // 更新或添加一个键值，无论是否已存在，带超时,单位s
	ReplaceWithExp(k string, v interface{}, exp time.Duration) error // 更新一个已存在且未过期的键值，不满足条件则报错，带超时,单位s
	Increment(k string, v int64) error                               // 自增int64
	Decrement(k string, v int64) error                               // 自减int64
	Delete(k string) bool                                            // 删除键值
}
