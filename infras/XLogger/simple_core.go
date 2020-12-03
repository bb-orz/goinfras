package XLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// 简单的日志记录器核心:只输出到stdout和file
func simpleCoreList(cfg *Config, format zapcore.EncoderConfig) zapcore.Core {
	var coreList []zapcore.Core
	if cfg.DebugLevelSwitch || cfg.InfoLevelSwitch || cfg.WarnLevelSwitch {
		coreList = append(coreList, simpleInfoCore(cfg, format))
	}
	if cfg.ErrorLevelSwitch || cfg.DPanicLevelSwitch || cfg.PanicLevelSwitch || cfg.FatalLevelSwitch {
		coreList = append(coreList, simpleErrorCore(cfg, format))
	}
	return zapcore.NewTee(coreList...)
}

// 简单非错误信息(debug/info/warn)日志记录器:只输出到stdout和file
func simpleInfoCore(cfg *Config, format zapcore.EncoderConfig) zapcore.Core {
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
	// 输出到文件
	if cfg.RotateLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(simpleFileLogHook("info", cfg)))
	}

	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(format),
		//日志异步输出配置
		zapcore.NewMultiWriteSyncer(writeSyncerList...),
		// 日志级别
		levelEnablerFunc,
	)
}

// 简单错误信息(error/dpanic/panic/fatal)日志记录器:只输出到stdout和file
func simpleErrorCore(cfg *Config, format zapcore.EncoderConfig) zapcore.Core {
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
	// 输出到文件
	if cfg.RotateLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(simpleFileLogHook("error", cfg)))
	}

	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(format),
		//日志异步输出配置
		zapcore.NewMultiWriteSyncer(writeSyncerList...),
		// 日志级别
		levelEnablerFunc,
	)
}

func simpleCore(cfg *Config, format zapcore.EncoderConfig) zapcore.Core {
	// 记录所有日志级别
	levelEnablerFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel && level <= zapcore.FatalLevel
	})

	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(format),
		// 日志异步输出配置
		zapcore.AddSync(os.Stdout),
		// 日志级别
		levelEnablerFunc,
	)
}

// 简单文件日志输出
func simpleFileLogHook(filename string, cfg *Config) io.Writer {
	var file io.Writer
	var err error
	fullFileName := cfg.LogDir + filename + ".log"
	file, err = os.OpenFile(fullFileName, os.O_RDWR, os.ModeAppend)
	if os.IsNotExist(err) {
		file, err = os.Create(fullFileName)
		if err != nil {
			panic(err)
		}
	}
	return file
}
