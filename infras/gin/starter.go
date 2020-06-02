package gin

import (
	"GoWebScaffold/infras"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
)

type GinStarter struct {
	infras.BaseStarter
	cfg *ginConfig
}

func (s *GinStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := ginConfig{}
	err := kvs.Unmarshal(configs, &define, "Gin")
	infras.FailHandler(err)
	s.cfg = &define
}

// 注册服务API
func (s *GinStarter) Setup(sctx *infras.StarterContext) {
	// 服务API注册后运行初始化
	for _, v := range GetApiInitializers() {
		v.Init()
	}
}

// 启动运行gin service
func (s *GinStarter) Start(sctx *infras.StarterContext) {

	err := runGinServer(s.cfg, zap.L())
	infras.FailHandler(err)
}

func (s *GinStarter) SetStartBlocking() bool {
	return true
}

func (s *GinStarter) Stop(sctx *infras.StarterContext) {
}
