package build

import (
	_ "GoWebScaffold/apis/restful" // 启动时自动注册apis/restful里所有的路由
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

	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"path/filepath"
)

// flag参数接收变量
var (
	flagConfigFile     string // 配置文件路径接收变量
	flagRemoteProvider string // 连接远程配置的类型（etcd/consul/firestore）
	flagRemoteEndpoint string // 连接远程配置的机器节点（etcd requires http://ip:port  consul requires ip:port）
	flagRemotePath     string // 连接远程配置的配置节点路径 (path is the path in the k/v store to retrieve configuration)
)

// 应用启动时注册资源组件启动器并按启动优先级进行排序
func init() {
	// 1.接收命令行参数
	bindingFlag()
	// 2.注册应用组件启动器
	registerComponent()
}

func main() {
	flag.Parse()

	// 实例化运行时viper配置
	runtimeViper := runtimeViper()

	// 创建应用程序启动管理器
	app := infras.NewApplication(runtimeViper)

	// 运行应用,启动已注册的资源组件
	app.Up()
	fmt.Println("Application Running  ......")
}

func runtimeViper() *viper.Viper {
	var err error
	var runtimeViper *viper.Viper
	var fileName, configFilePath, configFileName, configFileExt string

	runtimeViper = viper.New()

	// 1. 从环境变量导入配置项
	err = loadConfigFromEnv(runtimeViper)
	if err != nil {
		panic("Viper Loading ENV Error:" + err.Error())
	}

	// 2. 从配置文件导入配置项
	if flagConfigFile != "" {
		configFilePath, fileName = filepath.Split(flagConfigFile)
		configFileExt = path.Ext(fileName)
		configFileName = fileName[0 : len(fileName)-len(configFileExt)]
		configFileExt = configFileExt[1:]
		err = loadConfigFromFile(runtimeViper, configFilePath, configFileName, configFileExt)
		if err != nil {
			panic("Viper Loading Config File Error:" + err.Error())
		}
	}

	// 3. 从远程配置系统导入配置项
	if flagRemoteProvider != "" || flagRemoteEndpoint != "" || flagRemotePath != "" {
		err = loadConfigFromRemote(runtimeViper, flagRemoteProvider, flagRemoteEndpoint, flagRemotePath)
		if err != nil {
			panic("Viper Loading Remote Config Error:" + err.Error())
		}
	}

	return runtimeViper
}

// 绑定命令行参数
func bindingFlag() {
	// 启动时获取命令行flag参数
	flag.StringVar(&flagConfigFile, "f", "", "Config file,like: ./build/config.yaml")
	flag.StringVar(&flagRemoteProvider, "T", "", "Remote K/V config system provider，support etcd/consul")
	flag.StringVar(&flagRemoteEndpoint, "E", "", "Remote K/V config system endpoint，etcd requires http://ip:port  consul requires ip:port")
	flag.StringVar(&flagRemotePath, "P", "", "Remote K/V config path，path is the path in the k/v store to retrieve configuration,like: /configs/myapp.json")

}

// TODO Viper读取环境变量
func loadConfigFromEnv(runtimeViper *viper.Viper) error {
	// 若有需要读取环境变量
	// runtimeViper.SetEnvPrefix("env_prefix") // 添加需加载系统环境变量的前缀
	// runtimeViper.AllowEmptyEnv(true) // 是否允许环境变量为空值,默认为false
	// err := runtimeViper.BindEnv("your env var") // 绑定特定环境变量值到viper
	// runtimeViper.AutomaticEnv() // 自动载入所有环境变量，如设置SetEnvPrefix，则加载有特定前缀的所有环境变量

	return nil
}

// 读取配置渠道：本地文件读取（json/yaml/ini/...）或远程配置数据（etcd/consul/...）
// Viper 读取本地配置文件
func loadConfigFromFile(runtimeViper *viper.Viper, cfgPath, cfgName, cfgType string) error {
	var err error
	runtimeViper.AddConfigPath(cfgPath) // 设置配置文件读取路径，默认windows环境下为%GOPATH，linux环境下为$GOPATH
	runtimeViper.SetConfigName(cfgName) // 设置读取的配置文件名
	runtimeViper.SetConfigType(cfgType) // 设置配置文件类型

	if err = runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；可以读取当前目录,如果需要可以忽略

		} else {
			return err
		}
	}

	return nil
}

// Viper 读取远程配置系统
func loadConfigFromRemote(runtimeViper *viper.Viper, provider, endpoint, path string) error {
	var err error

	switch provider {
	case "etcd":
		err = runtimeViper.AddRemoteProvider(provider, endpoint, path)
		if err != nil {
			return err
		}

		// 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
		runtimeViper.SetConfigType("json")
		err = viper.ReadRemoteConfig()

		if err != nil {
			return err
		}

	case "consul":
		err = runtimeViper.AddRemoteProvider(provider, endpoint, path)
		if err != nil {
			return err
		}
		// 需要显示设置成json
		runtimeViper.SetConfigType("json")
		err = viper.ReadRemoteConfig()
		if err != nil {
			return err
		}
	default:
		return errors.New("Only Support etcd/consul K/V System. ")
	}

	return nil
}

// 注册应用组件启动器
func registerComponent() {
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	infras.Register(&logger.Starter{Writers: writers})
	// 注册mongodb启动器
	infras.Register(&mongoStore.Starter{})
	// 注册mysql启动器
	infras.Register(&sqlBuilderStore.Starter{})
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
