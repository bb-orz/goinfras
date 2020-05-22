package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)


func DemoWrapper1(j cron.Job) cron.Job {
	return cron.FuncJob(func() {
		// TODO before
		fmt.Println("DemoWrapper1 before")
		j.Run()
		fmt.Println("DemoWrapper1 after")
		// TODO after
	})
}


func DemoWrapper2(j cron.Job) cron.Job {
	return cron.FuncJob(func() {
		// TODO before
		fmt.Println("DemoWrapper2 before")
		j.Run()
		fmt.Println("DemoWrapper2 after")
		// TODO after
	})
}

