package XEtcd

import "go.etcd.io/etcd/clientv3"

/*KeepAlive*/
type EtcdCommonKeepAlive struct {
	client *clientv3.Client
}

func NewEtcdCommonKeepAlive() *EtcdCommonKeepAlive {
	common := new(EtcdCommonKeepAlive)
	common.client = XClient()
	return common
}
