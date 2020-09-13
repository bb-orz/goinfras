package global

import (
	"GoWebScaffold/infras"
)

// 全局配置
var cfg *GlobalConfig

func Config() *GlobalConfig {
	infras.Check(cfg)
	return cfg
}

type GlobalStarter struct {
	infras.BaseStarter
}

func (s *GlobalStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := GlobalConfig{}
	err := viper.UnmarshalKey("Global", &define)
	infras.FailHandler(err)
	cfg = &define
}
