package user

/*
Oauth 领域层：实现三方登录相关具体业务逻辑
封装领域层的错误信息并返回给调用者
*/
type OauthDomain struct {
	dao   *oauthDao
	cache *oauthCache
}

func NewOauthDomain() *OauthDomain {
	domain := new(OauthDomain)
	domain.dao = NewOauthDao()
	domain.cache = NewOauthCache()
	return domain
}
