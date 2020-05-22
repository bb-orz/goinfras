package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

/*
使用github.com/robfig/cron/v3包构建定时任务服务

一、关于spec参数：与linux cron 基本一样

 ┌─────────────second 范围 (0 - 59)，允许特殊字符：* / , -
 │ ┌───────────── min (0 - 59),允许特殊字符：* / , -
 │ │ ┌────────────── hour (0 - 23)，允许特殊字符：* / , -
 │ │ │ ┌─────────────── day of month (1 - 31)，允许特殊字符：* / , - ?
 │ │ │ │ ┌──────────────── month (1 - 12)，允许特殊字符：* / , -
 │ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
 │ │ │ │ │ │                  Saturday)，允许特殊字符： * / , - ?
 │ │ │ │ │ │
 │ │ │ │ │ │
 * * * * * *

Cron 特殊字符
星号(*) :表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
斜线(/):表示增长间隔，如第2个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后 每隔 15 分钟执行一次（即 3（3+0*15）、18（3+1*15）、33（3+2*15）、48（3+3*15） 这些时间点执行），这里也可以表示为：3/15
逗号(,):用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
连字号(-):表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
问号(?):只用于 日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替 *


Cron预定义时间的parser
输入						|	简述								|	相当于
---------------------------------------------------------------------------
@yearly 				|  	(or @annually)	1月1日午夜运行一次	|	0 0 0 1 1 *
@monthly				|	每个月的午夜，每个月的第一个月运行一次	|	0 0 0 1 * *
@weekly					|  	每周一次，周日午夜运行一次			|	0 0 0 * * 0
@daily (or @midnight)	|  	每天午夜运行一次					|	0 0 0 * * *
@hourly					|	每小时运行一次						|	0 0 * * * *



二、关于添加Jobs:

用法1：添加单独的定时任务
_, err = c.AddFunc("*\/5 * * * * *", func() {

})
if err != nil {
	return err
}

用法2：添加自定义的带参数任务,实现Job接口
type DemoJob1 struct {
	Args string // 可以理解为cmd的执行参数，传入AddJob()前设置
}

func (job *DemoJob1) Run() {
	fmt.Println("The demo job1's args is :",job.Args)
}
_, err = c.AddJob("*\/5 * * * * *", &DemoJob1{"arg1"})
if err != nil {
	return err
}

用法3：设置一个schedule执行时间，为这个schedule设置多个自定义任务

 (1)标准parser 不包含解析秒的schedule
 schedule, err := cron.ParseStandard("5 * * * *")

 (2)如需包含秒需NewParser
 schedule, err := cron.NewParser(
  cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor, ).Parse("5 * * * * *")

c.Schedule(schedule,new(DemoJob))
c.Schedule(schedule,cron.FuncJob(func() {fmt.Println("another func job 1...")}))
c.Schedule(schedule,cron.FuncJob(func() {fmt.Println("another func job 2...")}))
c.Schedule(schedule,cron.FuncJob(func() {fmt.Println("another func job 3...")}))


用法4：定义一个cron执行链
chainJobs := cron.NewChain(DemoWrapper1, DemoWrapper2).Then(&DemoJob{"arg"})
_, err = c.AddJob("1 * * * * *", chainJobs)
if err != nil {
	return err
}

*/
func Run() {
	c := cron.New(
		cron.WithSeconds(), // 提供秒字段的parser，如无该option秒字段不解析
		// cron.WithChain(DemoWrapper1, DemoWrapper2,cron.Recover(new(CronLogger))), // 使用自定义的全局cron执行链，在chain_cron.go定义
		// cron.WithLogger(new(CronLogger)),  // 使用自定义的cron执行日志，在logger_cron.go定义
	)

	// TODO Add Schedules and Jobs
	// 测试redisMq发布消息
	// c.AddJob("*/3 * * * * *",&RedisPublishJob1{})
	// c.AddJob("*/6 * * * * *",&RedisPublishJob2{})

	// 测试natsMq发布消息
	// c.AddJob("*/3 * * * * *",&NatsPublishJob1{}) // 发布普通消息
	// c.AddJob("*/6 * * * * *",&NatsPublishJob2{TestMsgData{"fun",19,1}}) //发布go数据类型消息
	// c.AddJob("*/3 * * * * *",&NatsRequestJob{Msg:"help me do something"}) // 发送伪同步消息并在超时前等待响应信息
	// c.AddJob("*/3 * * * * *",&NatsQueueJob{Msg:"queue message"}) // 基于队列发布消息，多组订阅
	// c.AddJob("*/3 * * * * *",&NatsBindChanJob{})


	if len(c.Entries()) > 0 {
		c.Start()
		fmt.Println("The Cron Jobs Running")
	}else {
		fmt.Println("No Cron Entries")
	}

}
