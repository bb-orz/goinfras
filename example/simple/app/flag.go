package main

import "flag"

// flag参数接收变量
var (
	flagConfigFile     string // 配置文件路径接收变量
	flagRemoteProvider string // 连接远程配置的类型（etcd/consul/firestore）
	flagRemoteEndpoint string // 连接远程配置的机器节点（etcd requires http://ip:port  consul requires ip:port）
	flagRemotePath     string // 连接远程配置的配置节点路径 (path is the path in the k/v store to retrieve configuration)
)

// 绑定命令行参数
func bindingFlag() {
	// 启动时获取命令行flag参数
	flag.StringVar(&flagConfigFile, "f", "", "Config file,like: ../config/config.yaml")
	flag.StringVar(&flagRemoteProvider, "T", "", "Remote K/V config system provider，support etcd/consul")
	flag.StringVar(&flagRemoteEndpoint, "E", "", "Remote K/V config system endpoint，etcd requires http://ip:port  consul requires ip:port")
	flag.StringVar(&flagRemotePath, "P", "", "Remote K/V config path，path is the path in the k/v store to retrieve configuration,like: /configs/myapp.json")

}
