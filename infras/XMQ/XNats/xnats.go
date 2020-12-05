package XNats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var natsMQPool *NatsPool

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
func XCommonNatsSubscribe() *commonNatsSubscribe {
	c := new(commonNatsSubscribe)
	c.pool = XPool()
	return c
}

// 基于请求响应方式的通用方法实例
func XCommonNatsReqResp() *commonNatsReqResp {
	c := new(commonNatsReqResp)
	c.pool = XPool()
	return c
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			Switch: true,
			NatsServers: []natsServer{
				{
					"127.0.0.1",
					4222,
					false,
					"",
					"",
				},
			},
		}

	}

	natsMQPool, err = NewPool(config, zap.L())
	return err
}
