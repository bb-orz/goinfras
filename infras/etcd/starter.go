package etcd

import (
	"GoWebScaffold/infras"
	"context"
	"go.etcd.io/etcd/clientv3"
)

var client *clientv3.Client

func ClientV3() *clientv3.Client {
	infras.Check(client)
	return client
}

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Etcd", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	client, err = NewEtcdClient(context.TODO(), s.cfg, nil)
	infras.FailHandler(err)
	sctx.Logger().Info("EtcdClientV3 Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	_ = ClientV3().Close()
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
	client, err = NewEtcdClient(context.TODO(), config, nil)
	return err
}
