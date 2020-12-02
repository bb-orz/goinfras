package redisStore

import (
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*实例化资源用于测试*/
func TestingInstantiation() error {
	var err error
	config := &Config{
		"127.0.0.1",
		6379,
		false,
		"",
		0,
		50,
		60,
	}

	pool, err = NewPool(config, zap.L())
	return err
}

func TestNewCommonRedisPool(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		err := TestingInstantiation()
		So(err, ShouldBeNil)
		Println("pool ActiveCount:", RedisComponent().Stats().ActiveCount, ",pool IdleCount:", RedisComponent().Stats().IdleCount)

		conn := RedisComponent().Get()
		Println("pool ActiveCount:", RedisComponent().Stats().ActiveCount, ",pool IdleCount:", RedisComponent().Stats().IdleCount)

		reply, err := conn.Do("Ping")
		So(err, ShouldBeNil)
		Println("Ping Reply", reply)

		err = conn.Close()
		So(err, ShouldBeNil)
		Println("pool ActiveCount:", RedisComponent().Stats().ActiveCount, ",pool IdleCount:", RedisComponent().Stats().IdleCount)

	})
}

func TestCommonRedisDao(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		err := TestingInstantiation()
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
