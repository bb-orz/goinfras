package mail

import (
	"GoWebScaffold/infras"
	"gopkg.in/gomail.v2"
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
	err := viper.UnmarshalKey("Mail", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var m *gomail.Dialer
	if s.cfg.NoAuth {
		m = NewNoAuthDialer(s.cfg.Server, s.cfg.Port)
	} else {
		m = NewAuthDialer(s.cfg.Server, s.cfg.User, s.cfg.Password, s.cfg.Port)
	}

	SetComponent(m)
}
