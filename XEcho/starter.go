package XEcho

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"goinfras"
)

type starter struct {
	goinfras.BaseStarter
	cfg         *Config
	middlewares []echo.MiddlewareFunc
}

func NewStarter(middlewares ...echo.MiddlewareFunc) *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	starter.middlewares = middlewares
	return starter
}

func (s *starter) Name() string {
	return "XEcho"
}

// 初始化时：加载配置
func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var ginDefine *EchoConfig
	var corsDefine *CorsConfig

	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Echo", &ginDefine)
		goinfras.ErrorHandler(err)
	}

	if viper != nil {
		err = viper.UnmarshalKey("Cors", &corsDefine)
		goinfras.ErrorHandler(err)
	}

	// 读配置为空时，默认配置
	if ginDefine == nil {
		s.cfg = DefaultConfig()

	} else {
		s.cfg = &Config{}
		s.cfg.EchoConfig = ginDefine
		s.cfg.CorsConfig = corsDefine
		sctx.Logger().Info("Print Cors Config:", zap.Any("CorsConfig", *corsDefine))
	}

	fmt.Println(*s.cfg.EchoConfig)
	// sctx.Logger().Info("Print Gin Config:", zap.Any("GinConfig", *ginDefine))

}

// 启动时：添加中间件，实例化应用，注册项目实现的API
func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// 1.配置gin中间件
	logger := sctx.Logger()

	middlewares := make([]echo.MiddlewareFunc, 0)
	middlewares = append(middlewares, LoggerMiddleware(logger), RecoveryMiddleware(logger, false))

	// 如开启cors限制，添加中间件
	if s.cfg.CorsConfig != nil && !s.cfg.CorsConfig.AllowAllHost {
		middlewares = append(middlewares, CORSMiddleware(s.cfg.CorsConfig))
	}

	// 其他由用户启动器传递的中间件
	middlewares = append(middlewares, s.middlewares...)

	// 2.New Gin Engine
	echoEngine = NewEchoEngine(s.cfg, middlewares...)

	// 3.API路由注册
	for _, v := range GetApis() {
		v.SetRoutes()
	}
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(echoEngine)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Echo Engine Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Echo Engine Setup Successful!", s.Name()))
	return true
}

// 启动时：运行echo engine
func (s *starter) Start(sctx *goinfras.StarterContext) {
	var addr string
	var err error
	addr = fmt.Sprintf("%s:%d", s.cfg.ListenHost, s.cfg.ListenPort)
	if s.cfg.Tls && s.cfg.CertFile != "" && s.cfg.KeyFile != "" {
		err = echoEngine.StartTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
		goinfras.ErrorHandler(err)
	} else {
		err = echoEngine.Start(addr)
		goinfras.ErrorHandler(err)
	}
}

func (s *starter) Stop() {}

// 默认设置阻塞启动
func (s *starter) SetStartBlocking() bool { return true }
