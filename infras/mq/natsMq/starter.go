package natsMq

import (
	"GoWebScaffold/infras"
	"go.uber.org/zap"
)

var natsMQPool *NatsPool

func NatsMQPool() *NatsPool {
	infras.Check(natsMQPool)
	return natsMQPool
}

type NatsMQStarter struct {
	infras.BaseStarter
	cfg *NatsMqConfig
}

func (s *NatsMQStarter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := NatsMqConfig{}
	err := viper.UnmarshalKey("NatsMq", &define)
	infras.FailHandler(err)

	s.cfg = &define
}

func (s *NatsMQStarter) Setup(sctx *infras.StarterContext) {
	var err error
	natsMQPool, err = NewNatsMqPool(s.cfg, sctx.Logger())
	infras.FailHandler(err)
	sctx.Logger().Info("NatsMQPool Setup Successful!")
}

func (s *NatsMQStarter) Stop(sctx *infras.StarterContext) {
	NatsMQPool().Close()
}

/*For testing*/
func RunForTesting(config *NatsMqConfig) error {
	var err error
	if config == nil {
		config = &NatsMqConfig{
			Switch: true,
			NatsServers: []natsServer{
				{
					"127.0.0.1",
					4222,
					false,
					"",
					"",
				},
			},
		}

	}

	natsMQPool, err = NewNatsMqPool(config, zap.L())
	return err
}
