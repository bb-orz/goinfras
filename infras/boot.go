package infras

import (
	"github.com/tietang/props/kvs"
)

// 应用程序启动管理器
type Boot struct {
	Sctx *StarterContext    // 应用启动器上下文
}

// 创建应用程序启动管理器
func NewBoot(conf kvs.ConfigSource) *Boot {
	// 创建启动管理器
	b := new(Boot)
	b.Sctx = &StarterContext{}
	b.Sctx.SetConfigs(conf)
	return b
}

// 启动应用程序所有基础资源 （初始化 -> 安装 -> 启动）
func (b *Boot) Up() {
	b.init()
	b.setup()
	b.start()
}

// 停止或销毁应用程序所有基础资源
func (b *Boot) Down() {
	for _, s := range StarterManager.GetAll() {
		s.Stop(b.Sctx)
	}
}

func (b *Boot) init() {
	for _, s := range StarterManager.GetAll() {
		s.Init(b.Sctx)
	}
}

func (b *Boot) setup() {
	for _, s := range StarterManager.GetAll() {
		s.Setup(b.Sctx)
	}
}

func (b *Boot) start() {
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
