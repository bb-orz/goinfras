package XLogger

import (
	"errors"
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
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

func TestNewSyncErrorLogger(t *testing.T) {
	Convey("Test SyncError Logger", t, func() {
		var err error
		syncErrorLogger, err = NewSyncErrorLogger(DefaultConfig())
		So(err, ShouldBeNil)
		So(syncErrorLogger, ShouldNotBeNil)
		XSyncError().Debug("Log Debug Message...")
		XSyncError().Info("Log Info Message...")
		XSyncError().Warn("Log Warn Message...")
		XSyncError().Error("Log Error Message...")
	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XLogger Starter", t, func() {
		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)

		XCommon().Info("Test XCommon Info", zap.Any("info", "information"))
		XSyncError().Error("Test XSyncError Error", zap.Error(errors.New("XSyncError")))

	})
}
