package XNats

// Nats Mq 消息系统
type Config struct {
	Switch      bool
	NatsServers []natsServer
}

// 可配集群
type natsServer struct {
	Host       string
	Port       int
	AuthSwitch bool
	UserName   string
	Password   string
}

func DefaultConfig() *Config {
	return &Config{
		Switch: true,
		NatsServers: []natsServer{
			{
				"127.0.0.1",
				4222,
				false,
				"",
				"",
			},
		},
	}
}
