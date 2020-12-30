package XLogger

import (
	"fmt"
	"github.com/bb-orz/goinfras"
)

type starter struct {
	goinfras.BaseStarter
	cfg     *Config
	Outputs []LoggerOutput
}

func NewStarter(outputs ...LoggerOutput) *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	starter.Outputs = outputs
	return starter
}

func (s *starter) Name() string {
	return "XLogger"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Logger", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v \n", define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	commonLogger, err = NewCommonLogger(s.cfg, s.Outputs...)
	if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().Info(s.Name(), goinfras.StepSetup, "Zap Commond Logger Setuped! ")
	}
	syncErrorLogger, err = NewSyncErrorLogger(s.cfg)
	if sctx.PassError(s.Name(), goinfras.StepSetup, err) {
		sctx.Logger().Info(s.Name(), goinfras.StepSetup, "Zap SyncError Logger Setuped! ")
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err error
	err = goinfras.Check(commonLogger)
	if !sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		return false
	}
	err = goinfras.Check(syncErrorLogger)
	if !sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		return false
	}
	sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Zap Logger Setup Successful!")
	return true
}

func (s *starter) Stop() {
	// 关闭前刷入日志数据
	commonLogger.Sync()
	syncErrorLogger.Sync()
}

func (s *starter) Priority() int { return goinfras.INT_MAX }

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.BasicGroup }
