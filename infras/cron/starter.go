package cron

import (
	"GoWebScaffold/infras"
	"fmt"
	"go.uber.org/zap"
)

var manager *Manager

// 任务管理器运行示例调用
func RuntimeManager() *Manager {
	infras.Check(manager)
	return manager
}

type Starter struct {
	infras.BaseStarter
	cfg     Config
	manager *Manager
	Tasks   []*Task
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Cron", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	// 1.获取Cron执行管理器
	manager := NewManager(&s.cfg, sctx.Logger())
	// 2.注册定时运行任务
	manager.RegisterTasks(s.Tasks...)
	sctx.Logger().Info("Cron Manager Setup Successful!")
}

func (s *Starter) Start(sctx *infras.StarterContext) {
	// 3.运行定时任务
	s.manager.RunTasks()
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	// 4.关闭定时任务
	s.manager.StopCron()
	sctx.Logger().Info("Cron Tasks Stopped!")

}

/*For testing*/
func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{Location: "Local"}
	}
	// 1.获取Cron执行管理器
	fmt.Println("创建任务执行管理器...")
	manager = NewManager(config, zap.L())
	return err
}
