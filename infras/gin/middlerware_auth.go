package gin

import (
	"GoWebScaffold/infras/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 用户鉴权中间件
func AuthMiddleware(tks *jwt.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从http头获取token string

		tkStr := c.GetHeader("Authorization")
		// fmt.Println("token string:",tkStr)
		if tkStr == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token string on http header is required!",
			})
			return
		}

		// 解码校验token是否合法
		customerClaim, err := tks.Decode(tkStr)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token string is invalid!",
			})
			return
		}

		fmt.Println("customerClaim:", customerClaim)

		// 在Redis查找token是否存在，不存在或过期则返回-1，还存在则返回token值
		key := tokenCache.UserTokenCacheKeyPrefix + strconv.Itoa(int(customerClaim.TokenUserClaim.Id))
		token := tokenCache.GetToken(key)
		// fmt.Println(key)
		// fmt.Println(tkStr)

		// 校验客户端token和服务端缓存的token
		if token != tkStr {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is not exist,please retry to sign in!",
			})
			return
		}

		// 鉴权通过后设置用户信息
		c.Set("tkStr", tkStr)
		c.Set("userClaim", customerClaim.TokenUserClaim)

		c.Next()
	}
}
