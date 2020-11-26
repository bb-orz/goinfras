package logger

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"
	"os"
	"testing"
)

func TestCommonLogger(t *testing.T) {
	Convey("Test Common Logger", t, func() {
		err := RunForTesting(nil)
		So(err, ShouldNotBeNil)
		CommonLogger().Debug("Log Debug Message...")
		CommonLogger().Info("Log Info Message...")
		CommonLogger().Warn("Log Warn Message...")
		CommonLogger().Error("Log Error Message...")
	})
}

func TestCommonLoggerOutLogFile(t *testing.T) {
	Convey("Test Common Logger", t, func() {
		config := Config{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Config:", config)

		// 注册日志记录启动器，并添加一个异步日志输出到文件
		file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		So(err, ShouldBeNil)
		logger := NewCommonLogger(&config, file)
		So(logger, ShouldNotBeNil)
		logger.Debug("Log Debug Message...")
		logger.Info("Log Info Message...")
		logger.Warn("Log Warn Message...")
		logger.Error("Log Error Message...")
	})
}

func TestNewSyncErrorLogger(t *testing.T) {
	Convey("Test Sync Error Logger", t, func() {
		config := Config{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Config:", config)
		logger := NewSyncErrorLogger(&config)
		So(logger, ShouldNotBeNil)
		logger.Debug("Log Debug Message...")
		logger.Info("Log Info Message...")
		logger.Warn("Log Warn Message...")
		logger.Error("Log Error Message...")
	})
}
