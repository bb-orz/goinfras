package validate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			true,
		}
	}

	if config.TransZh {
		validate, translator, err = NewZhValidator()
	} else {
		validate = NewValidator()
	}
	return err
}

type UserDemo struct {
	Name       string `validate:"required,alphanum"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,alphanumunicode"`
	RePassword string `validate:"required,alphanumunicode,eqfield=Password"`
}

func TestValidate(t *testing.T) {
	Convey("Test Validate DTO Struct", t, func() {
		err := TestingInstantiation(nil)
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
