package XOAuth

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
)

func TestQQOAuth(t *testing.T) {
	Convey("TestQQOAuth", t, func() {

		qqOM = NewQQOauthManager(nil)
		// TODO 先获取预授权码
		var qqprecode string
		qqOAuthResult := XQQOAuthManager().Authorize(qqprecode)
		Println("qqOAuthResult", qqOAuthResult)
	})
}

func TestWeiboOAuth(t *testing.T) {
	Convey("TestWeiboOAuth", t, func() {
		weiboOM = NewWeiboOAuthManager(nil)

		// TODO 先获取预授权码
		var weiboprecode string
		weiboOAuthResult := XWeiboOAuthManager().Authorize(weiboprecode)
		Println("weiboOAuthResult", weiboOAuthResult)
	})
}

func TestWechatOAuth(t *testing.T) {
	Convey("TestWechatOAuth", t, func() {
		wechatOM = NewWechatOAuthManager(nil)

		// TODO 先获取预授权码
		var wechatprecode string
		wechatOAuthResult := XWechatOAuthManager().Authorize(wechatprecode)
		Println("wechatOAuthResult", wechatOAuthResult)
	})
}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
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
