package cron

import (
	"fmt"
	"sync"
	"time"
)

/*
实现Job接口

如：
type DemoJob struct {
	Args string // 可以理解为cmd的执行参数，传入AddJob()前设置
}

func (job *DemoJob) Run() {
	fmt.Println("The demo job's args is :",job.Args)
}

*/

type RedisPublishJob1 struct {
	Args string // 可以理解为cmd的执行参数，传入AddJob()前设置
}

func (job *RedisPublishJob1) Run() {
	err := redisMq.Publish(redisMq.RedisTestchannel1, "[redis publish test1 msg]")
	if err != nil {
		fmt.Println("Redis Publish Error:", err.Error())
	}
}

type RedisPublishJob2 struct {
	Args string // 可以理解为cmd的执行参数，传入AddJob()前设置
}

func (job *RedisPublishJob2) Run() {
	err := redisMq.Publish(redisMq.RedisTestchannel2, "[redis publish test2 msg]")
	if err != nil {
		fmt.Println("Redis Publish Error:", err.Error())
	}
}

// 测试发布普通的nats消息
type NatsPublishJob1 struct {
	Args string // 可以理解为cmd的执行参数，传入AddJob()前设置
}

func (job *NatsPublishJob1) Run() {
	var err error
	err = natsMq.Publish(natsMq.NatsTestSubject1, "nats publish test1 msg")
	if err != nil {
		fmt.Println("Nats Publish Error:", err.Error())
	}
	fmt.Println("Nats Publish cron jobs test1")
}

// 测试发布go类型数据的消息
type TestMsgData struct {
	Name   string
	Age    int
	Gender int
}

type NatsPublishJob2 struct {
	msg TestMsgData // 传入需要发送的go类型数据
}

func (job *NatsPublishJob2) Run() {
	var err error
	err = natsMq.Publish(natsMq.NatsTestSubject2, job.msg)
	if err != nil {
		fmt.Println("Nats Publish Error:", err.Error())
	}
	fmt.Println("Nats Publish cron jobs test2")
}

// 测试发布伪同步的Request消息，5秒内等待别人订阅处理，并返回处理结果
type NatsRequestJob struct {
	Msg string // 传入需要发送的消息
}

func (job *NatsRequestJob) Run() {
	var err error
	var resp interface{} // 伪同步：传入等待接收的消息

	err = natsMq.Request(natsMq.NatsTestRequestSubject, job.Msg, &resp, 5*time.Second)
	if err != nil {
		fmt.Println("Nats Request Error:", err.Error())
	}
	fmt.Printf("Nats Request cron jobs test1,resp message:%v \n", resp)
}


// 测试发布基于队列订阅者的主题，多个订阅者只有一个能接收,
type NatsQueueJob struct {
	Msg string
}

func (job *NatsQueueJob) Run()  {
	var err error
	err = natsMq.Publish(natsMq.NatsTestQueueSubject, "queue message")
	if err != nil {
		fmt.Println("Nats Publish Error:", err.Error())
	}

	fmt.Println("Nats Publish cron job: NatsQueueJob")
}

type NatsBindChanJob struct {
	ChanMsg chan interface{}
}

func (job *NatsBindChanJob)Run()  {
	var err error
	job.ChanMsg = make(chan interface{},3)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			now := <- ticker.C
			job.ChanMsg <- now
			fmt.Println("send message to chan:",now.Unix())
			<- job.ChanMsg
		}
		wg.Done()
	}()

	err = natsMq.BindSendChan(natsMq.NatsTestBindChanSubject, job.ChanMsg)
	if err != nil {
		fmt.Println("Nats BindSendChan Error:", err.Error())
	}
	wg.Wait()
}