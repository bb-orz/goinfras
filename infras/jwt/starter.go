package jwt

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/store/redisStore"
)

var tku ITokenUtils

func TokenUtils() ITokenUtils {
	infras.Check(tku)
	return tku
}

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Jwt", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
}

func (s *Starter) Start(sctx *infras.StarterContext) {
	if redisStore.Pool() != nil {
		tku = NewTokenUtilsX([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds, redisStore.Pool())
	} else {
		tku = NewTokenUtils([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds)
	}
}

func (s *Starter) Stop(sctx *infras.StarterContext) {}

/*For testing*/
func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			PrivateKey: "ginger_key",
			ExpSeconds: 60,
		}

	}
	tku = NewTokenUtils([]byte(config.PrivateKey), config.ExpSeconds)
	return err
}
