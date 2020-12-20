package XLogger

import (
	"go.uber.org/zap/zapcore"
)

// 异步非错误信息(debug/info/warn)日志记录器
func NewUserOutputsCore(outputs ...LoggerOutput) zapcore.Core {
	var zCoreList []zapcore.Core

	for _, output := range outputs {
		zCore := zapcore.NewCore(
			// 日志格式配置
			zapcore.NewJSONEncoder(output.Format),
			// 日志异步输出配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(output.Writer)),
			// 日志级别
			output.LevelEnablerFunc,
		)
		zCoreList = append(zCoreList, zCore)
	}

	return zapcore.NewTee(zCoreList...)
}
