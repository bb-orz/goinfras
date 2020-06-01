package logger

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
	"io"
)

var commonLogger *zap.Logger
var syncErrorLogger *zap.Logger

func CommonLogger() *zap.Logger {
	infras.Check(commonLogger)
	return commonLogger
}

func SyncErrorLogger() *zap.Logger {
	infras.Check(syncErrorLogger)
	return syncErrorLogger
}

type LoggerStarter struct {
	infras.BaseStarter
	cfg     *loggerConfig
	Writers []io.Writer
}

func (s *LoggerStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := loggerConfig{}
	err := kvs.Unmarshal(configs, &define, "Logger")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *LoggerStarter) Setup(sctx *infras.StarterContext) {
	commonLogger = NewCommonLogger(s.cfg, s.Writers...)
	syncErrorLogger = NewSyncErrorLogger(s.cfg)
	sctx.SetLogger(commonLogger)
	sctx.Logger().Info("CommonLogger And SyncErrorLogger Setup Successful!")
}

func (s *LoggerStarter) Priority() int { return infras.INT_MAX }
