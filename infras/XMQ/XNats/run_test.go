package XNats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNatsMQComponent(t *testing.T) {
	Convey("Test Nats Component", t, func() {
		var err error
		err = TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO

	})
}
