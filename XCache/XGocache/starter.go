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
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Gocache", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	goCache, err = NewCacheForm(s.cfg)
	if sctx.PassWarning(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().Info(s.Name(), goinfras.StepSetup, "GoCache From DumpItems instance Setuped!")
	} else {
		goCache = NewCache(s.cfg)
		sctx.Logger().Info(s.Name(), goinfras.StepSetup, "GoCache New instance Setuped! ")
	}

	// 设置通用缓存操作
	XCache.SettingCommonCache(NewCommonGocache())
	sctx.Logger().Info(s.Name(), goinfras.StepSetup, "GoCache Common Cache Setuped! ")

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(goCache)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "GoCache instance Setup Successful! ")
		return true
	}
	return false
}

func (s *starter) Stop() error {
	goCache.Flush()
	if err := DumpItems(s.cfg); err != nil {
		return err
	}

	goCache = nil
	fmt.Println("GoCache Stopped!")
	return nil
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
