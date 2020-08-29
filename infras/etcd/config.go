package etcd

import "crypto/tls"

type EtcdConfig struct {
	Endpoints            []string `val:"localhost:2379"`
	TLS                  *tls.Config
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
