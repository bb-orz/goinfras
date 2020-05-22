package main

import (
	"GoWebScaffold/infras"
	"fmt"
)

func main() {
	fmt.Println("Application Starting  ......")

	// TODO 1.预先载入配置

	// TODO 2.注册启动器

	// 3.创建应用程序启动管理器
	app := infras.NewBoot()
	// 4.启动应用
	app.Up()

	fmt.Println("Application Running  ......")
}
