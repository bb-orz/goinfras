package XGocache

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
	return "XGocache"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Gocache", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %v \n", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	goCache, err = NewCacheForm(s.cfg)
	if sctx.PassWarning(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("GoCache From DumpItems instance Setuped! \n"))
	} else {
		goCache = NewCache(s.cfg)
		sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("GoCache New instance Setuped! \n"))
	}

	// 设置通用缓存操作
	XCache.SettingCommonCache(NewCommonGocache())
	sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("GoCache Common Cache Setuped! \n"))

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(goCache)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("GoCache instance Setup Successful! \n"))
		return true
	}
	return false
}

func (s *starter) Stop() {
	_ = DumpItems(s.cfg)
	goCache.Flush()
	goCache = nil
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
