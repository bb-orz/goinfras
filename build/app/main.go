package app

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
	"fmt"
	"github.com/tietang/props/kvs"
	"github.com/tietang/props/yam"
	"io"
	"os"

	"github.com/spf13/viper"
)

// TODO 测试infras各个资源组件，并整合gin框架，搭建基础脚手架
func main() {
	// 读取配置
	cfgSourse := loadConfigFile()

	// 创建应用程序启动管理器
	app := infras.NewBoot(cfgSourse)

	// 运行应用,启动已注册的资源组件
	app.Up()

	fmt.Println("Application Running  ......")
}

// 应用启动时注册资源组件启动器并按启动优先级进行排序
func init() {
	// 1注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	infras.Register(&logger.LoggerStarter{Writers: writers})
	// 2注册mongodb启动器
	infras.Register(&mongoStore.MongoDBStarter{})
	// 3注册mysql启动器
	infras.Register(&sqlbuilderStore.SqlBuilderStarter{})
	// 4 注册Redis连接池
	infras.Register(&redisStore.RedisStarter{})
	// 5 注册Oss
	infras.Register(&aliyunOss.AliyunOssStarter{})
	infras.Register(&qiniuOss.QiniuOssStarter{})
	// 6 注册Mq
	infras.Register(&redisPubSub.RedisPubSubStarter{})
	infras.Register(&natsMq.NatsMQStarter{})
	// 7 注册Oauth Manager
	infras.Register(&oauth.OauthStarter{})
	// 8 注册Cron定时任务
	infras.Register(&cron.CronStarter{})

	// 9 注册hook
	infras.Register(&hook.HookStarter{})

	// 对资源组件启动器进行排序
	infras.SortStarters()
}

// 读取配置文件
func loadConfigFile() kvs.ConfigSource {
	// 获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("config.yaml", 1)
	return yam.NewIniFileCompositeConfigSource(file)
}

// Viper 配置
func ViperLoadConfig() {
	viper.AddConfigPath("")
	v := viper.New()
	v.Get("")
}

// 日志记录器
