package restful

import (
	"GoWebScaffold/infras/ginger"
	"GoWebScaffold/infras/jwt"
	"GoWebScaffold/services"
	"github.com/gin-gonic/gin"
)

/*
API层，调用相关Service，封装响应返回，并记录日志
*/

func init() {
	// 初始化时注册该模块API
	Xgin.RegisterApi(new(UserApi))
}

type UserApi struct {
	userService services.IUserService
	mailService services.IMailService
	smsService  services.ISmsService
}

// 设置该模块的API Router
func (api *UserApi) SetRoutes() {
	api.userService = services.GetUserService()
	engine := Xgin.GinEngine()

	// 如TokenUtils服务已初始化，添加中间件
	var authMiddlerware gin.HandlerFunc
	if tku := XJWT.TokenUtils(); tku != nil {
		authMiddlerware = Xgin.JwtAuthMiddleware(tku)
	}

	engine.POST("/login", api.loginHandler)
	engine.POST("/logout", api.logoutHandler)

	registerGroup := engine.Group("/register")
	registerGroup.POST("/email", api.registerEmailHandler)
	registerGroup.POST("/phone", api.registerPhoneHandler)

	oauthGroup := engine.Group("/oauth")
	oauthGroup.GET("/qq", api.oauthQQHandler)
	oauthGroup.GET("/weixin", api.oauthWeixinHandler)
	oauthGroup.GET("/weibo", api.oauthWeiboHandler)

	userGroup := engine.Group("/user", authMiddlerware)
	userGroup.GET("/get", api.getUserInfoHandler)
	userGroup.POST("/set", api.setUserInfoHandler)
}

/*用户登录*/
func (api *UserApi) loginHandler(ctx *gin.Context) {

}

/*用户登出*/
func (api *UserApi) logoutHandler(ctx *gin.Context) {

}

/*邮箱注册注册*/
func (api *UserApi) registerEmailHandler(ctx *gin.Context) {

}

/*手机号码注册注册*/
func (api *UserApi) registerPhoneHandler(ctx *gin.Context) {

}

/*qq oauth 登录*/
func (api *UserApi) oauthQQHandler(ctx *gin.Context) {

}

/*微信oauth 登录*/
func (api *UserApi) oauthWeixinHandler(ctx *gin.Context) {

}

/*微博oauth登录*/
func (api *UserApi) oauthWeiboHandler(ctx *gin.Context) {

}

/*设置用户信息*/

func (api *UserApi) setUserInfoHandler(ctx *gin.Context) {

}

/*获取用户信息*/
func (api *UserApi) getUserInfoHandler(ctx *gin.Context) {

}
