package XMongo

// MongoDB 配置
type Config struct {
	DbHosts               []string
	DbUser                string
	DbPasswd              string
	Database              string
	ReplicaSet            string   // 指定群集的副本集名称。如果指定，集群将被视为副本集，驱动程序将自动发现集中的所有服务器，从通过ApplyURI或SetHosts指定的节点开始。副本集中的所有节点必须具有相同的副本集名称，否则客户端不会将它们视为该集的一部分。
	PasswordSet           bool     // 对于GSSAPI，如果指定了密码，则此值必须为true，即使密码是空字符串，并且 如果未指定密码，则为false，表示应从运行的上下文中获取密码 过程。对于其他机制，此字段将被忽略。
	LocalThreshold        int      // 指定“延迟窗口”的宽度：在为一个操作选择多个合适的服务器时，这是最短和最长平均往返时间之间可接受的非负增量。延迟窗口中的服务器是随机选择的。默认值为15毫秒。
	Compressors           []string // 通信数据压缩器,可多选
	Direct                bool     // 是否单机直连
	HeartbeatInterval     int      // 定期连接心跳检查,不设置默认10s
	MinPoolSize           uint64   // 最小连接池连接数
	MaxPoolSize           uint64   // 最大连接池连接数
	MaxConnIdleTime       uint64   // 连接池闲置连结束最大保持时间,0时表示无限制保持闲置连接状态
	AutoEncryptionOptions bool     // 作用于collection的自动加密
	ConnectTimeout        int      // 接超时时间,单位秒
	RetryReads            bool     // 指定是否应在某些错误（如网络）上重试一次受支持的读操作
	RetryWrites           bool     // 指定是否应在某些错误（如网络）上重试一次受支持的写入操作

}

func DefaultConfig() *Config {
	return &Config{
		[]string{"127.0.0.1:27017"},
		"",
		"",
		"",
		"",
		true,
		15,
		nil,
		true,
		10,
		100,
		1000,
		120,
		false,
		20,
		true,
		true,
	}
}
