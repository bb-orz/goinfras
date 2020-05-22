package core

import (
	"GoWebScaffold/infras/config"
	"GoWebScaffold/infras/logger/hook"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// 简单的日志记录器核心:只输出到stdout和file
func RotateCoreList(appConf *config.AppConfig, format zapcore.EncoderConfig) zapcore.Core {
	var coreList []zapcore.Core
	if appConf.LogConf.DebugLevelSwitch || appConf.LogConf.InfoLevelSwitch || appConf.LogConf.WarnLevelSwitch {
		coreList = append(coreList, RotateInfoCore(appConf, format))
	}
	if appConf.LogConf.ErrorLevelSwitch || appConf.LogConf.DPanicLevelSwitch || appConf.LogConf.PanicLevelSwitch || appConf.LogConf.FatalLevelSwitch {
		coreList = append(coreList, RotateErrorCore(appConf, format))
	}
	return zapcore.NewTee(coreList...)
}

//简单非错误信息(debug/info/warn)日志记录器:只输出到stdout和file
func RotateInfoCore(appConf *config.AppConfig, format zapcore.EncoderConfig) zapcore.Core {
	// 记录所有非错误日志级别
	levelEnablerFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel && level <= zapcore.WarnLevel
	})
	// 异步写输出列表,可多个
	var writeSyncerList []zapcore.WriteSyncer

	// 默认输出到stdout
	if appConf.LogConf.StdoutLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(os.Stdout))
	}
	// 输出到可归档文件
	if appConf.LogConf.RetateLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(hook.RotateFileLogHook("info", appConf)))
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

//简单错误信息(error/dpanic/panic/fatal)日志记录器:只输出到stdout和file
func RotateErrorCore(appConf *config.AppConfig, format zapcore.EncoderConfig) zapcore.Core {
	// 记录所有非错误日志级别
	levelEnablerFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel && level <= zapcore.FatalLevel
	})
	// 异步写输出列表,可多个
	var writeSyncerList []zapcore.WriteSyncer

	// 默认输出到stdout
	if appConf.LogConf.StdoutLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(os.Stdout))
	}
	// 输出到可归档文件
	if appConf.LogConf.RetateLogSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(hook.RotateFileLogHook("error", appConf)))
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
