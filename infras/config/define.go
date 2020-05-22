package config

//配置
type AppConfig struct {
	BaseConf  *Base
	LogConf   *Log
	MysqlConf *Mysql
	RedisConf *Redis
	MongoConf *Mongodb
	MqConf    *Mq
	CorsConf  *Cors
	OssConf   *Oss
	EtcdConf  *Etcd
}

//基础配置
type Base struct {
	Env        string `yaml:"Env"`
	ListenPort int64  `yaml:"Listen"`
}

// 日志配置
type Log struct {
	DebugLevelSwitch   bool   `yaml:"DebugLevelSwitch"`
	InfoLevelSwitch    bool   `yaml:"InfoLevelSwitch"`
	WarnLevelSwitch    bool   `yaml:"WarnLevelSwitch"`
	ErrorLevelSwitch   bool   `yaml:"ErrorLevelSwitch"`
	DPanicLevelSwitch  bool   `yaml:"DPanicLevelSwitch"`
	PanicLevelSwitch   bool   `yaml:"PanicLevelSwitch"`
	FatalLevelSwitch   bool   `yaml:"FatalLevelSwitch"`
	SimpleZapCore      bool   `yaml:"SimpleZapCore"`
	SyncZapCore        bool   `yaml:"SyncZapCore"`
	SyncLogSwitch      bool   `yaml:"SyncLogSwitch"`
	StdoutLogSwitch    bool   `yaml:"StdoutLogSwitch"`
	RotateLogSwitch    bool   `yaml:"RotateLogSwitch"`
	LogDir             string `yaml:"LogDir"`
	WithRotationTime   int    `yaml:"WithRotationTime"`
	MaxDayCount        int    `yaml:"MaxDayCount"`
	RotateZapCore      bool   `yaml:"RotateZapCore"`
	MongoLogSwitch     bool   `yaml:"MongoLogSwitch"`
	MongoLogCollection string `yaml:"MongoLogCollection"`
	MongoLogExpire     int    `yaml:"MongoLogExpire"`
}

// MysqlDB 配置
type Mysql struct {
	DbHost                  string `yaml:"DbHost"`
	DbPort                  int64  `yaml:"DbPort"`
	DbUser                  string `yaml:"DbUser"`
	DbPasswd                string `yaml:"DbPasswd"`
	DbName                  string `yaml:"DbName"`
	ConnMaxLifetime         int64  `yaml:"ConnMaxLifetime"`
	MaxIdleConns            int64  `yaml:"MaxIdleConns"`
	MaxOpenConns            int64  `yaml:"MaxOpenConns"`
	ChartSet                string `yaml:"Charset"`
	AllowCleartextPasswords bool   `yaml:"AllowCleartextPasswords"`
	InterpolateParams       bool   `yaml:"InterpolateParams"`
	Timeout                 int64  `yaml:"Timeout"`
	ReadTimeout             int64  `yaml:"ReadTimeout"`
	ParseTime               bool   `yaml:"ParseTime"`
	PING                    bool   `yaml:"Ping"`
}

// RedisDB配置
type Redis struct {
	DbHost      string `yaml:"DbHost"`
	DbPort      int64  `yaml:"DbPort"`
	DbAuth      bool   `yaml:"DbAuth"`
	DbPasswd    string `yaml:"DbPasswd"`
	MaxActive   int64  `yaml:"MaxActive"`
	MaxIdle     int64  `yaml:"MaxIdle"`
	IdleTimeout int64  `yaml:"IdleTimeout"`
}

// MongoDB 配置
type Mongodb struct {
	DbHosts  []string `yaml:"DbHosts"`
	DbUser   string   `yaml:"DbUser"`
	DbPasswd string   `yaml:"DbPasswd"`
	Database string   `yaml:"Database"`

	ReplicaSet            string   `yaml:"ReplicaSet"`
	RetryWrites           bool     `yaml:"RetryWrites"`
	PasswordSet           bool     `yaml:"PasswordSet"`
	LocalThreshold        int      `yaml:"LocalThreshold"`
	Compressors           []string `yaml:"Compressors"`
	Direct                bool     `yaml:"Direct"`
	HeartbeatInterval     int      `yaml:"HeartbeatInterval"`
	MaxPoolSize           uint64   `yaml:"MaxPoolSize"`
	MaxConnIdleTime       uint64   `yaml:"MaxConnIdleTime"`
	AutoEncryptionOptions bool     `yaml:"AutoEncryptionOptions"`
	ConnectTimeout        int      `yaml:"ConnectTimeout"`
	MinPoolSize           uint64   `yaml:"MinPoolSize"`
	RetryReads            bool     `yaml:"RetryReads"`
}

// 消息系统配置
type Mq struct {
	RedisMq `yaml:"RedisMq"`
	NatsMq  `yaml:"NatsMq"`
}

// Redis Pubsub 消息系统
type RedisMq struct {
	Switch      bool   `yaml:"Switch"`
	MaxActive   int    `yaml:"MaxActive"`
	MaxIdle     int    `yaml:"MaxIdle"`
	IdleTimeout int    `yaml:"IdleTimeout"`
	DbHost      string `yaml:"DbHost"`
	DbPort      int    `yaml:"DbPort"`
	DbAuth      bool   `yaml:"DbAuth"`
	DbPasswd    int    `yaml:"DbPasswd"`
}

// Nats Mq 消息系统
type NatsMq struct {
	Switch      bool         `yaml:"Switch"`
	NatsServers []NatsServer `yaml:"NatsServer"`
}

// 可配集群
type NatsServer struct {
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	AuthSwitch bool   `yaml:"AuthSwitch"`
	UserName   string `yaml:"UserName"`
	Password   string `yaml:"Password"`
}

// Cors配置
type Cors struct {
	AllowHeaders     []string `yaml:"AllowHeaders"`
	AllowCredentials bool     `yaml:"AllowCredentials"`
	ExposeHeaders    []string `yaml:"ExposeHeaders"`
	MaxAge           int      `yaml:"MaxAge"`
	AllowAllOrigins  bool     `yaml:"AllowAllOrigins"`
	AllowOrigins     []string `yaml:"AllowOrigins"`
	AllowMethods     []string `yaml:"AllowMethods"`
}

// OSS对象存储配置
type Oss struct {
	Qiniu  `yaml:"Qiniu"`
	Aliyun `yaml:"Aliyun"`
}

type Qiniu struct {
	Switch           bool   `yaml:"Switch"`
	AccessKey        string `yaml:"AccessKey"`
	SecretKey        string `yaml:"SecretKey"`
	Bucket           string `yaml:"Bucket"`
	UseHTTPS         bool   `yaml:"UseHTTPS"`
	UseCdnDomains    bool   `yaml:"UseCdnDomains"`
	UpTokenExpires   int    `yaml:"UpTokenExpires"`
	CallbackURL      string `yaml:"CallbackURL"`
	CallbackBodyType string `yaml:"CallbackBodyType"`
	EndUser          string `yaml:"EndUser"`
	FsizeMin         int    `yaml:"FsizeMin"`
	FsizeMax         int    `yaml:"FsizeLimit"`
	MimeLimit        string `yaml:"MimeLimit"`
}

type Aliyun struct {
	Switch          bool   `yaml:"Switch"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	ConnTimeout     int    `yaml:"ConnTimeout"`
	RWTimeout       int    `yaml:"RwTimeout"`
	EnableMD5       bool   `yaml:"EnableMD5"`
	EnableCRC       bool   `yaml:"EnableCRC"`
	AuthProxy       string `yaml:"AuthProxy"`
	Proxy           string `yaml:"Proxy"`
	AccessKeyId     string `yaml:"AccessKeyId"`
	BucketName      string `yaml:"BucketName"`
	Endpoint        string `yaml:"Endpoint"`
	UseCname        bool   `yaml:"UseCname"`
	SecurityToken   string `yaml:"SecurityToken"`
}

type Etcd struct {
	Endpoints            []string `yaml:"Endpoints"`
	TLS                  string   `yaml:"TLS"`
	Username             string   `yaml:"Username"`
	Password             string   `yaml:"Password"`
	PermitWithoutStream  bool     `yaml:"PermitWithoutStream"`
	RejectOldCluster     bool     `yaml:"RejectOldCluster"`
	AutoSyncInterval     uint     `yaml:"AutoSyncInterval"`
	DialTimeout          uint     `yaml:"DialTimeout"`
	DialKeepAliveTime    uint     `yaml:"DialKeepAliveTime"`
	DialKeepAliveTimeout uint     `yaml:"DialKeepAliveTimeout"`
	MaxCallRecvMsgSize   int      `yaml:"MaxCallRecvMsgSize"`
	MaxCallSendMsgSize   int      `yaml:"MaxCallSendMsgSize"`
}
