package cron

import "GoWebScaffold/infras"

/*资源组件化调用*/

var manager *Manager

// 设置组件资源
func SetComponent(m *Manager) {
	manager = m
}

// 组件化使用
func CronComponent() *Manager {
	infras.Check(manager)
	return manager
}
