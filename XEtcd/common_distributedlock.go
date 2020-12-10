package XEtcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
)

/*实现分布式锁*/
type EtcdCommonDistributedLock struct {
	client *clientv3.Client
}

func NewEtcdCommonDistributedLock() *EtcdCommonDistributedLock {
	common := new(EtcdCommonDistributedLock)
	common.client = XClient()
	return common
}

func (c *EtcdCommonDistributedLock) DistributedLock() {
	var err error
	if err != nil {
		log.Fatal(err)
	}

	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(c.client)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(c.client)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
}
