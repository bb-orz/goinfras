package hook

import (
	"GoWebScaffold/infras/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

// 按日期归档记录的日志输出
func RotateFileLogHook(filename string, appConf *config.AppConfig) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	rotateLogHook, err := rotatelogs.New(
		appConf.LogConf.LogDir+filename+"[%Y-%m-%d %H:%M:%S].log",
		rotatelogs.WithLinkName(filename),
		// 最多保留多久
		rotatelogs.WithMaxAge(time.Hour*time.Duration(appConf.LogConf.MaxDayCount*24)),
		// 多久做一次归档
		rotatelogs.WithRotationTime(time.Hour*24*time.Duration(appConf.LogConf.WithRotationTime)),
	)

	if err != nil {
		panic(err)
	}
	return rotateLogHook
}
