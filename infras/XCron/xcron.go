package XCron

import "go.uber.org/zap"

// 实例变量
var manager *Manager

// 创建一个默认配置的Manager
func CreateDefaultManager(config *Config, logger *zap.Logger) {
	if config == nil {
		config = DefaultConfig()
	}
	manager = NewManager(config, logger)
}

// 资源组件实例调用
func XManager() *Manager {
	return manager
}

// 资源组件闭包执行
func XFManager(f func(m *Manager) error) error {
	return f(manager)
}
