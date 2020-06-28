package mysqlStore

// MysqlDB 配置
type MysqlConfig struct {
	DbHost                  string `val:"127.0.0.1"`
	DbPort                  int64  `val:"3308"`
	DbUser                  string `val:"dev"`
	DbPasswd                string `val:"123456"`
	DbName                  string `val:"dev_db"`
	ConnMaxLifetime         int64
	MaxIdleConns            int64
	MaxOpenConns            int64
	ChartSet                string
	AllowCleartextPasswords bool
	InterpolateParams       bool
	Timeout                 int64
	ReadTimeout             int64
	ParseTime               bool
	PING                    bool
}
