package gin

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/jwt"
	"github.com/gin-gonic/gin"
	"github.com/tietang/props/kvs"
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
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(sctx.Logger()), ZapRecoveryMiddleware(sctx.Logger(), false))

	// 如开启cors限制，添加中间件
	if !s.cfg.cors.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(s.cfg.cors))
	}

	// 如TokenUtils服务已初始化，添加中间件
	if tku := jwt.TokenUtils(); tku != nil {
		middlewares = append(middlewares, JwtAuthMiddleware(tku))
	}

	err := GinServerRun(s.cfg, sctx.Logger(), middlewares...)
	infras.FailHandler(err)
}

func (s *GinStarter) SetStartBlocking() bool {
	return true
}

func (s *GinStarter) Stop(sctx *infras.StarterContext) {
}
