package XLogger

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"go.uber.org/zap"
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
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print Logger Config:", zap.Any("EtcdConfig", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	commonLogger = NewCommonLogger(s.cfg, s.Writers...)
	syncErrorLogger = NewSyncErrorLogger(s.cfg)
	sctx.SetLogger(commonLogger)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err1, err2 error
	err1 = goinfras.Check(commonLogger)
	err2 = goinfras.Check(syncErrorLogger)
	if err1 != nil || err2 != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Zap Logger Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Zap Logger Setup Successful!", s.Name()))
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
