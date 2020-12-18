package XEtcd

import (
	"context"
	"fmt"
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

const (
	KeyTest   = "/test"
	KeyLevel1 = "/test/level1"
	KeyLevel2 = "/test/level1/lever2"
)

// GET SET 设置
func TestEtcdGetSet(t *testing.T) {
	Convey("ETCD Client Test", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		Println("Put Key...")
		withTimeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*3)
		putResponse1, err := XClient().Put(withTimeoutCtx, KeyTest, "test value", func(op *clientv3.Op) {
			isPut := op.IsPut()
			So(isPut, ShouldBeTrue)
		})
		So(err, ShouldBeNil)
		Println(*putResponse1)

		putResponse2, err := XClient().Put(withTimeoutCtx, KeyLevel1, "level1 value", func(op *clientv3.Op) {
			isPut := op.IsPut()
			So(isPut, ShouldBeTrue)
		})
		So(err, ShouldBeNil)
		Println(*putResponse2)

		putResponse3, err := XClient().Put(withTimeoutCtx, KeyLevel2, "level2 value", func(op *clientv3.Op) {
			isPut := op.IsPut()
			So(isPut, ShouldBeTrue)
		})
		So(err, ShouldBeNil)
		Println(*putResponse3)

		Println("Get KeyLevel2 Value...")
		response1, err := XClient().Get(context.Background(), KeyLevel2)
		So(err, ShouldBeNil)
		values1 := response1.Kvs
		Println(values1)

		Println("Get KeyTest All Values...")
		response2, err := XClient().Get(context.Background(), KeyTest, clientv3.WithPrefix())
		So(err, ShouldBeNil)
		values2 := response2.Kvs
		Println(values2)

	})
}

// 监听Key测试
func TestEtcdWatch(t *testing.T) {
	Convey("TestEtcdWatch", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		Println("Watch Key:", KeyLevel1)
		go func() {
			watchChans := XClient().Watch(context.Background(), KeyLevel1)
			fmt.Println("Watch Key:", KeyLevel1)
			for wv := range watchChans {
				for _, w := range wv.Events {
					fmt.Println("Key:", string(w.Kv.Key))
					fmt.Println("Value:", string(w.Kv.Value))
					fmt.Println("Version:", w.Kv.Version)
					fmt.Println("ModRevision:", w.Kv.ModRevision)
					fmt.Println("CreateRevision:", w.Kv.CreateRevision)
				}
			}
		}()

		time.Sleep(time.Second)

		Println("Put KeyLevel1 ...")
		withTimeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*3)
		_, err = XClient().Put(withTimeoutCtx, KeyLevel1, "update KeyLevel1", func(op *clientv3.Op) {
			isPut := op.IsPut()
			So(isPut, ShouldBeTrue)
		})
		So(err, ShouldBeNil)

		time.Sleep(time.Second * 3)

	})
}

// 租约Lease测试
func TestEtcdLease(t *testing.T) {
	Convey("TestEtcdWatch", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		// 申请一个5秒的租约
		lease, err := XClient().Grant(context.Background(), 5)
		So(err, ShouldBeNil)

		// 使用租约放置配置键值
		Println("Put KeyLevel1 ...")
		withTimeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*3)
		_, err = XClient().Put(withTimeoutCtx, KeyLevel1, "KeyLevel1ValueWithLease", clientv3.WithLease(lease.ID))
		So(err, ShouldBeNil)

		// 租约内获取键值
		Println("Get KeyLevel1 Value In Lease...")
		response1, err := XClient().Get(context.Background(), KeyLevel1)
		So(err, ShouldBeNil)
		value1 := response1.Kvs
		Println(value1)

		// 租约超时后获取键值
		time.Sleep(time.Second * 6)
		Println("Get KeyLevel1 Value TimeOut Lease...")
		response2, err := XClient().Get(context.Background(), KeyLevel1)
		So(err, ShouldBeNil)
		value2 := response2.Kvs
		Println(value2)

	})
}

// 租约延续测试
func TestEtcdKeepAlive(t *testing.T) {
	Convey("TestEtcdWatch", t, func() {
		err := CreateDefaultClient(nil)
		So(err, ShouldBeNil)

		// 申请一个5秒的租约
		lease, err := XClient().Grant(context.Background(), 5)
		So(err, ShouldBeNil)

		// 使用租约放置配置键值
		Println("Put KeyLevel1 ...")
		withTimeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*3)
		_, err = XClient().Put(withTimeoutCtx, KeyLevel1, "KeyLevel1ValueWithLease", clientv3.WithLease(lease.ID))
		So(err, ShouldBeNil)

		// 租约内获取键值
		Println("Get KeyLevel1 Value In Lease...")
		response1, err := XClient().Get(context.Background(), KeyLevel1)
		So(err, ShouldBeNil)
		value1 := response1.Kvs
		Println(value1)

		// 租约延期
		// the key 'foo' will be kept forever
		ch, err := XClient().KeepAlive(context.TODO(), lease.ID)
		So(err, ShouldBeNil)
		go func() {
			for i := 0; i < 2; i++ {
				ka := <-ch
				fmt.Println("ttl:", ka.TTL)
			}
		}()

		// 原租约超时后再测试获取键值
		time.Sleep(time.Second * 6)
		Println("Get KeyLevel1 Value TimeOut Lease...")
		response2, err := XClient().Get(context.Background(), KeyLevel1)
		So(err, ShouldBeNil)
		value2 := response2.Kvs
		Println(value2)
	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XEtcd Starter", t, func() {
		s := NewStarter()
		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)

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
