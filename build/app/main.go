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
	"flag"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	configFile string //  配置文件路径接收变量

	remoteProvider string // 连接远程配置的类型（etcd/consul/firestore）
	remoteEndpoint string // 连接远程配置的机器节点（etcd requires http://ip:port  consul requires ip:port）
	remotePath     string // 连接远程配置的配置节点路径 (path is the path in the k/v store to retrieve configuration)

)

// TODO 测试infras各个资源组件，并整合gin框架，搭建基础脚手架
func main() {
	var err error
	var runtimeViper *viper.Viper
	var fileName, configFilePath, configFileName, configFileExt string

	flag.Parse()
	if configFile != "" {
		configFilePath, fileName = filepath.Split(configFile)
		configFileExt = path.Ext(fileName)
		configFileName = fileName[0 : len(fileName)-len(configFileExt)]
		configFileExt = configFileExt[1:]
		runtimeViper, err = viperLoadConfigFile(configFilePath, configFileName, configFileExt)
		if err != nil {
			panic("Viper Read Config File Error:" + err.Error())
		}
	}

	if remoteProvider != "" || remoteEndpoint != "" || remotePath != "" {
		runtimeViper, err = viperLoadRemoteEtcdConfig(remoteProvider, remoteEndpoint, remotePath)
		if err != nil {
			panic("Viper Read Remote Config Error:" + err.Error())
		}
	}

	if runtimeViper == nil {
		runtimeViper = viper.New()
	}

	// TODO kvs库替换为viper库读取配置

	// kvs包读取配置
	// cfgSourse := yam.NewIniFileCompositeConfigSource(kvs.GetCurrentFilePath("config.yaml", 1))

	// 创建应用程序启动管理器
	// app := infras.NewApplication(cfgSourse)

	// 运行应用,启动已注册的资源组件
	// app.Up()

	// fmt.Println("Application Running  ......")
}

// 应用启动时注册资源组件启动器并按启动优先级进行排序
func init() {
	// 1.接收命令行参数
	receiveFlag()
	// 2.注册应用组件启动器
	// registerComponent()
}

// 接收命令行参数
func receiveFlag() {
	// 启动时获取命令行flag参数
	flag.StringVar(&configFile, "f", "", "Config file,like: ./build/config.yaml")
	flag.StringVar(&remoteProvider, "T", "", "Remote K/V config system provider，support etcd/consul")
	flag.StringVar(&remoteEndpoint, "E", "", "Remote K/V config system endpoint，etcd requires http://ip:port  consul requires ip:port")
	flag.StringVar(&remotePath, "P", "", "Remote K/V config path，path is the path in the k/v store to retrieve configuration,like: /configs/myapp.json")
}

// 注册应用组件启动器
func registerComponent() {
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	infras.Register(&logger.LoggerStarter{Writers: writers})
	// 注册mongodb启动器
	infras.Register(&mongoStore.MongoDBStarter{})
	// 注册mysql启动器
	infras.Register(&sqlbuilderStore.SqlBuilderStarter{})
	// 注册Redis连接池
	infras.Register(&redisStore.RedisStarter{})
	// 注册Oss
	infras.Register(&aliyunOss.AliyunOssStarter{})
	infras.Register(&qiniuOss.QiniuOssStarter{})
	// 注册Mq
	infras.Register(&redisPubSub.RedisPubSubStarter{})
	infras.Register(&natsMq.NatsMQStarter{})
	// 注册Oauth Manager
	infras.Register(&oauth.OauthStarter{})
	// 注册Cron定时任务
	infras.Register(&cron.CronStarter{})
	// 注册hook
	infras.Register(&hook.HookStarter{})
	// 对资源组件启动器进行排序
	infras.SortStarters()
}

// TODO 读取配置渠道：远程配置数据（etcd/...）、或本地文件读取（json/yaml/ini/...）
// Viper 读取本地配置文件
func viperLoadConfigFile(cfgPath, cfgName, cfgType string) (*viper.Viper, error) {
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
func viperLoadRemoteEtcdConfig(provider, endpoint, path string) (*viper.Viper, error) {

	return &viper.Viper{}, nil
}
