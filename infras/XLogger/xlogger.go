package XLogger

import (
	"go.uber.org/zap"
)

var commonLogger *zap.Logger
var syncErrorLogger *zap.Logger

func XCommon() *zap.Logger {
	return commonLogger
}

func XSyncError() *zap.Logger {
	return syncErrorLogger
}
