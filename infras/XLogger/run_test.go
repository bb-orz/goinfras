package XLogger

import (
	"GoWebScaffold/infras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"os"
	"testing"
	"time"
)

func TestCommonLogger(t *testing.T) {
	Convey("Test Common Logger", t, func() {
		var err error
		err = CreateDefaultLogger(nil)
		So(err, ShouldBeNil)

		XCommon().Debug("Log Debug Message...")
		XCommon().Info("Log Info Message...")
		XCommon().Warn("Log Warn Message...")
		XCommon().Error("Log Error Message...")
	})
}

func TestCommonLoggerOutLogFile(t *testing.T) {
	Convey("Test Common Logger", t, func() {
		var err error
		// 注册日志记录启动器，并添加一个异步日志输出到文件
		fileWriter, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		So(err, ShouldBeNil)
		err = CreateDefaultLogger(nil, fileWriter)
		So(err, ShouldBeNil)

		XCommon().Debug("Log Debug Message...")
		XCommon().Info("Log Info Message...")
		XCommon().Warn("Log Warn Message...")
		XCommon().Error("Log Error Message...")
	})
}

func TestNewSyncErrorLogger(t *testing.T) {
	Convey("Test Sync Error Logger", t, func() {
		syncErrorLogger = NewSyncErrorLogger(DefaultConfig())
		So(syncErrorLogger, ShouldNotBeNil)
		XSyncError().Debug("Log Debug Message...")
		XSyncError().Info("Log Info Message...")
		XSyncError().Warn("Log Warn Message...")
		// SyncErrorLogger只会记录错误级别的日志
		XSyncError().Error("Log Error Message...")
	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XLogger Starter", t, func() {

		s := NewStarter()
		sctx := infras.CreateDefaultSystemContext()
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")
		s.Start(sctx)
		Println("Starter Start Successful!")
		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

		XCommon().Info("Test Logger Info", zap.Any("info", "information"))

		s.Stop()
		time.Sleep(time.Second * 3)
		XCommon().Info("Test Logger Info", zap.Any("info", "information"))

	})
}
