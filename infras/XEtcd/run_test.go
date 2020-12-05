package XEtcd

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEtcdClientV3(t *testing.T) {
	Convey("ETCD Client Test", t, func() {
		var err error
		config := Config{}
		err = TestingInstantiation(&config)
		So(err, ShouldBeNil)

		Println("Cron Config:", config)

		response, err := XClient().Get(context.TODO(), "demo.a")
		So(err, ShouldBeNil)

		values := response.Kvs
		_, err = Println(values)
		So(err, ShouldBeNil)

	})

}
