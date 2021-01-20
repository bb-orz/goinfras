package XRedis

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XCache"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"

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

		logger := goinfras.NewCommandLineStarterLogger("debug")
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
	})
}

func TestXCommonWithRedisCache(t *testing.T) {
	Convey("TestStarter", t, func() {
		var err error
		err = CreateDefaultPool(nil)
		So(err, ShouldBeNil)
		XCache.SettingCommonCache(NewCommonRedisCache())

		XCache.XCommon().Delete("a")
		XCache.XCommon().Delete("b")
		XCache.XCommon().Delete("i")

		Println("Add key a as val aa")
		err = XCache.XCommon().Add("a", "aa")
		So(err, ShouldBeNil)

		if v, b := XCache.XCommon().Get("a"); b {
			Printf("Get key a:%s \n", v)
		}

		Println("Set key a as val b")
		err = XCache.XCommon().Set("a", "b")
		So(err, ShouldBeNil)
		if v, b := XCache.XCommon().Get("a"); b {
			Printf("Get key a:%s \n", v)
		}

		Println("Replace key a as val c")
		err = XCache.XCommon().Replace("a", "c")
		So(err, ShouldBeNil)

		if v, b := XCache.XCommon().Get("a"); b {
			Printf("Get key a:%s \n", v)
		}

		Println("AddWithExp key b as val bb,exp:5s")
		err = XCache.XCommon().AddWithExp("b", "bb", 5)
		So(err, ShouldBeNil)

		var t time.Time
		var v interface{}
		var b bool
		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with exp: [value]:%s,[exp]:%v \n", v, t)
		}

		dur := t.Unix() - time.Now().Unix()
		time.Sleep(time.Duration(dur+1) * time.Second)
		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with exp on timeout: [value]:%s,[exp]:%v \n", v, t)
		} else {
			Println("After Sleep,key b Timeout")
		}

		Println("SetWithExp key b as val bbb,exp:5s")
		err = XCache.XCommon().SetWithExp("b", "bbb", 5)
		So(err, ShouldBeNil)

		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with: [value]:%s,[exp]:%v \n", v, t)
		} else {
			Println("key b Timeout")
		}

		Println("ReplaceWithExp key b as val ccc,exp:3s")
		err = XCache.XCommon().ReplaceWithExp("b", "ccc", 3)
		So(err, ShouldBeNil)

		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with exp: [value]:%s,[exp]:%v \n", v, t)
		} else {
			Println("key b Timeout")
		}

		Println("Set int key i as 5")
		err = XCache.XCommon().Set("i", 5)
		So(err, ShouldBeNil)

		Println("Increment key i + 3")
		err = XCache.XCommon().Increment("i", 3)
		So(err, ShouldBeNil)

		Println("Decrement key i -2")
		err = XCache.XCommon().Decrement("i", 2)
		So(err, ShouldBeNil)

		if i, b := XCache.XCommon().Get("i"); b {
			Printf("Print i: %s \n", i)
		}

		Println("delete key i")
		XCache.XCommon().Delete("i")
		if i, b := XCache.XCommon().Get("i"); b {
			Printf("i:%#v \n", i)
		} else {
			Println("i was deleted")
		}
	})
}
