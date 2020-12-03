package cron

import (
	"GoWebScaffold/infras"
)

// 实例化资源存储变量
var manager *Manager

/* 资源启动器 */
type Starter struct {
	infras.BaseStarter
	cfg   Config
	Tasks []*Task
}

// 应用注册启动器时创建
func NewStarter() *Starter {
	starter := new(Starter)
	starter.cfg = Config{}
	return starter
}

// 应用初始化时加载配置数据
func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Cron", &define)
	infras.FailHandler(err)
	s.cfg = define
}

// 应用安装阶段创建Cron管理器，并注册为应用组件
func (s *Starter) Setup(sctx *infras.StarterContext) {
	// 1.创建Cron执行管理器并设置资源
	manager = NewManager(&s.cfg, sctx.Logger())

	// 2.创建后可立即注册定时运行任务
	manager.RegisterTasks(s.Tasks...)
	sctx.Logger().Info("Cron Manager Setup Successful!")
}

// 应用开始运行时，执行任务
func (s *Starter) Start(sctx *infras.StarterContext) {
	// 3.运行定时任务
	manager.RunTasks()
}

// 应用停机时，优雅关闭
func (s *Starter) Stop(sctx *infras.StarterContext) {
	// 4.关闭定时任务
	manager.StopCron()
	sctx.Logger().Info("Cron Tasks Stopped!")

}
