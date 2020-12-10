# ETCD v3 分布式配置管理客户端

> 基于 go.etcd.io/etcd/clientv3 包构建

### Etcd/clientv3 Documentation
> Documentation https://pkg.go.dev/go.etcd.io/etcd/clientv3

> 使用示例：
https://pkg.go.dev/go.etcd.io/etcd/clientv3#pkg-examples


### Starter Usage
```
// ...

goinfras.Register(XEtcd.NewStarter()

// ...
```

### ETCD Config Setting
```
Endpoints            []string    // etcd服务节点列表
TLS                  *tls.Config // 加密配置
Username             string      // 用户名
Password             string      // 用户密码
PermitWithoutStream  bool        // 如为true则设置后将允许客户端在没有任何活动流（RPC）的情况下向服务器发送keepalive ping。
RejectOldCluster     bool        // 如果true,则拒绝针对过时的群集创建客户端。
AutoSyncInterval     uint        // 更新其最新成员端点的时间间隔。 0禁用自动同步。 默认情况下，自动同步被禁用。
DialTimeout          uint        // 未能建立连接超时。
DialKeepAliveTime    uint        // 客户端ping服务器以查看传输是否活动的时间。
DialKeepAliveTimeout uint        // 客户端等待keep-alive探测响应的时间。如果此时未收到响应，则连接将关闭。
MaxCallRecvMsgSize   int         // 客户端响应接收限制。如果为0，则默认为“math.MaxInt32”，因为范围响应可能会明显超过请求发送限制。请确保“MaxCallRecvMsgSize”>=服务器端默认发送/接收限制。（--etcd的“max request bytes”标志或“embed.Config.MaxRequestBytes”）。
MaxCallSendMsgSize   int         // 客户端请求发送限制（字节）。如果为0，则默认为2.0 MiB（2*1024*1024）。请确保“MaxCallSendMsgSize”<服务器端默认发送/接收限制。 （“--max request bytes”标记为etcd或“embed.Config.MaxRequestBytes”）。
```


### XEtcd Usage
```
sr, err := XEtcd.XClient().Put(context.Background(), "mykeya", "aaaaaaa")
So(err, ShouldBeNil)
Println("Set Key Response:", sr)

gr, err := XEtcd.XClient().Get(context.Background(), "mykeya")
So(err, ShouldBeNil)
Println("Get Key Response:", gr)


Other Usage... 
```