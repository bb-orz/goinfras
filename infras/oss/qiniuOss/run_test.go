package qiniuOss

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQiniuOssClient(t *testing.T) {
	Convey("Qiniu OSS Testing:", t, func() {
		err := RunForTesting(nil)
		So(err, ShouldBeNil)
	})

}
