package goinfras

import (
	"github.com/spf13/viper"
	"io"
)

// 应用运行实例
var appRuntime *Application

// 获取应用实例的方法,通过App实例可在项目内部获取启动器上下文，包括启动配置信息等
func XApp() *Application {
	return appRuntime
}

// 应用程序启动管理器
type Application struct {
	Sctx *StarterContext // 应用启动器上下文
}

// 创建应用程序启动管理器
func NewApplication(vpcfg *viper.Viper) *Application {
	// 创建启动管理器
	appRuntime = new(Application)
	appRuntime.Sctx = &StarterContext{}
	appRuntime.Sctx.SetConfigs(vpcfg)
	global := NewGlobal(vpcfg)
	appRuntime.Sctx.SetGlobal(global)
	env := global.GetEnv()
	appRuntime.Sctx.SetLogger(NewCommandLineStarterLogger(env))
	return appRuntime
}

// 创建一个带输出启动日志的应用管理器
func NewApplicationWithStarterLoggerWriter(vpcfg *viper.Viper, logWriters ...io.Writer) *Application {
	// 创建启动管理器
	appRuntime = new(Application)
	appRuntime.Sctx = &StarterContext{}
	appRuntime.Sctx.SetConfigs(vpcfg)
	global := NewGlobal(vpcfg)
	appRuntime.Sctx.SetGlobal(global)
	env := global.GetEnv()
	appRuntime.Sctx.SetLogger(NewStarterLoggerWithWriters(env, logWriters...))
	return appRuntime
}

// 启动应用程序所有基础资源 （初始化 -> 安装 -> 检查 -> 启动 -> 监听系统退出信号）
func (app *Application) Up() {
	app.init()         // 加载所有注册启动器配置
	app.setup()        // 安装所有注册启动器组件
	app.check()        // 检查所有注册组件
	app.start()        // 启动组件实例
	app.listenSignal() // 监听退出信号，实现优雅关闭所有启动器 ,阻塞
}

// 停止或销毁应用程序所有基础资源
func (app *Application) Down() {
	for _, s := range StarterManager.GetAll() {
		s.Stop()
	}
}

func (app *Application) init() {
	for _, s := range StarterManager.GetAll() {
		s.Init(app.Sctx)
	}
}

func (app *Application) setup() {
	// 安装所有注册启动器组件
	for _, s := range StarterManager.GetAll() {
		s.Setup(app.Sctx)
	}
}

// 检查组件实例
func (app *Application) check() {
	for _, s := range StarterManager.GetAll() {
		s.Check(app.Sctx)
	}
}

func (app *Application) start() {
	for _, s := range StarterManager.GetAll() {
		if s.StartBlocking() { // 阻塞的starter另开go程启动
			go s.Start(app.Sctx)
		} else { // 非阻塞启动
			s.Start(app.Sctx)
		}
	}
}

func (app *Application) listenSignal() {
	// 应用启动时开始监听系统信号
	NotifySignal(app.Sctx.Logger())
}
