package XGocache

import (
	"github.com/bb-orz/goinfras"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGocache(t *testing.T) {
	Convey("TestStarter", t, func() {
		var err error
		CreateDefaultCache(nil)
		So(err, ShouldBeNil)

		X().SetDefault("aaa", "aaaValue")
		X().SetDefault("bbb", "bbbValue")
		X().SetDefault("ccc", "cccValue")
		X().SetDefault("ddd", "dddValue")

		akey, aexp, afound := X().GetWithExpiration("aaa")
		Printf("Key:%s,Exp:%v,Found:%t \n", akey, aexp, afound)
		bkey, bexp, bfound := X().GetWithExpiration("bbb")
		Printf("Key:%s,Exp:%v,Found:%t \n", bkey, bexp, bfound)
		ckey, cexp, cfound := X().GetWithExpiration("ccc")
		Printf("Key:%s,Exp:%v,Found:%t \n", ckey, cexp, cfound)

		err = DumpItems(DefaultConfig())
		So(err, ShouldBeNil)

		X().Flush()
		Println("After Flush ... ")
		aakey, aaexp, aafound := X().GetWithExpiration("aaa")
		Printf("Key:%s,Exp:%v,Found:%t \n", aakey, aaexp, aafound)
		bbkey, bbexp, bbfound := X().GetWithExpiration("bbb")
		Printf("Key:%s,Exp:%v,Found:%t \n", bbkey, bbexp, bbfound)
		cckey, ccexp, ccfound := X().GetWithExpiration("ccc")
		Printf("Key:%s,Exp:%v,Found:%t \n", cckey, ccexp, ccfound)

		// 重新载入cache实例
		err = CreateDefaultCacheFrom(DefaultConfig())
		So(err, ShouldBeNil)
		Println("After NewCacheForm ... ")
		aaakey, aaaexp, aaafound := X().GetWithExpiration("aaa")
		Printf("Key:%s,Exp:%v,Found:%t \n", aaakey, aaaexp, aaafound)
		bbbkey, bbbexp, bbbfound := X().GetWithExpiration("bbb")
		Printf("Key:%s,Exp:%v,Found:%t \n", bbbkey, bbbexp, bbbfound)
		ccckey, cccexp, cccfound := X().GetWithExpiration("ccc")
		Printf("Key:%s,Exp:%v,Found:%t \n", ccckey, cccexp, cccfound)
	})
}

func TestStarter(t *testing.T) {
	Convey("TestStarter", t, func() {
		var err error
		CreateDefaultCache(nil)
		So(err, ShouldBeNil)

		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)

		time.Sleep(time.Second * 5)
		s.Stop()
	})
}
