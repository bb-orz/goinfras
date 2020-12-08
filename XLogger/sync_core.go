package XLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// 异步日志记录器核心
func syncCoreList(cfg *Config, format zapcore.EncoderConfig, syncWriters ...io.Writer) zapcore.Core {
	var coreList []zapcore.Core
	if cfg.DebugLevelSwitch || cfg.InfoLevelSwitch || cfg.WarnLevelSwitch {
		coreList = append(coreList, syncInfoCore(cfg, format, syncWriters...))
	}
	if cfg.ErrorLevelSwitch || cfg.DPanicLevelSwitch || cfg.PanicLevelSwitch || cfg.FatalLevelSwitch {
		coreList = append(coreList, syncErrorCore(cfg, format, syncWriters...))
	}
	return zapcore.NewTee(coreList...)
}

// 异步非错误信息(debug/info/warn)日志记录器
func syncInfoCore(cfg *Config, format zapcore.EncoderConfig, syncWriters ...io.Writer) zapcore.Core {
	// 记录所有非错误日志级别
	levelEnablerFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel && level <= zapcore.WarnLevel
	})
	// 异步写输出列表,可多个
	var writeSyncerList []zapcore.WriteSyncer

	// 默认输出到stdout
	if cfg.StdoutLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(os.Stdout))
	}
	// 添加异步输出
	if cfg.SyncLogSwitch {
		for _, sw := range syncWriters {
			writeSyncerList = append(writeSyncerList, zapcore.AddSync(sw))
		}
	}

	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(format),
		// 日志异步输出配置
		zapcore.NewMultiWriteSyncer(writeSyncerList...),
		// 日志级别
		levelEnablerFunc,
	)
}

// 异步错误信息(error/dpanic/panic/fatal)日志记录器:
func syncErrorCore(cfg *Config, format zapcore.EncoderConfig, syncWriters ...io.Writer) zapcore.Core {
	// 记录所有非错误日志级别
	levelEnablerFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel && level <= zapcore.FatalLevel
	})
	// 异步写输出列表,可多个
	var writeSyncerList []zapcore.WriteSyncer

	// 默认输出到stdout
	if cfg.StdoutLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(os.Stdout))
	}
	// 添加异步输出
	if cfg.SyncLogSwitch {
		for _, sw := range syncWriters {
			writeSyncerList = append(writeSyncerList, zapcore.AddSync(sw))
		}
	}

	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(format),
		// 日志异步输出配置
		zapcore.NewMultiWriteSyncer(writeSyncerList...),
		// 日志级别
		levelEnablerFunc,
	)
}
