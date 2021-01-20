package XGocache

import (
	"github.com/bb-orz/goinfras"
	"github.com/bb-orz/goinfras/XCache"
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

		logger := goinfras.NewCommandLineStarterLogger("debug")
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)

		time.Sleep(time.Second * 5)
		s.Stop()
	})
}

func TestXCommonWithGoCache(t *testing.T) {
	Convey("TestStarter", t, func() {
		var err error
		CreateDefaultCache(nil)
		XCache.SettingCommonCache(NewCommonGocache())

		Println("Add key a as val aa")
		err = XCache.XCommon().Add("a", "aa")
		So(err, ShouldBeNil)

		if v, b := XCache.XCommon().Get("a"); b {
			Println("Get key a:", v)
		}

		Println("Set key a as val b")
		err = XCache.XCommon().Set("a", "b")
		So(err, ShouldBeNil)
		if v, b := XCache.XCommon().Get("a"); b {
			Println("Get key a:", v)
		}

		Println("Replace key a as val c")
		err = XCache.XCommon().Replace("a", "c")
		So(err, ShouldBeNil)

		if v, b := XCache.XCommon().Get("a"); b {
			Println("Get key a:", v)
		}

		Println("AddWithExp key b as val bb,exp:5s")
		err = XCache.XCommon().AddWithExp("b", "bb", 5)
		So(err, ShouldBeNil)

		var t time.Time
		var v interface{}
		var b bool
		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with exp: [value]:%v,[exp]:%v \n", v, t)
		}

		dur := t.Unix() - time.Now().Unix()
		time.Sleep(time.Duration(dur+1) * time.Second)
		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with exp on timeout: [value]:%v,[exp]:%v \n", v, t)
		} else {
			Println("After Sleep,key b Timeout")
		}

		Println("SetWithExp key b as val bbb,exp:5s")
		err = XCache.XCommon().SetWithExp("b", "bbb", 5)
		So(err, ShouldBeNil)

		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with: [value]:%v,[exp]:%v \n", v, t)
		} else {
			Println("key b Timeout")
		}

		Println("ReplaceWithExp key b as val ccc,exp:3s")
		err = XCache.XCommon().ReplaceWithExp("b", "ccc", 3)
		So(err, ShouldBeNil)

		if v, t, b = XCache.XCommon().GetWithExp("b"); b {
			Printf("Get b with exp: [value]:%v,[exp]:%v \n", v, t)
		} else {
			Println("key b Timeout")
		}

		Println("Set int key i as 5")
		err = XCache.XCommon().Set("i", 5)
		So(err, ShouldBeNil)

		Println("Increment key i + 3")
		err = XCache.XCommon().Increment("i", 3)
		So(err, ShouldBeNil)

		Println("Decrement key i -2")
		err = XCache.XCommon().Decrement("i", 2)
		So(err, ShouldBeNil)

		if i, b := XCache.XCommon().Get("i"); b {
			Println("Print i:", i)
		}

		Println("delete key i")
		XCache.XCommon().Delete("i")
		if i, b := XCache.XCommon().Get("i"); b {
			Println("i:", i)
		} else {
			Println("i was deleted")

		}
	})
}
