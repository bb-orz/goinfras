package XEtcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
)

var client *clientv3.Client

// 创建一个默认配置的Manager
func CreateDefaultClient(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	client, err = NewEtcdClient(context.TODO(), config, nil)
	return err
}

// 资源组件实例调用
func XClient() *clientv3.Client {
	return client
}

// 资源组件闭包执行
func XFClient(f func(c *clientv3.Client) error) error {
	return f(client)
}