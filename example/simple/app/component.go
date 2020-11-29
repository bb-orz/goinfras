package main

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/cron"
	"GoWebScaffold/infras/hook"
	"GoWebScaffold/infras/logger"
	"GoWebScaffold/infras/mq/natsMq"
	"GoWebScaffold/infras/mq/redisPubSub"
	"GoWebScaffold/infras/oauth"
	"GoWebScaffold/infras/oss/aliyunOss"
	"GoWebScaffold/infras/oss/qiniuOss"
	"GoWebScaffold/infras/store/mongoStore"
	"GoWebScaffold/infras/store/redisStore"
	"GoWebScaffold/infras/store/sqlbuilderStore"
	"io"
	"os"
)

// 注册应用组件启动器
func registerComponent() {
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	// 注册zap日志记录启动器
	infras.Register(&logger.Starter{Writers: writers})
	// 注册mongodb启动器
	infras.Register(&mongoStore.Starter{})
	// 注册mysql启动器
	infras.Register(&sqlbuilderStore.Starter{})
	// 注册Redis连接池
	infras.Register(&redisStore.Starter{})
	// 注册Oss
	infras.Register(&aliyunOss.Starter{})
	infras.Register(&qiniuOss.Starter{})
	// 注册Mq
	infras.Register(&redisPubSub.Starter{})
	infras.Register(&natsMq.Starter{})
	// 注册Oauth Manager
	infras.Register(&oauth.Starter{})
	// 注册Cron定时任务
	infras.Register(&cron.Starter{})
	// 注册hook
	infras.Register(&hook.Starter{})
	// 对资源组件启动器进行排序
	infras.SortStarters()
}
