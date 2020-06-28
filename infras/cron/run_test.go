package cron

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
	"testing"
	"time"
)

type JobA struct{}

func (j JobA) Run() {
	fmt.Println("Running Job A ...")
}

func TestCron(t *testing.T) {
	Convey("Test Cron", t, func() {
		config := CronConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Cron Config:", config)

		manager := NewCronManager(&config, zap.L())
		// 注册任务
		Println("Register Tasks...")
		task1 := NewTask("*/2 * * * * *", &JobA{})
		manager.RegisterTasks(task1)

		Println("Run Tasks...")
		manager.RunTasks()

		time.Sleep(10 * time.Second)
		Println("Stop Cron...")
		manager.StopCron()
	})
}
