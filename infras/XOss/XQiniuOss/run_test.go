package XQiniuOss

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQiniuOssClient(t *testing.T) {
	Convey("Qiniu OSS Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO
	})

}
