package XEtcd

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"go.etcd.io/etcd/clientv3"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	var c *clientv3.Client
	if config == nil {
		config = &Config{
			Endpoints: []string{"localhost:2379"},
		}
	}
	c, err = NewEtcdClient(context.TODO(), config, nil)
	SetComponent(c)
	return err
}

func TestEtcdClientV3(t *testing.T) {
	Convey("ETCD Client Test", t, func() {
		var err error
		config := Config{}
		err = TestingInstantiation(&config)
		So(err, ShouldBeNil)

		Println("Cron Config:", config)

		response, err := EtcdComponent().Get(context.TODO(), "demo.a")
		So(err, ShouldBeNil)

		values := response.Kvs
		Println(values)
	})

}
