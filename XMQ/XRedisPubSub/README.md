# XRedisPubSub Starter

> 基于  包

### redis Documentation

> Documentation https://godoc.org/github.com/gomodule/redigo/redis#hdr-Publish_and_Subscribe

> Example 

Publish and Subscribe

Use the Send, Flush and Receive methods to implement Pub/Sub subscribers.
```
c.Send("SUBSCRIBE", "example")
c.Flush()
for {
    reply, err := c.Receive()
    if err != nil {
        return err
    }
    // process pushed message
}

```

The PubSubConn type wraps a Conn with convenience methods for implementing subscribers. The Subscribe, PSubscribe, Unsubscribe and PUnsubscribe methods send and flush a subscription management command. The receive method converts a pushed message to convenient types for use in a type switch.
``` 

psc := redis.PubSubConn{Conn: c}
psc.Subscribe("example")
for {
    switch v := psc.Receive().(type) {
    case redis.Message:
        fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
    case redis.Subscription:
        fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
    case error:
        return v
    }
}
```


### XRedisPubSub Starter Usage
```
goinfras.RegisterStarter(X.NewStarter())

```

### XRedisPubSub Config Setting

```
Switch      bool   // 开关
DbHost      string // 主机地址
DbPort      int    // 主机端口
DbAuth      bool   // 权限认证开关
DbPasswd    string // 权限密码
MaxActive   int64  // 最大活动连接数，0为无限
MaxIdle     int64  // 最大闲置连接数，0为无限
IdleTimeout int64  // 闲置超时时间，0位无限
```

### XRedisPubSub Usage

```
// 发布...
err := XRedisPubSub.XRedisPublisher().Pulbish("channel", "msg")
// 处理err...
```

```
// 订阅...
recSubMsgFuncs := make(map[string]RecSubMsgFunc)
// ChannelName1 订阅频道消息的处理函数
recSubMsgFuncs[ChannelName1] = func(channel string, msg interface{}) error {
    logger.Info("Receive Message:", zap.String("channel", channel), zap.Any("message", msg))
    fmt.Println(msg)
    return nil
}
// ChannelName2 订阅频道消息的处理函数
recSubMsgFuncs[ChannelName2] = func(channel string, msg interface{}) error {
    logger.Info("Receive Message:", zap.String("channel", channel), zap.Any("message", msg))
    fmt.Println(msg)
    return nil
}

// 取消订阅通道信号，传入需要取消订阅的频道名称
unSubCh := make(chan string, 1)

go func() {
    // 10s后发送取消订阅信号
    time.Sleep(10 * time.Second)
    unSubCh <- ChannelName1
    unSubCh <- ChannelName2
}()

// 开始订阅，go程阻塞
err = XRedisSubscriber(logger).Subscribe(recSubMsgFuncs, unSubCh)
// 处理err...

```