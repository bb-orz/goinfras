package logger

// 日志配置
type loggerConfig struct {
	AppName            string
	AppVersion 		   string
	DevEnv 			   bool	  `val:"true"`
	AddCaller		   bool	  `val:"true"`
	DebugLevelSwitch   bool   `val:"false"`
	InfoLevelSwitch    bool   `val:"true"`
	WarnLevelSwitch    bool   `val:"true"`
	ErrorLevelSwitch   bool   `val:"true"`
	DPanicLevelSwitch  bool   `val:"true"`
	PanicLevelSwitch   bool   `val:"false"`
	FatalLevelSwitch   bool   `val:"true"`
	SimpleZapCore      bool   `val:"true"`
	SyncZapCore        bool   `val:"false"`
	SyncLogSwitch      bool   `val:"true"`
	StdoutLogSwitch    bool   `val:"true"`
	RotateLogSwitch    bool   `val:"false"`
	LogDir             string `val:"../../log"`
	WithRotationTime   int
	MaxDayCount        int
	RotateZapCore      bool
	MongoLogSwitch     bool
	MongoLogCollection string
	MongoLogExpire     int
}
