package XAliyunSms

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCommonSms(t *testing.T) {
	Convey("Aliyun SMS Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO
	})
}
