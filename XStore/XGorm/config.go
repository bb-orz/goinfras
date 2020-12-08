package XGorm

// MysqlDB 配置
type Config struct {
	Dialect       string // 数据库驱动类型
	DbHost        string // 主机地址
	DbPort        int64  // 主机端口
	DbUser        string // 用户名
	DbPasswd      string // 密码
	DbName        string // 数据库名
	ChartSet      string // 字符集
	ParseTime     bool   // 将数据库的datetime时间格式转换为go time包数据类型
	Local         string // 本地时区设置
	SSLMode       string // 加密传输
	SingularTable bool   // 是否设置全局复数表名
}

func DefaultConfig() *Config {
	return &Config{
		"mysql",
		"127.0.0.1",
		3306,
		"dev",
		"123456",
		"dev_db",
		"utf8",
		true,
		"Local",
		"disable",
		false,
	}
}
