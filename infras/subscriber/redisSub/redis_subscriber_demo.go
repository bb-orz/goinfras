package redisSub

import (
	"GoWebScaffold/infras/mq/redisMq"
	"fmt"
)

func Run() {
	// 订阅需开启redismq服务
	if !config.MqConf.RedisMq.Switch {
		return
	}

	fmt.Println("The RedisMq Subscriber Running")
	var err error

	// 整合所有的订阅消息处理函数给订阅器接收消息
	// 演示订阅两个topic的接收消息处理
	var recMsgFuncs = make(map[string]redisPubSub.RecSubMsgFunc)
	recMsgFuncs[redisPubSub.RedisTestchannel1] = func(topic string, msg interface{}) error {
		logger.InfoLog("Subscriber receive", fmt.Sprintf("Redis Topic: %s,Message:%s", topic, msg))
		return nil
	}
	recMsgFuncs[redisPubSub.RedisTestchannel2] = func(topic string, msg interface{}) error {
		logger.InfoLog("Subscriber receive", fmt.Sprintf("Redis Topic: %s,Message:%s", topic, msg))
		return nil
	}

	// 订阅频道，可多个,并接收处理消息
	err = redisPubSub.Subscribe(recMsgFuncs, redisPubSub.RedisTestchannel1, redisPubSub.RedisTestchannel2)
	if err != nil {
		logger.WarmLog(err.Error())
	}

}
