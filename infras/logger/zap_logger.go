package logger

import (
	"GoWebScaffold/infras/config"
	"GoWebScaffold/infras/constant"
	"GoWebScaffold/infras/logger/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

/*
通用zap日志记录器
*/
func CommonLogger(appConf *config.AppConfig, syncWriters ...io.Writer) *zap.Logger {
	var optionList []zap.Option

	// Option：基本日志字段
	appName := zap.Fields(zap.String("app", constant.AppName))
	version := zap.Fields(zap.String("version", constant.AppVersion))

	// Option：注释每条信息所在文件名和行号
	caller := zap.AddCaller()
	optionList = append(optionList, appName, version, caller)
	// Option：开发环境时进入开发模式，使其具有良好的动态性能,记录死机而不是简单地记录错误。
	if os.Getenv(constant.OsEnvVarName) == constant.DevEnv {
		optionList = append(optionList, zap.Development())
	}

	// Option:配置日志记录器核心列表
	var zCore zapcore.Core
	if appConf.LogConf.SimpleZapCore {
		zCore = zapcore.NewTee(core.SimpleCoreList(appConf, commonFormatConfig()))
	} else if appConf.LogConf.RotateZapCore {
		zCore = zapcore.NewTee(core.RotateCoreList(appConf, commonFormatConfig()))
	} else if appConf.LogConf.SyncLogSwitch {
		zCore = zapcore.NewTee(core.SyncCoreList(appConf, commonFormatConfig(), syncWriters...))
	} else {
		zCore = core.SimpleCore(appConf, commonFormatConfig())
	}
	return zap.New(zCore, optionList...)
}

/*
记录zap hook 异步日志中的日志信息，一般在记录远程日志记录出错时使用，记录异步信息出现的问题，如mongo记录异常，消息队列记录异常等
该类别的日志信息只能输出到std和日志文件
*/
func SyncErrorLogger(appConf *config.AppConfig) *zap.Logger {
	var optionList []zap.Option

	// Option：基本日志选项
	appName := zap.Fields(zap.String("app", constant.AppName))
	version := zap.Fields(zap.String("version", constant.AppVersion))

	// Option：注释每条信息所在文件名和行号
	caller := zap.AddCaller()
	optionList = append(optionList, appName, version, caller)

	// Option：开发环境时进入开发模式，使其具有良好的动态性能,记录死机而不是简单地记录错误。
	if os.Getenv(constant.OsEnvVarName) == constant.DevEnv {
		optionList = append(optionList, zap.Development())
	}

	// 配置核心
	c := core.SimpleErrorCore(appConf, commonFormatConfig())
	return zap.New(c, optionList...)
}
