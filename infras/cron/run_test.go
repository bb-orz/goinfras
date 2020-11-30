package cron

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
	"time"
)

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	var m *Manager
	if config == nil {
		config = &Config{Location: "Local"}
	}
	// 1.获取Cron执行管理器
	fmt.Println("创建任务执行管理器...")
	m = NewManager(config, zap.L())
	SetComponent(m)
	return err
}

type JobA struct{}

func (j JobA) Run() {
	fmt.Println("Running Job A ...")
}

type JobB struct{}

func (j JobB) Run() {
	fmt.Println("Running Job B ...")
}

func TestCron(t *testing.T) {
	Convey("Test Cron", t, func() {
		err := TestingInstantiation(nil)
		So(err, ShouldBeNil)

		// 1.定义定时任务
		fmt.Println("定义第一个定时任务...")
		tasks := make([]*Task, 0)
		task1 := NewTask("*/2 * * * * *", &JobA{})
		tasks = append(tasks, task1)

		// 2.注册定时运行任务
		fmt.Println("注册第一个定时任务...")
		CronComponent().RegisterTasks(tasks...)

		// 3.运行定时任务
		fmt.Println("开始运行定时任务...")
		CronComponent().RunTasks()

		// 主协程运行5s
		time.Sleep(time.Second * 5)

		// 4.定义定时任务
		fmt.Println("定义第二个定时任务...")
		task2 := NewTask("*/1 * * * * *", &JobB{})

		// 5.注册定时运行任务
		fmt.Println("注册第二个定时任务...")
		CronComponent().RegisterTasks(task2)

		// 添加新的任务后主协程再运行10s
		time.Sleep(time.Second * 10)
	})
}
