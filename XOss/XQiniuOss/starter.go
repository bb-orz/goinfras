package XQiniuOss

import (
	"fmt"
	"github.com/bb-orz/goinfras"
)

type starter struct {
	goinfras.BaseStarter
	cfg *Config
}

func NewStarter() *starter {
	starter := new(starter)
	starter.cfg = &Config{}
	return starter
}

func (s *starter) Name() string {
	return "XQiniuOss"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("QiniuOss", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	fmt.Printf("XQiniuOss Starter Init: [Config] %v \n", *define)
}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	qiniuOssClient = NewQnClient(s.cfg)
	sctx.Logger().Info("QiniuOss Setup Successful!")
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(qiniuOssClient)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: QiniuOss Client Setup Fail!", s.Name()))
		return false
	}
	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: QiniuOss Client Setup Successful!", s.Name()))
	return true
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
