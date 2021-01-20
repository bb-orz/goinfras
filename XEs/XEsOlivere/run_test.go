package XEsOlivere

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
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
		logger := goinfras.NewCommandLineStarterLogger("debug")
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
	})
}
