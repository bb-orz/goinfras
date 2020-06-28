package validate

import (
	"GoWebScaffold/infras"
	"github.com/go-playground/universal-translator"
	"github.com/tietang/props/kvs"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate
var translator ut.Translator

// 验证器
func Validater() *validator.Validate {
	infras.Check(validate)
	return validate
}

// 验证信息翻译器
func Translater() ut.Translator {
	infras.Check(translator)
	return translator
}

type ValidatorStarter struct {
	infras.BaseStarter
	cfg *ValidateConfig
}

func (s *ValidatorStarter) Init(sctx infras.StarterContext) {
	configs := sctx.Configs()
	define := ValidateConfig{}
	err := kvs.Unmarshal(configs, &define, "Validate")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *ValidatorStarter) Setup(sctx infras.StarterContext) {
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
func RunForTesting(config *ValidateConfig) error {
	var err error
	if config == nil {
		config = &ValidateConfig{}
		p := kvs.NewEmptyCompositeConfigSource()
		err = p.Unmarshal(config)
		if err != nil {
			return err
		}
	}

	if config.TransZh {
		validate, translator, err = NewZhValidator()
	} else {
		validate = NewValidator()
	}
	return err
}
