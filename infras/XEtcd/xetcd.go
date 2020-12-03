package XEtcd

import (
	"go.etcd.io/etcd/clientv3"
)

var client *clientv3.Client

// 资源组件实例调用
func XClient() *clientv3.Client {
	return client
}

// 资源组件闭包执行
func XFClient(f func(c *clientv3.Client) error) error {
	return f(client)
}
