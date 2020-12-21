package XLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 根据配置设置可输出日志信息的级别
func SettingLevelEnableFunc(cfg *Config) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		switch level {
		case zap.DebugLevel:
			if cfg.DisableDebugLevelSwitch {
				return false
			}
		case zap.InfoLevel:
			if cfg.DisableInfoLevelSwitch {
				return false
			}
		case zap.WarnLevel:
			if cfg.DisableWarnLevelSwitch {
				return false
			}
		case zap.ErrorLevel:
			if cfg.DisableErrorLevelSwitch {
				return false
			}
		case zap.DPanicLevel:
			if cfg.DisableDPanicLevelSwitch {
				return false
			}
		case zap.PanicLevel:
			if cfg.DisablePanicLevelSwitch {
				return false
			}
		case zap.FatalLevel:
			if cfg.DisableFatalLevelSwitch {
				return false
			}
		}
		return true
	}
}
