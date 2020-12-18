package XAliyunOss

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCommonOss(t *testing.T) {
	Convey("TestAliyunOssClient", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		// 一些通用的简单操作
		err = XCommonOss().UploadString("", "", "")
		So(err, ShouldBeNil)

		err = XCommonOss().AppendUpload("", "", "")
		So(err, ShouldBeNil)

		err = XCommonOss().Uploadfile("", "", "")
		So(err, ShouldBeNil)

		err = XCommonOss().DownLoadFile("", "", "")
		So(err, ShouldBeNil)

		err = XCommonOss().LimitConditionDownload("", "", "")
		So(err, ShouldBeNil)

		err = XCommonOss().CompressDownload("", "", "")
		So(err, ShouldBeNil)

		rs1, err := XCommonOss().RangeDownload("", "", 0, 10)
		So(err, ShouldBeNil)
		Println(rs1)

		rs2, err := XCommonOss().StreamDownload("", "")
		So(err, ShouldBeNil)
		Println(rs2)

	})
}

func TestBreakPointOss(t *testing.T) {
	Convey("TestAliyunOssClient", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		// 一些通用的简单操作
		err = XBreakPointOss().BreakPointUpload("", "", "")
		So(err, ShouldBeNil)
		err = XBreakPointOss().BreakPointDownload("", "", "")
		So(err, ShouldBeNil)

	})
}

func TestMultipartOss(t *testing.T) {
	Convey("TestAliyunOssClient", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		// 一些通用的简单操作
		result, err := XMultipartOss().MultipartUpload("", "", "")
		So(err, ShouldBeNil)
		Println(result)

		err = XMultipartOss().CancelMultipartUpload("", "")
		So(err, ShouldBeNil)

	})
}

func TestProgressOss(t *testing.T) {
	Convey("TestAliyunOssClient", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		// 一些通用的简单操作
		err = XProgressOss().ProgressUpload("", "", "")
		So(err, ShouldBeNil)

		err = XProgressOss().ProgressDownload("", "", "")
		So(err, ShouldBeNil)

	})
}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)

		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
	})
}
