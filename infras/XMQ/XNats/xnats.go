package XNats

import (
	"GoWebScaffold/infras"
)

var natsMQPool *NatsPool

// 资源组件实例调用
func XPool() (*NatsPool, error) {
	err := infras.Check(natsMQPool)
	if err != nil {
		return nil, err
	}

	return natsMQPool, nil
}

// 资源组件闭包执行
func XFPool(f func(c *NatsPool) error) error {
	return f(natsMQPool)
}
