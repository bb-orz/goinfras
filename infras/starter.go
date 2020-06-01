package infras

import (
	"sort"
)

// 启动器接口
type Starter interface {
	// 初始化：资源组件读取配置信息
	Init(*StarterContext)
	// 安装：检查该组件的前置依赖
	Setup(*StarterContext)
	// 启动：该资源组件的连接或启动以供应用程序后续使用
	Start(ctx *StarterContext)
	// 阻塞启动：设置需要后置启动的资源组件，默认为false
	SetStartBlocking() bool
	//资源停止：
	// 通常在启动时遇到异常时或者启用远程管理时，用于释放资源和终止资源的使用，
	// 通常要优雅的释放，等待正在进行的任务继续，但不再接受新的任务
	Stop(ctx *StarterContext)
	// 优先组：从高到低分：系统级别、基本资源级别、应用级别三组
	PriorityGroup() PriorityGroup
	// 设置该资源组件的启动优先级，默认为DEFAULT_PRIORITY，最大为INT_MAX
	Priority() int
}

// 启动组件顺序的优先组
type PriorityGroup int

const (
	SystemGroup         PriorityGroup = 30 // 系统级别优先组
	BasicResourcesGroup PriorityGroup = 20 // 基本资源级别优先组
	AppGroup            PriorityGroup = 10 // 应用级别优先组

	INT_MAX          = int(^uint(0) >> 1) // 最优先启动级别
	INT_MIN          = 0                  // 最末位启动级别
	DEFAULT_PRIORITY = 10000              // 默认
)

// 基础空启动器，可便于被其他具体的基础资源嵌入
type BaseStarter struct{}

// 载入启动器配置信息
func (*BaseStarter) Init(*StarterContext) {}

// 安装作为系统资源组件
func (*BaseStarter) Setup(*StarterContext) {}

// 如需随应用运行需Start启动
func (*BaseStarter) Start(*StarterContext) {}

// 阻塞型组件需设置为true
func (*BaseStarter) SetStartBlocking() bool { return false }

// 关闭应用时的资源组件清理工作
func (*BaseStarter) Stop(*StarterContext) {}

// 为资源组件设置优先组
func (s *BaseStarter) PriorityGroup() PriorityGroup { return BasicResourcesGroup }

// 为资源组件设置启动优先值
func (s *BaseStarter) Priority() int { return DEFAULT_PRIORITY }

// 接口实现检查
var _ Starter = new(BaseStarter)

// 启动器注册管理器
type starterManager struct {
	starters []Starter
}

// 注册启动器
func (m *starterManager) Register(s Starter) {
	m.starters = append(m.starters, s)
}

// 返回所有已注册的启动器
func (m *starterManager) GetAll() []Starter {
	return m.starters
}

// 实现可排序接口
func (m *starterManager) Len() int {
	return len(m.starters)
}
func (m starterManager) Swap(i, j int) { m.starters[i], m.starters[j] = m.starters[j], m.starters[i] }
func (m starterManager) Less(i, j int) bool {
	return m.starters[i].PriorityGroup() > m.starters[j].PriorityGroup() && m.starters[i].Priority() > m.starters[j].Priority()
}

var StarterManager = new(starterManager)

// 开放启动注册器
func Register(s Starter) {
	StarterManager.Register(s)
}

// 组件启动器排序
func SortStarters() {
	sort.Sort(StarterManager)
}
