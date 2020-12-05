package restful

import (
	"GoWebScaffold/example/simple/services"
	"GoWebScaffold/infras/Xgin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化时自动注册该API到Gin Engine
	Xgin.RegisterApi(new(SimpleApi))
}

type SimpleApi struct {
	service1 services.IService1
}

// SetRouter由Gin Engine 启动时调用
func (s *SimpleApi) SetRoutes() {
	s.service1 = services.GetService1()

	engine := Xgin.XEngine()

	engine.POST("simple/foo", s.Foo)
	engine.POST("simple/bar", s.Bar)
}

func (s *SimpleApi) Foo(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := s.service1.Foo(services.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}

func (s *SimpleApi) Bar(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := s.service1.Bar(services.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}
