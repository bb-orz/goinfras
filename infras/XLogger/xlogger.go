package XLogger

import (
	"go.uber.org/zap"
	"io"
)

var commonLogger *zap.Logger
var syncErrorLogger *zap.Logger

func XCommon() *zap.Logger {
	return commonLogger
}

func XSyncError() *zap.Logger {
	return syncErrorLogger
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config, syncWriters ...io.Writer) error {

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
