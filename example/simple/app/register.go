package main

import (
	_ "GoWebScaffold/example/simple/apis" // 初始化时自动注册apis层的所有接口
	"GoWebScaffold/hub"
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/cron"
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

	// TODO 把各基础设施的资源组件化，如cron组件

	// 注册zap日志记录启动器
	writers := []io.Writer{file}
	loggerStarter := new(logger.Starter)
	loggerStarter.Writers = writers
	infras.RegisterStarter(loggerStarter)

	// 注册hook
	hookStarter := new(hook.Starter)
	infras.RegisterStarter(hookStarter)

	// 注册Cron定时任务
	infras.RegisterStarter(new(cron.Starter))

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
	infras.RegisterStarter(&ginger.Starter{})
	// 注册验证器
	infras.RegisterStarter(&validate.Starter{})

	// 对资源组件启动器进行排序
	infras.SortStarters()
}
