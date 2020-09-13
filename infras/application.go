package infras

import (
	"github.com/spf13/viper"
)

// 应用程序启动管理器
type Application struct {
	Sctx *StarterContext // 应用启动器上下文
}

// 创建应用程序启动管理器
func NewApplication(vpcfg *viper.Viper) *Application {
	// 创建启动管理器
	b := new(Application)
	b.Sctx = &StarterContext{}
	b.Sctx.SetConfigs(vpcfg)
	return b
}

// 启动应用程序所有基础资源 （初始化 -> 安装 -> 启动）
func (b *Application) Up() {
	b.init()
	b.setup()
	b.start()
}

// 停止或销毁应用程序所有基础资源
func (b *Application) Down() {
	for _, s := range StarterManager.GetAll() {
		s.Stop(b.Sctx)
	}
}

func (b *Application) init() {
	for _, s := range StarterManager.GetAll() {
		s.Init(b.Sctx)
	}
}

func (b *Application) setup() {
	for _, s := range StarterManager.GetAll() {
		s.Setup(b.Sctx)
	}
}

func (b *Application) start() {
	for i, s := range StarterManager.GetAll() {
		if s.SetStartBlocking() { // 阻塞启动
			// 如果是最后一个可阻塞的，直接启动并阻塞
			if i == StarterManager.Len()-1 {
				s.Start(b.Sctx)
			} else {
				// 如果不是，使用goroutine来异步启动，防止阻塞后面starter
				go s.Start(b.Sctx)
			}
		} else { // 非阻塞启动
			s.Start(b.Sctx)
		}
	}
}
