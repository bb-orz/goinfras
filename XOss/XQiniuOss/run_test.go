package XQiniuOss

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"goinfras"
	"testing"
)

// 客户端上传获取token测试
func TestQiniuOssClientUpload(t *testing.T) {
	Convey("TestQiniuOssClient", t, func() {
		CreateDefaultClient(nil)

		// TODO
		upToken := XClient().SimpleUpload("")
		Println("Client Upload Token:", upToken)

		token := XClient().OverwriteUpload("", "")
		Println("Client Overwrite Upload Token:", token)

		callbackUploadToken := XClient().CallbackUpload("")
		Println("Client Callback Upload Token:", callbackUploadToken)

	})
}

// 服务端断点上传测试
func TestQiniuOssServerBreakPointUpload(t *testing.T) {
	Convey("TestQiniuOssServerBreakPointUpload", t, func() {
		CreateDefaultClient(nil)

		putRet, err := XClient().BreakPointUpload("", "", "", "")
		So(err, ShouldBeNil)
		Println("BreakPointUpload Key:", putRet.Key)
		Println("BreakPointUpload PersistentID:", putRet.PersistentID)
		Println("BreakPointUpload Hash:", putRet.Hash)

	})
}

// 服务端断点上传测试
func TestQiniuOssServerFormUpload(t *testing.T) {
	Convey("TestQiniuOssServerFormUpload", t, func() {
		CreateDefaultClient(nil)

		// 服务器表单上传
		putRet1, err := XClient().FormUploadWithLocalFile("", "", "")
		So(err, ShouldBeNil)
		Println("FormUploadWithLocalFile Key:", putRet1.Key)
		Println("FormUploadWithLocalFile PersistentID:", putRet1.PersistentID)
		Println("FormUploadWithLocalFile Hash:", putRet1.Hash)

		// 服务器字节数组上传
		var data []byte
		putRet2, err := XClient().FormUploadWithByteSlice("", "", data)
		So(err, ShouldBeNil)
		Println("FormUploadWithByteSlice Key:", putRet2.Key)
		Println("FormUploadWithByteSlice PersistentID:", putRet2.PersistentID)
		Println("FormUploadWithByteSlice Hash:", putRet2.Hash)

	})
}

// 服务端断点上传测试
func TestQiniuOssServerMultipartUpload(t *testing.T) {
	Convey("TestQiniuOssServerMultipartUpload", t, func() {
		CreateDefaultClient(nil)

		putRet, err := XClient().MultipartUpload("", "", "")
		So(err, ShouldBeNil)
		Println("MultipartUpload Key:", putRet.Key)
		Println("MultipartUpload PersistentID:", putRet.PersistentID)
		Println("MultipartUpload Hash:", putRet.Hash)

	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		CreateDefaultClient(nil)

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