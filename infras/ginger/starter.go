package ginger

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tietang/props/kvs"
)

var ginEngine *gin.Engine

func GinEngine() *gin.Engine {
	infras.Check(ginEngine)
	return ginEngine
}

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

	// 1.配置gin中间件
	log := logger.CommonLogger()
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(log), ZapRecoveryMiddleware(log, false))

	// 如开启cors限制，添加中间件
	if !s.cfg.cors.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(s.cfg.cors))
	}

	// 2.New Gin Engine
	ginEngine = NewGinEngine(s.cfg, middlewares...)

	// 3.服务API 模块路由注册
	for _, v := range GetApiModules() {
		v.SetRoutes()
	}
}

// 启动运行gin service
func (s *GinStarter) Start(sctx *infras.StarterContext) {
	var addr string
	var err error
	addr = fmt.Sprintf("%s:%d", s.cfg.ListenHost, s.cfg.ListenPort)
	if s.cfg.tls && s.cfg.certFile != "" && s.cfg.keyFile != "" {
		err = GinEngine().RunTLS(addr, s.cfg.certFile, s.cfg.keyFile)
		infras.FailHandler(err)
	} else {
		err = GinEngine().Run(addr)
		infras.FailHandler(err)
	}
}

func (s *GinStarter) SetStartBlocking() bool {
	return true
}

func (s *GinStarter) Stop(sctx *infras.StarterContext) {
}
