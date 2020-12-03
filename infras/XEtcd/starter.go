package XEtcd

import (
	"GoWebScaffold/infras"
	"context"
)

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *Starter {
	starter := new(Starter)
	starter.cfg = Config{}
	return starter
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Etcd", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	client, err = NewEtcdClient(context.TODO(), &s.cfg, nil)
	infras.FailHandler(err)
	sctx.Logger().Info("EtcdClientV3 Setup Successful!")
}

func (s *Starter) Stop() {
	_ = client.Close()
}
