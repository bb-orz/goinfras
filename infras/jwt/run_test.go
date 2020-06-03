package jwt

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNewTokenService(t *testing.T) {
	Convey("Test JWT Token Service", t, func() {
		tks := NewTokenUtils([]byte("key"), 5)
		userClaim := UserClaim{Id: "qwertwerhadfsgsadfg", Name: "joker"}

		Println("Token Service Encode:")
		token, err := tks.Encode(userClaim)
		So(err, ShouldBeNil)
		Println("Token String", token)

		Println("Token Service Decode:")
		claim, err := tks.Decode(token)
		So(err, ShouldBeNil)
		Println("Token Claim:", claim)

		time.Sleep(6 * time.Second)
		Println("Token Decode ExpTime:")
		claim, err = tks.Decode(token)
		So(err, ShouldNotBeNil)

		Println("Token Service Decode Expired Error:", err)

	})
}
