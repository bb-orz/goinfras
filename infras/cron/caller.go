package cron

import "GoWebScaffold/infras"

// 闭包执行
func Execute(f func(m *Manager) error) error {
	return f(manager)
}

// 提供与包同名的函数，实例调用
func Cron() (*Manager, error) {
	err := infras.Check(manager)
	if err != nil {
		return nil, err
	}

	return manager, nil
}
