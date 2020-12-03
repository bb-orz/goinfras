package XAliyunSms

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			"https",
			"dysmsapi.aliyuncs.com",
			"",
			"",
			"",
			"",
			"SendSms",
			"",
			"",
		}
	}
	aliyunSmsClient, err = NewAliyunSmsClient(config)
	return err
}

func TestCommonSms(t *testing.T) {
	Convey("Aliyun SMS Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)
	})
}
