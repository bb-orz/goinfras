package starter

import (
	"GoWebScaffold/infras"
	"GoWebScaffold/infras/logger"
	"io"
)

type LoggerStarter struct {
	infras.BaseStarter
	Writers []io.Writer
}

// 初始化日志记录器并设置到应用启动器上下文
func (s *LoggerStarter) Init(sctx *StarterContext) {
	sctx.SetCommonLogger(logger.CommonLogger(sctx.GetConfig(), s.writers...))
	sctx.SetSyncErrorLogger(logger.SyncErrorLogger(sctx.GetConfig()))
}
