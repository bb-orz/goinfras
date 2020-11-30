package cron

import (
	"GoWebScaffold/infras"
)

/* 资源启动器 */

type Starter struct {
	infras.BaseStarter
	cfg   Config
	Tasks []*Task
}

func NewStarter() *Starter {
	starter := new(Starter)
	starter.cfg = Config{}
	return starter
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Cron", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var m *Manager
	// 1.创建Cron执行管理器并设置资源
	m = NewManager(&s.cfg, sctx.Logger())
	SetComponent(m)

	// 2.创建后可立即注册定时运行任务
	CronComponent().RegisterTasks(s.Tasks...)
	sctx.Logger().Info("Cron Manager Setup Successful!")
}

func (s *Starter) Start(sctx *infras.StarterContext) {
	// 3.运行定时任务
	CronComponent().RunTasks()
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	// 4.关闭定时任务
	CronComponent().StopCron()
	sctx.Logger().Info("Cron Tasks Stopped!")

}
