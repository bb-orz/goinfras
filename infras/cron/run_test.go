package cron

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
	"testing"
)

type JobA struct{}

func (j JobA) Run() {
	fmt.Println("Running Job A ...")
}

func TestCron(t *testing.T) {
	Convey("Test Cron", t, func() {
		RegisterTask(&Task{spec: "*/2 * * * * *", job: &JobA{}})

		config := cronConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err := p.Unmarshal(&config)
		So(err, ShouldBeNil)
		Println("Cron Config:", config)

		logger, err := zap.NewDevelopment()
		So(err, ShouldBeNil)
		Do(&config, logger)

	})
}
