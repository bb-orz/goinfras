package XEtcd

import (
	"GoWebScaffold/infras"
	"context"
	"fmt"
)

type starter struct {
	infras.BaseStarter
	cfg Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = Config{}
	return starter
}

func (s *starter) Name() string {
	return "XEtcd"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Etcd", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	var err error
	client, err = NewEtcdClient(context.TODO(), &s.cfg, nil)
	infras.FailHandler(err)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	err := infras.Check(client)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: ETCD Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: ETCD Client Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop() {
	_ = client.Close()
}
