package etcd

import (
	"GoWebScaffold/infras"
	"context"
)

type EtcdStarter struct {
	infras.BaseStarter
}

func (s *EtcdStarter) Init(sctx *StarterContext) {
	client, err := NewEtcdClient(context.TODO(), sctx.GetConfig(), nil)
	sctx.SetEtcdClient(client)
}
