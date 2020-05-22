package starter

import (
	"GoWebScaffold/infras/config"
	"GoWebScaffold/infras/mq/natsMq"
	"database/sql"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/garyburd/redigo/redis"
	redigo "github.com/garyburd/redigo/redis"
	qiuniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
	"go.etcd.io/etcd/clientv3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// 系统启动器上下文
type StarterContext struct {
	appConfig       *config.AppConfig
	commonLogger    *zap.Logger
	syncErrorLogger *zap.Logger
	mongoClient     *mongo.Client
	mysqlClient     *sql.DB
	redisPool       *redis.Pool
	etcdClient      *clientv3.Client
	natsMQPool      *natsMq.NatsPool
	redisPubSubPool *redigo.Pool
	qiniuOSS        *qiuniuOss.Mac
	aliyunOSS       *aliOss.Client
}

// 设置配置信息到启动器上下文
func (s *StarterContext) SetConfig(appConf *config.AppConfig) {
	if s.appConfig == nil {
		s.appConfig = appConf
	}
}

//从启动器上下文获取配置信息
func (s *StarterContext) GetConfig() *config.AppConfig {
	if s.appConfig == nil {
		panic("配置还为被初始化")
	}
	return s.appConfig
}

//
// // 设置通用日志记录器到启动器上下文
// func (s *StarterContext) SetCommonLogger(logger *zap.Logger) {
// 	if s.commonLogger == nil {
// 		s.commonLogger = logger
// 	}
// }
//
// //从启动器上下文获取通用日志记录器实例
// func (s *StarterContext) GetCommonLogger() *zap.Logger {
// 	if s.commonLogger == nil {
// 		panic("系统通用日志记录器还未启动")
// 	}
// 	return s.commonLogger
// }
//
// // 设置异步错误日志记录器到启动器上下文
// func (s *StarterContext) SetSyncErrorLogger(logger *zap.Logger) {
// 	if s.syncErrorLogger == nil {
// 		s.syncErrorLogger = logger
// 	}
// }
//
// //从启动器上下文获取异步错误日志记录器实例
// func (s *StarterContext) GetSyncErrorLogger() *zap.Logger {
// 	if s.syncErrorLogger == nil {
// 		panic("系统异步错误日志记录器还未启动")
// 	}
// 	return s.syncErrorLogger
// }
//
// // 设置mongodb连接客户端
// func (s *StarterContext) SetMongoClient(mp *mongoDao.MongoClient) {
// 	if s.mongoClient == nil {
// 		s.mongoClient = mp
// 	}
// }
//
// // 获取mongodb连接客户端
// func (s *StarterContext) GetMongoClient() *mongoDao.MongoClient {
// 	if s.mongoClient == nil {
// 		panic("Mongodb Client 未初始化")
// 	}
// 	return s.mongoClient
// }
//
// // 设置mysql连接客户端
// func (s *StarterContext) SetMysqlClient(m *mysqlDao.MysqlClient) {
// 	if s.mysqlClient == nil {
// 		s.mysqlClient = m
// 	}
// }
//
// // 获取mysql连接客户端
// func (s *StarterContext) GetMysqlClient() *mysqlDao.MysqlClient {
// 	if s.mysqlClient == nil {
// 		panic("Mysql Client 未初始化")
// 	}
// 	return s.mysqlClient
// }
//
// // 设置redis连接池
// func (s *StarterContext) SetRedisPool(p *RedisDao.RedisPool) {
// 	if s.redisPool == nil {
// 		s.redisPool = p
// 	}
// }
//
// // 获取redis连接客户端
// func (s *StarterContext) GetRedisConn() redis.Conn {
// 	if s.redisPool == nil {
// 		panic("Redis 连接池未初始化")
// 	}
// 	return s.redisPool.GetRedisConn()
// }
//
// // 设置etcd连接客户端
// func (s *StarterContext) SetEtcdClient(e *etcd.EtcdClient) {
// 	if s.etcdClient == nil {
// 		s.etcdClient = e
// 	}
// }
//
// // 获取etcd连接客户端
// func (s *StarterContext) GetEtcdClient() *etcd.EtcdClient {
// 	if s.etcdClient == nil {
// 		panic("ETCD Client 未初始化")
// 	}
// 	return s.etcdClient
// }
//
// // 设置natsMQ连接客户端
// func (s *StarterContext) SetNatsMQPool(mq *natsMq.NatsPool) {
// 	if s.natsMQPool == nil {
// 		s.natsMQPool = mq
// 	}
// }
//
// // 获取NatsMQ连接客户端
// func (s *StarterContext) GetNatsMQConn() (*nats.Conn, error) {
// 	if s.natsMQPool == nil {
// 		panic("NatsMQ Pool 未初始化")
// 	}
// 	return s.natsMQPool.Get()
// }
//
// // 设置redisPubSub连接池
// func (s *StarterContext) SetRedisPubSubMPool(mq *redigo.Pool) {
// 	if s.redisPubSubPool == nil {
// 		s.redisPubSubPool = mq
// 	}
// }
//
// // 获取redisPubSub连接
// func (s *StarterContext) GetRedisPubSubConn() redigo.Conn {
// 	if s.redisPubSubPool == nil {
// 		panic("redisPubSub Pool 未初始化")
// 	}
// 	return s.redisPubSubPool.Get()
// }
//
// // 设置aliyunOSS连接客户端
// func (s *StarterContext) SetAliyunOSSClient(c *aliOss.Client) {
// 	if s.aliyunOSS == nil {
// 		s.aliyunOSS = c
// 	}
// }
//
// // 获取aliyunOss连接
// func (s *StarterContext) GetAliyunOSSConn() *aliOss.Client {
// 	if s.aliyunOSS == nil {
// 		panic("AliyunOSS Client 未初始化")
// 	}
// 	return s.aliyunOSS
// }
//
// // 设置qiuniuOSS连接客户端
// func (s *StarterContext) SetQiniuOSS(c *qiuniuOss.Mac) {
// 	if s.qiniuOSS == nil {
// 		s.qiniuOSS = c
// 	}
// }
//
// // 获取qiniuOss连接
// func (s *StarterContext) GetQiniuOSS() *qiuniuOss.Mac {
// 	if s.qiniuOSS == nil {
// 		panic("QiniuOSS Client 未初始化")
// 	}
// 	return s.qiniuOSS
// }
