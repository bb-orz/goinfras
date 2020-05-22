package mongoDao

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/config"
	"GoWebScaffold/infras/constant"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoClient(appConf *config.AppConfig) (mc *mongo.Client, err error) {

	// TODO 创建选项,改为通过配置值设置
	opt := options.Client()
	opt.SetAppName(constant.AppName) //设置应用名
	opt.SetAuth(options.Credential{
		// AuthMechanism:           "",  //用于身份验证的机制。支持的值包括“SCRAM-SHA-256”、“SCRAM-SHA-1”， “MONGODB-CR”、“PLAIN”、“GSSAPI”和“MONGODB-X509”。
		// AuthMechanismProperties: nil, //身份认证机制的附加参数
		// AuthSource:              "", //用于身份验证的数据库的名称。
		Username:    appConf.MongoConf.DbUser,
		Password:    appConf.MongoConf.DbPasswd,
		PasswordSet: appConf.MongoConf.PasswordSet, //对于GSSAPI，如果指定了密码，则此值必须为true，即使密码是空字符串，并且 如果未指定密码，则为false，表示应从运行的上下文中获取密码 过程。对于其他机制，此字段将被忽略。
	}) //设置权限认证
	if appConf.MongoConf.AutoEncryptionOptions {
		opt.SetAutoEncryptionOptions(options.AutoEncryption()) //作用于collection的自动加密
	}
	opt.SetCompressors(appConf.MongoConf.Compressors)                                          // 通信数据压缩器
	opt.SetConnectTimeout(time.Duration(appConf.MongoConf.ConnectTimeout) * time.Second)       //连接超时时间
	opt.SetDirect(appConf.MongoConf.Direct)                                                    // 设置单机直连,不会连接到集群
	opt.SetHosts(appConf.MongoConf.DbHosts)                                                    //设置mongo集群
	opt.SetHeartbeatInterval(time.Duration(appConf.MongoConf.HeartbeatInterval) * time.Second) //定期连接心跳检查,不设置默认10s
	opt.SetLocalThreshold(time.Duration(appConf.MongoConf.LocalThreshold) * time.Microsecond)  //指定“延迟窗口”的宽度：在为一个操作选择多个合适的服务器时，这是最短和最长平均往返时间之间可接受的非负增量。延迟窗口中的服务器是随机选择的。默认值为15毫秒。
	// opt.SetPoolMonitor(&event.PoolMonitor{Event: func(poolEvent *event.PoolEvent) {
	//
	// }}) //设置连接池监视器的事件处理器
	opt.SetMinPoolSize(appConf.MongoConf.MinPoolSize)                                      //连接池最小连接数,启动时最少保持可用的连接数
	opt.SetMaxPoolSize(appConf.MongoConf.MaxPoolSize)                                      // 连接池最大连接数
	opt.SetMaxConnIdleTime(time.Duration(appConf.MongoConf.MaxConnIdleTime) * time.Second) //连接池闲置连结束最大保持时间,0时表示无限制保持闲置连接状态
	opt.SetReplicaSet(appConf.MongoConf.ReplicaSet)                                        //指定群集的副本集名称。如果指定，集群将被视为副本集，驱动程序将自动发现集中的所有服务器，从通过ApplyURI或SetHosts指定的节点开始。副本集中的所有节点必须具有相同的副本集名称，否则客户端不会将它们视为该集的一部分。
	opt.SetRetryWrites(appConf.MongoConf.RetryWrites)                                      //指定是否应在某些错误（如网络）上重试一次受支持的写入操作
	opt.SetRetryReads(appConf.MongoConf.RetryReads)                                        //指定是否应在某些错误（如网络）上重试一次受支持的读操作
	// 创建连接客户端
	mc, err = mongo.NewClient(opt)
	infras.FailHandler(err)

	// 设置默认使用的db
	mc.Database(appConf.MongoConf.Database)

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

func (mp *MongoClient) M(colName string, f func(c *mongo.Collection) error) error {
	collection := mp.defaultDb.Collection(colName)
	return f(collection)
}

func (mp *MongoClient) DM(dbName, colName string, f func(c *mongo.Collection) error) error {
	collection := mp.client.Database(dbName).Collection(colName)

	return f(collection)
}
