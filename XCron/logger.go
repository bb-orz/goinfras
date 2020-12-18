package XCron

import (
	"github.com/bb-orz/goinfras/XLogger"
	"go.uber.org/zap"
)

/*
如需自定义cron执行日志，可在此编写
*/
type cronLogger struct{}

func NewCronLogger() *cronLogger {
	if XLogger.XCommon() == nil {
		XLogger.CreateDefaultLogger(nil)
	}
	return new(cronLogger)
}

func (l *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	XLogger.XCommon().Info("[Cron Log]:", zap.String("msg", msg), zap.Any("KV", keysAndValues))
}

func (l *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	XLogger.XCommon().Error("[Cron Error]:", zap.String("msg", msg), zap.Any("KV", keysAndValues))

}
