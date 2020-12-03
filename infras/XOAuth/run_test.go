package XOAuth

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var om *OAuthManager

	if config == nil {
		config = &Config{
			false,
			"",
			"",
			false,
			"",
			"",
			false,
			"",
			"",
		}
	}

	om = new(OAuthManager)
	if config.QQSignSwitch {
		om.QQ = NewQQOauthManager(config)
	}
	if config.WechatSignSwitch {
		om.Wechat = NewWechatOAuthManager(config)
	}
	if config.WeiboSignSwitch {
		om.Weibo = NewWeiboOAuthManager(config)
	}

	SetComponent(om)
	return nil
}

func TestOAuthComponent(t *testing.T) {
	Convey("OAuthManager Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO

	})
}
