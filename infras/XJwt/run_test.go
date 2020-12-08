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

		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: "", Gender: 1}

		Println("Token Service Encode:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode:")
		claim, err := XTokenUtils().Decode(token)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		exp := DefaultConfig().ExpSeconds + 1
		time.Sleep(time.Duration(exp) * time.Second)
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
		// 测试前先启动XRedis 组件的默认连接池
		logger, err := zap.NewDevelopment()
		err = XRedis.CreateDefaultPool(nil, logger)
		So(err, ShouldBeNil)
		err = CreateDefaultTkuX(nil)
		So(err, ShouldBeNil)

		// 打印redis pool 状态
		Println("pool ActiveCount:", XRedis.XPool().Stats().ActiveCount, ",pool IdleCount:", XRedis.XPool().Stats().IdleCount)

		// 启动带redis缓存的token加解码工具
		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: "", Gender: 1}

		Println("Token Service Encode And Save:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode And Validate:")

		claim, err := XTokenUtils().Decode(token)
		Println("Validate Error:", err)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		exp := DefaultConfig().ExpSeconds + 1
		time.Sleep(time.Duration(exp) * time.Second)

		claim, err = XTokenUtils().Decode(token)
		So(err, ShouldNotBeNil)
		Println("Exp Token Claim:", claim)

	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XJWT Starter", t, func() {
		// 创建启动器上下文所需要的viper
		newViper := viper.New()
		newViper.Set("Jwt.PrivateKey", DefaultConfig().PrivateKey)
		newViper.Set("Jwt.ExpSeconds", DefaultConfig().ExpSeconds)
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := infras.CreateDefaultStarterContext(newViper, logger)
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

		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: "", Gender: 1}

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
