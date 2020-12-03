package XAliyunOss

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
			60,
			60,
			false,
			false,
			"",
			"",
			"",
			"",
			"http://oss-cn-shenzhen.aliyuncs.com",
			false,
			"",
		}
	}

	aliyunOssClient, err = NewClient(config)
	return err
}

func TestAliyunOssClient(t *testing.T) {
	Convey("Aliyun OSS Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)
	})

	// TODO

}
