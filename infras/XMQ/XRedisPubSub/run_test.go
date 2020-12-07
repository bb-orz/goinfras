package XRedisPubSub

import (
	"GoWebScaffold/infras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
)

const (
	ChannelName1 = "test1"
	ChannelName2 = "test2"
)

func TestRedisPubsubPool(t *testing.T) {
	Convey("TestRedisPubsubPool", t, func() {
		CreateDefaultPool(nil, zap.L())

		var subHandle1 = func(channel string, msg interface{}) error {

			return nil
		}

		var subHandle2 = func(channel string, msg interface{}) error {

			return nil
		}

		recSubMsgFuncs := make(map[string]RecSubMsgFunc)
		recSubMsgFuncs[ChannelName1] = subHandle1
		recSubMsgFuncs[ChannelName2] = subHandle2
		// TODO
		XRedisSubscriber().Subscribe(recSubMsgFuncs, ChannelName1, ChannelName2)

	})
}

func TestStarter(t *testing.T) {
	s := NewStarter()
	logger, err := zap.NewDevelopment()
	So(err, ShouldBeNil)
	sctx := infras.CreateDefaultStarterContext(nil, logger)
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
}
