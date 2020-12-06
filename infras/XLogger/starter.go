package XLogger

import (
	"GoWebScaffold/infras"
	"fmt"
	"go.uber.org/zap"
	"io"
)

type starter struct {
	infras.BaseStarter
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

func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Logger", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	sctx.Logger().Info("Print Logger Config:", zap.Any("EtcdConfig", *define))
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	commonLogger = NewCommonLogger(s.cfg, s.Writers...)
	syncErrorLogger = NewSyncErrorLogger(s.cfg)
	sctx.SetLogger(commonLogger)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	var err1, err2 error
	err1 = infras.Check(commonLogger)
	err2 = infras.Check(syncErrorLogger)
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

func (s *starter) Priority() int { return infras.INT_MAX }
