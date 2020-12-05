package XRedisPubSub

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRedisPubsubPool(t *testing.T) {
	Convey("Test Redis PubSub Component", t, func() {
		var err error
		err = TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO

	})
}
