package XEtcd

import "go.etcd.io/etcd/clientv3"

/*watch用来获取未来更改的通知*/
type EtcdCommonWatch struct {
	client *clientv3.Client
}

func NewEtcdCommonWatch() *EtcdCommonWatch {
	common := new(EtcdCommonWatch)
	common.client = XClient()
	return common
}
