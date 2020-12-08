package XOAuth

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"goinfras"
	"testing"
)

func TestQQOAuth(t *testing.T) {
	Convey("TestOAuthManager", t, func() {
		CreateDefaultManager(nil)

		// TODO 先获取预授权码
		var qqprecode string
		qqOAuthResult := XManager().QQOAuthManager.Authorize(qqprecode)
		Println("qqOAuthResult", qqOAuthResult)
	})
}

func TestWeiboOAuth(t *testing.T) {
	Convey("TestOAuthManager", t, func() {
		CreateDefaultManager(nil)

		// TODO 先获取预授权码
		var weiboprecode string
		weiboOAuthResult := XManager().WeiboOAuthManager.Authorize(weiboprecode)
		Println("weiboOAuthResult", weiboOAuthResult)
	})
}

func TestWechatOAuth(t *testing.T) {
	Convey("TestOAuthManager", t, func() {
		CreateDefaultManager(nil)

		// TODO 先获取预授权码
		var wechatprecode string
		wechatOAuthResult := XManager().WechatOAuthManager.Authorize(wechatprecode)
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
