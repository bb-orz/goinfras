package XOAuth

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
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

	oAuthManager = new(OAuthManager)
	if config.QQSignSwitch {
		oAuthManager.QQOAuthManager = NewQQOauthManager(config)
	}
	if config.WechatSignSwitch {
		oAuthManager.WechatOAuthManager = NewWechatOAuthManager(config)
	}
	if config.WeiboSignSwitch {
		oAuthManager.WeiboOAuthManager = NewWeiboOAuthManager(config)
	}

	return nil
}

func TestOAuthComponent(t *testing.T) {
	Convey("OAuthManager Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO

	})
}
