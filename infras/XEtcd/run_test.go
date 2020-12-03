package XEtcd

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			Endpoints: []string{"localhost:2379"},
		}
	}
	client, err = NewEtcdClient(context.TODO(), config, nil)
	return err
}

func TestEtcdClientV3(t *testing.T) {
	Convey("ETCD Client Test", t, func() {
		var err error
		config := Config{}
		err = TestingInstantiation(&config)
		So(err, ShouldBeNil)

		Println("Cron Config:", config)

		response, err := client.Get(context.TODO(), "demo.a")
		So(err, ShouldBeNil)

		values := response.Kvs
		Println(values)
	})

}
