package restful

import (
	"GoWebScaffold/infras/ginger"
	"GoWebScaffold/infras/jwt"
	"GoWebScaffold/services"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化时注册该模块API
	ginger.RegisterApiModule(new(UserApi))
}

type UserApi struct {
	service services.IUserService
}

// 设置该模块的API Router
func (a *UserApi) SetRoutes() {
	a.service = services.GetUserService()
	engine := ginger.GinEngine()

	// 如TokenUtils服务已初始化，添加中间件
	var authMiddlerware gin.HandlerFunc
	if tku := jwt.TokenUtils(); tku != nil {
		authMiddlerware = ginger.JwtAuthMiddleware(tku)
	}

	engine.POST("/register", registerHandler)
	engine.POST("/login", loginHandler)
	engine.POST("/logout", logoutHandler)

	oauthGroup := engine.Group("/oauth")
	oauthGroup.GET("/qq", oauthQQHandler)
	oauthGroup.GET("/weixin", oauthWeixinHandler)
	oauthGroup.GET("/weibo", oauthWeiboHandler)

	userGroup := engine.Group("/user", authMiddlerware)
	userGroup.GET("/get", getUserInfoHandler)
	userGroup.POST("/set_info", setUserInfoHandler)
}

/*用户注册*/
func registerHandler(ctx *gin.Context) {

}

/*用户登录*/
func loginHandler(ctx *gin.Context) {

}

/*用户登出*/
func logoutHandler(ctx *gin.Context) {

}

/*qq oauth 登录*/
func oauthQQHandler(ctx *gin.Context) {

}

/*微信oauth 登录*/
func oauthWeixinHandler(ctx *gin.Context) {

}

/*微博oauth登录*/
func oauthWeiboHandler(ctx *gin.Context) {

}

/*设置用户信息*/

func setUserInfoHandler(ctx *gin.Context) {

}

/*获取用户信息*/
func getUserInfoHandler(ctx *gin.Context) {

}
