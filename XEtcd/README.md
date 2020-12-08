# ETCD v3 分布式配置管理客户端

> 基于 go.etcd.io/etcd/clientv3 包构建

### Etcd/clientv3包基本用法
> Documentation https://pkg.go.dev/go.etcd.io/etcd/clientv3

> 使用示例：
https://pkg.go.dev/go.etcd.io/etcd/clientv3#pkg-examples


### 启动资源组件时注册
```
// ...

goinfras.Register(XEtcd.NewStarter()

// ...
```