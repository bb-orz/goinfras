package XJwt

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XStore/XRedis"
	"github.com/spf13/viper"
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
	return "XJWT"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config

	// 先从viper读取配置信息
	viperConfig := sctx.Configs()
	if viperConfig != nil {
		err = viper.UnmarshalKey("Jwt", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	// Viper读取不到配置时，default设置
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %v \n", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// 如果redis 连接池组件已安装，则缓存token到redis服务器
	if s.cfg.UseCache && XRedis.CheckPool() {
		tku = NewTokenUtilsX(s.cfg)
	} else {
		tku = NewTokenUtils(s.cfg)
	}
	sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("JWT TokenUtils Steuped!  \n"))
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(tku)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("JWT TokenUtils Setup Successful! \n"))
		return true
	}
	return false
}

func (s *starter) Stop() {}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
