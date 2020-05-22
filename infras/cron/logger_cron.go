package cron

import "fmt"

/*
如需自定义cron执行日志，可在此编写
*/
type CronLogger struct {}

func (logger *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	fmt.Println("MSG:",msg," KV:",keysAndValues)
}

func (logger *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	fmt.Println("Error:",err.Error()," MSG:",msg," KV:",keysAndValues)
}
