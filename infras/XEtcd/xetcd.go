package XEtcd

import (
	"GoWebScaffold/infras"
	"go.etcd.io/etcd/clientv3"
)

// 资源组件实例调用
func X() (*clientv3.Client, error) {
	err := infras.Check(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// 资源组件闭包执行
func XF(f func(c *clientv3.Client) error) error {
	return f(client)
}
