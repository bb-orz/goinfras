package gin

type ginConfig struct {
	listenPort int `val:"8090"`
	cors       *corsConfig
}

// Cors配置
type corsConfig struct {
	AllowAllOrigins  bool
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	MaxAge           int
	AllowOrigins     []string
	AllowMethods     []string
}
