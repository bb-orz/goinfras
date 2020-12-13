# Echo Web Engine Starter

> 基于 https://github.com/labstack/echo/v4 包

### Echo Documentation
> Documentation https://echo.labstack.com/guide

> 中文文档 http://echo.topgoer.com


### Starter Usage

1、实现IApi接口，以模块方式注册实现的API

模块接口需实现 IApi interface
```
// 每个模块服务应该实现的接口
type IApi interface {
	SetRoutes() // 模块服务应该实现的方法，各模块启动器设置相应路由
}
```

以下为API模块简单示例：

```
// 包初始化时注册API模块
func init() {
	// 初始化时自动注册该API到Gin Engine
	XGin.RegisterApi(new(SimpleApi))
}

/*定义一个简单的API实现IApi接口，注册到gin引擎*/
type SimpleApi struct {
	// TODO binding service
}

// SetRouter由Gin Engine 启动时调用
func (s *SimpleApi) SetRoutes() {
	// TODO set api routes

	XEngine().GET("simple/foo", s.Foo)
	XEngine().GET("simple/bar", s.Bar)
}

func (s *SimpleApi) Foo(ctx *echo.Context) {
	// TODO call service's method to doing biz logic
	fmt.Println("Call Foo service's method to complete the biz implementation")
	ctx.JSON(200, gin.H{
		"status":  "ok",
		"message": "Call Foo service's method to complete the biz implementation",
	})
}

func (s *SimpleApi) Bar(ctx *echo.Context) {
	// TODO call service's method to doing biz logic
	fmt.Println("Call Bar service's method to complete the biz implementation")
	ctx.JSON(200, gin.H{
		"status":  "ok",
		"message": "Call Bar service's method to complete the biz implementation",
	})
}

```

 2、应用层中定义需要的中间件，并注册启动器
```
...
echoStarter := XEcho.NewStarter()

// 添加路由前中间件
preMW := make([]echo.MiddlewareFunc,0)
// ...
echoStarter.SettingPreMiddleware(preMW...)

// 添加路由后中间件
useMW := make([]echo.MiddlewareFunc,0)
// ...
echoStarter.SettingUseMiddleware(useMW...)

goinfras.RegisterStarter(echoStarter)

```


以下为echo框架内部提供的中间件，可做参考：
```
/* echo内置的中间件 */
/* =======  Root Level (Before router) ======= */
// Add trailing slash 会在在请求的 URI 后加上反斜杠
engine.Pre(middleware.AddTrailingSlash())

// Remove trailing slash 会在在请求的 URI 后去掉反斜杆
engine.Pre(middleware.RemoveTrailingSlash())

// Method Override 检查从请求中重写的方法，并使用它来代替原来的方法。出于安全原因，只有 POST 方法可以被重写。
engine.Pre(middleware.MethodOverride())

/* ======= Root Level (After router) =======*/
// Body Limit 中间件用于设置请求体的最大长度，如果请求体的大小超过了限制值，则返回 "413 － Request Entity Too Large" 响应。
// 该限制的判断是根据 Content-Length 请求标头和实际内容确定的，这使其尽可能的保证安全。
engine.Use(middleware.BodyLimit("2M"))

// Body dump 中间件通常在调试 / 记录的情况下被使用，它可以捕获请求并调用已注册的处理程序 (handler) 响应有效负载。
// 然而，当您的请求 / 响应有效负载很大时（例如上传 / 下载文件）需避免使用它；但如果避免不了，可在 skipper 函数中为端点添加异常。
engine.Use(middleware.BodyDump(nil))

// Logger 中间件记录有关每个 HTTP 请求的信息。
engine.Use(middleware.Logger())

// Gzip 中间件使用 gzip 方案来对 HTTP 响应进行压缩
engine.Use(middleware.Gzip())

// Recover 中间件从 panic 链中的任意位置恢复程序， 打印堆栈的错误信息，并将错误集中交给 HTTPErrorHandler 处理。
engine.Use(middleware.Recover())

// JWT 提供了一个 JSON Web Token (JWT) 认证中间件。
// 对于有效的 token，它将用户置于上下文中并调用下一个处理程序。
// 对于无效的 token，它会发送 "401 - Unauthorized" 响应。
// 对于丢失或无效的 Authorization 标头，它会发送 "400 - Bad Request" 。
engine.Use(middleware.JWT("Secret Key"))

// Secure 中间件用于阻止跨站脚本攻击(XSS)，内容嗅探，点击劫持，不安全链接等其他代码注入攻击
engine.Use(middleware.Secure())

// CORS (Cross-origin resource sharing) 中间件实现了 CORS 的标准。
// CORS为Web服务器提供跨域访问控制，从而实现安全的跨域数据传输。
engine.Use(middleware.CORS())

// CSRF (Cross-site request forgery) 跨域请求伪造，也被称为 one-click attack 或者 session riding，
// 通常缩写为 CSRF 或者 XSRF， 是一种挟制用户在当前已登录的Web应用程序上执行非本意的操作的攻击方法。
// 跟跨网站脚本 (XSS) 相比，XSS 利用的是用户对指定网站的信任，CSRF 利用的是网站对用户网页浏览器的信任
engine.Use(middleware.CSRF())

// Static 中间件可已被用于服务从根目录中提供的静态文件。
engine.Use(middleware.Static("/static"))

```