package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
)

type CommonEtcd struct {
	client *clientv3.Client
	ctx    context.Context
}

func NewCommonEtcd() *CommonEtcd {
	c := new(CommonEtcd)
	c.client = EtcdClientV3()
	c.ctx = context.TODO()
	return c
}

// TODO 编写常用的etcd操作
