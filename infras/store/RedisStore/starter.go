package RedisStore

import (
	"GoWebScaffold/infras"
	"github.com/garyburd/redigo/redis"
	"github.com/tietang/props/kvs"
)

var rPool *redis.Pool

func RedisPool() *redis.Pool {
	infras.Check(rPool)
	return rPool
}

type RedisStarter struct {
	infras.BaseStarter
	cfg *redisConfig
}

func (s *RedisStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := redisConfig{}
	err := kvs.Unmarshal(configs, &define, "Redis")
	if err != nil {
		panic(err.Error())
	}
	s.cfg = &define
}

// 检查该组件的前置依赖
func (s *RedisStarter) Setup(sctx *infras.StarterContext) {}

// 启动该资源组件
func (s *RedisStarter) Start(sctx *infras.StarterContext) {
	var err error
	rPool, err = NewRedisPool(s.cfg)
	if err != nil {
		panic(err.Error())
	}
}

// 停止服务
func (s *RedisStarter) Stop(sctx *infras.StarterContext) {

}
