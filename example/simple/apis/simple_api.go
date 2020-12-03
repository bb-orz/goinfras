package apis

import (
	"GoWebScaffold/example/simple/services"
	"GoWebScaffold/infras/ginger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化时注册该模块API
	Xgin.RegisterApi(new(SimpleApi))
}

type SimpleApi struct {
	service1 services.IService1
}

func (api *SimpleApi) SetRoutes() {
	api.service1 = services.GetService1()

	engine := Xgin.GinComponent()

	engine.POST("simple/foo", api.Foo)
	engine.POST("simple/bar", api.Bar)
}

func (api *SimpleApi) Foo(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := api.service1.Foo(services.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}

func (api *SimpleApi) Bar(ctx *gin.Context) {
	email := ctx.Param("email")
	// 调用服务
	err := api.service1.Bar(services.InDTO{Email: email})

	// 处理错误
	fmt.Println(err)
}
