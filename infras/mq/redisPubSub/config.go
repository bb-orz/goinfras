package redisPubSub

type redisPubSubConfig struct {
	Switch      bool   `yaml:"Switch"`
	MaxActive   int    `yaml:"MaxActive"`
	MaxIdle     int    `yaml:"MaxIdle"`
	IdleTimeout int    `yaml:"IdleTimeout"`
	DbHost      string `yaml:"DbHost"`
	DbPort      int    `yaml:"DbPort"`
	DbAuth      bool   `yaml:"DbAuth"`
	DbPasswd    int    `yaml:"DbPasswd"`
}
