package ginger

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

func GinEngine() *gin.Engine {
	infras.Check(ginEngine)
	return ginEngine
}

type GinStarter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *GinStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Gin", &define)
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
	if !s.cfg.Cors.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(s.cfg.Cors))
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
	if s.cfg.Tls && s.cfg.CertFile != "" && s.cfg.KeyFile != "" {
		err = GinEngine().RunTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
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

/*For testing*/
func RunForTesting(config *Config, apis []IApiModule) error {
	var err error
	if config == nil {
		config = &Config{
			ListenHost: "127.0.0.1",
			ListenPort: 8090,
		}

	}

	// 1.配置gin中间件
	log := logger.CommonLogger()
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(log), ZapRecoveryMiddleware(log, false))

	// 如开启cors限制，添加中间件
	if !config.Cors.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(config.Cors))
	}

	// 2.New Gin Engine
	ginEngine = NewGinEngine(config, middlewares...)

	// 3.服务API 模块路由注册
	for _, v := range apis {
		v.SetRoutes()
	}

	// 4.启动
	var addr string
	addr = fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort)
	if config.Tls && config.CertFile != "" && config.KeyFile != "" {
		err = GinEngine().RunTLS(addr, config.CertFile, config.KeyFile)
		infras.FailHandler(err)
	} else {
		err = GinEngine().Run(addr)
		infras.FailHandler(err)
	}

	return err
}
