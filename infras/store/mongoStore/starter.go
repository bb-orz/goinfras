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
	cfg *MongoConfig
}

func (s *MongoDBStarter) Init(sctx *infras.StarterContext) {
	configs := sctx.Configs()
	define := MongoConfig{}
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

func RunForTesting(config *MongoConfig) error {
	var err error
	if config == nil {
		config = &MongoConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(&config)
		if err != nil {
			return err
		}
		config.DbHosts = []string{"127.0.0.1:27017"}
	}

	mClient, err = NewMongoClient(config)
	return err
}
