package main

import (
	_ "GoWebScaffold/example/simple/restful" // 初始化时自动注册restful apis层的所有接口
	// _ "GoWebScaffold/example/simple/rpc" // 初始化时自动注册rpc apis层的所有接口
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/XCron"
	"GoWebScaffold/infras/XEtcd"
	"GoWebScaffold/infras/XLogger"
	"GoWebScaffold/infras/XValidate"
	"GoWebScaffold/infras/Xgin"

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
	infras.RegisterStarter(XLogger.NewStarter(writers...))

	// 注册Cron定时任务
	// 可以自定义一些定时任务给starter启动
	infras.RegisterStarter(XCron.NewStarter())

	// 注册ETCD
	infras.RegisterStarter(XEtcd.NewStarter())

	// 注册mongodb启动器
	// mongoStarter := new(mongoStore.Starter)
	// infras.RegisterStarter(mongoStarter)

	// 注册mysql启动器
	// infras.RegisterStarter(new(sqlBuilderStore.Starter{})
	// 注册Redis连接池
	// infras.RegisterStarter(new(redisStore.Starter{})
	// 注册Oss
	// infras.RegisterStarter(new(aliyunOss.Starter{})
	// infras.RegisterStarter(new(qiniuOss.Starter{})
	// 注册Mq
	// infras.RegisterStarter(new(redisPubSub.Starter{})
	// infras.RegisterStarter(new(natsMq.Starter{})
	// 注册Oauth Manager
	// infras.RegisterStarter(new(oauth.Starter))

	// 注册gin web 服务
	infras.RegisterStarter(Xgin.NewStarter())
	// 注册验证器
	infras.RegisterStarter(XValidate.NewStarter())

	// 对资源组件启动器进行排序
	infras.SortStarters()
}
