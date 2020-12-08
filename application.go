package goinfras

import (
	"fmt"
	"github.com/spf13/viper"
)

// 应用程序启动管理器
type Application struct {
	Sctx *StarterContext // 应用启动器上下文
}

// 创建应用程序启动管理器
func NewApplication(vpcfg *viper.Viper) *Application {
	// 创建启动管理器
	app := new(Application)
	app.Sctx = &StarterContext{}
	app.Sctx.SetConfigs(vpcfg)
	return app
}

// 启动应用程序所有基础资源 （初始化 -> 安装 -> 启动）
func (app *Application) Up() {
	app.init()
	app.setup()
	app.check()
	app.start()
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

	// 注册资源组件的关闭回调
	RegisterStopFunc(app.Sctx.Logger())
}

// 检查组件实例
func (app *Application) check() {
	for _, s := range StarterManager.GetAll() {
		if !s.Check(app.Sctx) {
			fmt.Printf("%s Starter Setup Fail：An error was found during the check! ", s.Name())
		} else {
			fmt.Printf("%s Starter：Setup Successful!", s.Name())
		}
	}
}

func (app *Application) start() {
	for i, s := range StarterManager.GetAll() {
		if s.SetStartBlocking() { // 阻塞启动
			// 如果是最后一个可阻塞的，直接启动并阻塞
			if i == StarterManager.Len()-1 {
				s.Start(app.Sctx)
			} else {
				// 如果不是，使用goroutine来异步启动，防止阻塞后面starter
				go s.Start(app.Sctx)
			}
		} else { // 非阻塞启动
			s.Start(app.Sctx)
		}
	}

	// 应用启动时开始监听系统信号
	NotifySignal(app.Sctx.Logger())
}
