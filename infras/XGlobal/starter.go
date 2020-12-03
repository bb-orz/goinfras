package XGlobal

import (
	"GoWebScaffold/infras"
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
	err := viper.UnmarshalKey("Global", &define)
	infras.FailHandler(err)
	SetComponent(define)
}

func (s *Starter) Start(sctx *infras.StarterContext) {
	// 把全局配置变量设置进viper
	// sctx.Configs().Set("AppName", GlobalConfig().AppName)
	// sctx.Configs().Set("ServerName", GlobalConfig().ServerName)
	// sctx.Configs().Set("Env", GlobalConfig().Env)
}
