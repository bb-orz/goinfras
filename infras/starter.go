package infras

import "GoWebScaffold/infras/starter"

// 启动器接口
type Starter interface {
	// 初始化基础资源
	Init(*starter.StarterContext)
	// 安装
	Setup(*starter.StarterContext)
	// 启动
	Start(*starter.StarterContext)
	// 阻塞启动
	SetStartBlocking() bool
	// 停止或销毁基础资源
	Stop()
}

// 基础空启动器，可便于被其他具体的基础资源嵌入
type BaseStarter struct{}

func (*BaseStarter) Init(*starter.StarterContext) {}

func (*BaseStarter) Setup(*starter.StarterContext) {}

func (*BaseStarter) Start(*starter.StarterContext) {}

func (*BaseStarter) SetStartBlocking() bool { return false }

func (*BaseStarter) Stop() {}

// 接口实现检查
var _ Starter = new(BaseStarter)

// 启动器注册管理器
type starterManager struct {
	starters []Starter
}

// 注册
func (m *starterManager) Register(s Starter) {
	m.starters = append(m.starters, s)
}

// 返回所有已注册的启动器
func (m *starterManager) GetAll() []Starter {
	return m.starters
}

func (m *starterManager) Len() int {
	return len(m.starters)
}

var StarterManager = new(starterManager)

// 开放启动注册器
func Register(s Starter) {
	StarterManager.Register(s)
}
