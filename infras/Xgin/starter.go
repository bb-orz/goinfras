package Xgin

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/XLogger"
	"GoWebScaffold/infras/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *Starter {
	starter := new(Starter)
	starter.cfg = Config{}
	return starter
}

// 初始化时：加载配置
func (s *Starter) Init(sctx *infras.StarterContext) {
	var err error
	viper := sctx.Configs()
	ginDefine := GinConfig{}
	err = viper.UnmarshalKey("Gin", &ginDefine)
	infras.FailHandler(err)
	s.cfg.GinConfig = ginDefine

	corsDefine := CorsConfig{}
	err = viper.UnmarshalKey("Cors", &corsDefine)
	infras.FailHandler(err)
	s.cfg.CorsConfig = corsDefine
}

// 启动时：添加中间件，实例化应用，注册项目实现的API
func (s *Starter) Setup(sctx *infras.StarterContext) {
	var engine *gin.Engine
	// 1.配置gin中间件
	log := XLogger.CLogger()
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(log), ZapRecoveryMiddleware(log, false))

	// 如开启cors限制，添加中间件
	if !s.cfg.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(&s.cfg.CorsConfig))
	}

	// 2.New Gin Engine
	engine = NewGinEngine(&s.cfg, middlewares...)
	SetComponent(engine)

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
		err = GinComponent().RunTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
		infras.FailHandler(err)
	} else {
		err = GinComponent().Run(addr)
		infras.FailHandler(err)
	}
}

func (s *Starter) Stop() {}

// 默认设置阻塞启动
func (s *Starter) SetStartBlocking() bool { return true }
