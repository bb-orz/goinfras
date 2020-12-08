package main

import (
	"github.com/gin-gonic/gin"
	"goinfras"
	"goinfras/XCron"
	"goinfras/XLogger"
	"goinfras/XOAuth"
	"goinfras/XStore/XGorm"
	"goinfras/XStore/XRedis"
	"goinfras/XValidate"
	_ "goinfras/example/simple/restful" // 初始化时自动注册restful apis层的所有接口

	"goinfras/XGin"
	"io"
	"os"
)

// 注册应用组件启动器，把基础设施各资源组件化
func registerStarter() {
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	goinfras.RegisterStarter(XLogger.NewStarter(writers...))

	// 注册Cron定时任务
	// 可以自定义一些定时任务给starter启动
	goinfras.RegisterStarter(XCron.NewStarter())

	// 注册ETCD
	// goinfras.RegisterStarter(XEtcd.NewStarter())

	// 注册mongodb启动器
	// goinfras.RegisterStarter(XMongo.NewStarter())

	// 注册mysql启动器
	goinfras.RegisterStarter(XGorm.NewStarter())
	// 注册Redis连接池
	goinfras.RegisterStarter(XRedis.NewStarter())
	// 注册Oss
	// goinfras.RegisterStarter(XAliyunOss.NewStarter())
	// goinfras.RegisterStarter(XQiniuOss.NewStarter())
	// 注册Mq
	// goinfras.RegisterStarter(XNats.NewStarter())
	// goinfras.RegisterStarter(XRedisPubSub.NewStarter())
	// 注册Oauth Manager
	goinfras.RegisterStarter(XOAuth.NewStarter())

	// 注册gin web 服务
	middlewares := make([]gin.HandlerFunc, 0)
	// TODO add your gin middlewares
	goinfras.RegisterStarter(XGin.NewStarter(middlewares...))
	// 注册验证器
	goinfras.RegisterStarter(XValidate.NewStarter())

	// 对资源组件启动器进行排序
	goinfras.SortStarters()
}
