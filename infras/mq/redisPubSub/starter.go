package redisPubSub

import (
	"GoWebScaffold/infras"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/tietang/props/kvs"
)

var redispsPool *redigo.Pool

func RedisPubSubPool() *redigo.Pool {
	infras.Check(redispsPool)
	return redispsPool
}

// 从Redis连接池获取一个连接
func GetRedisConn() redigo.Conn {
	conn := RedisPubSubPool().Get()
	return conn
}

// 从Redis连接池获取一个PubSub连接
func GetRedisPubSubConn() *redigo.PubSubConn {
	conn := RedisPubSubPool().Get()
	psConn := redigo.PubSubConn{Conn: conn}
	return &psConn
}

type RedisPubSubStarter struct {
	infras.BaseStarter
	cfg *redisPubSubConfig
}

func (s *RedisPubSubStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := redisPubSubConfig{}
	err := kvs.Unmarshal(configs, &define, "RedisPubSub")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *RedisPubSubStarter) Setup(sctx *infras.StarterContext) {}

func (s *RedisPubSubStarter) Start(sctx *infras.StarterContext) {
	redispsPool = GetRedisPubsubPool(s.cfg, sctx.Logger())
	sctx.Logger().Info("Redis PubSub Pool Start Up ...")

}

func (s *RedisPubSubStarter) Stop(sctx *infras.StarterContext) {
	_ = RedisPubSubPool().Close()
}
