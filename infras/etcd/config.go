package etcd

type etcdConfig struct {
	Endpoints            []string
	TLS                  string
	Username             string
	Password             string
	PermitWithoutStream  bool
	RejectOldCluster     bool
	AutoSyncInterval     uint
	DialTimeout          uint
	DialKeepAliveTime    uint
	DialKeepAliveTimeout uint
	MaxCallRecvMsgSize   int
	MaxCallSendMsgSize   int
}
