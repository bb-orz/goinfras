package user

import (
	"GoWebScaffold/infras/store/ormStore"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"
	"testing"
)

func getOrmDb() *gorm.DB {
	config := ormStore.OrmConfig{}
	p := kvs.NewEmptyCompositeConfigSource()
	err := p.Unmarshal(&config)
	So(err, ShouldBeNil)
	Println("ORM Config:", config)

	gormDb, err := ormStore.NewORMDb(&config)
	So(err, ShouldBeNil)
	return gormDb
}

func TestUserService_CreateUser(t *testing.T) {
	Convey("User Service Create User Testing:", t, func() {
		err := ormStore.RunForTesting()
		So(err, ShouldBeNil)
	})
}
