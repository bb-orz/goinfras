package XEtcd

import (
	"context"
	"fmt"
	"github.com/bb-orz/goinfras"
	"go.uber.org/zap"
)

type starter struct {
	goinfras.BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XEtcd"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Etcd", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	sctx.Logger().Info("Print ETCD Config:", zap.Any("EtcdConfig", *define))
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	client, err = NewEtcdClient(context.TODO(), s.cfg, nil)
	goinfras.ErrorHandler(err)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	var err error
	err = goinfras.Check(client)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: ETCD Client Setup Fail!", s.Name()))
		return false
	}
	status, err := client.Status(context.TODO(), s.cfg.Endpoints[0])
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Check ETCD Client Status Error:%s", s.Name(), err.Error()))
		return false
	} else {
		sctx.Logger().Info(fmt.Sprintf("[%s Starter]: ETCD Client Setup Successful!", s.Name()))
		sctx.Logger().Info(fmt.Sprintf("[%s Starter]: ETCD Client Status: %v", s.Name(), *status))
		return true
	}
}

func (s *starter) Stop() {
	_ = client.Close()
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.BasicGroup }
