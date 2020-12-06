package XGin

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/XLogger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type starter struct {
	infras.BaseStarter
	cfg         *Config
	middlewares []gin.HandlerFunc
}

func NewStarter(middlewares ...gin.HandlerFunc) *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	starter.middlewares = middlewares
	return starter
}

func (s *starter) Name() string {
	return "Xgin"
}

// 初始化时：加载配置
func (s *starter) Init(sctx *infras.StarterContext) {
	var err error
	var ginDefine *GinConfig
	var corsDefine *CorsConfig

	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Gin", &ginDefine)
		infras.ErrorHandler(err)
	}

	if viper != nil {
		err = viper.UnmarshalKey("Cors", &corsDefine)
		infras.ErrorHandler(err)
	}

	// 读配置为空时，默认配置
	if ginDefine == nil {
		s.cfg = DefaultConfig()

	} else {
		s.cfg = &Config{}
		s.cfg.GinConfig = ginDefine
		s.cfg.CorsConfig = corsDefine
		sctx.Logger().Info("Print Cors Config:", zap.Any("CorsConfig", *corsDefine))
	}

	fmt.Println(*s.cfg.GinConfig)
	// sctx.Logger().Info("Print Gin Config:", zap.Any("GinConfig", *ginDefine))

}

// 启动时：添加中间件，实例化应用，注册项目实现的API
func (s *starter) Setup(sctx *infras.StarterContext) {
	// 1.配置gin中间件
	logger := XLogger.XCommon()
	// 如若logger组件没有注册，先启用默认启动器的简单logger
	if logger == nil {
		logger = sctx.Logger()
	}
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(logger), ZapRecoveryMiddleware(logger, false))

	// 如开启cors限制，添加中间件
	if s.cfg.CorsConfig != nil && !s.cfg.CorsConfig.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(s.cfg.CorsConfig))
	}

	// 其他由用户启动器传递的中间件
	middlewares = append(middlewares, s.middlewares...)

	// 2.New Gin Engine
	ginEngine = NewGinEngine(s.cfg, middlewares...)

	// 3.API路由注册
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
		infras.ErrorHandler(err)
	} else {
		err = ginEngine.Run(addr)
		infras.ErrorHandler(err)
	}
}

func (s *starter) Stop() {}

// 默认设置阻塞启动
func (s *starter) SetStartBlocking() bool { return true }
