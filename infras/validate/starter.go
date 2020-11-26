package validate

import (
	"GoWebScaffold/infras"
	"github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate
var translator ut.Translator

// 验证器
func Validator() *validator.Validate {
	infras.Check(validate)
	return validate
}

// 验证信息翻译器
func Translator() ut.Translator {
	infras.Check(translator)
	return translator
}

type Starter struct {
	infras.BaseStarter
	cfg *Config
}

func (s *Starter) Init(sctx infras.StarterContext) {
	viper := sctx.Configs()
	define := Config{}
	err := viper.UnmarshalKey("Validate", &define)
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *Starter) Setup(sctx infras.StarterContext) {
	var err error
	if s.cfg.TransZh {
		validate, translator, err = NewZhValidator()
	} else {
		validate = NewValidator()
	}
	if err != nil {
		sctx.Logger().Error("Validator Error:", zap.Error(err))
	}
}

/*For testing*/
func RunForTesting(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			true,
		}
	}

	if config.TransZh {
		validate, translator, err = NewZhValidator()
	} else {
		validate = NewValidator()
	}
	return err
}
