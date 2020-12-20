package XLogger

import (
	"go.uber.org/zap/zapcore"
	"os"
)

func NewStdOutCore(cfg *Config) zapcore.Core {
	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewConsoleEncoder(defaultFormatConfig()),
		// 日志异步输出配置
		zapcore.AddSync(os.Stdout),
		// 日志级别
		SettingLevelEnableFunc(cfg),
	)
}
