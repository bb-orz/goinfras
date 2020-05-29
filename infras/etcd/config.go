package etcd

type etcdConfig struct {
	Endpoints            []string `yaml:"Endpoints"`
	TLS                  string   `yaml:"TLS"`
	Username             string   `yaml:"Username"`
	Password             string   `yaml:"Password"`
	PermitWithoutStream  bool     `yaml:"PermitWithoutStream"`
	RejectOldCluster     bool     `yaml:"RejectOldCluster"`
	AutoSyncInterval     uint     `yaml:"AutoSyncInterval"`
	DialTimeout          uint     `yaml:"DialTimeout"`
	DialKeepAliveTime    uint     `yaml:"DialKeepAliveTime"`
	DialKeepAliveTimeout uint     `yaml:"DialKeepAliveTimeout"`
	MaxCallRecvMsgSize   int      `yaml:"MaxCallRecvMsgSize"`
	MaxCallSendMsgSize   int      `yaml:"MaxCallSendMsgSize"`
}
