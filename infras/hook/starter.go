package hook

import (
	"GoWebScaffold/infras"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

/*
应用运行监听系统信号钩子组件，该组件运行后可让系统资源连接随应用一起优雅退出。
*/

var callbacks []func()

func Register(fn func()) {
	callbacks = append(callbacks, fn)
}

type Starter struct {
	infras.BaseStarter
}

func NewStarter() *Starter {
	starter := new(Starter)
	return starter
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			c := <-sigs
			sctx.Logger().Info("System signal notify:", zap.String("signal", c.String()))
			for _, fn := range callbacks {
				fn()
			}
			break
			os.Exit(0)
		}
	}()

}

func (s *Starter) Start(sctx *infras.StarterContext) {
	starters := infras.StarterManager.GetAll()

	for _, s := range starters {
		typ := reflect.TypeOf(s)
		sctx.Logger().Info("【Register Notify Stop】:%s.Stop()", zap.String("Resource Component", typ.String()))
		Register(func() {
			s.Stop(sctx)
		})
	}
}

// 默认设置优先级为最末位启动
func (s *Starter) Priority() int { return infras.INT_MIN }
