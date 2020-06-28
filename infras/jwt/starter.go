package jwt

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/store/redisStore"
	"github.com/tietang/props/kvs"
)

var tku ITokenUtils

func TokenUtils() ITokenUtils {
	infras.Check(tku)
	return tku
}

type JwtStarter struct {
	infras.BaseStarter
	cfg *JwtConfig
}

func (s *JwtStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := JwtConfig{}
	err := kvs.Unmarshal(configs, &define, "Jwt")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *JwtStarter) Setup(sctx *infras.StarterContext) {
}

func (s *JwtStarter) Start(sctx *infras.StarterContext) {
	if redisStore.RedisPool() != nil {
		tku = NewTokenUtilsX([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds, redisStore.RedisPool())
	} else {
		tku = NewTokenUtils([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds)
	}
}

func (s *JwtStarter) Stop(sctx *infras.StarterContext) {}

/*For testing*/
func RunForTesting(config *JwtConfig) error {
	var err error
	if config == nil {
		config = &JwtConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}
	tku = NewTokenUtils([]byte(config.PrivateKey), config.ExpSeconds)
	return err
}
