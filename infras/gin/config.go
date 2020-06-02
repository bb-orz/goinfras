package gin

type ginConfig struct {
	listenPort int `val:"8090"`
}

// Cors配置
type corsConfig struct {
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	MaxAge           int
	AllowAllOrigins  bool
	AllowOrigins     []string
	AllowMethods     []string
}
