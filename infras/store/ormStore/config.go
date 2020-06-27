package ormStore

// MysqlDB 配置
type ormConfig struct {
	Dialect       string `val:"mysql"` // 数据库驱动类型
	DbHost        string `val:"127.0.0.1"`
	DbPort        int64  `val:"3306"`
	DbUser        string `val:"dev"`
	DbPasswd      string `val:"123456"`
	DbName        string `val:"dev_db"`
	ChartSet      string `val:"utf8"`
	ParseTime     bool   `val:"true"`
	Local         string `val:"Local"`
	SSLMode       string `val:"disable"` // 加密传输
	SingularTable bool   `val："false"`   // 是否设置全局复数表名
}
