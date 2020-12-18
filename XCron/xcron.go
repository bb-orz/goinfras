package XCron

// 资源组件实例调用
func XManager() *Manager {
	return manager
}

// 资源组件闭包执行
func XFManager(f func(m *Manager) error) error {
	return f(manager)
}
