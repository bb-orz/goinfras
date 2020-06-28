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
	cfg     *LoggerConfig
	Writers []io.Writer
}

func (s *LoggerStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := LoggerConfig{}
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

func (s *LoggerStarter) Stop(sctx *infras.StarterContext) {
	// 关闭前刷入日志数据
	CommonLogger().Sync()
	SyncErrorLogger().Sync()
}

func (s *LoggerStarter) Priority() int { return infras.INT_MAX }

/*For testing*/
func RunForTesting(config *LoggerConfig) error {
	var err error
	if config == nil {
		config = &LoggerConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}

	commonLogger = NewCommonLogger(config)
	return err
}
