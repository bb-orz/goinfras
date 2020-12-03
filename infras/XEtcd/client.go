package XEtcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"time"
)

var client *clientv3.Client

func NewEtcdClient(ctx context.Context, cfg *Config, zapLoggerConf *zap.Config) (cli *clientv3.Client, err error) {
	EtcdConfig := clientv3.Config{
		Endpoints:            cfg.Endpoints,                                         // 单机或集群主机地址
		AutoSyncInterval:     time.Duration(cfg.AutoSyncInterval) * time.Second,     // 更新其最新成员端点的时间间隔。 0禁用自动同步。 默认情况下，自动同步被禁用。
		DialTimeout:          time.Duration(cfg.DialTimeout) * time.Second,          // 未能建立连接超时。
		DialKeepAliveTime:    time.Duration(cfg.DialKeepAliveTime) * time.Second,    // 客户端ping服务器以查看传输是否活动的时间。
		DialKeepAliveTimeout: time.Duration(cfg.DialKeepAliveTimeout) * time.Second, // 客户端等待keep-alive探测响应的时间。如果此时未收到响应，则连接将关闭。
		MaxCallSendMsgSize:   cfg.MaxCallSendMsgSize,                                // 客户端请求发送限制（字节）。如果为0，则默认为2.0 MiB（2*1024*1024）。请确保“MaxCallSendMsgSize”<服务器端默认发送/接收限制。 （“--max request bytes”标记为etcd或“embed.Config.MaxRequestBytes”）。
		MaxCallRecvMsgSize:   cfg.MaxCallRecvMsgSize,                                // 客户端响应接收限制。如果为0，则默认为“math.MaxInt32”，因为范围响应可能会明显超过请求发送限制。请确保“MaxCallRecvMsgSize”>=服务器端默认发送/接收限制。（--etcd的“max request bytes”标志或“embed.Config.MaxRequestBytes”）。
		TLS:                  cfg.TLS,                                               // TLS保存客户端安全凭据（如果有）。
		Username:             cfg.Username,                                          // 认证用户名
		Password:             cfg.Password,                                          // 认证密码
		RejectOldCluster:     cfg.RejectOldCluster,                                  // 如果true,则拒绝针对过时的群集创建客户端。
		LogConfig:            zapLoggerConf,                                         // LogConfig配置客户端记录器。如果为nil，则使用默认记录器。
		Context:              ctx,                                                   // 连接设置上下文
		PermitWithoutStream:  cfg.PermitWithoutStream,                               // 如为true则设置后将允许客户端在没有任何活动流（RPC）的情况下向服务器发送keepalive ping。
	}

	return clientv3.New(EtcdConfig)
}
