package XLogger

import (
	"GoWebScaffold/infras"
	"go.uber.org/zap"
)

var commonLogger *zap.Logger
var syncErrorLogger *zap.Logger

func CLogger() *zap.Logger {
	infras.Check(commonLogger)
	return commonLogger
}

func SELogger() *zap.Logger {
	infras.Check(syncErrorLogger)
	return syncErrorLogger
}

func SetComponentForCommonLogger(l *zap.Logger) {
	commonLogger = l
}

func SetComponentForSyncErrorLogger(l *zap.Logger) {
	syncErrorLogger = l
}
