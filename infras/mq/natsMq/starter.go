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

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("NatsMq", &define)
	infras.FailHandler(err)

	s.cfg = &define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	natsMQPool, err = NewNatsMqPool(s.cfg, sctx.Logger())
	infras.FailHandler(err)
	sctx.Logger().Info("NatsMQPool Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	NatsMQPool().Close()
}

/*For testing*/
func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
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
