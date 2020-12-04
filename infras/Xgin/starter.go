package Xgin

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/XLogger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type starter struct {
	infras.BaseStarter
	cfg         Config
	middlewares []gin.HandlerFunc
}

func NewStarter(middlewares ...gin.HandlerFunc) *starter {
	starter := new(starter)
	starter.cfg = Config{}
	starter.middlewares = middlewares
	return starter
}

func (s *starter) Name() string {
	return "Xgin"
}

// 初始化时：加载配置
func (s *starter) Init(sctx *infras.StarterContext) {
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
func (s *starter) Setup(sctx *infras.StarterContext) {
	// 1.配置gin中间件
	log := XLogger.XCommon()
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(log), ZapRecoveryMiddleware(log, false))

	// 如开启cors限制，添加中间件
	if !s.cfg.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(&s.cfg.CorsConfig))
	}

	// 其他由用户启动器传递的中间件
	middlewares = append(middlewares, s.middlewares...)

	// 2.New Gin Engine
	ginEngine = NewGinEngine(&s.cfg, middlewares...)

	// 3.API 路由注册
	for _, v := range GetApis() {
		v.SetRoutes()
	}
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(ginEngine)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Gin Engine Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Gin Engine Setup Successful!", s.Name()))
	return true
}

// 启动时：运行gin engine
func (s *starter) Start(sctx *infras.StarterContext) {
	var addr string
	var err error
	addr = fmt.Sprintf("%s:%d", s.cfg.ListenHost, s.cfg.ListenPort)
	if s.cfg.Tls && s.cfg.CertFile != "" && s.cfg.KeyFile != "" {
		err = ginEngine.RunTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
		infras.FailHandler(err)
	} else {
		err = ginEngine.Run(addr)
		infras.FailHandler(err)
		sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Gin Engine Running Successful!", s.Name()))
	}
}

func (s *starter) Stop() {}

// 默认设置阻塞启动
func (s *starter) SetStartBlocking() bool { return true }
