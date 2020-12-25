package XLogger

// 日志配置
type Config struct {
	AppName    string // 记录日志的应用名称
	AppVersion string // 版本
	DevEnv     bool   // 是否开发环境运行
	AddCaller  bool   // 是否注释每条信息所在文件名和行号

	// 开启可用的日志级别
	DisableDebugLevelSwitch  bool // 禁用Debug级别日志记录，默认false
	DisableInfoLevelSwitch   bool // 禁用Info级别日志记录，默认false
	DisableWarnLevelSwitch   bool // 禁用Warn级别日志记录，默认false
	DisableErrorLevelSwitch  bool // 禁用Error级别日志记录，默认false
	DisableDPanicLevelSwitch bool // 禁用DPanic级别日志记录，默认false
	DisablePanicLevelSwitch  bool // 禁用Panic级别日志记录，默认false
	DisableFatalLevelSwitch  bool // 禁用Fatal级别日志记录，默认false

	// 标准日志记录核心
	EnableStdZapCore bool // 是否启用标准输出核心,默认false

	// 文件日志记录核心
	EnableFileZapCore bool   // 是否启用简单文件日志记录器核心,默认false
	FileLogName       string // 日志记录文件路径
	SyncErrorLogName  string // 异步日志错误记录器日志文件路径

	// 归档文件记录核心
	EnableRotateZapCore bool   // 是否启用归档文件日志核心,默认false
	RotateLogDir        string // 归档日志记录目录
	RotateLogBaseName   string // 归档日志记录基本文件名
	WithRotationTime    uint   // 日志多久做一次归档,以天为单位
	MaxDayCount         uint   // 归档日志最多保留多久,以天为单位

	// 异步输出日志到mongo数据库记录核心
	EnableMongoLogZapCore bool   // 是否启用异步日志记录器核心：输出到外部储存系统,默认false
	MongoLogDbName        string // mongo数据库名称
	MongoLogCollection    string // mongo集合名称
	MongoLogExpire        int    // mongo日志超时时间,以天为单位
}

// 默认最小启动配置
func DefaultConfig() *Config {
	return &Config{
		AppName:    "",
		AppVersion: "",
		DevEnv:     true,
		AddCaller:  true,

		// 禁用的日志级别
		DisableDebugLevelSwitch:  false, // 禁用Debug级别日志记录，默认false
		DisableInfoLevelSwitch:   false, // 禁用Info级别日志记录，默认false
		DisableWarnLevelSwitch:   false, // 禁用Warn级别日志记录，默认false
		DisableErrorLevelSwitch:  false, // 禁用Error级别日志记录，默认false
		DisableDPanicLevelSwitch: false, // 禁用DPanic级别日志记录，默认false
		DisablePanicLevelSwitch:  false, // 禁用Panic级别日志记录，默认false
		DisableFatalLevelSwitch:  false, // 禁用Fatal级别日志记录，默认false

		// 标准日志记录核心
		EnableStdZapCore: true, // 是否启用标准输出核心,默认false

		// 文件日志记录核心
		EnableFileZapCore: false,                  // 是否启用简单文件日志记录器核心,默认false
		FileLogName:       "./log/common.log",     // 日志记录文件路径
		SyncErrorLogName:  "./log/sync_error.log", // 日志记录文件路径

		// 归档文件记录核心
		EnableRotateZapCore: false,    // 是否启用归档文件日志核心,默认false
		RotateLogDir:        "./log/", // 归档日志记录目录
		WithRotationTime:    1,        // 日志多久做一次归档,以天为单位
		MaxDayCount:         365,      // 归档日志最多保留多久,以天为单位

		// 异步输出日志记录核心
		EnableMongoLogZapCore: false,     // 是否启用异步日志记录器核心：输出到外部储存系统,默认false
		MongoLogDbName:        "zap_log", // mongo存储日志数据库名称
		MongoLogCollection:    "zap_log", // mongo集合名称
		MongoLogExpire:        365,       // mongo日志超时时间,以天为单位

	}
}
