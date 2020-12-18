package XNats

import (
	"fmt"
	"github.com/bb-orz/goinfras"
	"github.com/nats-io/nats.go"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

const (
	TestChanSubjectName    = "testChan"
	TestQueueSubjectName   = "testQueueSubject"
	TestQueueName          = "testQueue"
	TestReqRespSubjectName = "testReqRespSubject"
)

// 发送消息秒级消息方法for testing
func sendTickerMsg(subjectName string) {
	// 发送
	var err error
	var sendCh chan interface{}
	sendCh = make(chan interface{}, 1)
	Println("Sending...")
	go func() {
		for {
			msg := <-time.NewTicker(time.Second).C
			sendCh <- msg
			fmt.Println("----------------------------------")
			fmt.Println("Send===>", msg)
			fmt.Println("----------------------------------")
		}
	}()

	err = XCommonNatsChan().BindSendChan(subjectName, sendCh)
	So(err, ShouldBeNil)

}

// 通道发送接收模式测试
func TestNatsMQChan(t *testing.T) {
	Convey("TestNatsMQChan", t, func() {
		var err error

		err = CreateDefaultPool(nil)
		So(err, ShouldBeNil)

		// 接收
		var recevCh chan interface{}
		recevCh = make(chan interface{}, 1)
		Println("Receiving...")
		go func() {
			for {
				msg := <-recevCh
				fmt.Println("Receive:", msg)
			}
		}()
		err = XCommonNatsChan().BindRecvChan(TestChanSubjectName, recevCh)
		So(err, ShouldBeNil)

		// 发送消息
		sendTickerMsg(TestChanSubjectName)

		time.Sleep(time.Second * 10)

	})
}

type person struct {
	Name string `json:"name,omitempty"`
	Age  uint   `json:"age,omitempty"`
}

// 队列组订阅器测试
func TestNatsMQQueueSubscribe(t *testing.T) {
	Convey("TestNatsMQQueueSubscribe", t, func() {
		var err error
		err = CreateDefaultPool(nil)
		So(err, ShouldBeNil)

		// 4个go程订阅特定消息person的队列组
		for i := 0; i < 4; i++ {
			go func(x int) {
				err = XCommonNatsQueue().QueueSubscribe(TestQueueSubjectName, TestQueueName, func(subject, reply string, p *person) {
					fmt.Println("==============================")
					fmt.Println("Queue ：", TestQueueName)
					fmt.Println("Worker ：", x)
					fmt.Println("subject:", subject)
					fmt.Println("reply:", reply)
					fmt.Println("person:", p)
					fmt.Println("==============================")
				})

			}(i)
		}
		So(err, ShouldBeNil)

		// 发送消息
		var sendCh chan interface{}
		sendCh = make(chan interface{}, 1)
		Println("Sending...")
		go func() {
			for {
				<-time.NewTicker(time.Second).C
				p := &person{
					Name: "person1",
					Age:  1,
				}
				sendCh <- p
				fmt.Println("----------------------------------")
				fmt.Println("Send===>", *p)
				fmt.Println("----------------------------------")
			}
		}()

		err = XCommonNatsChan().BindSendChan(TestQueueSubjectName, sendCh)
		So(err, ShouldBeNil)

		time.Sleep(time.Second * 10)

	})
}

// 通道队列组接收测试
func TestNatsMQQueueChanRecv(t *testing.T) {
	Convey("TestNatsMQQueueChanRecv", t, func() {
		var err error
		err = CreateDefaultPool(nil)
		So(err, ShouldBeNil)

		// 4个go程接收统一队列组消息
		for i := 0; i < 4; i++ {
			go func(x int) {
				// 接收队列消息
				var recevCh chan interface{}
				recevCh = make(chan interface{}, 1)
				for {
					err = XCommonNatsQueue().BindRecvQueueChan(TestQueueSubjectName, TestQueueName, recevCh)
					msg := <-recevCh
					fmt.Println("==============================")
					fmt.Println("QueueName:", TestQueueName)
					fmt.Println("Worker:", x)
					fmt.Println("Message:", msg)
					fmt.Println("==============================")
				}
			}(i)
		}
		So(err, ShouldBeNil)

		// 发送消息
		sendTickerMsg(TestQueueSubjectName)
		time.Sleep(time.Second * 10)

	})
}

// request/reply模式测试
func TestNatsMQReqSub(t *testing.T) {
	t.Run("mmm", func(t *testing.T) {
		testNatsMQSubscribeReply(t)
		time.Sleep(time.Second)
		testNatsMQRequest(t)
	})
}

func testNatsMQRequest(t *testing.T) {
	Convey("TestNatsMQRequest", t, func() {
		var err error
		err = CreateDefaultPool(nil)
		So(err, ShouldBeNil)

		// Request
		msg := "help me" // 生成时间消息
		var reply string // 接收订阅者响应的消息
		var exp = time.Second * 10
		// 请求时阻塞等待回执
		err = XCommonNatsRequest().Request(TestReqRespSubjectName, msg, &reply, exp)
		fmt.Println("Request----------------------------------")
		fmt.Println("Subject:", TestReqRespSubjectName)
		fmt.Println("Msg:", msg)
		fmt.Println("Received Reply Message From Subscriber:", reply)
		fmt.Println("----------------------------------")

	})

}

func testNatsMQSubscribeReply(t *testing.T) {
	Convey("TestNatsMQSubscribeReply", t, func() {
		var err error
		var subscriber *nats.Subscription
		err = CreateDefaultPool(nil)
		So(err, ShouldBeNil)

		// 订阅接收到消息后发送回执
		subscriber, err = XCommonNatsPubSub().Subscribe(TestReqRespSubjectName, func(msg *nats.Msg) {
			fmt.Println("Subscriber Receive Message:", string(msg.Data))
			err = XCommonNatsPubSub().Publish(msg.Reply, fmt.Sprintf("I Receive Message :%s", string(msg.Data)))
		})
		So(err, ShouldBeNil)

		time.Sleep(time.Second * 10)
		err = subscriber.Unsubscribe()
		So(err, ShouldBeNil)

	})
}

func TestStarter(t *testing.T) {
	Convey("Test XNats Starter", t, func() {
		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)

		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)

		s.Stop()
		time.Sleep(time.Second * 3)
		Println("Starter Stop Successful!")

	})
}
