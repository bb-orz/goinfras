package mongoStore

import (
	"GoWebScaffold/infras"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

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
	var c *mongo.Client
	c, err = NewClient(&s.cfg)
	infras.FailHandler(err)
	SetComponent(c)
	sctx.Logger().Info("MongoClient Setup Successful!")
}

func (s *Starter) Stop(sctx *infras.StarterContext) {
	_ = MongoComponent().Disconnect(context.TODO())
}
