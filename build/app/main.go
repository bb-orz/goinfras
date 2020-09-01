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
	"fmt"
	"github.com/tietang/props/kvs"
	"github.com/tietang/props/yam"
	"io"
	"os"

	"github.com/spf13/viper"
)

// TODO 测试infras各个资源组件，并整合gin框架，搭建基础脚手架
func main() {

	// TODO 启动时获取命令行flag参数、读取环境变量

	// TODO 读取配置渠道：远程配置数据（etcd/...）、或本地文件读取（json/yaml/ini/...）

	// kvs包读取配置
	cfgSourse := yam.NewIniFileCompositeConfigSource(kvs.GetCurrentFilePath("config.yaml", 1))

	// 创建应用程序启动管理器
	app := infras.NewApplication(cfgSourse)

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

// Viper 读取本地配置文件
func ViperLoadConfigFile(cfgPath, cfgName, cfgType string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(cfgPath) // 设置配置文件读取路径，默认windows环境下为%GOPATH，linux环境下为$GOPATH
	v.SetConfigName(cfgName) // 设置读取的配置文件名
	v.SetConfigType(cfgType) // 设置配置文件类型

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

// Viper 读取远程配置系统
func ViperLoadRemoteEtcdConfig() {

}

// 日志记录器
