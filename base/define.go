package base

//基础配置
type Base struct {
	Env        string `yaml:"Env"`
	ListenPort int64  `yaml:"Listen"`
}



// Cors配置
type Cors struct {
	AllowHeaders     []string `yaml:"AllowHeaders"`
	AllowCredentials bool     `yaml:"AllowCredentials"`
	ExposeHeaders    []string `yaml:"ExposeHeaders"`
	MaxAge           int      `yaml:"MaxAge"`
	AllowAllOrigins  bool     `yaml:"AllowAllOrigins"`
	AllowOrigins     []string `yaml:"AllowOrigins"`
	AllowMethods     []string `yaml:"AllowMethods"`
}

