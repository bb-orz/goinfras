package XLogger

import (
	"GoWebScaffold/infras"
	"fmt"
	"io"
)

type Starter struct {
	infras.BaseStarter
	cfg     Config
	Writers []io.Writer
}

func NewStarter() *Starter {
	starter := new(Starter)
	starter.cfg = Config{}
	return starter
}

func (s *Starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Logger", &define)
	infras.FailHandler(err)
	fmt.Println(s.cfg)
	s.cfg = define
}

func (s *Starter) Setup(sctx *infras.StarterContext) {
	cl := NewCommonLogger(&s.cfg, s.Writers...)
	SetComponentForCommonLogger(cl)
	sel := NewSyncErrorLogger(&s.cfg)
	SetComponentForSyncErrorLogger(sel)
	sctx.SetLogger(commonLogger)
	sctx.Logger().Info("CommonLogger And SyncErrorLogger Setup Successful!")
}

func (s *Starter) Stop() {
	// 关闭前刷入日志数据
	CLogger().Sync()
	SELogger().Sync()
}

func (s *Starter) Priority() int { return infras.INT_MAX }
