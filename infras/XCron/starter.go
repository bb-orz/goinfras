package XCron

import (
	"GoWebScaffold/infras"
	"fmt"
	"go.uber.org/zap"
)

// 实例化资源存储变量

/* 资源启动器 */
type starter struct {
	infras.BaseStarter
	cfg   *Config
	Tasks []*Task
}

// 应用注册启动器时创建
func NewStarter(tasks ...*Task) *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	starter.Tasks = tasks
	return starter
}

func (s *starter) Name() string {
	return "XCron"
}

// 应用初始化时加载配置数据
func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Cron", &define)
		infras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	sctx.Logger().Info("Print Cron Config:", zap.Any("CronConfig", *define))
	s.cfg = define
}

// 应用安装阶段创建Cron管理器，并注册为应用组件
func (s *starter) Setup(sctx *infras.StarterContext) {
	// 1.创建Cron执行管理器并设置资源
	manager = NewManager(s.cfg, sctx.Logger())
	// 2.创建后可立即注册定时运行任务
	manager.RegisterTasks(s.Tasks...)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(manager)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Cron Manager Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Cron Manager Setup Successful!", s.Name()))
	return true
}

// 应用开始运行时，执行任务
func (s *starter) Start(sctx *infras.StarterContext) {
	// 3.运行定时任务
	manager.RunTasks()
}

// 应用停机时，优雅关闭
func (s *starter) Stop() {
	// 4.关闭定时任务
	manager.StopCron()
}
