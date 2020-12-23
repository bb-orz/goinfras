package XGocache

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		var err error
		CreateDefaultCache(nil)
		So(err, ShouldBeNil)

		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
	})
}
