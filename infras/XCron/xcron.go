package XCron

import (
	"fmt"
	"go.uber.org/zap"
)

// 实例变量
var manager *Manager

// 资源组件实例调用
func XManager() *Manager {
	return manager
}

// 资源组件闭包执行
func XFManager(f func(m *Manager) error) error {
	return f(manager)
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{Location: "Local"}
	}
	// 1.获取Cron执行管理器
	fmt.Println("创建任务执行管理器...")
	manager = NewManager(config, zap.L())
	return err
}
