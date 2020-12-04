package XSQLBuilder

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			"127.0.0.1",
			3306,
			"",
			"",
			"",
			60,
			100,
			200,
			"uft8",
			true,
			true,
			5,
			30,
			true,
			true,
		}
	}
	db, err = NewDB(config)
	return err
}

// 测试使用mysql client
func TestMysqlDB(t *testing.T) {
	Convey("测试使用mysql client", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		err = db.Ping()
		So(err, ShouldBeNil)

	})

}

// 测试通用的mysql存储
func TestNewCommonMysqlStore(t *testing.T) {
	Convey("测试使用Common Mysql Store", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		commonStore := XCommon()
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
		err := TestingInstantiation(nil)
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