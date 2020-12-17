package goinfras

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

/*
应用运行监听系统信号钩子，该组件运行后可让系统资源连接随应用一起优雅退出。
*/

var callbacks []func()

func Register(fn func()) {
	callbacks = append(callbacks, fn)
}

// 应用安装时注册组件关闭函数
func RegisterStarterStopFunc(logger *zap.Logger) {
	starters := StarterManager.GetAll()
	for _, s := range starters {
		Register(func() {
			s.Stop()
		})
		logger.Info(fmt.Sprintf("【%s Starter】: Stop Function Registered.", s.Name()))
	}
}

// 应用启动时监听系统信号：停止和退出时只需关闭回调
func NotifySignal(logger *zap.Logger) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		c := <-sigs
		logger.Info("System signal notify:", zap.String("signal", c.String()))
		for _, fn := range callbacks {
			fn()
		}
		os.Exit(0)
	}
}
