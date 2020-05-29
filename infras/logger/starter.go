package logger

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
	"io"
)

var commonLogger *zap.Logger
var syncErrorLogger *zap.Logger

func CommonLogger() *zap.Logger  {
	infras.Check(commonLogger)
	return commonLogger
}

func SyncErrorLogger() *zap.Logger  {
	infras.Check(syncErrorLogger)
	return syncErrorLogger
}

type LoggerStarter struct {
	infras.BaseStarter
	cfg *loggerConfig
	Writers []io.Writer
}

// 初始化时读取有关日志记录器的配置信息
func (s *LoggerStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := loggerConfig{}
	err := kvs.Unmarshal(configs, &define, "Logger")
	if err != nil {
		panic(err.Error())
	}
	s.cfg = &define
}

// 启动日志记录器
func (s *LoggerStarter) Start(sctx *infras.StarterContext) {
	commonLogger = NewCommonLogger(s.cfg,s.Writers...)
	syncErrorLogger = NewSyncErrorLogger(s.cfg)
}

// 优先级为最优先启动
func (s *LoggerStarter) Priority() int { return infras.INT_MAX }