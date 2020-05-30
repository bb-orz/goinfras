package cron

import (
	"GoWebScaffold/infras"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"
)

type Task struct {
	spec string
	job  cron.Job
}

var taskList []*Task

func RegisterTask(t ...*Task) {
	taskList = append(taskList, t...)
}

func Do(cfg *cronConfig, logger *zap.Logger) {
	cronLogger := &LoggerCron{logger: logger}
	location, err := time.LoadLocation(cfg.Location)
	infras.FailHandler(err)

	c := cron.New(
		cron.WithSeconds(),          // 提供秒字段的parser，如无该option秒字段不解析
		cron.WithLocation(location), // 本地时间设置
		cron.WithLogger(cronLogger), // 使用自定义的cron执行日志，在logger_cron.go定义
		cron.WithChain(cron.DelayIfStillRunning(cronLogger), cron.Recover(cronLogger)), // 全局cron执行链
	)

	// Add Schedules and Jobs
	for _, t := range taskList {
		entryID, err := c.AddJob(t.spec, t.job)
		if infras.ErrorHandler(err) {
			logger.Info("[Cron Add Task]", zap.Int("Entry ID", int(entryID)))
		}
	}

	if len(c.Entries()) > 0 {
		c.Start()
		logger.Info("The Cron Jobs Running")
	} else {
		logger.Info("No Cron Entries")
	}

}
