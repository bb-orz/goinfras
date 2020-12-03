package redisPubSub

import (
	"GoWebScaffold/infras/logger"
	redigo "github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

type RedisPublisher struct{}

func NewRedisPublisher() *RedisPublisher {
	return new(RedisPublisher)
}

/*
发布消息
@param channel string 发布频道
@param msg interface{} 发布的消息
*/
func (*RedisPublisher) Publish(channel string, msg interface{}) error {
	conn := getRedisConn()
	defer conn.Close()

	receiveNum, err := redigo.Int(conn.Do("PUBLISH", channel, redigo.Args{}.AddFlat(msg)))
	XLogger.CLogger().Info("Redis Publish Topic Message:", zap.String("channel", channel), zap.String("message", msg.(string)), zap.Int("receive count", receiveNum))
	if err != nil {
		return err
	}

	if receiveNum == 0 {
		// 订阅并接收到该channel的数量为receiveNum
		XLogger.CLogger().Warn("No subscriber subscribe or receive this channel", zap.String("channel", channel))
	}

	return nil
}
