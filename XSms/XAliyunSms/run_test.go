package XAliyunSms

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
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

		logger := goinfras.NewCommandLineStarterLogger("debug")
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
	})
}
