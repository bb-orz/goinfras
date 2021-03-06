package XGin

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/gin-gonic/gin"
)

type starter struct {
	goinfras.BaseStarter
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
	return "XGin"
}

// 初始化时：加载配置
func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config

	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Gin", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}

	s.cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))
}

// 启动时：添加中间件，实例化应用，注册项目实现的API
func (s *starter) Setup(sctx *goinfras.StarterContext) {
	// 1.配置gin中间件
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(), ZapRecoveryMiddleware(s.cfg.RecoveryDebugStack))

	// 其他由用户启动器传递的中间件
	middlewares = append(middlewares, s.middlewares...)

	// 2.New Gin Engine
	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "Gin Engine Creating ...  ")
	ginEngine = NewGinEngine(s.cfg, middlewares...)

	// 3.API路由注册
	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "Gin Engine Register Api Routes...  ")
	for _, v := range GetApis() {
		v.SetRoutes()
	}
	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "Gin Engine Setuped! ")

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(ginEngine)
	if sctx.PassError(s.Name(), goinfras.StepCheck, err) {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "Gin Engine Setup Successful! ")
		return true
	}
	return false
}

// 启动时：运行gin engine
func (s *starter) Start(sctx *goinfras.StarterContext) {
	var addr string
	var err error
	addr = fmt.Sprintf("%s:%d", s.cfg.ListenHost, s.cfg.ListenPort)
	sctx.Logger().OK(s.Name(), goinfras.StepStart, "Gin Server Starting ... ")
	if s.cfg.Tls && s.cfg.CertFile != "" && s.cfg.KeyFile != "" {
		err = ginEngine.RunTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
	} else {
		err = ginEngine.Run(addr)
	}
	sctx.PassError(s.Name(), goinfras.StepStart, err)
}

func (s *starter) Stop() error {
	ginEngine = nil
	return nil
}

// 默认设置阻塞启动
func (s *starter) StartBlocking() bool { return true }

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
