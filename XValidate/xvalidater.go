package XValidate

import (
	ut "github.com/go-playground/universal-translator"
	"goinfras"
	"gopkg.in/go-playground/validator.v9"
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

// 验证器
func XValidater() *validator.Validate {
	return validater
}

// 验证信息翻译器
func XTranslater() ut.Translator {
	return translater
}
