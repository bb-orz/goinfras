package XOAuth

// OAuth2服务的通用鉴权接口
type OAuth interface {
	Authorize(code string) OAuthResult                                    // 根据微信返回的accessTokenCode开始鉴权并获取用户信息
	getAccessToken(code string) map[string]interface{}                    // 从平台回调拿到accessTokenCode后调用此方法可拿到accessToken
	getUserInfo(accessToken string, openId string) map[string]interface{} // 拿到accessToken和openId后获取用户信息
}

// OAuth鉴权结果
type OAuthResult struct {
	Result   bool
	UserInfo *OAuthAccountInfo
	Error    error
}

// OAuth授权获取的第三方账号信息
type OAuthAccountInfo struct {
	AccessToken string
	OpenId      string
	UnionId     string
	NickName    string
	Gender      uint
	AvatarUrl   string
}
