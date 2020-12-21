package XLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

// 定义日志输出
type LoggerOutput struct {
	Format           zapcore.EncoderConfig
	Writer           io.Writer
	LevelEnablerFunc zap.LevelEnablerFunc
}
