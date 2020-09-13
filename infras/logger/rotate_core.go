package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

// 归档日志记录器核心:只输出到stdout和归档日期file
func RotateCoreList(cfg *LoggerConfig, format zapcore.EncoderConfig) zapcore.Core {
	var coreList []zapcore.Core
	if cfg.DebugLevelSwitch || cfg.InfoLevelSwitch || cfg.WarnLevelSwitch {
		coreList = append(coreList, RotateInfoCore(cfg, format))
	}
	if cfg.ErrorLevelSwitch || cfg.DPanicLevelSwitch || cfg.PanicLevelSwitch || cfg.FatalLevelSwitch {
		coreList = append(coreList, RotateErrorCore(cfg, format))
	}
	return zapcore.NewTee(coreList...)
}

// 归档非错误信息(debug/info/warn)日志记录器:只输出到stdout和归档日期file
func RotateInfoCore(cfg *LoggerConfig, format zapcore.EncoderConfig) zapcore.Core {
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
	// 输出到可归档文件
	if cfg.RotateLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(RotateFileLogHook("info", cfg)))
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

// 归档错误信息(error/dpanic/panic/fatal)日志记录器:只输出到stdout和归档日期file
func RotateErrorCore(cfg *LoggerConfig, format zapcore.EncoderConfig) zapcore.Core {
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
	// 输出到可归档文件
	if cfg.RotateLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(RotateFileLogHook("error", cfg)))
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

// 按日期归档记录的日志输出
func RotateFileLogHook(filename string, cfg *LoggerConfig) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	rotateLogHook, err := rotatelogs.New(
		cfg.LogDir+filename+"[%Y-%m-%d %H:%M:%S].log",
		rotatelogs.WithLinkName(filename),
		// 最多保留多久
		rotatelogs.WithMaxAge(time.Hour*time.Duration(cfg.MaxDayCount*24)),
		// 多久做一次归档
		rotatelogs.WithRotationTime(time.Hour*24*time.Duration(cfg.WithRotationTime)),
	)

	if err != nil {
		panic(err)
	}
	return rotateLogHook
}
