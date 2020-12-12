package XCron

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"goinfras"
	"time"
)

// 实例变量
var manager *Manager

// 创建一个默认配置的Manager
func CreateDefaultManager(config *Config, logger *zap.Logger) {
	if config == nil {
		config = DefaultConfig()
	}
	manager = NewManager(config, logger)
}

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
type Manager struct {
	client *cron.Cron
	logger *zap.Logger
	tasks  []*Task
}

func NewManager(cfg *Config, logger *zap.Logger) *Manager {
	cronLogger := &cronLogger{logger: logger}
	location, err := time.LoadLocation(cfg.Location)
	goinfras.ErrorHandler(err)

	c := cron.New(
		cron.WithSeconds(),          // 提供秒字段的parser，如无该option秒字段不解析
		cron.WithLocation(location), // 本地时间设置
		cron.WithLogger(cronLogger), // 使用自定义的cron执行日志，在logger_cron.go定义
		cron.WithChain(cron.DelayIfStillRunning(cronLogger), cron.Recover(cronLogger)), // 全局cron执行链
	)

	manager := new(Manager)
	manager.client = c
	manager.logger = logger
	return manager
}

// 注册任务
func (m *Manager) RegisterTasks(tasks ...*Task) {
	// Add Schedules and Jobs
	for _, t := range tasks {
		entryID, err := m.client.AddJob(t.spec, t.job)
		if err != nil {
			m.logger.Error("[Cron Add Task Error]", zap.Error(err))
		}
		m.logger.Info("[Cron Add Task]", zap.Int("Entry ID", int(entryID)))
	}
	m.logger.Info("Cron Register Tasks Finish!")
}

// 运行所有任务
func (m *Manager) RunTasks() {
	if len(manager.client.Entries()) > 0 {
		m.client.Start()
		m.logger.Info("The Cron Tasks Running...")
		m.logger.Info("Cron Entries:", zap.Any("Tasks", manager.client.Entries()))
	} else {
		m.logger.Info("No Cron Entries")
	}
}

// 停止所有任务
func (m *Manager) StopCron() {
	entries := m.client.Entries()
	if len(entries) > 0 {
		m.client.Stop()
		for _, e := range entries {
			m.client.Remove(e.ID)
		}
	}
}

// 重启所有任务
func (m *Manager) RestartCron() {
	m.logger.Info("The Cron Tasks Stopping...")
	m.StopCron()
	m.logger.Info("The Cron Tasks Restarting...")
	m.RunTasks()
}
