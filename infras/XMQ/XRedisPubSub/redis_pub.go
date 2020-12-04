package XRedisPubSub

import (
	"GoWebScaffold/infras/XLogger"
	redigo "github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

type redisPublisher struct {
	pool *redigo.Pool
}

/*
发布消息
@param channel string 发布频道
@param msg interface{} 发布的消息
*/
func (c *redisPublisher) Publish(channel string, msg interface{}) error {
	conn := c.pool.Get()
	defer func() {
		conn.Close()
	}()

	receiveNum, err := redigo.Int(conn.Do("PUBLISH", channel, redigo.Args{}.AddFlat(msg)))
	XLogger.XCommon().Info("Redis Publish Topic Message:", zap.String("channel", channel), zap.String("message", msg.(string)), zap.Int("receive count", receiveNum))
	if err != nil {
		return err
	}

	if receiveNum == 0 {
		// 订阅并接收到该channel的数量为receiveNum
		XLogger.XCommon().Warn("No subscriber subscribe or receive this channel", zap.String("channel", channel))
	}

	return nil
}
