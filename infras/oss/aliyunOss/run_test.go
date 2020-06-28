package aliyunOss

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAliyunOssClient(t *testing.T) {
	Convey("Aliyun OSS Testing:", t, func() {
		err := RunForTesting(nil)
		So(err, ShouldBeNil)
	})

}
