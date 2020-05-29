package mongoStore

import (
	"GoWebScaffold/infras"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoClient(cfg *mongoConfig) (mc *mongo.Client, err error) {
	opt := options.Client()
	opt.SetAuth(options.Credential{
		// AuthMechanism:           "",  //用于身份验证的机制。支持的值包括“SCRAM-SHA-256”、“SCRAM-SHA-1”， “MONGODB-CR”、“PLAIN”、“GSSAPI”和“MONGODB-X509”。
		// AuthMechanismProperties: nil, //身份认证机制的附加参数
		// AuthSource:              "", //用于身份验证的数据库的名称。
		Username:    cfg.DbUser,
		Password:    cfg.DbPasswd,
		PasswordSet: cfg.PasswordSet, //对于GSSAPI，如果指定了密码，则此值必须为true，即使密码是空字符串，并且 如果未指定密码，则为false，表示应从运行的上下文中获取密码 过程。对于其他机制，此字段将被忽略。
	}) //设置权限认证
	if cfg.AutoEncryptionOptions {
		opt.SetAutoEncryptionOptions(options.AutoEncryption()) //作用于collection的自动加密
	}
	opt.SetCompressors(cfg.Compressors)                                          // 通信数据压缩器
	opt.SetConnectTimeout(time.Duration(cfg.ConnectTimeout) * time.Second)       //连接超时时间
	opt.SetDirect(cfg.Direct)                                                    // 设置单机直连,不会连接到集群
	opt.SetHosts(cfg.DbHosts)                                                    //设置mongo集群
	opt.SetHeartbeatInterval(time.Duration(cfg.HeartbeatInterval) * time.Second) //定期连接心跳检查,不设置默认10s
	opt.SetLocalThreshold(time.Duration(cfg.LocalThreshold) * time.Microsecond)  //指定“延迟窗口”的宽度：在为一个操作选择多个合适的服务器时，这是最短和最长平均往返时间之间可接受的非负增量。延迟窗口中的服务器是随机选择的。默认值为15毫秒。
	// opt.SetPoolMonitor(&event.PoolMonitor{Event: func(poolEvent *event.PoolEvent) {
	//
	// }}) //设置连接池监视器的事件处理器
	opt.SetMinPoolSize(cfg.MinPoolSize)                                      //连接池最小连接数,启动时最少保持可用的连接数
	opt.SetMaxPoolSize(cfg.MaxPoolSize)                                      // 连接池最大连接数
	opt.SetMaxConnIdleTime(time.Duration(cfg.MaxConnIdleTime) * time.Second) //连接池闲置连结束最大保持时间,0时表示无限制保持闲置连接状态
	opt.SetReplicaSet(cfg.ReplicaSet)                                        //指定群集的副本集名称。如果指定，集群将被视为副本集，驱动程序将自动发现集中的所有服务器，从通过ApplyURI或SetHosts指定的节点开始。副本集中的所有节点必须具有相同的副本集名称，否则客户端不会将它们视为该集的一部分。
	opt.SetRetryWrites(cfg.RetryWrites)                                      //指定是否应在某些错误（如网络）上重试一次受支持的写入操作
	opt.SetRetryReads(cfg.RetryReads)                                        //指定是否应在某些错误（如网络）上重试一次受支持的读操作
	// 创建连接客户端
	mc, err = mongo.NewClient(opt)
	infras.FailHandler(err)

	// 连接并Ping测试
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	err = mc.Connect(ctx)
	if infras.ErrorHandler(err) {
		// 测试Ping
		err = mc.Ping(ctx, nil)
		if infras.ErrorHandler(err) {
			return mc, nil
		}
	}
	return nil, err
}
