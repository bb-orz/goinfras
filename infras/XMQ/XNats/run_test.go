package XNats

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
	"time"
)

const (
	TestChanSubjectName    = "testChan"
	TestQueueSubjectName   = "testQueueSubject"
	TestQueueName          = "testQueue"
	TestReqRespSubjectName = "testReqRespSubject"
)

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

func TestNatsMQChan(t *testing.T) {
	Convey("TestNatsMQChan", t, func() {
		var err error
		err = CreateDefaultPool(nil, zap.L())
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

func TestNatsMQQueueSubscribe(t *testing.T) {
	Convey("TestNatsMQQueueSubscribe", t, func() {
		var err error
		err = CreateDefaultPool(nil, zap.L())
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

func TestNatsMQQueueChanRecv(t *testing.T) {
	Convey("TestNatsMQQueueChanRecv", t, func() {
		var err error
		err = CreateDefaultPool(nil, zap.L())
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

func TestNatsMQReqResp(t *testing.T) {
	Convey("TestNatsMQReqResp", t, func() {
		var err error
		err = CreateDefaultPool(nil, zap.L())
		So(err, ShouldBeNil)

		// 发布等待响应的请求
		go func() {
			for {
				msg := <-time.NewTicker(time.Second).C // 生成时间消息
				var reply interface{}                  // 接收订阅者响应的消息
				err = XCommonNatsReqResp().Request(TestReqRespSubjectName, msg, reply, time.Second)
				fmt.Println("Request----------------------------------")
				fmt.Println("Subject:", TestReqRespSubjectName)
				fmt.Println("Msg:", msg)
				fmt.Println("Receive Reply Message:", reply)
				fmt.Println("----------------------------------")
			}
		}()
		So(err, ShouldBeNil)

		// 订阅一个请求并发送一个reply消息
		err = XCommonNatsReqResp().SubscribeForRequest(TestReqRespSubjectName, func(subject, reply string, msg interface{}) {
			fmt.Println("Subscribe Receive==============================")
			fmt.Println("subject:", subject)
			fmt.Println("Request Reply Box:", reply)
			fmt.Println("Receive Message:", msg)

			// 返回一个消息给请求者收件箱
			err = XCommonNatsReqResp().PublishReply(reply, "I can help you!")
			fmt.Println("==============================")
		})
		So(err, ShouldBeNil)
		time.Sleep(time.Second * 10)

	})

}
