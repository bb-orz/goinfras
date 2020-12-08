package XRedis

import (
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"goinfras"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCommonRedisPool(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		err = CreateDefaultPool(nil, logger)
		So(err, ShouldBeNil)
		Println("pool ActiveCount:", pool.Stats().ActiveCount, ",pool IdleCount:", pool.Stats().IdleCount)

		conn := pool.Get()
		Println("pool ActiveCount:", pool.Stats().ActiveCount, ",pool IdleCount:", pool.Stats().IdleCount)

		reply, err := conn.Do("Ping")
		So(err, ShouldBeNil)
		Println("Ping Reply", reply)

		err = conn.Close()
		So(err, ShouldBeNil)
		Println("pool ActiveCount:", pool.Stats().ActiveCount, ",pool IdleCount:", pool.Stats().IdleCount)

	})
}

func TestCommonRedisDao(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		err = CreateDefaultPool(nil, logger)
		So(err, ShouldBeNil)

		common := XCommon()

		reply1, err := common.R("Set", "name", "joker")
		So(err, ShouldBeNil)
		Println("Set reply:", reply1)

		reply2, err := redis.String(common.R("Get", "name"))
		So(err, ShouldBeNil)
		Println("Get reply:", reply2)
	})
}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		err = CreateDefaultPool(nil, logger)
		So(err, ShouldBeNil)

		s := NewStarter()
		sctx := CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")

		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

	})
}
