package XLogger

import (
	"go.uber.org/zap"
	"io"
)

func XCommon() *zap.Logger {
	return commonLogger
}

func XSyncError() *zap.Logger {
	return syncErrorLogger
}
