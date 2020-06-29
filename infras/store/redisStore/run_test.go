package redisStore

import (
	"github.com/garyburd/redigo/redis"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCommonRedisPool(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		err := RunForTesting()
		So(err, ShouldBeNil)
		Println("pool ActiveCount:", RedisPool().Stats().ActiveCount, ",pool IdleCount:", RedisPool().Stats().IdleCount)

		conn := RedisPool().Get()
		Println("pool ActiveCount:", RedisPool().Stats().ActiveCount, ",pool IdleCount:", RedisPool().Stats().IdleCount)

		reply, err := conn.Do("Ping")
		So(err, ShouldBeNil)
		Println("Ping Reply", reply)

		err = conn.Close()
		So(err, ShouldBeNil)
		Println("pool ActiveCount:", RedisPool().Stats().ActiveCount, ",pool IdleCount:", RedisPool().Stats().IdleCount)

	})
}

func TestCommonRedisDao(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		err := RunForTesting()
		So(err, ShouldBeNil)

		commonRedisDao := NewCommonRedisDao()

		reply1, err := commonRedisDao.R("Set", "name", "joker")
		So(err, ShouldBeNil)
		Println("Set reply:", reply1)

		reply2, err := redis.String(commonRedisDao.R("Get", "name"))
		So(err, ShouldBeNil)
		Println("Get reply:", reply2)
	})
}
