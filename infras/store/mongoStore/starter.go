package mongoStore

import (
	"GoWebScaffold/infras"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func Client() *mongo.Client {
	infras.Check(client)
	return client
}

type Starter struct {
	infras.BaseStarter
	cfg Config
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Mongodb", &define)
	infras.FailHandler(err)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	var err error
	client, err = NewClient(&s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("MongoClient Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	_ = Client().Disconnect(context.TODO())
}

func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			[]string{"127.0.0.1:27017"},
			"",
			"",
			"",
			"",
			true,
			15,
			nil,
			true,
			10,
			100,
			1000,
			120,
			false,
			20,
			true,
			true,
		}
	}

	client, err = NewClient(config)
	return err
}
