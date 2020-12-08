package XEtcd

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"goinfras"
	"testing"
	"time"
)

func TestEtcdClientV3(t *testing.T) {
	Convey("ETCD Client Test", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		Println("Put Key...")
		putResponse, err := XClient().Put(context.Background(), "demo.a", "some value", func(op *clientv3.Op) {
			// some operation ...
		})
		So(err, ShouldBeNil)
		Println(*putResponse)

		Println("Get Key...")
		response, err := XClient().Get(context.Background(), "demo.a")
		So(err, ShouldBeNil)
		values := response.Kvs
		Println(values)

	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XEtcd Starter", t, func() {

		s := NewStarter()
		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		Println("Starter Init Successful!")
		s.Setup(sctx)
		Println("Starter Setup Successful!")
		s.Start(sctx)
		Println("Starter Start Successful!")
		if s.Check(sctx) {
			Println("Component Check Successful!")
		} else {
			Println("Component Check Fail!")
		}

		// 尝试设置和获取键值对
		sr, err := XClient().Put(context.Background(), "mykeya", "aaaaaaa")
		So(err, ShouldBeNil)
		Println("Set Key Response:", sr)

		gr, err := XClient().Get(context.Background(), "mykeya")
		So(err, ShouldBeNil)
		Println("Get Key Response:", gr)

		time.Sleep(time.Second * 5)
		s.Stop()
		Println("Component Stopped!")

	})
}
