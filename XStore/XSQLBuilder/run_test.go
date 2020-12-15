package XSQLBuilder

import (
	"context"
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
)

// 测试使用mysql client
func TestMysqlDB(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		err = db.Ping()
		So(err, ShouldBeNil)

	})

}

// 测试通用的mysql存储
func TestNewCommonMysqlStore(t *testing.T) {
	Convey("测试使用Common Mysql Store", t, func() {
		err := CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		lastedId, err := XCommon().Insert("user", []map[string]interface{}{
			{"name": "aaaa", "age": 18, "gender": 1}, {"name": "bbbb", "age": 20, "gender": 0},
		})
		So(err, ShouldBeNil)
		Println("Lasted Insert Id:", lastedId)

		count, err := XCommon().GetCount("user", nil)
		So(err, ShouldBeNil)
		Println("User Count:", count)

		rs := UserSchema{}
		err = XCommon().GetOne("user", map[string]interface{}{"name": "aaaa"}, nil, &rs)
		So(err, ShouldBeNil)
		Println("GetOne:", rs)

		rsList := make([]interface{}, 0)
		XCommon().GetMulti("user", map[string]interface{}{"name": "aaaa"}, nil, rsList)
		So(err, ShouldBeNil)
		Println("GetMulti:", rsList)

		update, err := XCommon().Update("user", map[string]interface{}{"age": 18}, map[string]interface{}{"age": 28})
		So(err, ShouldBeNil)
		Println("Update Lasted Id:", update)

		deleteId, err := XCommon().Delete("user", map[string]interface{}{"name": "ken"})
		So(err, ShouldBeNil)
		Println("Delete Id:", deleteId)
	})
}

type UserSchema struct {
	Id     int    `ddb:"id"`
	Name   string `ddb:"name"`
	Age    int    `ddb:"age"`
	Gender int    `ddb:"gender"`
}

func (u UserSchema) TableName() string {
	return "user"
}

// 测试事务的mysql存储
func TestBaseDaoTx(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		tx, err := XCommon().NewTx(context.Background(), nil)
		lastedId, err := tx.Insert("user", []map[string]interface{}{
			{"name": "fff", "age": 18, "gender": 1}, {"name": "kkk", "age": 20, "gender": 0},
		})
		So(err, ShouldBeNil)
		Println("Lasted Insert Id:", lastedId)

		rs := UserSchema{}
		err = tx.GetOne("user", map[string]interface{}{"name": "fff"}, nil, &rs)
		So(err, ShouldBeNil)
		Println("GetOne:", rs)

		rsList := make([]interface{}, 0)
		tx.GetMulti("user", map[string]interface{}{"name": "fff"}, nil, rsList)
		So(err, ShouldBeNil)
		Println("GetMulti:", rsList)

		update, err := tx.Update("user", map[string]interface{}{"age": 18}, map[string]interface{}{"age": 28})
		So(err, ShouldBeNil)
		Println("Update Lasted Id:", update)

		deleteId, err := tx.Delete("user", map[string]interface{}{"name": "fff"})
		So(err, ShouldBeNil)
		Println("Delete Id:", deleteId)

		tx.tx.Commit()

	})
}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		err := CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
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
