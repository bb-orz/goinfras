package XMail

import (
	"GoWebScaffold/infras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
)

func TestCommonMail(t *testing.T) {
	Convey("Test Common Mail", t, func() {
		CreateDefaultManager(nil)

		err := XCommonMail().SendSimpleMail(
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			[]string{""},
		)

		So(err, ShouldBeNil)
	})
}

func TestStarter(t *testing.T) {
	Convey("Test XMail Starter", t, func() {
		s := NewStarter()
		sctx := infras.CreateDefaultStarterContext(nil, zap.L())
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")

		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

		err := XCommonMail().SendSimpleMail(
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			[]string{""},
		)

		So(err, ShouldBeNil)

	})
}
