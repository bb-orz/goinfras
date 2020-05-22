package redisPubSub

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
)

/*
发布消息
@param channel string 发布频道
@param msg interface{} 发布的消息
*/
func Publish(channel string, msg interface{}) error {
	conn := GetRedisConn()
	defer conn.Close()

	receiveNum, err := redigo.Int(conn.Do("PUBLISH", channel, redigo.Args{}.AddFlat(msg)))
	logger.InfoLog("Redis Publish Topic Message", fmt.Sprintf("[channel：%s],[message:%v],[receiveNum:%d]", channel, msg, receiveNum))
	if err != nil {
		return err
	}

	if receiveNum == 0 {
		// 订阅并接收到该channel的数量为receiveNum
		logger.WarmLog(fmt.Sprintf("No subscriber subscribe or receive %s channel  ", channel))
	}

	return nil
}
