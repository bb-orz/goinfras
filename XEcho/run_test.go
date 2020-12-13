package XEcho

import (
	"fmt"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"goinfras"
	"testing"
)

func TestEchoEngine(t *testing.T) {
	Convey("Echo Server Run Test", t, func() {
		config := DefaultConfig()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		// 初始化默认引擎
		CreateDefaultEngine(nil, logger)

		// 注册API接口
		RegisterApi(new(SimpleApi))

		// 以API为模块设置路由
		for _, v := range GetApis() {
			// 路由注册
			v.SetRoutes()
		}

		// 启动
		var addr string
		addr = fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort)
		if config.Tls && config.CertFile != "" && config.KeyFile != "" {
			err = XEngine().StartTLS(addr, config.CertFile, config.KeyFile)
			So(err, ShouldBeNil)
		} else {
			err = XEngine().Start(addr)
			So(err, ShouldBeNil)
		}
	})

}

/*定义一个简单的API实现IApi接口，注册到echo引擎*/
type SimpleApi struct {
	// TODO binding service
}

// SetRouter由Gin Engine 启动时调用
func (s *SimpleApi) SetRoutes() {
	// TODO binding service

	XEngine().GET("simple/foo", s.Foo)
	XEngine().GET("simple/bar", s.Bar)
}

func (s *SimpleApi) Foo(ctx echo.Context) error {
	// TODO call service's method to doing biz logic
	fmt.Println("Call Foo service's method to complete the biz implementation")
	return ctx.JSON(200, echo.Map{
		"status":  "ok",
		"message": "Call Foo service's method to complete the biz implementation",
	})
}

func (s *SimpleApi) Bar(ctx echo.Context) error {
	// TODO call service's method to doing biz logic
	fmt.Println("Call Bar service's method to complete the biz implementation")
	return ctx.JSON(200, echo.Map{
		"status":  "ok",
		"message": "Call Bar service's method to complete the biz implementation",
	})

}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XEcho Starter", t, func() {
		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")

		// 注册API接口
		RegisterApi(new(SimpleApi))

		s.Setup(sctx)
		Println("Starter Setup Successful!")
		s.Start(sctx)
		Println("Starter Start Successful!")
		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

	})
}
