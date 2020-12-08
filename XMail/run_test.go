package XMail

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"goinfras"
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
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")

		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

		err = XCommonMail().SendSimpleMail(
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
