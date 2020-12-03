package common

/*
常量级参数配置项，编译前设置值
*/

const (
	// 环境相关常量设置
	OsEnvVarName = "CURRENT_ENV" // 检测环境时系统设置的环境变量名
	DefaultEnv   = "dev"         // 环境变量未设置时的默认值，默认使用开发环境配置
	DevEnv       = "dev"
	TestingEnv   = "testing"
	ProductEnv   = "product"

	// jwt编码时的私钥字符串
	TokenPrivateKey = "myapp"
)

// 自定义资源组件名称
// const (
// 	CronComponent        = "cron"
// 	EtcdComponent        = "etcd"
// 	GinComponent         = "gin"
// 	GlobalComponent      = "global"
// 	HookComponent        = "hook"
// 	JwtComponent         = "jwt"
// 	LoggerComponent      = "logger"
// 	MailComponent        = "mail"
// 	NatsComponent        = "nats"
// 	RedisPubSubComponent = "redis_pub_sub"
// 	OAuthComponent       = "oauth"
// 	AliyunOssComponent   = "aliyun_oss"
// 	QiniuOssComponent    = "qiniu_oss"
// 	AliyunSMSComponent   = "aliyun_sms"
// 	MongoStoreComponent  = "mongo_store"
// 	ORMComponent         = "orm_store"
// 	RedisComponent       = "redis_store"
// 	SQLBuilderComponent  = "sql_builder_store"
// 	ValidateComponent    = "validate"
// )
