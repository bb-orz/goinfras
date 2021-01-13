package XLogger

import (
	"github.com/bb-orz/goinfras"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// 文件记录核心
func NewFileSyncCore(cfg *Config) (zapcore.Core, error) {
	var err error
	var file io.Writer
	// 创建日志文件
	fileLogName := cfg.FileLogName

	if file, err = goinfras.OpenFile(fileLogName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm); err != nil {
		return nil, err
	}

	// 转成符合zapcore的输出接口类型
	fileSyncer := zapcore.AddSync(file)
	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(defaultFormatConfig()),
		// 日志异步输出配置
		zapcore.NewMultiWriteSyncer(fileSyncer),
		// 日志级别
		SettingLevelEnableFunc(cfg),
	), nil
}

// 异步日志记录错误文件记录核心
func NewFileSyncErrorCore(cfg *Config) (zapcore.Core, error) {
	var err error
	var file io.Writer
	// 创建日志文件
	syncErrorLogName := cfg.SyncErrorLogName

	if file, err = goinfras.OpenFile(syncErrorLogName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm); err != nil {
		return nil, err
	}

	// 转成符合zapcore的输出接口类型
	fileSyncer := zapcore.AddSync(file)

	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(defaultFormatConfig()),
		// 日志异步输出配置
		zapcore.NewMultiWriteSyncer(fileSyncer),
		// 日志级别
		SettingLevelEnableFunc(cfg),
	), nil
}
