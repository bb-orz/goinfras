package logger

// 日志配置
type LoggerConfig struct {
	AppName            string // 记录日志的应用名称
	AppVersion         string // 版本
	DevEnv             bool   // 是否开发环境运行
	AddCaller          bool   // 是否注释每条信息所在文件名和行号
	DebugLevelSwitch   bool   // Debug级别日志记录开关
	InfoLevelSwitch    bool   // Info级别日志记录开关
	WarnLevelSwitch    bool   // Warn级别日志记录开关
	ErrorLevelSwitch   bool   // Error级别日志记录开关
	DPanicLevelSwitch  bool   // DPanic级别日志记录开关
	PanicLevelSwitch   bool   // Panic级别日志记录开关
	FatalLevelSwitch   bool   // Fatal级别日志记录开关
	LogDir             string // 日志记录目录
	SimpleZapCore      bool   // 是否启用简单的日志记录器核心:只输出到stdout和file
	RotateZapCore      bool   // 是否启用归档日志核心
	SyncZapCore        bool   // 是否启用异步日志记录器核心：输出到外部储存系统
	SyncLogSwitch      bool   // 添加异步输出的开关
	StdoutLogSwitch    bool   // 标准输出日志开关
	RotateLogSwitch    bool   // 是否启用归档日志记录器核心:只输出到stdout和归档日期file
	WithRotationTime   int    // 日志多久做一次归档
	MaxDayCount        int    // 归档日志最多保留多久
	MongoLogSwitch     bool   // mongo存储日志开关
	MongoLogCollection string // mongo集合名称
	MongoLogExpire     int    // mongo日志超时时间
}
