package main

import (
	"goinfras"
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
	RegisterStarter(NewStarter(writers...))

	// 注册Cron定时任务
	// 可以自定义一些定时任务给starter启动
	RegisterStarter(NewStarter())

	// 注册ETCD
	RegisterStarter(NewStarter())

	// 注册mongodb启动器
	RegisterStarter(NewStarter())

	// 注册mysql启动器
	RegisterStarter(NewStarter())
	// 注册Redis连接池
	RegisterStarter(NewStarter())
	// 注册Oss
	RegisterStarter(NewStarter())
	RegisterStarter(NewStarter())
	// 注册Mq
	RegisterStarter(NewStarter())
	RegisterStarter(NewStarter())
	// 注册Oauth Manager
	RegisterStarter(NewStarter())

	// 注册gin web 服务
	RegisterStarter(XGin.NewStarter())
	// 注册验证器
	RegisterStarter(NewStarter())

	// 对资源组件启动器进行排序
	SortStarters()
}
