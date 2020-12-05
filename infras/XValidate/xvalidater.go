package XValidate

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

var validater *validator.Validate
var translater ut.Translator

// 验证器
func XValidater() *validator.Validate {
	return validater
}

// 验证信息翻译器
func XTranslater() ut.Translator {
	return translater
}

/*实例化资源用于测试*/
func TestingInstantiation(config *Config) error {
	var err error
	if config == nil {
		config = &Config{
			true,
		}
	}

	if config.TransZh {
		validater, translater, err = NewZhValidater()
	} else {
		validater = NewValidater()
	}
	return err
}
