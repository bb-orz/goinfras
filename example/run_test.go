package example

import (
	_ "GoWebScaffold/example/simple/apis" // 运行时自动注册api路由
	"GoWebScaffold/infras/ginger"
	"GoWebScaffold/infras/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestGinger(t *testing.T) {
	Convey("Test Ginger...", t, func() {
		var err error
		config := &ginger.Config{
			ListenHost: "127.0.0.1",
			ListenPort: 8090,
			Cors: &ginger.CorsConfig{
				AllowAllOrigins: true,
			},
		}

		// 1.配置gin中间件
		log := logger.CommonLogger()
		middlewares := make([]gin.HandlerFunc, 0)
		middlewares = append(middlewares, ginger.ZapLoggerMiddleware(log), ginger.ZapRecoveryMiddleware(log, false))

		// 如开启cors限制，添加中间件
		if !config.Cors.AllowAllOrigins {
			middlewares = append(middlewares, ginger.CORSMiddleware(config.Cors))
		}

		// 2.New Gin Engine
		ginEngine := ginger.NewGinEngine(config, middlewares...)
		ginger.SetGinEngine(ginEngine)

		// 3.Restful API 模块注册
		for _, v := range ginger.GetApis() {
			// 路由注册
			v.SetRoutes()
		}

		// 4.启动
		var addr string
		addr = fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort)
		if config.Tls && config.CertFile != "" && config.KeyFile != "" {
			err = ginEngine.RunTLS(addr, config.CertFile, config.KeyFile)
			So(err, ShouldBeNil)
		} else {
			err = ginEngine.Run(addr)
			So(err, ShouldBeNil)
		}

		fmt.Print("Ginger Running...")

	})
}
