package XEcho

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/labstack/echo/v4"
)

type starter struct {
	goinfras.BaseStarter
	cfg            *Config
	preMiddlewares []echo.MiddlewareFunc
	useMiddlewares []echo.MiddlewareFunc
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

// 设置 Root Level (Before router) 中间件
func (s *starter) SettingPreMiddleware(pmw ...echo.MiddlewareFunc) {
	s.preMiddlewares = pmw
}

// 设置Root Level (After router) 中间件
func (s *starter) SettingUseMiddleware(umw ...echo.MiddlewareFunc) {
	s.useMiddlewares = umw
}

func (s *starter) Name() string {
	return "XEcho"
}

// 初始化时：加载配置
func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config

	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Echo", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().SDebug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v \n", define))
}

// 启动时：添加中间件，实例化应用，注册项目实现的API
func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// New Gin Engine
	echoEngine = NewEchoEngine(s.cfg)
	// 添加路由前中间件
	if len(s.preMiddlewares) > 0 {
		echoEngine.Pre(s.preMiddlewares...)
	}
	// 默认添加必要的中间件
	if !s.cfg.UseSelfMiddleware {
		// 自定义日志记录
		echoEngine.Use(LoggerMiddleware())
		// 自定义错误处理
		echoEngine.Use(ErrorMiddleware())
		// 自定义panic恢复
		echoEngine.Use(RecoveryMiddleware(true))
	}
	// 添加路由后中间件
	if len(s.useMiddlewares) > 0 {
		echoEngine.Use(s.useMiddlewares...)
	}
	// API路由注册
	for _, v := range GetApis() {
		v.SetRoutes()
	}
	sctx.Logger().SInfo(s.Name(), goinfras.StepSetup, fmt.Sprintf("Echo Engine Setuped! \n"))
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(echoEngine)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().SInfo(s.Name(), goinfras.StepCheck, fmt.Sprintf("Echo Engine Setup Successful! \n "))
		return true
	}
	return false
}

// 启动时：运行echo engine
func (s *starter) Start(sctx *goinfras.StarterContext) {
	var addr string
	var err error
	addr = fmt.Sprintf("%s:%d", s.cfg.ListenHost, s.cfg.ListenPort)
	sctx.Logger().SInfo(s.Name(), goinfras.StepStart, fmt.Sprintf("Echo Server Starting ... \n"))
	if s.cfg.Tls && s.cfg.CertFile != "" && s.cfg.KeyFile != "" {
		err = echoEngine.StartTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
	} else {
		err = echoEngine.Start(addr)
	}
	sctx.PassError(s.Name(), goinfras.StepStart, err)
}

func (s *starter) Stop() {}

// 默认设置阻塞启动
func (s *starter) StartBlocking() bool { return true }

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
