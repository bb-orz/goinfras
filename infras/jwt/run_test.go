package jwt

import (
	"GoWebScaffold/infras/store/redisStore"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"

	"go.uber.org/zap"
	"testing"
	"time"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	var t ITokenUtils

	if config == nil {
		config = &Config{
			PrivateKey: "ginger_key",
			ExpSeconds: 60,
		}

	}
	t = NewTokenUtils([]byte(config.PrivateKey), config.ExpSeconds)
	SetComponent(t)
	return err
}

// 无缓存加解码
func TestNewTokenUtils(t *testing.T) {
	Convey("Test JWT Token Utils", t, func() {
		var err error
		err = TestingInstantiation(nil)
		So(err, ShouldBeNil)

		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker"}

		Println("Token Service Encode:")
		token, err := JWTComponent().Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode:")
		claim, err := JWTComponent().Decode(token)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		time.Sleep(6 * time.Second)
		Println("Token Decode ExpTime:")
		claim, err = JWTComponent().Decode(token)
		So(err, ShouldNotBeNil)

		Println("Token Service Decode Expired Error:", err)

	})
}

// Redis 缓存加解码
func TestTokenUtilsX(t *testing.T) {
	Convey("Test JWT Token Utils Cache", t, func() {
		var err error
		err = TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO

		// 启动Redis存储
		config := redisStore.RedisConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Redis Config:", config)

		pool, err := redisStore.NewRedisPool(&config, zap.L())
		Println("pool ActiveCount:", pool.Stats().ActiveCount, ",pool IdleCount:", pool.Stats().IdleCount)

		// 启动带redis缓存的token加解码工具
		tks := NewTokenUtilsX([]byte("key"), 5, pool)
		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker"}

		Println("Token Service Encode And Save:")
		token, err := tks.Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode And Validate:")

		claim, err := tks.Decode(token)
		Println("Validate Error:", err)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		time.Sleep(6 * time.Second)
		Println("Token Decode ExpTime:")
		claim, err = tks.Decode(token)
		So(err, ShouldNotBeNil)
	})
}
