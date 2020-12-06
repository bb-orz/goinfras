package XJwt

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/XStore/XRedis"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"testing"
	"time"
)

// 无缓存加解码
func TestNewTokenUtils(t *testing.T) {
	Convey("Test JWT Token Utils", t, func() {
		var err error
		CreateDefaultTkuX(nil)

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
		CreateDefaultTku(nil)

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

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XJWT Starter", t, func() {
		// 创建启动器上下文所需要的viper
		newViper := viper.New()
		newViper.Set("Jwt.PrivateKey", DefaultConfig().PrivateKey)
		newViper.Set("Jwt.ExpSeconds", DefaultConfig().ExpSeconds)
		sctx := infras.CreateDefaultStarterContext(newViper, zap.L())
		sctx.SetConfigs(newViper)

		s := NewStarter()
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")
		s.Start(sctx)
		Println("Starter Start Successful!")
		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker"}

		Println("Token Service Encode:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode:")
		claim, err := XTokenUtils().Decode(token)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		// 测试超时1s
		exp := sctx.Configs().GetInt64("Jwt.ExpSeconds")
		time.Sleep(time.Duration(exp+1) * time.Second)
		Println("Token Decode ExpTime:")
		claim, err = XTokenUtils().Decode(token)
		So(err, ShouldNotBeNil)
		Println("Token Service Decode Expired Error:", err)
	})
}
