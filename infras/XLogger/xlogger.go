package XLogger

import (
	"go.uber.org/zap"
	"io"
)

var commonLogger *zap.Logger
var syncErrorLogger *zap.Logger

// 创建一个默认配置的Logger
func CreateDefaultLogger(config *Config, syncWriters ...io.Writer) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	// 有异步写入器，打开配置开关
	if len(syncWriters) != 0 {
		config.SimpleZapCore = false
		config.SyncLogSwitch = true
		config.SyncZapCore = true
	}
	commonLogger = NewCommonLogger(config, syncWriters...)
	return err
}

func XCommon() *zap.Logger {
	return commonLogger
}

func XSyncError() *zap.Logger {
	return syncErrorLogger
}
