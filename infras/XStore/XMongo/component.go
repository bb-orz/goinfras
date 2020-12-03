package XMongo

import (
	"GoWebScaffold/infras"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func MongoComponent() *mongo.Client {
	infras.Check(client)
	return client
}

func SetComponent(c *mongo.Client) {
	client = c
}
