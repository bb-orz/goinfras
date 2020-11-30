package jwt

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/store/redisStore"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *Starter {
	starter := new(Starter)
	starter.cfg = Config{}
	return starter
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
	var t ITokenUtils
	if redisStore.Pool() != nil {
		t = NewTokenUtilsX([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds, redisStore.Pool())
	} else {
		t = NewTokenUtils([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds)
	}
	SetComponent(t)
}

func (s *Starter) Stop(sctx *infras.StarterContext) {}
