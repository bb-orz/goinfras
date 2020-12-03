package XJwt

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/store/redisStore"
	"fmt"
)

type starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = Config{}
	return starter
}

func (s *starter) Name() string {
	return "XJWT"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Jwt", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	if redisStore.Pool() != nil {
		tku = NewTokenUtilsX([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds, redisStore.Pool())
	} else {
		tku = NewTokenUtils([]byte(s.cfg.PrivateKey), s.cfg.ExpSeconds)
	}
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(tku)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: JWT TokenUtils Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: JWT TokenUtils Setup Successful!", s.Name()))
	return true
}

func (s *starter) Start(sctx *infras.StarterContext) {

}

func (s *starter) Stop(sctx *infras.StarterContext) {}
