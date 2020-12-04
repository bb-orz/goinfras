package XGlobal

import (
	"GoWebScaffold/infras"
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
	return "XGlobal"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Global", &define)
	infras.FailHandler(err)
	global = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	// 把全局配置变量设置进viper
	sctx.Configs().Set("AppName", global.AppName)
	sctx.Configs().Set("ServerName", global.ServerName)
	sctx.Configs().Set("Env", global.Env)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(global)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Global Config And Common Function Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Global Config And Common Function Setup Successful!", s.Name()))
	return true
}
