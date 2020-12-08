package XAliyunOss

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
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

		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := CreateDefaultStarterContext(nil, logger)
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
