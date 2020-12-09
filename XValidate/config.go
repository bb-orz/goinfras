package XValidate

type Config struct {
	TransZh bool // 是否开启验证结果信息的中文翻译
}

func DefaultConfig() *Config {
	return &Config{
		true,
	}
}
