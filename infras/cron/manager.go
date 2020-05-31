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

func NewTask(spec string, job cron.Job) *Task {
	task := new(Task)
	task.spec = spec
	task.job = job
	return task
}

// 定时任务管理器
type CronManager struct {
	client *cron.Cron
	logger *zap.Logger
	tasks  []*Task
}

func NewCronManager(cfg *cronConfig, logger *zap.Logger) *CronManager {
	cronLogger := &LoggerCron{logger: logger}
	location, err := time.LoadLocation(cfg.Location)
	infras.FailHandler(err)

	c := cron.New(
		cron.WithSeconds(),          // 提供秒字段的parser，如无该option秒字段不解析
		cron.WithLocation(location), // 本地时间设置
		cron.WithLogger(cronLogger), // 使用自定义的cron执行日志，在logger_cron.go定义
		cron.WithChain(cron.DelayIfStillRunning(cronLogger), cron.Recover(cronLogger)), // 全局cron执行链
	)

	manager := new(CronManager)
	manager.client = c
	manager.logger = logger
	return manager
}

func (manager *CronManager) RegisterTasks(tasks ...*Task) {
	// Add Schedules and Jobs
	for _, t := range tasks {
		entryID, err := manager.client.AddJob(t.spec, t.job)
		if err != nil {
			manager.logger.Error("[Cron Add Task Error]", zap.Error(err))
		}
		manager.logger.Info("[Cron Add Task]", zap.Int("Entry ID", int(entryID)))
	}
}

func (manager *CronManager) RunTasks() {
	if len(manager.client.Entries()) > 0 {
		manager.client.Start()
		manager.logger.Info("The Cron Jobs Start Up...")
		manager.logger.Info("Cron Entries:", zap.Any("Tasks", manager.client.Entries()))
	} else {
		manager.logger.Info("No Cron Entries")
	}
}

func (manager *CronManager) StopCron() {
	if len(manager.client.Entries()) > 0 {
		manager.client.Stop()
	}
}
