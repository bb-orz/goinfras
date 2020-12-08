package XAliyunSms

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"goinfras"
	"testing"
)

func TestCommonSms(t *testing.T) {
	Convey("Aliyun SMS Testing:", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		sms := XCommonSms(DefaultConfig())

		response, err := sms.SendSmsMsg("", "")
		So(err, ShouldBeNil)
		Println("Send Sms Status:", response.IsSuccess())

	})
}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

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

	})
}
