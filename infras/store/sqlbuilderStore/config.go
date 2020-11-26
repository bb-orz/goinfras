package sqlbuilderStore

// MysqlDB 配置
type Config struct {
	DbHost                  string // 主机地址
	DbPort                  int64  // 主机端口
	DbUser                  string // 用户名
	DbPasswd                string // 密码
	DbName                  string // 数据库名
	ConnMaxLifetime         int64  // 每个连接最长生命周期，单位秒
	MaxIdleConns            int64  // 连接池最大闲置连接数
	MaxOpenConns            int64  // 连接池最大连接数
	ChartSet                string // 传输字符编码
	AllowCleartextPasswords bool   // 开发环境中设置允许明文密码通信
	InterpolateParams       bool   // 设置允许占位符参数
	Timeout                 int64  // 连接请求的超时时间，单位秒
	ReadTimeout             int64  // 读超时时间，单位秒
	ParseTime               bool   // 将数据库的datetime时间格式转换为go time包数据类型
	PING                    bool   // 连接时PING测试
}
