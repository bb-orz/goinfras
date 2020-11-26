package etcd

import (
	"GoWebScaffold/infras"
	"context"
	"go.etcd.io/etcd/clientv3"
)

var etcdClient *clientv3.Client

func EtcdClientV3() *clientv3.Client {
	infras.Check(etcdClient)
	return etcdClient
}

type EtcdStarter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *EtcdStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Etcd", &define)
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

/*For testing*/
func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			Endpoints: []string{"localhost:2379"},
		}
	}
	etcdClient, err = NewEtcdClient(context.TODO(), config, nil)
	return err
}
