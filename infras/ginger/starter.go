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

func SetGinEngine(engine *gin.Engine) {
	ginEngine = engine
}

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

// 初始化时：加载配置
func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Gin", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

// 启动时：添加中间件，实例化应用，注册项目实现的API
func (s *Starter) Setup(sctx *infras.StarterContext) {

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

	// 3.API 路由注册
	for _, v := range GetApis() {
		v.SetRoutes()
	}
}

// 启动时：运行gin engine
func (s *Starter) Start(sctx *infras.StarterContext) {
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

func (s *Starter) SetStartBlocking() bool {
	return true
}

func (s *Starter) Stop(sctx *infras.StarterContext) {

}

/*For testing*/
func RunForTesting(config *Config, apis []IApi) error {
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

	// 3.Restful API 模块注册
	for _, v := range apis {
		// 路由注册
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
