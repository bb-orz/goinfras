package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

/*
通用zap日志记录器
*/
func NewCommonLogger(cfg *Config, syncWriters ...io.Writer) *zap.Logger {
	var optionList []zap.Option
	var appName, version, caller zap.Option
	// Option：基本日志字段
	if cfg.AppName != "" {
		appName = zap.Fields(zap.String("app", cfg.AppName))
		optionList = append(optionList, appName)
	}
	if cfg.AppVersion != "" {
		version = zap.Fields(zap.String("version", cfg.AppVersion))
		optionList = append(optionList, version)
	}

	// Option：注释每条信息所在文件名和行号
	if cfg.AddCaller {
		caller = zap.AddCaller()
		optionList = append(optionList, caller)
	}

	// Option：开发环境时进入开发模式，使其具有良好的动态性能,记录死机而不是简单地记录错误。
	if cfg.DevEnv {
		optionList = append(optionList, zap.Development())
	}

	// Option:配置日志记录器核心列表
	var zCore zapcore.Core
	if cfg.SimpleZapCore {
		zCore = zapcore.NewTee(SimpleCoreList(cfg, commonFormatConfig()))
	} else if cfg.RotateZapCore {
		zCore = zapcore.NewTee(RotateCoreList(cfg, commonFormatConfig()))
	} else if cfg.SyncLogSwitch {
		zCore = zapcore.NewTee(SyncCoreList(cfg, commonFormatConfig(), syncWriters...))
	} else {
		zCore = SimpleCore(cfg, commonFormatConfig())
	}
	return zap.New(zCore, optionList...)
}

/*
记录zap hook 异步日志中的日志信息，一般在记录远程日志记录出错时使用，记录异步信息出现的问题，如mongo记录异常，消息队列记录异常等
该类别的日志信息只能输出到std和日志文件
*/
func NewSyncErrorLogger(cfg *Config) *zap.Logger {
	var optionList []zap.Option
	var appName, version, caller zap.Option
	// Option：基本日志字段
	if cfg.AppName != "" {
		appName = zap.Fields(zap.String("app", cfg.AppName))
		optionList = append(optionList, appName)
	}
	if cfg.AppVersion != "" {
		version = zap.Fields(zap.String("version", cfg.AppVersion))
		optionList = append(optionList, version)
	}

	// Option：注释每条信息所在文件名和行号
	if cfg.AddCaller {
		caller = zap.AddCaller()
		optionList = append(optionList, caller)
	}

	// Option：开发环境时进入开发模式，使其具有良好的动态性能,记录死机而不是简单地记录错误。
	if cfg.DevEnv {
		optionList = append(optionList, zap.Development())
	}

	// 配置核心
	c := SimpleErrorCore(cfg, commonFormatConfig())
	return zap.New(c, optionList...)
}
