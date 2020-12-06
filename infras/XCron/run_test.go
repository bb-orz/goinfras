package XCron

import (
	"GoWebScaffold/infras"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"testing"
	"time"
)

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
		CreateDefaultManager()

		// 1.定义定时任务
		fmt.Println("定义第一个定时任务...")
		tasks := make([]*Task, 0)
		task1 := NewTask("*/2 * * * * *", &JobA{})
		tasks = append(tasks, task1)

		// 2.注册定时运行任务
		fmt.Println("注册第一个定时任务...")
		XManager().RegisterTasks(tasks...)

		// 3.运行定时任务
		fmt.Println("开始运行定时任务...")
		XManager().RunTasks()

		// 主协程运行5s
		time.Sleep(time.Second * 5)

		// 4.定义定时任务
		fmt.Println("定义第二个定时任务...")
		task2 := NewTask("*/1 * * * * *", &JobB{})

		// 5.注册定时运行任务
		fmt.Println("注册第二个定时任务...")
		XManager().RegisterTasks(task2)

		// 添加新的任务后主协程再运行10s
		time.Sleep(time.Second * 10)
	})
}

// 测试启动器
func TestStarter(t *testing.T) {
	Convey("Test XCron Starter", t, func() {

		// 1.定义定时任务
		fmt.Println("定义第一个定时任务...")
		tasks := make([]*Task, 0)
		task1 := NewTask("*/1 * * * * *", &JobA{})
		tasks = append(tasks, task1)

		s := NewStarter(tasks...)
		sctx := infras.CreateDefaultStarterContext(nil, zap.L())
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
		time.Sleep(time.Second * 20)
		s.Stop()
		if len(XManager().client.Entries()) == 0 {
			Println("Starter Stop Successful!")
		} else {
			Println("Starter Stop Exception!")
		}

	})
}
