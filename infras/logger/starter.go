package logger

import (
	"GoWebScaffold/infras"
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

type Starter struct {
	infras.BaseStarter
	cfg     *Config
	Writers []io.Writer
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Logger", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	commonLogger = NewCommonLogger(s.cfg, s.Writers...)
	syncErrorLogger = NewSyncErrorLogger(s.cfg)
	sctx.SetLogger(commonLogger)
	sctx.Logger().Info("CommonLogger And SyncErrorLogger Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	// 关闭前刷入日志数据
	CommonLogger().Sync()
	SyncErrorLogger().Sync()
}

func (s *Starter) Priority() int { return infras.INT_MAX }

/*For testing*/
func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			AppName:           "",
			AppVersion:        "",
			DevEnv:            true,
			AddCaller:         true,
			DebugLevelSwitch:  false,
			InfoLevelSwitch:   true,
			WarnLevelSwitch:   true,
			ErrorLevelSwitch:  true,
			DPanicLevelSwitch: true,
			PanicLevelSwitch:  false,
			FatalLevelSwitch:  true,
			SimpleZapCore:     true,
			SyncZapCore:       false,
			SyncLogSwitch:     true,
			StdoutLogSwitch:   true,
			RotateLogSwitch:   false,
			LogDir:            "../../log",
		}
	}

	commonLogger = NewCommonLogger(config)
	return err
}
