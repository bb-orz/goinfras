package XCron

import (
	"github.com/robfig/cron/v3"
	"time"
)

// 实例变量
var manager *Manager

// 创建一个默认配置的Manager
func CreateDefaultManager(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}
	manager, err = NewManager(config)
	return err
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
	tasks  []*Task
}

func NewManager(cfg *Config) (*Manager, error) {
	cronLogger := &cronLogger{}
	location, err := time.LoadLocation(cfg.Location)
	if err != nil {
		return nil, err
	}

	c := cron.New(
		cron.WithSeconds(),          // 提供秒字段的parser，如无该option秒字段不解析
		cron.WithLocation(location), // 本地时间设置
		cron.WithLogger(cronLogger), // 使用自定义的cron执行日志，在logger_cron.go定义
		cron.WithChain(cron.DelayIfStillRunning(cronLogger), cron.Recover(cronLogger)), // 全局cron执行链
	)

	manager := new(Manager)
	manager.client = c
	return manager, nil
}

// 注册任务
func (m *Manager) RegisterTasks(tasks ...*Task) ([]cron.EntryID, error) {
	var err error
	var entryIDs = make([]cron.EntryID, 0)
	// Add Schedules and Jobs
	for _, t := range tasks {
		var entryID cron.EntryID
		entryID, err = m.client.AddJob(t.spec, t.job)
		if err != nil {
			return entryIDs, err
		}
		entryIDs = append(entryIDs, entryID)
	}
	return entryIDs, nil
}

// 运行所有任务
func (m *Manager) RunTasks() {
	if len(manager.client.Entries()) > 0 {
		m.client.Start()
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
	m.StopCron()
	m.RunTasks()
}
