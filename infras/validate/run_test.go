package validate

import (
	. "github.com/smartystreets/goconvey/convey"
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
		err := RunForTesting(nil)
		So(err, ShouldBeNil)

		userDemo1 := UserDemo{
			Name:       "abc",
			Email:      "123456@qq.com",
			Password:   "123456",
			RePassword: "123456",
		}

		err = Validate(userDemo1)
		So(err, ShouldBeNil)

		userDemo2 := UserDemo{
			Name:       "abc",
			Email:      "123456",
			Password:   "123456fff",
			RePassword: "123456ddd",
		}

		err = Validate(userDemo2)
		So(err, ShouldNotBeNil)
		Println("Validate Error:", err)
	})

}
