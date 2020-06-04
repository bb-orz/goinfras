package redisStore

import (
	"github.com/garyburd/redigo/redis"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCommonRedisPool(t *testing.T) {
	Convey("Redis Dao Test", t, func() {
		config := redisConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Redis Config:", config)

		pool, err := NewRedisPool(&config, zap.L())
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
		config := redisConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Redis Config:", config)

		rPool, err = NewRedisPool(&config, zap.L())
		So(err, ShouldBeNil)

		commonRedisDao := NewCommonRedisDao(rPool)

		reply1, err := commonRedisDao.R("Set", "name", "joker")
		So(err, ShouldBeNil)
		Println("Set reply:", reply1)

		reply2, err := redis.String(commonRedisDao.R("Get", "name"))
		So(err, ShouldBeNil)
		Println("Get reply:", reply2)
	})
}
