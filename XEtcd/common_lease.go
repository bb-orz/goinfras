package XEtcd

import "go.etcd.io/etcd/clientv3"

/*lease 租约常用操作*/
type EtcdCommonLease struct {
	client *clientv3.Client
}

func NewEtcdCommonLease() *EtcdCommonLease {
	common := new(EtcdCommonLease)
	common.client = XClient()
	return common
}
