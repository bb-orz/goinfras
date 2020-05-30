package hook

import (
	"GoWebScaffold/infras"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

var callbacks []func()

func Register(fn func()) {
	callbacks = append(callbacks, fn)
}

type HookStarter struct {
	infras.BaseStarter
}

func (s *HookStarter) Setup(ctx *infras.StarterContext) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			c := <-sigs
			// TODO log something...
			fmt.Printf("notify: ", c)
			for _, fn := range callbacks {
				fn()
			}
			break
			os.Exit(0)
		}
	}()

}

func (s *HookStarter) Start(ctx *infras.StarterContext) {
	starters := infras.StarterManager.GetAll()

	for _, s := range starters {
		typ := reflect.TypeOf(s)
		// TODO log something...
		fmt.Printf("【Register Notify Stop】:%s.Stop()", typ.String())
		Register(func() {
			s.Stop(ctx)
		})
	}
}

// 优先级为最末位启动
func (s *HookStarter) Priority() int { return infras.INT_MIN }
