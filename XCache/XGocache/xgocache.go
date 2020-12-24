package XGocache

import (
	"github.com/pmylund/go-cache"
)

func X() *cache.Cache {
	return goCache
}

// 资源组件闭包执行
func XF(f func(c *cache.Cache) error) error {
	return f(goCache)
}
