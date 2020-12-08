package XValidate

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"goinfras"
	"testing"
)

type UserDemo struct {
	Name       string `validate:"required,alphanum"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,alphanumunicode"`
	RePassword string `validate:"required,alphanumunicode,eqfield=Password"`
}

func TestValidate(t *testing.T) {
	Convey("Test Validate DTO Struct", t, func() {
		err := CreateDefaultValidater(nil)
		So(err, ShouldBeNil)

		userDemo1 := UserDemo{
			Name:       "abc",
			Email:      "123456@qq.com",
			Password:   "123456",
			RePassword: "123456",
		}

		err = V(userDemo1)
		So(err, ShouldBeNil)

		userDemo2 := UserDemo{
			Name:       "abc",
			Email:      "123456",
			Password:   "123456fff",
			RePassword: "123456ddd",
		}

		err = V(userDemo2)
		So(err, ShouldNotBeNil)
		Println("Validate Error:", err)
	})

}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		err := CreateDefaultValidater(nil)
		So(err, ShouldBeNil)

		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := CreateDefaultStarterContext(nil, logger)
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
