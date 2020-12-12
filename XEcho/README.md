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

middlewares := make([]echo.MiddlewareFunc, 0)
// TODO add your echo middlewares
// ...
// ...
goinfras.RegisterStarter(XEcho.NewStarter(middlewares...))

```