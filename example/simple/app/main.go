package main

import (
	"flag"
	"fmt"
	"goinfras"
	_ "goinfras/example/simple/restful" // 运行时自动注册api路由
)

// 应用启动时注册资源组件启动器并按启动优先级进行排序
func init() {
	// 1.接收命令行参数
	bindingFlag()

}

func main() {
	// 命令行启动参数解析
	flag.Parse()
	fmt.Println("Flag parameter parsed  ......")

	// 实例化运行时viper配置加载器
	runtimeViper := viperLoader()
	fmt.Println("Viper config loaded  ......")

	// 注册应用组件启动器
	fmt.Println("Register component starter  ......")
	registerStarter()

	// 创建应用程序启动管理器
	app := NewApplication(runtimeViper)

	// 运行应用,启动已注册的资源组件
	app.Up()

	fmt.Println("Application Running  ......")

}
