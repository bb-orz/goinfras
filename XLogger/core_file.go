package XLogger

import (
	"go.uber.org/zap/zapcore"
	"os"
)

// 文件记录核心
func NewFileSyncCore(cfg *Config) (zapcore.Core, error) {
	var err error
	var file *os.File
	// 创建日志文件
	fileLogName := cfg.FileLogName

	// 追加方式打开
	file, err = os.OpenFile(fileLogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
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
	var file *os.File
	// 创建日志文件
	syncErrorLogName := cfg.SyncErrorLogName

	// 追加方式打开
	file, err = os.OpenFile(syncErrorLogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
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
