package validate

import (
	"GoWebScaffold/infras"
	"github.com/go-playground/universal-translator"
	"github.com/tietang/props/kvs"
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
	cfg *validateConfig
}

func (s *ValidatorStarter) Init(sctx infras.StarterContext) {
	configs := sctx.Configs()
	define := validateConfig{}
	err := kvs.Unmarshal(configs, &define, "Validate")
	infras.FailHandler(err)
	s.cfg = &define
}

func (s *ValidatorStarter) Setup(sctx infras.StarterContext) {
	if s.cfg.transZh {
		validate, translator = NewZhValidator(sctx.Logger())
	} else {
		validate = NewValidator()
	}
}
