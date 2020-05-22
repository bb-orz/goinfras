package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/etcd"
	"context"
)

type EtcdStarter struct {
	infras.BaseStarter
}

func (s *EtcdStarter) Init(sctx *StarterContext) {
	client, err := etcd.NewEtcdClient(context.TODO(), sctx.GetConfig(), nil)
	sctx.SetEtcdClient(client)
}
