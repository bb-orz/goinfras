# ZAP Logger Starter

> 基于 go.uber.org/zap 包

### Zap Logger Documentation
> Documentation https://pkg.go.dev/go.uber.org/zap#section-documentation

### Starter Usage
```
// 异步输出Writer
var writers []io.Writer
// 该启动器已内置文件输出、日期归档输出及mongo数据库输出，只需配置打开开关即可
// 若有其它输出需求，可与启动器注册时传递

var yourWriter io.Writer
// TODO your writer
writers = append(writers,yourWriter)
goinfras.RegisterStarter(XLogger.NewStarter(writers...))

```

### XLogger Config Setting
日志启动器定义的信息及异步输出开关，启动前可设置一下
```
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
LogDir             string // 日志记录文件目录
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
```