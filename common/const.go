package common

/*
常量级参数配置项，编译前设置值
*/
const (

	// 基本设置
	AppName    = "MyApp"
	AppVersion = "v1.0.0"

	//环境相关常量设置
	OsEnvVarName = "CURRENT_ENV" // 检测环境时系统设置的环境变量名
	DefaultEnv   = "dev"         // 环境变量未设置时的默认值，默认使用开发环境配置
	DevEnv       = "dev"
	TestingEnv   = "testing"
	ProductEnv   = "product"

	// jwt编码时的私钥字符串
	TokenPrivateKey = "myapp"
)
