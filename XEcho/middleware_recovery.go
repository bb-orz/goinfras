package XEcho

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func RecoveryMiddleware(logger *zap.Logger, stack bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					var brokenPipe bool
					if ne, ok := err.(*net.OpError); ok {
						if se, ok := ne.Err.(*os.SyscallError); ok {
							if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
								brokenPipe = true
							}
						}
					}

					httpRequest, _ := httputil.DumpRequest(c.Request(), false)
					if brokenPipe {
						logger.Error(c.Request().URL.Path,
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
						)
						c.Error(echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
							"message": "[Server Error]: Broken Pipe",
						}))
						return
					}

					if stack {
						logger.Error("[Recovery from panic]",
							zap.Time("time", time.Now()),
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
							zap.String("stack", string(debug.Stack())),
						)
					} else {
						logger.Error("[Recovery from panic]",
							zap.Time("time", time.Now()),
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
						)
					}
					c.Error(echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
						"message": "Server Error",
					}))
				}
			}()

			return next(c)
		}
	}
}
