package user

import (
	"GoWebScaffold/infras/store/ormStore"
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
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
		var err error
		err = validate.RunForTesting(nil)
		So(err, ShouldBeNil)
		err = ormStore.RunForTesting(nil)
		So(err, ShouldBeNil)

		dto := services.CreateUserDTO{
			Username:   "fun",
			Email:      "123456@qq.com",
			Password:   "123456",
			RePassword: "123456",
		}
		service := new(userService)
		userDTO, err := service.CreateUser(dto)
		So(err, ShouldBeNil)
		Println(err)

		Println("New User:", userDTO)
	})
}
