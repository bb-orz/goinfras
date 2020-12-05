package Xgin

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGinEngine(t *testing.T) {
	Convey("Gin Server Run Test", t, func() {
		var err error
		config := Config{}
		err = TestingInstantiation(&config, nil)
		So(err, ShouldBeNil)
	})

	// TODO
}
