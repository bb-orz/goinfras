package XEtcd

import (
	"go.etcd.io/etcd/clientv3"
)

type EtcdCommon struct {
	client *clientv3.Client
}

/*简单的get set 操作*/
func NewEtcdCommon() *EtcdCommon {
	common := new(EtcdCommon)
	common.client = XClient()
	return common
}
