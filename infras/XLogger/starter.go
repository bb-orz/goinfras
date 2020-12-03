package XLogger

import (
	"GoWebScaffold/infras"
	"fmt"
	"io"
)

type starter struct {
	infras.BaseStarter
	cfg     Config
	Writers []io.Writer
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = Config{}
	return starter
}

func (s *starter) Name() string {
	return "XLogger"
}

func (s *starter) Init(sctx *infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Logger", &define)
	infras.FailHandler(err)
	fmt.Println(s.cfg)
	s.cfg = define
}

func (s *starter) Setup(sctx *infras.StarterContext) {
	commonLogger = NewCommonLogger(&s.cfg, s.Writers...)
	sctx.Logger().Info("CommonLogger Setup Successful!")

	syncErrorLogger = NewSyncErrorLogger(&s.cfg)
	sctx.Logger().Info("SyncErrorLogger Setup Successful!")

	sctx.SetLogger(commonLogger)
}

func (s *starter) Check(sctx *infras.StarterContext) bool {
	var err1, err2 error
	err1 = infras.Check(commonLogger)
	err2 = infras.Check(syncErrorLogger)

	if err1 != nil || err2 != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: Zap Logger Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: Zap Logger Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop() {
	// 关闭前刷入日志数据
	commonLogger.Sync()
	syncErrorLogger.Sync()
}

func (s *starter) Priority() int { return infras.INT_MAX }
