package XLogger

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"io"
)

type starter struct {
	goinfras.BaseStarter
	cfg     *Config
	Writers []io.Writer
}

func NewStarter(writer ...io.Writer) *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	starter.Writers = writer
	return starter
}

func (s *starter) Name() string {
	return "XLogger"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Logger", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %v \n", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	commonLogger = NewCommonLogger(s.cfg, s.Writers...)
	syncErrorLogger = NewSyncErrorLogger(s.cfg)
	sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("Zap Logger Steuped!  \n"))
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
	sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("Zap Logger Setup Successful! \n"))
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
