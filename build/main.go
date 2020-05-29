package main

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/store/mongoStore"
	"GoWebScaffold/infras/store/mysqlStore"
	"GoWebScaffold/infras/logger"
	"fmt"
	"github.com/tietang/props/kvs"
	"github.com/tietang/props/yam"
	"io"
	"os"
)

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
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./log/info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	infras.Register(&logger.LoggerStarter{Writers: writers})
	// 注册mongodb启动器
	infras.Register(&mongoStore.MongoDBStarter{})
	// 注册mysql启动器
	infras.Register(&mysqlStore.MysqlStarter{})


	// 对资源组件启动器进行排序
	infras.SortStarters()
}

// 读取配置文件
func loadConfigFile() kvs.ConfigSource  {
	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("config.yaml", 1)
	return yam.NewIniFileCompositeConfigSource(file)
}

// 日志记录器