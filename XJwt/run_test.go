package XJwt

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XCache"
	"github.com/bb-orz/goinfras/XCache/XGocache"
	"github.com/bb-orz/goinfras/XCache/XRedis"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
	"time"
)

// 无缓存加解码
func TestNewTokenUtils(t *testing.T) {
	Convey("Test JWT Token Utils", t, func() {
		var err error
		CreateDefaultTku(nil)

		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: ""}

		Println("Token Service Encode:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Printf("Token String:%+v \n", token)

		Println("Token Service Decode:")
		claim, err := XTokenUtils().Decode(token)
		So(err, ShouldBeNil)
		Printf("Token Claim:%+v \n", claim)

		exp := DefaultConfig().ExpSeconds + 1
		time.Sleep(time.Duration(exp) * time.Second)
		Println("Token Decode ExpTime:")
		claim, err = XTokenUtils().Decode(token)
		So(err, ShouldNotBeNil)

		Println("Token Service Decode Expired Error:", err)

	})
}

// Redis 缓存加解码
func TestTokenUtilsRedisCache(t *testing.T) {
	Convey("Test JWT Token Utils Cache", t, func() {
		var err error
		// 测试前先启动XRedis 组件的默认连接池
		err = XRedis.CreateDefaultPool(nil)
		So(err, ShouldBeNil)
		XCache.SettingCommonCache(XRedis.NewCommonRedisCache())

		CreateDefaultTkuWithCache(nil)

		// 打印redis pool 状态
		Println("pool ActiveCount:", XRedis.XPool().Stats().ActiveCount, ",pool IdleCount:", XRedis.XPool().Stats().IdleCount)

		// 启动带redis缓存的token加解码工具
		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: ""}

		Println("Token Service Encode And Save:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode And Validate:")

		claim, err := XTokenUtils().Decode(token)
		Println("Validate Error:", err)
		So(err, ShouldBeNil)
		Printf("Token Claim:%+v \n", claim)

		exp := DefaultConfig().ExpSeconds + 1
		time.Sleep(time.Duration(exp) * time.Second)

		claim, err = XTokenUtils().Decode(token)
		So(err, ShouldNotBeNil)
		Printf("Exp Token Claim:%+v \n", claim)

	})
}

// go-cache 缓存加解码
func TestTokenUtilsGoCache(t *testing.T) {
	Convey("Test JWT Token Utils Cache", t, func() {
		var err error
		// 测试前先启动XRedis 组件的默认连接池
		XGocache.CreateDefaultCache(nil)
		XCache.SettingCommonCache(XGocache.NewCommonGocache())

		CreateDefaultTkuWithCache(nil)

		// 启动带redis缓存的token加解码工具
		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: ""}

		Println("Token Service Encode And Save:")
		token, err := XTokenUtils().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode And Validate:")

		claim, err := XTokenUtils().Decode(token)
		Println("Validate Error:", err)
		So(err, ShouldBeNil)
		Printf("Token Claim:%+v \n", claim)

		exp := DefaultConfig().ExpSeconds + 1
		time.Sleep(time.Duration(exp) * time.Second)

		claim, err = XTokenUtils().Decode(token)
		So(err, ShouldNotBeNil)
		Printf("Exp Token Claim:%+v \n", claim)

	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XJWT Starter", t, func() {
		// 创建启动器上下文所需要的viper
		newViper := viper.New()
		newViper.Set("Jwt.PrivateKey", DefaultConfig().PrivateKey)
		newViper.Set("Jwt.ExpSeconds", DefaultConfig().ExpSeconds)

		logger := goinfras.NewCommandLineStarterLogger("debug")
		sctx := goinfras.CreateDefaultStarterContext(newViper, logger)
		sctx.SetConfigs(newViper)

		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
		s.Start(sctx)

		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker", Avatar: ""}

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
