package validate

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

// 默认验证器
func NewValidator() *validator.Validate {
	return validator.New()
}

// 中文翻译验证器
func NewZhValidator() (*validator.Validate, ut.Translator, error) {
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
