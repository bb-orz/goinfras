package XLogger

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"time"
)

// 归档记录核心
func NewRotateLogCore(cfg *Config) (zapcore.Core, error) {
	// 归档实例
	rotateWriter, err := rotateLogs.New(
		cfg.RotateLogDir+cfg.RotateLogBaseName+"[%Y-%m-%d %H:%M:%S].log",
		// 最多保留多久
		rotateLogs.WithMaxAge(time.Hour*time.Duration(cfg.MaxDayCount*24)),
		// 多久做一次归档
		rotateLogs.WithRotationTime(time.Hour*time.Duration(cfg.WithRotationTime*24)),
	)
	if err != nil {
		return nil, err
	}

	// 返回归档记录核心
	return zapcore.NewCore(
		// 日志格式配置
		zapcore.NewJSONEncoder(defaultFormatConfig()),
		// 日志异步输出配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(rotateWriter)),
		// 日志级别
		SettingLevelEnableFunc(cfg),
	), nil
}
