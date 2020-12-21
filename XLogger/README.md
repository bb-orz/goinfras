# ZAP Logger Starter

> 基于 go.uber.org/zap 包

### Zap Logger Documentation
> Documentation https://pkg.go.dev/go.uber.org/zap#section-documentation

### Starter Usage
该启动器内置了标准输出，文件输出，归档文件输出，mongo日志输出，只需在配置文件打开定义即可，若有其它输出需求，用户可自定义传入启动器。
```
// TODO define your outputs
// 定义日志输出
//type LoggerOutput struct {
//	Format           zapcore.EncoderConfig
//	Writer           io.Writer
//	LevelEnablerFunc zap.LevelEnablerFunc
//}


outputs = append(outputs,youroutputs...)


goinfras.RegisterStarter(XLogger.NewStarter(outputs...))

```

### XLogger Config Setting
日志启动器定义的信息及异步输出开关，启动前可设置一下
```
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
```

### XLogger Usage
该启动器提供两个基本的日志记录实例
- XLogger.XCommon()  =>  通用日志记录器
- XLogger.XSyncError()  => 异步输出错误记录器，用于记录异步输出时出现的错误

```
// 通用记录器用法
XLogger.XCommon().Debug("Log Debug Message...")
XLogger.XCommon().Info("Log Info Message...")
XLogger.XCommon().Warn("Log Warn Message...")
XLogger.XCommon().Error("Log Error Message...")
... 
```