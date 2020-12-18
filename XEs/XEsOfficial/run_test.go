package XEsOfficial

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XEs Starter", t, func() {
		s := NewStarter(nil)
		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
	})
}
