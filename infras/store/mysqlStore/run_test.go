package mysqlStore

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"
	"testing"
)

// 测试使用mysql client
func TestMysqlClient(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := RunForTesting(nil)
		So(err, ShouldBeNil)

		err = MysqlClient().Ping()
		So(err, ShouldBeNil)

	})

}

// 测试通用的mysql存储
func TestNewCommonMysqlStore(t *testing.T) {
	Convey("测试使用Common Mysql Store", t, func() {
		err := RunForTesting(nil)
		So(err, ShouldBeNil)

		commonStore := NewCommonMysqlStore()
		lastedId, err := commonStore.Insert("user", []map[string]interface{}{
			{"name": "aaaa", "age": 18, "gender": 1}, {"name": "bbbb", "age": 20, "gender": 0},
		})
		So(err, ShouldBeNil)
		Println("Lasted Insert Id:", lastedId)

		count, err := commonStore.GetCount("user", nil)
		So(err, ShouldBeNil)
		Println("User Count:", count)

		rs := UserSchema{}
		err = commonStore.GetOne("user", map[string]interface{}{"name": "joker"}, nil, &rs)
		So(err, ShouldBeNil)
		Println("GetOne:", rs)

		rsList := make([]UserSchema, 0)
		commonStore.GetMulti("user", map[string]interface{}{"name": "aaaa"}, nil, &rsList)
		So(err, ShouldBeNil)
		Println("GetMulti:", rsList)

		update, err := commonStore.Update("user", map[string]interface{}{"age": 18}, map[string]interface{}{"age": 28})
		So(err, ShouldBeNil)
		Println("Update Lasted Id:", update)

		deleteId, err := commonStore.Delete("user", map[string]interface{}{"name": "ken"})
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
		err := RunForTesting(nil)
		So(err, ShouldBeNil)

		tx, err := NewCommonMysqlStore().NewTx(context.Background(), nil)
		lastedId, err := tx.Insert("user", []map[string]interface{}{
			{"name": "fff", "age": 18, "gender": 1}, {"name": "kkk", "age": 20, "gender": 0},
		})
		So(err, ShouldBeNil)
		Println("Lasted Insert Id:", lastedId)

		rs := UserSchema{}
		err = tx.GetOne("user", map[string]interface{}{"name": "fff"}, nil, &rs)
		So(err, ShouldBeNil)
		Println("GetOne:", rs)

		rsList := make([]UserSchema, 0)
		tx.GetMulti("user", map[string]interface{}{"name": "fff"}, nil, &rsList)
		So(err, ShouldBeNil)
		Println("GetMulti:", rsList)

		update, err := tx.Update("user", map[string]interface{}{"age": 18}, map[string]interface{}{"age": 28})
		So(err, ShouldBeNil)
		Println("Update Lasted Id:", update)

		deleteId, err := tx.Delete("user", map[string]interface{}{"name": "fff"})
		So(err, ShouldBeNil)
		Println("Delete Id:", deleteId)

		// TODO 解决没有commit也保存的问题
		//tx.tx.Commit()

	})
}
