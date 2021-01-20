package XRedisPubSub

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XLogger"
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
		XLogger.CreateDefaultLogger(nil)

		CreateDefaultPool(nil)
		recSubMsgFuncs := make(map[string]RecSubMsgFunc)
		// ChannelName1 订阅频道消息的处理函数
		recSubMsgFuncs[ChannelName1] = func(channel string, msg interface{}) error {
			XLogger.XCommon().Info("Receive Message:", zap.String("channel", channel), zap.Any("message", msg))
			fmt.Println(msg)
			return nil
		}
		// ChannelName2 订阅频道消息的处理函数
		recSubMsgFuncs[ChannelName2] = func(channel string, msg interface{}) error {
			XLogger.XCommon().Info("Receive Message:", zap.String("channel", channel), zap.Any("message", msg))
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
		err := XRedisSubscriber().Subscribe(recSubMsgFuncs, unSubCh)
		So(err, ShouldBeNil)

	})
}

func TestPublisher(t *testing.T) {
	Convey("TestRedisSubscriber", t, func() {
		XLogger.CreateDefaultLogger(nil)

		CreateDefaultPool(nil)
		publisher := XRedisPublisher()
		var err error
		err = publisher.Publish(ChannelName1, "this a message from TestPublisher To ChannelName1")
		So(err, ShouldBeNil)
		err = publisher.Publish(ChannelName2, "this a message from TestPublisher To ChannelName2")
		So(err, ShouldBeNil)
	})
}

func TestStarter(t *testing.T) {
	Convey("Test XRedisPubSub Starter", t, func() {
		logger := goinfras.NewCommandLineStarterLogger("debug")
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
		s.Start(sctx)
	})
}
