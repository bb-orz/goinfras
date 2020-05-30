package cron

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
)

type CronStarter struct {
	infras.BaseStarter
	cfg   *cronConfig
	Tasks []*Task
}

func (s *CronStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := cronConfig{}
	err := kvs.Unmarshal(configs, &define, "Cron")
	infras.FailHandler(err)
	s.cfg = &define

	// 注册定时任务
	RegisterTask(s.Tasks...)
}

func (s *CronStarter) Setup(sctx *infras.StarterContext) {}

func (s *CronStarter) Start(sctx *infras.StarterContext) {
	Do(s.cfg, sctx.Logger())
}

func (s *CronStarter) Stop(sctx *infras.StarterContext) {
}
