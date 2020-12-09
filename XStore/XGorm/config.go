package XGorm

import "time"

// MysqlDB 配置
type Config struct {
	Dialect                   string        // 数据库驱动类型:mysql/postgres
	DbHost                    string        // 主机地址
	DbPort                    int64         // 主机端口
	DbUser                    string        // 用户名
	DbPasswd                  string        // 密码
	DbName                    string        // 数据库名
	ChartSet                  string        // 字符集
	ParseTime                 bool          // 将数据库的datetime时间格式转换为go time包数据类型
	Local                     string        // 本地时区设置
	ConnMaxLifetime           time.Duration // 设置了连接可复用的最大时间
	MaxOpenConns              int           // 设置打开数据库连接的最大数量
	MaxIdleConns              int           // 设置连接池中空闲连接的最大数量
	DefaultStringSize         uint          // For Mysql，字符串字段默认长度
	DisableDatetimePrecision  bool          // For Mysql，是否禁用datetime precision，这在MySQL5.6之前不受支持
	DontSupportRenameIndex    bool          // For Mysql，是否重命名索引，重命名索引在MySQL 5.7和MariaDB之前不支持
	DontSupportRenameColumn   bool          // For Mysql，是否可以重命名列，MySQL 8和MariaDB之前不支持重命名列
	SkipInitializeWithVersion bool          // For Mysql，是否基于当前MySQL版本自动配置
	PreferSimpleProtocol      bool          // For PostgresSQL，禁用隐式准备语句用法
	TimeZone                  string        // For PostgresSQL，时区
	SslMode                   string        // For PostgresSQL，SSL模式
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
		time.Hour,
		100,
		10,
		256,
		true,
		true,
		true,
		false,
		true,
		"Asia/Shanghai",
		"disable",
	}
}
