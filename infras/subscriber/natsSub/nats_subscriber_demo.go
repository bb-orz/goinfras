package natsSub

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"sync"
)

func Run() {
	// 订阅需开启NatsMq服务
	if !config.MqConf.NatsMq.Switch {
		return
	}
	fmt.Println("The NatsMq Subscriber Running")
	wg := sync.WaitGroup{}

	// 一、演示订阅并接收普通的字符串消息
	wg.Add(1)
	go func() {
		err := natsMq.Subscribe(natsMq.NatsTestSubject1, func(msg *nats.Msg) {
			fmt.Printf("Nats Subscribe subject:%s,receive massage:%s\n",msg.Subject,msg.Data)
		})
		if err != nil {
			logger.WarmLog(err.Error())
		}
		wg.Done()
	}()

	// 二、演示订阅并接收go数据类型的消息
	wg.Add(1)
	go func() {
		err := natsMq.SubscribeForEncodedMsg(natsMq.NatsTestSubject2, func(subject string, goDataMsg interface{}) {
			fmt.Printf("Nats Subscribe subject:%s,receive massage:%v\n",subject,goDataMsg)
		})

		if err != nil {
			logger.WarmLog(err.Error())
		}
		wg.Done()
	}()

	// 三、演示订阅并在超时时间内处理伪同步的Request消息，并传回处理结果到reply主题信箱
	wg.Add(1)
	go func() {
		// 订阅一个Nats Request 主题
		err := natsMq.SubscribeForRequest(natsMq.NatsTestRequestSubject, func(subj, reply string, msg interface{}){
			fmt.Printf("Nats Subscribe request subject:%s,receive massage:%s,reply subject:%s\n",subj,msg,reply)

			// TODO 处理消息
			data := "do something and return data"

			// 伪同步响应：接收到请求消息后需向响应收件箱发送一条消息作为回应
			err := natsMq.Publish(reply, map[string]interface{}{"res":"ok","data":data})
			if err != nil {
				logger.WarmLog(err.Error())
			}
		})

		if err != nil {
			logger.WarmLog(err.Error())
		}
		wg.Done()
	}()


	// 四、演示队列订阅者接收消息,分两组订阅主题，每组三个订阅成员
	for i:=0;i<3 ;i++  {
		wg.Add(1)
		go func(n int) {
			err := natsMq.QueueSubscribe(natsMq.NatsTestQueueSubject, natsMq.QueueGroup1, func(msg *nats.Msg) {
				fmt.Printf("Testing Queue Subscribe,subject:[%s],QueueGroup1:[%s],message:[%v],receive No.[%d] \n",msg.Subject,natsMq.QueueGroup1,string(msg.Data),n)
			})
			if err != nil {
				logger.WarmLog(err.Error())
			}
			wg.Done()
		}(i)
	}
	for i:=0;i<3 ;i++  {
		wg.Add(1)
		go func(n int) {
			err := natsMq.QueueSubscribe(natsMq.NatsTestQueueSubject, natsMq.QueueGroup2, func(msg *nats.Msg) {
				fmt.Printf("Testing Queue Subscribe,subject:[%s],QueueGroup1:[%s],message:[%v],receive No.[%d] \n",msg.Subject,natsMq.QueueGroup2,string(msg.Data),n)
			})
			if err != nil {
				logger.WarmLog(err.Error())
			}
			wg.Done()
		}(i)
	}


	// 五、演示基于管道接收订阅的消息
	wg.Add(1)
	go func() {
		recMsgChan := make(chan interface{},6)
		err := natsMq.BindRecvChan(natsMq.NatsTestBindChanSubject,recMsgChan )
		if err != nil {
			logger.WarmLog(err.Error())
		}

		for {
			msg := <-recMsgChan
			fmt.Printf("Testing Bind Chan,subject:[%s],receive message:[%v] \n",natsMq.NatsTestBindChanSubject,msg)
		}

		wg.Done()
	}()

	wg.Wait()

}
