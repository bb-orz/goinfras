package main

import (
	_ "GoWebScaffold/example/simple/apis" // 初始化时自动注册apis层的所有接口
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/ginger"
	"GoWebScaffold/infras/hook"
	"GoWebScaffold/infras/logger"
	"GoWebScaffold/infras/validate"
	"io"
	"os"
)

// 注册应用组件启动器
func registerStarter() {
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	// 注册zap日志记录启动器
	writers := []io.Writer{file}
	loggerStarter := new(logger.Starter)
	loggerStarter.Writers = writers
	infras.Register(loggerStarter)

	// 注册hook
	hookStarter := new(hook.Starter)
	infras.Register(hookStarter)

	// 注册mongodb启动器
	// mongoStarter := new(mongoStore.Starter)
	// infras.Register(mongoStarter)

	// 注册mysql启动器
	// infras.Register(new(sqlBuilderStore.Starter{})
	// 注册Redis连接池
	// infras.Register(new(redisStore.Starter{})
	// 注册Oss
	// infras.Register(new(aliyunOss.Starter{})
	// infras.Register(new(qiniuOss.Starter{})
	// 注册Mq
	// infras.Register(new(redisPubSub.Starter{})
	// infras.Register(new(natsMq.Starter{})
	// 注册Oauth Manager
	// infras.Register(new(oauth.Starter))
	// 注册Cron定时任务
	// infras.Register(new(cron.Starter))
	// 注册gin web 服务
	infras.Register(&ginger.Starter{})
	// 注册验证器
	infras.Register(&validate.Starter{})

	// 对资源组件启动器进行排序
	infras.SortStarters()
}
