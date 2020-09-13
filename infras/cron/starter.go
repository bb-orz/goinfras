package cron

import (
	"GoWebScaffold/infras"
	"go.uber.org/zap"
)

type CronStarter struct {
	infras.BaseStarter
	cfg     *CronConfig
	manager *CronManager
	Tasks   []*Task
}

func (s *CronStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := CronConfig{}
	err := viper.UnmarshalKey("Cron", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *CronStarter) Setup(sctx *infras.StarterContext) {
	// 1.获取Cron执行管理器
	manager := NewCronManager(s.cfg, sctx.Logger())
	// 2.注册定时运行任务
	manager.RegisterTasks(s.Tasks...)
	sctx.Logger().Info("Cron Manager Setup Successful!")
}

func (s *CronStarter) Start(sctx *infras.StarterContext) {
	// 3.运行定时任务
	s.manager.RunTasks()
}

func (s *CronStarter) Stop(sctx *infras.StarterContext) {
	// 4.关闭定时任务
	s.manager.StopCron()
	sctx.Logger().Info("Cron Tasks Stopped!")

}

/*For testing*/
func RunForTesting(config *CronConfig, tasks []*Task) error {
	var err error
	if config == nil {
		config = &CronConfig{Location: "Local"}
	}
	// 1.获取Cron执行管理器
	manager := NewCronManager(config, zap.L())
	// 2.注册定时运行任务
	manager.RegisterTasks(tasks...)

	// 3.运行定时任务
	manager.RunTasks()

	return err
}
