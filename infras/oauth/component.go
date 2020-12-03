package oauth

import (
	"GoWebScaffold/hub"
	"GoWebScaffold/infras"
)

var oAuthManager *OAuthManager

type OAuthManager struct {
	Wechat *WechatOAuthManager
	Weibo  *WeiboOAuthManager
	QQ     *QQOAuthManager
}

func SetComponent(m *OAuthManager) {
	oAuthManager = m
}

func OAuthComponent() *OAuthManager {
	infras.Check(oAuthManager)
	return oAuthManager
}
