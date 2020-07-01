package global

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
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
	configs := sctx.Configs()
	define := GlobalConfig{}
	err := kvs.Unmarshal(configs, &define, "Global")
	infras.FailHandler(err)
	cfg = &define
}
