package XGlobal

import (
	"fmt"
	"go.uber.org/zap"
	"goinfras"
)

type starter struct {
	goinfras.BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XGlobal"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Global", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Global Config:", zap.Any("GlobalConfig", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// 把全局配置变量设置进viper
	sctx.Configs().Set("AppName", global.AppName)
	sctx.Configs().Set("ServerName", global.ServerName)
	sctx.Configs().Set("Env", global.Env)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(global)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Global Config And Common Function Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Global Config And Common Function Setup Successful!", s.Name()))
	return true
}
