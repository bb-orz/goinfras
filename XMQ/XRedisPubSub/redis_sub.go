package XRedisPubSub

import (
	"github.com/bb-orz/goinfras/XLogger"
	redigo "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

type RedisSubscriber struct {
	pool *redigo.Pool
}

// 订阅模式下的消息处理函数类型
type RecSubMsgFunc func(channel string, msg interface{}) error

// 订阅并接收消息，该方法阻塞
func (c *RedisSubscriber) Subscribe(recMsgFuncs map[string]RecSubMsgFunc, unSubChannel <-chan string) error {
	var err error
	conn := c.pool.Get()
	defer func() {

		conn.Close()
	}()

	var channels []interface{}
	for c := range recMsgFuncs {
		channels = append(channels, c)
	}

	// 订阅
	psConn := redigo.PubSubConn{Conn: conn}
	err = psConn.Subscribe(channels...)
	if err != nil {
		return err
	}

	// 开新协程接收消息
	var done = make(chan error, 1)
	go func() {
		var receiveTimes = 0
		for {
			receiveTimes++
			XLogger.XCommon().Info("receiveTimes:", zap.Int("times", receiveTimes))
			switch res := psConn.Receive().(type) {
			case redigo.Message:
				// 每接收一个已发布消息开一个协程执行消息处理函数
				go func() {
					err := recMsgFuncs[res.Channel](res.Channel, string(res.Data))
					if err != nil {
						done <- err
					}
				}()
			case redigo.Subscription:
				// 订阅与取消订阅的消息
				XLogger.XCommon().Info("redis SubReceiver:", zap.String("receive kind", res.Kind), zap.String("receive Channel", res.Channel), zap.Int("receive Count", res.Count))
				if res.Count == 0 {
					done <- nil
				}
			case error:
				// 接收到错误信息
				done <- res
			}
		}
	}()

	// 如有接收到error或检查链接断开则退出
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()
	for {
		select {
		case err := <-done:
			// 接收到错误信息
			return err
		case channelName := <-unSubChannel:
			// 取消订阅
			if _, ok := recMsgFuncs[channelName]; ok {
				if err := psConn.Unsubscribe(channelName); err != nil {
					return err
				}
			}
		case <-tick.C:
			// 定时检测连接
			if err := psConn.Ping("test conn"); err != nil {
				return err
			}
		}
	}
}
