package XValidate

import (
	"errors"
	"github.com/bb-orz/goinfras"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

var validater *validator.Validate
var translater ut.Translator

// 创建一个默认配置的Manager
func CreateDefaultValidater(config *Config) error {
	var err error
	if config == nil {
		config = DefaultConfig()
	}

	if config.TransZh {
		validater, translater, err = NewZhValidater()
	} else {
		validater = NewValidater()
	}
	return err
}

// 默认验证器
func NewValidater() *validator.Validate {
	return validator.New()
}

// 中文翻译验证器
func NewZhValidater() (*validator.Validate, ut.Translator, error) {
	valid := validator.New()
	// 创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	var trans ut.Translator
	trans, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(valid, trans)
		if err != nil {
			return valid, nil, err
		}
	} else {
		return valid, nil, errors.New("Not found translator: zh! ")
	}

	return valid, trans, nil
}

// 设置启动组级别
func (s *starter) PriorityGroup() goinfras.PriorityGroup { return goinfras.AppGroup }
