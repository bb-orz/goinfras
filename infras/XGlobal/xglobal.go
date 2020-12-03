package XGlobal

import "GoWebScaffold/infras"

// 全局配置
var global Config

func SetComponent(cfg Config) {
	global = cfg
}

func GlobalConfig() *Config {
	infras.Check(global)
	return &global
}
