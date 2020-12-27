package XJwt

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XCache"
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
	var define Config

	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Jwt", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	s.cfg = &define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v \n", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// 如果redis 连接池组件已安装，则缓存token到redis服务器
	if !s.cfg.UseCache {
		tku = NewTokenUtils(s.cfg)
		sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("JWT TokenUtils Not Cache Setuped!  \n"))
	} else {
		// 检查通用缓存
		if XCache.CheckXCommon() {
			tku = NewTokenUtilsWithCache(s.cfg)
			sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("JWT TokenUtils With Cache Setuped!  \n"))
		}
	}
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
