package etcd

import (
	"GoWebScaffold/infras"
	"context"
	"github.com/tietang/props/kvs"
	"go.etcd.io/etcd/clientv3"
)

var etcdClient *clientv3.Client

func EtcdClientV3() *clientv3.Client {
	infras.Check(etcdClient)
	return etcdClient
}

type EtcdStarter struct {
	infras.BaseStarter
	cfg *etcdConfig
}

func (s *EtcdStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := etcdConfig{}
	err := kvs.Unmarshal(configs, &define, "Etcd")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *EtcdStarter) Setup(sctx *infras.StarterContext) {
	var err error
	etcdClient, err = NewEtcdClient(context.TODO(), s.cfg, nil)
	infras.FailHandler(err)
	sctx.Logger().Info("EtcdClientV3 Setup Successful!")
}

func (s *EtcdStarter) Stop(sctx *infras.StarterContext) {
	_ = EtcdClientV3().Close()
	sctx.Logger().Info("EtcdClientV3 Closed!")
}
