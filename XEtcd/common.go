package XEtcd

import (
	"context"
	"github.com/bb-orz/goinfras/XLogger"
	"go.etcd.io/etcd/clientv3"
)

type EtcdCommon struct {
	client *clientv3.Client
}

// 简单设置键值
func (c *EtcdCommon) Put(key, value string) error {
	var err error
	_, err = c.client.Put(context.Background(), key, value, func(op *clientv3.Op) {
		if !op.IsPut() {
			XLogger.XCommon().Error("Put Error ")
		}
	})
	return err
}

// 设置键值携带context
func (c *EtcdCommon) PutWithContext(ctx context.Context, key, value string) error {
	var err error
	_, err = c.client.Put(ctx, key, value, func(op *clientv3.Op) {
		if !op.IsPut() {
			XLogger.XCommon().Error("Put Error ")
		}
	})
	return err
}

// 使用租约设置键值
func (c *EtcdCommon) PutWithLease(ctx context.Context, key, value string, ttl int64) (clientv3.LeaseID, error) {
	var err error

	lease, err := c.client.Grant(ctx, ttl)
	if err != nil {
		return -1, err
	}
	_, err = c.client.Put(ctx, key, value, clientv3.WithLease(lease.ID), func(op *clientv3.Op) {
		if !op.IsPut() {
			XLogger.XCommon().Error("Put Error ")
		}
	})

	if err != nil {
		return -1, err
	}

	return lease.ID, nil
}

// 简单获取键值
func (c *EtcdCommon) Get(key string) (*clientv3.GetResponse, error) {
	return c.client.Get(context.Background(), key)
}

// 获取键值携带context
func (c *EtcdCommon) GetWithContext(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	return c.client.Get(ctx, key)
}

// 获取该键下所有的子健值
func (c *EtcdCommon) GetWithPrefix(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	return c.client.Get(ctx, key, clientv3.WithPrefix())
}

// 获取该键包含子健的个数
func (c *EtcdCommon) GetCount(key string) (*clientv3.GetResponse, error) {
	return c.client.Get(context.Background(), key, clientv3.WithCountOnly(), clientv3.WithPrefix())
}
