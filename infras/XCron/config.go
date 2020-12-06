package XCron

type Config struct {
	Location string // 定时器时区设置
}

func DefaultConfig() *Config {
	return &Config{Location: "Local"}
}
