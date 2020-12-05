package XOAuth

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOAuthComponent(t *testing.T) {
	Convey("OAuthManager Testing:", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO
		// XManager()....
	})
}
