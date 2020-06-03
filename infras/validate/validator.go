package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

// 默认验证器
func NewValidator() *validator.Validate {
	return validator.New()
}

// 中文翻译验证器
func NewZhValidator(logger *zap.Logger) (*validator.Validate, ut.Translator) {
	valid := validator.New()
	// 创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	var trans ut.Translator
	trans, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logger.Error("Translate Error", zap.Error(err))
		}
	} else {
		logger.Error("Not found translator: zh")
	}

	return valid, trans
}