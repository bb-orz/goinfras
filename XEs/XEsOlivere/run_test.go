package XEsOlivere

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"goinfras"
	"testing"
)

func TestEsCommon(t *testing.T) {
	Convey("Test TestEsCommon", t, func() {

	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XEsOlivere Starter", t, func() {
		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")
		if s.Check(sctx) {
			Println("Starter Check Successful!")
		} else {
			Println("Starter Check Fail!")
		}

	})
}
