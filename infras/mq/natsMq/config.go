package natsMq

// Nats Mq 消息系统
type natsMqConfig struct {
	Switch      bool         `yaml:"Switch"`
	NatsServers []natsServer `yaml:"NatsServer"`
}

// 可配集群
type natsServer struct {
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	AuthSwitch bool   `yaml:"AuthSwitch"`
	UserName   string `yaml:"UserName"`
	Password   string `yaml:"Password"`
}
