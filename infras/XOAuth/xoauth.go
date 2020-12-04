package XOAuth

var oAuthManager *OAuthManager

type OAuthManager struct {
	WechatOAuthManager *WechatOAuthManager
	WeiboOAuthManager  *WeiboOAuthManager
	QQOAuthManager     *QQOAuthManager
}

func XManager() *OAuthManager {
	return oAuthManager
}

// 资源组件闭包执行
func XFManager(f func(m *OAuthManager) error) error {
	return f(oAuthManager)
}
