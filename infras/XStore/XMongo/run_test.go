package XMongo

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestMongoClient(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		err = client.Ping(context.TODO(), nil)
		So(err, ShouldBeNil)
	})

}

func TestNewCommonMongoDao(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := TestingInstantiation(nil)

		So(err, ShouldBeNil)

		// 通用操作：增删改查
		commonMongoDao := XCommon("dev_test")
		// 增
		insertID, err := commonMongoDao.InsertOne(context.TODO(), "demo", bson.M{"name": "joker"})
		So(err, ShouldBeNil)
		Println("InsertID", insertID)

		// 查
		result := commonMongoDao.FindOne(context.TODO(), "demo", bson.M{"name": "joker"})
		So(result.Err(), ShouldBeNil)
		res := bson.M{}
		err = result.Decode(res)
		So(err, ShouldBeNil)
		Println("FindOne Result", res)

		// 改
		opts := options.Update().SetUpsert(true)
		filter := bson.D{{"_id", insertID}}
		update := bson.D{{"$set", bson.D{{"name", "ken"}}}}
		updateResult, err := commonMongoDao.UpdateOne(context.TODO(), "demo", filter, update, opts)
		So(err, ShouldBeNil)
		Println("Update Result", updateResult)

		// 删
		deleteCount, err := commonMongoDao.DeleteOne(context.TODO(), "demo", bson.D{{"_id", insertID}})
		So(err, ShouldBeNil)
		Println("Delete Count", deleteCount)

		// ...
	})
}
