package Xgin

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config, apis []IApi) error {
	var err error
	if config == nil {
		config = &Config{
			GinConfig{
				ListenHost: "127.0.0.1",
				ListenPort: 8090,
			},
			CorsConfig{},
		}
	}

	// 1.配置gin中间件
	log := XLogger.CLogger()
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, ZapLoggerMiddleware(log), ZapRecoveryMiddleware(log, false))

	// 如开启cors限制，添加中间件
	if !config.AllowAllOrigins {
		middlewares = append(middlewares, CORSMiddleware(&config.CorsConfig))
	}

	// 2.New Gin Engine
	ginEngine = NewGinEngine(config, middlewares...)

	// 3.Restful API 模块注册
	for _, v := range apis {
		// 路由注册
		v.SetRoutes()
	}

	// 4.启动
	var addr string
	addr = fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort)
	if config.Tls && config.CertFile != "" && config.KeyFile != "" {
		err = GinComponent().RunTLS(addr, config.CertFile, config.KeyFile)
		infras.FailHandler(err)
	} else {
		err = GinComponent().Run(addr)
		infras.FailHandler(err)
	}

	return err
}

func TestGinEngine(t *testing.T) {
	Convey("ETCD Client Test", t, func() {
		var err error
		config := Config{}
		err = TestingInstantiation(&config, nil)
		So(err, ShouldBeNil)
	})

	// TODO
}
