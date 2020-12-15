package XMongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func XClient() *mongo.Client {
	return client
}

// 资源组件闭包执行
func XFClient(f func(c *mongo.Client) error) error {
	return f(client)
}

// mongodb 通用操作实例
func XCommon(dbName string) *CommonMongoDao {
	c := new(CommonMongoDao)
	c.client = XClient()
	c.defaultDb = c.client.Database(dbName)
	return c
}
