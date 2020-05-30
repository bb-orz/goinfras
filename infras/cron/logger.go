package cron

import (
	"go.uber.org/zap"
)

/*
如需自定义cron执行日志，可在此编写
*/
type LoggerCron struct {
	logger *zap.Logger
}

func (l *LoggerCron) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info("[Cron Log]:", zap.String("msg", msg), zap.Any("KV", keysAndValues))
}

func (l *LoggerCron) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error("[Cron Error]:", zap.String("msg", msg), zap.Any("KV", keysAndValues))

}
