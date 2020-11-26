package redisPubSub

import (
	"GoWebScaffold/infras"
	redigo "github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

var redispsPool *redigo.Pool

func Pool() *redigo.Pool {
	infras.Check(redispsPool)
	return redispsPool
}

// 从Redis连接池获取一个连接
func GetRedisConn() redigo.Conn {
	conn := Pool().Get()
	return conn
}

// 从Redis连接池获取一个PubSub连接
func GetRedisPubSubConn() *redigo.PubSubConn {
	conn := Pool().Get()
	psConn := redigo.PubSubConn{Conn: conn}
	return &psConn
}

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("RedisPubSub", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	redispsPool = NewRedisPubsubPool(s.cfg, sctx.Logger())
	sctx.Logger().Info("RedisPubSubPool Setup Successful ...")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	_ = Pool().Close()
}

/*For testing*/
func RunForTesting(config *Config) error {
	if config == nil {
		config = &Config{
			true,
			"127.0.0.1",
			6380,
			false,
			"",
			0,
			50,
			60,
		}

	}

	redispsPool = NewRedisPubsubPool(config, zap.L())
	return nil
}
