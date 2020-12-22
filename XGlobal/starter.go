package XGlobal

import (
	"fmt"
	"github.com/bb-orz/goinfras"
)

type starter struct {
	goinfras.BaseStarter
	cfg Global
}

func NewStarter() *starter {
	starter := new(starter)
	return starter
}

func (s *starter) Name() string {
	return "XGlobal"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define map[string]interface{}
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Global", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
		sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %v \n", define))
	}
	s.cfg = Global(define)
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// 把全局配置变量设置进viper
	_g = s.cfg
	sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("Global Constant Setuped! \n"))
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(_g)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("Global function and Constant Setup Successful! \n"))
		return true
	}
	return false
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.BasicGroup }
