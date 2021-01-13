package XLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var commonLogger *zap.Logger
var syncErrorLogger *zap.Logger

// 创建一个默认配置的Logger
func CreateDefaultLogger(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	commonLogger, err = NewCommonLogger(config)
	return err
}

/*
通用zap日志记录器
*/
func NewCommonLogger(cfg *Config, outputs ...LoggerOutput) (*zap.Logger, error) {
	var optionList []zap.Option
	var appName, version, caller zap.Option
	// Option：基本日志字段
	if cfg.AppName != "" {
		appName = zap.Fields(zap.String("App", cfg.AppName))
		optionList = append(optionList, appName)
	}
	if cfg.AppVersion != "" {
		version = zap.Fields(zap.String("Version", cfg.AppVersion))
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

	// 日志记录器核心列表
	var zCore zapcore.Core
	var zCoreList []zapcore.Core

	// 添加标准输出记录核心
	if cfg.EnableStdZapCore {
		stdCore := NewStdOutCore(cfg)
		zCoreList = append(zCoreList, stdCore)
	}

	// 添加归档文件输出或单文件日志记录核心
	if cfg.EnableRotateZapCore {
		rotateLogCore, err := NewRotateLogCore(cfg)
		if err != nil {
			return nil, err
		}
		zCoreList = append(zCoreList, rotateLogCore)

	} else if cfg.EnableFileZapCore {
		fileSyncCore, err := NewFileSyncCore(cfg)
		if err != nil {
			return nil, err
		}
		zCoreList = append(zCoreList, fileSyncCore)
	}

	// mongo日志记录核心
	if cfg.EnableMongoLogZapCore {
		mongoLogCore, err := NewMongoLogCore(cfg)
		if err != nil {
			return nil, err
		}
		zCoreList = append(zCoreList, mongoLogCore)
	}

	// 添加其他用户自定义的记录核心
	if len(outputs) > 0 {
		core := NewUserOutputsCore(outputs...)
		zCoreList = append(zCoreList, core)
	}

	zCore = zapcore.NewTee(zCoreList...)

	return zap.New(zCore, optionList...), nil
}

/*
记录zap hook 异步日志中的日志信息，一般在记录远程日志记录出错时使用，记录异步信息出现的问题，如mongo记录异常，消息队列记录异常等
该类别的日志信息只能输出到std和日志文件
*/
func NewSyncErrorLogger(cfg *Config) (*zap.Logger, error) {
	var optionList []zap.Option
	var appName, version, caller zap.Option
	// Option：基本日志字段
	if cfg.AppName != "" {
		appName = zap.Fields(zap.String("App", cfg.AppName))
		optionList = append(optionList, appName)
	}
	if cfg.AppVersion != "" {
		version = zap.Fields(zap.String("Version", cfg.AppVersion))
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

	// 日志记录器核心列表
	var zCore zapcore.Core
	var zCoreList []zapcore.Core
	// 添加标准输出记录核心
	if cfg.EnableStdZapCore {
		stdCore := NewStdOutCore(cfg)
		zCoreList = append(zCoreList, stdCore)
	}

	if cfg.EnableFileZapCore {
		fileSyncCore, err := NewFileSyncErrorCore(cfg)
		if err != nil {
			return nil, err
		}
		zCoreList = append(zCoreList, fileSyncCore)
	}

	zCore = zapcore.NewTee(zCoreList...)

	return zap.New(zCore, optionList...), nil
}
