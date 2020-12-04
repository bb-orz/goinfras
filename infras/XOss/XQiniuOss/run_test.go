package XQiniuOss

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			"",
			"",
			"",
			false,
			false,
			7200,
			"",
			"",
			"",
			1024,
			10485760,
			"",
		}
	}
	qiniuOssClient = NewQnClient(config)
	return err
}

func TestQiniuOssClient(t *testing.T) {
	Convey("Qiniu OSS Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)
	})

}
