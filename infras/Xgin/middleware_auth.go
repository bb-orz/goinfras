package Xgin

import (
	"GoWebScaffold/infras/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户鉴权中间件
func JwtAuthMiddleware(tku XJWT.ITokenUtils) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.从http头获取token string
		tkStr := c.GetHeader("Authorization")
		// fmt.Println("token string:",tkStr)
		if tkStr == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token string on http header is required!",
			})
			return
		}

		// 2.解码校验token是否合法
		customerClaim, err := tku.Decode(tkStr)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		// 鉴权通过后设置用户信息
		c.Set("tkStr", tkStr)
		c.Set("userClaim", customerClaim.UserClaim)

		c.Next()
	}
}