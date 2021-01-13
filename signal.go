package goinfras

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
应用运行监听系统信号钩子，该组件运行后可让系统资源连接随应用一起优雅退出。
*/

var callbacks = make(map[string]func())

func Register(name string, fn func()) {
	callbacks[name] = fn
}

// 应用启动时监听系统信号：停止和退出时只需关闭回调
func NotifySignal(logger IStarterLogger) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	for {
		c := <-sigs
		logger.Info("Application", StepStop, fmt.Sprintf("System signal notify: %s ", c.String()))

		starters := StarterManager.GetAll()
		for _, s := range starters {
			err := s.Stop()
			if err != nil {
				logger.Warning("Application", StepStop, fmt.Sprintf("%s Starter Stop Error:%s", s.Name(), err.Error()))
			}
			logger.OK("Application", StepStop, fmt.Sprintf("%s Starter Stopped ", s.Name()))
		}

		os.Exit(0)
	}
}
