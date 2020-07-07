package user

import (
	"GoWebScaffold/core"
	"GoWebScaffold/infras/store/ormStore"
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	Convey("User Service Create User Testing:", t, func() {
		var err error
		err = validate.RunForTesting(nil)
		So(err, ShouldBeNil)
		err = ormStore.RunForTesting(nil)
		So(err, ShouldBeNil)

		dto := services.CreateUserWithEmailDTO{
			Name:       "fun",
			Email:      "123456@qq.com",
			Password:   "123456",
			RePassword: "123456",
		}
		service := new(core.UserService)
		userDTO, err := service.CreateUserWithEmail(dto)
		So(err, ShouldBeNil)

		Println("New User:", userDTO)
	})
}

func TestUserService_GetUserInfo(t *testing.T) {
	Convey("User Service Get User Info Testing:", t, func() {
		var err error
		err = validate.RunForTesting(nil)
		So(err, ShouldBeNil)
		err = ormStore.RunForTesting(nil)
		So(err, ShouldBeNil)

		service := new(core.UserService)
		userDTO, err := service.GetUserInfo(12)
		So(err, ShouldBeNil)
		Println("Get User Info:", userDTO)
	})
}
