package etcd

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"
	"testing"
)

func TestEtcdClientV3(t *testing.T) {
	Convey("ETCD Client Test", t, func() {
		config := EtcdConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Cron Config:", config)

		client, err := NewEtcdClient(context.TODO(), &config, nil)
		So(err, ShouldBeNil)

		response, err := client.Get(context.TODO(), "demo.a")
		So(err, ShouldBeNil)

		values := response.Kvs
		Println(values)
	})

}
