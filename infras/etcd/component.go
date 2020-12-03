package etcd

import (
	"GoWebScaffold/infras"
	"go.etcd.io/etcd/clientv3"
)

/*资源组件化调用*/

var client *clientv3.Client

// 设置组件资源
func SetComponent(c *clientv3.Client) {
	client = c
}

// 组件化使用
func EtcdComponent() *clientv3.Client {
	_ = infras.Check(client)
	return client
}
