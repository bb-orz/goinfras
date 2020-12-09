# XNats Starter

> 基于 https://github.com/nats-io/nats.go 包

### Nats Documentation

> Documentation  https://docs.nats.io



### XNats Starter Usage
```
goinfras.RegisterStarter(XNats.NewStarter())

```

### XNats Config Setting

```
// Nats Mq 消息系统
type Config struct {
	Switch      bool
	NatsServers []natsServer
}

// 可配集群
type natsServer struct {
	Host       string
	Port       int
	AuthSwitch bool
	UserName   string
	Password   string
}

```

### XNats Usage

详细可查看run_test.go
```
// 连接池资源实例调用
XNats.XPool()

// 资源组件闭包执行
err := XNats.XF(func(c *nats.Conn) error {
    // TODO
})

// 通用管道方法实例
XNats.XCommonNatsChan()

// 通用发布订阅方法实例
XNats.XCommonNatsPubSub()

// 基于队列组的主题订阅方法实例
XNats.XCommonNatsQueue()

// 基于请求响应方式的通用方法实例
XNats.XCommonNatsRequest()

```