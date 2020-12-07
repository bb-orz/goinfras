package XMongo

import (
	"GoWebScaffold/infras"
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"testing"
)

func TestMongoClient(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		err = client.Ping(context.TODO(), nil)
		So(err, ShouldBeNil)
	})

}

func TestNewCommonMongoDao(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := CreateDefaultDB(nil)

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

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		err := CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := infras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")

		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

	})
}
