package hook

import (
	"GoWebScaffold/infras/config"
	"io"
	"os"
)

//简单文件日志输出
func SimpleFileLogHook(filename string, appConf *config.AppConfig) io.Writer {
	var file io.Writer
	var err error
	fullFileName := appConf.LogConf.LogDir + filename + ".log"
	file, err = os.OpenFile(fullFileName, os.O_RDWR, os.ModeAppend)
	if os.IsNotExist(err) {
		file, err = os.Create(fullFileName)
		if err != nil {
			panic(err)
		}
	}
	return file
}
