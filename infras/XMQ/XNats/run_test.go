package XNats

import (
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			Switch: true,
			NatsServers: []natsServer{
				{
					"127.0.0.1",
					4222,
					false,
					"",
					"",
				},
			},
		}

	}

	natsMQPool, err = NewPool(config, zap.L())
	return err
}

func TestNatsMQComponent(t *testing.T) {
	Convey("Test NatsMQ Component", t, func() {
		var err error
		err = TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// TODO

	})
}
