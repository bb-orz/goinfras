package XMongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

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

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
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
