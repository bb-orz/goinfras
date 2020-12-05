package XEtcd

import (
	"context"
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

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			Endpoints: []string{"localhost:2379"},
		}
	}
	client, err = NewEtcdClient(context.TODO(), config, nil)
	return err
}
