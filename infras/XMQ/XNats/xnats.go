package XNats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var natsMQPool *NatsPool

func CreateDefaultPool(config *Config, logger *zap.Logger) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}

	natsMQPool, err = NewPool(config, logger)
	return err
}

// 资源组件实例调用
func XPool() *NatsPool {
	return natsMQPool
}

// 资源组件闭包执行
func XF(f func(c *nats.Conn) error) error {
	var err error
	var conn *nats.Conn
	// 获取连接
	conn, err = natsMQPool.Get()
	if err != nil {
		return err
	}

	// 放回连接池
	defer func() {
		natsMQPool.Put(conn)
	}()

	// 执行用户操作
	err = f(conn)
	if err != nil {
		return err
	}

	return nil
}

// 通用管道方法实例
func XCommonNatsChan() *commonNatsChan {
	c := new(commonNatsChan)
	c.pool = XPool()
	return c
}

// 通用发布订阅方法实例
func XCommonNatsPubSub() *commonNatsPubSub {
	c := new(commonNatsPubSub)
	c.pool = XPool()
	return c
}

// 基于队列组的主题订阅方法实例
func XCommonNatsQueue() *commonNatsQueue {
	c := new(commonNatsQueue)
	c.pool = XPool()
	return c
}

// 基于请求响应方式的通用方法实例
func XCommonNatsReqResp() *commonNatsReqResp {
	c := new(commonNatsReqResp)
	c.pool = XPool()
	return c
}
