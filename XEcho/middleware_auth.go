package XEcho

import (
	"github.com/labstack/echo/v4"
	"goinfras/XJwt"
	"net/http"
)

// 用户鉴权中间件
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1.从http头获取token string
			tkStr := c.Request().Header.Get("Authorization")
			// fmt.Println("token string:",tkStr)
			if tkStr == "" {
				c.Error(echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
					"message": "token string on http header is required!",
				}))
				return nil
			}

			// 2.解码校验token是否合法
			customerClaim, err := XJwt.XTokenUtils().Decode(tkStr)
			if err != nil {
				c.Error(echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
					"message": err.Error(),
				}))
				return nil
			}

			// 鉴权通过后设置用户信息
			c.Set("tkStr", tkStr)
			c.Set("userClaim", customerClaim.UserClaim)

			return next(c)
		}
	}
}
