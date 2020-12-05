package XAliyunOss

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAliyunOssClient(t *testing.T) {
	Convey("Aliyun OSS Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)
	})

	// TODO

}
