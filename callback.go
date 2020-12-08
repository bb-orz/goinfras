package goinfras

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"reflect"
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
func RegisterStopFunc(logger *zap.Logger) {
	starters := StarterManager.GetAll()

	for _, s := range starters {
		typ := reflect.TypeOf(s)
		logger.Info("【Register Notify Stop】:%s.Stop()", zap.String("Resource Component", typ.String()))
		Register(func() {
			s.Stop()
		})
	}
}

// 应用启动时监听系统信号：停止和退出时只需关闭回调
func NotifySignal(logger *zap.Logger) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			c := <-sigs
			logger.Info("System signal notify:", zap.String("signal", c.String()))
			for _, fn := range callbacks {
				fn()
			}
			break
			os.Exit(0)
		}
	}()

}
