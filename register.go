package main

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/starter"
	"io"
	"os"
)

// 主程序启动时引入
func init() {
	// 注册配置启动器
	infras.Register(&starter.ConfigStarter{})
	// 注册日志记录启动器，并添加一个异步日志输出到文件
	file, err := os.OpenFile("./log/info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writers := []io.Writer{file}
	infras.Register(&starter.LoggerStarter{Writers: writers})
	// 注册mongodb启动器
	infras.Register(&starter.MongoDBStarter{})
	// 注册mysql启动器
	infras.Register(&starter.MysqlStarter{})
}
