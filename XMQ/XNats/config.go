package XNats

// Nats Mq 消息系统
type Config struct {
	NatsServers []natsServer
}

// 可配集群
type natsServer struct {
	Host     string
	Port     int
	UserName string
	Password string
}

func DefaultConfig() *Config {
	return &Config{
		NatsServers: []natsServer{
			{
				"127.0.0.1",
				4222,
				"",
				"",
			},
		},
	}
}
