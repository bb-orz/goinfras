package mongoStore

import (
	"GoWebScaffold/infras"
	"context"
	"github.com/tietang/props/kvs"
	"go.mongodb.org/mongo-driver/mongo"
)

var mClient *mongo.Client

func MongoClient() *mongo.Client {
	infras.Check(mClient)
	return mClient
}

type MongoDBStarter struct {
	infras.BaseStarter
	cfg *mongoConfig
}

func (s *MongoDBStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := mongoConfig{}
	err := kvs.Unmarshal(configs, &define, "Mongodb")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *MongoDBStarter) Setup(sctx *infras.StarterContext) {
	var err error
	mClient, err = NewMongoClient(s.cfg)
	infras.FailHandler(err)
	sctx.Logger().Info("MongoClient Setup Successful!")
}

func (s *MongoDBStarter) Stop(sctx *infras.StarterContext) {
	_ = MongoClient().Disconnect(context.TODO())
}
