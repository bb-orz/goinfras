package XCron

import "GoWebScaffold/infras"

// 资源组件实例调用
func X() (*Manager, error) {
	err := infras.Check(manager)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

// 资源组件闭包执行
func XF(f func(m *Manager) error) error {
	err := infras.Check(manager)
	if err != nil {
		return err
	}
	return f(manager)
}
