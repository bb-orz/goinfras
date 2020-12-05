package XJwt

import (
	"GoWebScaffold/infras/XStore/XRedis"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

// 无缓存加解码
func TestNewTokenUtils(t *testing.T) {
	Convey("Test JWT Token Utils", t, func() {
		var err error
		err = TestingInstantiation(nil)
		So(err, ShouldBeNil)

		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker"}

		Println("Token Service Encode:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode:")
		claim, err := XTokenUtils().Decode(token)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		time.Sleep(6 * time.Second)
		Println("Token Decode ExpTime:")
		claim, err = XTokenUtils().Decode(token)
		So(err, ShouldNotBeNil)

		Println("Token Service Decode Expired Error:", err)

	})
}

// Redis 缓存加解码
func TestTokenUtilsX(t *testing.T) {
	Convey("Test JWT Token Utils Cache", t, func() {
		var err error
		err = TestingInstantiationForRedisCache(nil)
		So(err, ShouldBeNil)

		// 打印redis pool 状态
		Println("pool ActiveCount:", XRedis.XPool().Stats().ActiveCount, ",pool IdleCount:", XRedis.XPool().Stats().IdleCount)

		// 启动带redis缓存的token加解码工具
		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker"}

		Println("Token Service Encode And Save:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode And Validate:")

		claim, err := XTokenUtils().Decode(token)
		Println("Validate Error:", err)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		time.Sleep(6 * time.Second)
		Println("Token Decode ExpTime:")
		claim, err = XTokenUtils().Decode(token)
		So(err, ShouldNotBeNil)
	})
}
