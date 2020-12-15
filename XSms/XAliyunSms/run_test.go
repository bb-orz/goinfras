package XAliyunSms

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
)

func TestCommonSms(t *testing.T) {
	Convey("Aliyun SMS Testing:", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		response, err := XCommonSms().SendSmsMsg("", "", "", "", "")
		So(err, ShouldBeNil)
		Println("Send Sms Status:", response.IsSuccess())

		smsResponse, err := XCommonSms().SendBatchSmsMsg("", "", "", []string{""}, []string{""})
		So(err, ShouldBeNil)
		Println("Send Batch Sms Status:", smsResponse.IsSuccess())
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
