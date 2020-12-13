package XEcho

import (
	"fmt"
	"github.com/labstack/echo/middleware"
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

	middlewares := make([]echo.MiddlewareFunc, 0)

	/* TODO 使用echo内置的中间件 */
	/* TODO Root Level (Before router) */
	// 1.Add trailing slash 中间件会在在请求的 URI 后加上反斜杠
	middleware.AddTrailingSlash()

	// 2.Remove trailing slash 中间件会在在请求的 URI 后去掉反斜杆
	middleware.RemoveTrailingSlash()

	// 3.Method Override 中间件检查从请求中重写的方法，并使用它来代替原来的方法。出于安全原因，只有 POST 方法可以被重写。
	middleware.MethodOverride()

	/* TODO Root Level (After router)*/
	// 1.Body Limit 中间件用于设置请求体的最大长度，如果请求体的大小超过了限制值，则返回 "413 － Request Entity Too Large" 响应。
	// 该限制的判断是根据 Content-Length 请求标头和实际内容确定的，这使其尽可能的保证安全。
	middleware.BodyLimit("2M")

	// 2.Logger 中间件记录有关每个 HTTP 请求的信息。
	middleware.Logger()

	// 3.Gzip 中间件使用 gzip 方案来对 HTTP 响应进行压缩
	middleware.Gzip()

	// 4.Recover 中间件从 panic 链中的任意位置恢复程序， 打印堆栈的错误信息，并将错误集中交给 HTTPErrorHandler 处理。
	middleware.Recover()

	// 5.JWT 提供了一个 JSON Web Token (JWT) 认证中间件。
	// 对于有效的 token，它将用户置于上下文中并调用下一个处理程序。
	// 对于无效的 token，它会发送 "401 - Unauthorized" 响应。
	// 对于丢失或无效的 Authorization 标头，它会发送 "400 - Bad Request" 。
	middleware.JWT("Secret Key")

	// 6.Secure 中间件用于阻止跨站脚本攻击(XSS)，内容嗅探，点击劫持，不安全链接等其他代码注入攻击
	middleware.Secure()

	// 7.CORS (Cross-origin resource sharing) 中间件实现了 CORS 的标准。
	// CORS为Web服务器提供跨域访问控制，从而实现安全的跨域数据传输。
	middleware.CORS()

	// 8.Static 中间件可已被用于服务从根目录中提供的静态文件。
	middleware.Static("/static")

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

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
