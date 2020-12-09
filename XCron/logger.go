package XCron

import (
	"go.uber.org/zap"
)

/*
如需自定义cron执行日志，可在此编写
*/
type cronLogger struct {
	logger *zap.Logger
}

func (l *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info("[Cron Log]:", zap.String("msg", msg), zap.Any("KV", keysAndValues))
}

func (l *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error("[Cron Error]:", zap.String("msg", msg), zap.Any("KV", keysAndValues))

}
