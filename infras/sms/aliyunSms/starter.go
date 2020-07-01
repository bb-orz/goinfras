package aliyunSms

import "GoWebScaffold/infras"

type aliyunSmsStarter struct {
	infras.BaseStarter
	cfg *AliyunSmsConfig
}

func (s *aliyunSmsStarter) Init(sctx *infras.StarterContext) {
	panic("implement me")
}

func (s *aliyunSmsStarter) Setup(sctx *infras.StarterContext) {
	panic("implement me")
}

func (s *aliyunSmsStarter) Start(sctx *infras.StarterContext) {
	panic("implement me")
}

func (s *aliyunSmsStarter) SetStartBlocking() bool {
	panic("implement me")
}

func (s *aliyunSmsStarter) Stop(sctx *infras.StarterContext) {
	panic("implement me")
}

func (s *aliyunSmsStarter) PriorityGroup() infras.PriorityGroup {
	panic("implement me")
}

func (s *aliyunSmsStarter) Priority() int {
	panic("implement me")
}
