package XGorm

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
	return "XGorm"
}

// 读取配置
func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define *Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("Gorm", &define)
		goinfras.ErrorHandler(err)
	}
	if define == nil {
		define = DefaultConfig()
	}
	s.cfg = define
	fmt.Printf("XGORM Starter Init: [Config] %v \n", *define)
}

// 连接数据库
func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	db, err = NewORMDb(s.cfg)
	goinfras.ErrorHandler(err)
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	err := goinfras.Check(db)
	if err != nil {
		sctx.Logger().Error(fmt.Sprintf("[%s Starter]: GORM DB Setup Fail!", s.Name()))
		return false
	}

	sctx.Logger().Info(fmt.Sprintf("[%s Starter]: GORM DB Setup Successful!", s.Name()))
	return true
}

func (s *starter) Stop() {
	d, _ := db.DB()
	d.Close()
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.ResourcesGroup }
