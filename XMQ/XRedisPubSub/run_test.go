package XRedisPubSub

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
	"time"
)

const (
	ChannelName1 = "test1"
	ChannelName2 = "test2"
)

func TestRedisSubscriber(t *testing.T) {
	Convey("TestRedisSubscriber", t, func() {
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		CreateDefaultPool(nil, logger)
		recSubMsgFuncs := make(map[string]RecSubMsgFunc)
		// ChannelName1 订阅频道消息的处理函数
		recSubMsgFuncs[ChannelName1] = func(channel string, msg interface{}) error {
			logger.Info("Receive Message:", zap.String("channel", channel), zap.Any("message", msg))
			fmt.Println(msg)
			return nil
		}
		// ChannelName2 订阅频道消息的处理函数
		recSubMsgFuncs[ChannelName2] = func(channel string, msg interface{}) error {
			logger.Info("Receive Message:", zap.String("channel", channel), zap.Any("message", msg))
			fmt.Println(msg)
			return nil
		}

		// 取消订阅通道信号，传入需要取消订阅的频道名称
		unSubCh := make(chan string, 1)

		go func() {
			// 10s后发送取消订阅信号
			time.Sleep(10 * time.Second)
			unSubCh <- ChannelName1
			unSubCh <- ChannelName2
		}()
		err = XRedisSubscriber(logger).Subscribe(recSubMsgFuncs, unSubCh)
		So(err, ShouldBeNil)

	})
}

func TestPublisher(t *testing.T) {
	Convey("TestRedisSubscriber", t, func() {
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		CreateDefaultPool(nil, logger)
		publisher := XRedisPublisher(logger)
		err = publisher.Publish(ChannelName1, "this a message from TestPublisher To ChannelName1")
		So(err, ShouldBeNil)
		err = publisher.Publish(ChannelName2, "this a message from TestPublisher To ChannelName2")
		So(err, ShouldBeNil)
	})
}

func TestStarter(t *testing.T) {
	Convey("Test XRedisPubSub Starter", t, func() {

		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")

		s.Setup(sctx)
		Println("Starter Setup Successful!")
		s.Start(sctx)
		Println("Starter Start Successful!")
		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}
	})
}
