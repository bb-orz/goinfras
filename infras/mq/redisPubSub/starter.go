package redisPubSub

import (
	"GoWebScaffold/infras"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
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
	cfg *RedisPubSubConfig
}

func (s *RedisPubSubStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := RedisPubSubConfig{}
	err := kvs.Unmarshal(configs, &define, "RedisPubSub")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *RedisPubSubStarter) Setup(sctx *infras.StarterContext) {
	redispsPool = NewRedisPubsubPool(s.cfg, sctx.Logger())
	sctx.Logger().Info("RedisPubSubPool Setup Successful ...")
}

func (s *RedisPubSubStarter) Stop(sctx *infras.StarterContext) {
	_ = RedisPubSubPool().Close()
}

/*For testing*/
func RunForTesting(config *RedisPubSubConfig) error {
	var err error
	if config == nil {
		config = &RedisPubSubConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}

	redispsPool = NewRedisPubsubPool(config, zap.L())
	return nil
}
