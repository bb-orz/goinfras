package XCron

import (
	"fmt"
	"github.com/bb-orz/goinfras"
)

// 实例化资源存储变量

/* 资源启动器 */
type starter struct {
	goinfras.BaseStarter
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
func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Cron", &define)
		sctx.PassWarning(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().SDebug(fmt.Sprintf("[XCron Init] Config: %v \n", *define))
}

// 应用安装阶段创建Cron管理器，并注册为应用组件
func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	// 1.创建Cron执行管理器并设置资源
	manager, err = NewManager(s.cfg)
	// 2.创建后可立即注册定时运行任务
	entryIDS, err := manager.RegisterTasks(s.Tasks...)
	if sctx.PassError(err) {
		sctx.Logger().SInfo(fmt.Sprintf("[XCron Setup] Tasks EntryIDs: %v \n", entryIDS))
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(manager)
	if sctx.PassError(err) {
		sctx.Logger().SInfo(fmt.Sprintf("[XCron Check]: Cron Manager Setup Successful! \n"))
	}
	return false
}

// 应用开始运行时，执行任务
func (s *starter) Start(sctx *goinfras.StarterContext) {
	// 3.运行定时任务
	manager.RunTasks()
	sctx.Logger().SInfo(fmt.Sprintf("[XCron Start]: Cron Running Tasks... \n"))
}

// 应用停机时，优雅关闭
func (s *starter) Stop() {
	// 4.关闭定时任务
	manager.StopCron()
}

// 设置启动组级别:
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
