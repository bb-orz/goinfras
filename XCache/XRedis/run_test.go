package XRedis

import (
	"github.com/bb-orz/goinfras"
	"github.com/gomodule/redigo/redis"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCommonRedisPool(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		err := CreateDefaultPool(nil)
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
		err := CreateDefaultPool(nil)
		So(err, ShouldBeNil)

		command := XCommand()
		reply1, err := command.R("Set", "name", "joker")
		So(err, ShouldBeNil)
		Println("Set reply:", reply1)

		reply2, err := redis.String(command.R("Get", "name"))
		So(err, ShouldBeNil)
		Println("Get reply:", reply2)
	})
}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		var err error
		err = CreateDefaultPool(nil)
		So(err, ShouldBeNil)

		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
	})
}
