package etcd

import (
	"go.etcd.io/etcd/clientv3"
)

type EtcdCommonClient struct {
	client *clientv3.Client
}

func NewEtcdCommonClient() *EtcdCommonClient {
	c := new(EtcdCommonClient)
	c.client = EtcdClientV3()
	return c
}

// TODO 编写常用的etcd操作
