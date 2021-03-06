package XValidate

import (
	"github.com/bb-orz/goinfras/XLogger"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

// 验证器
func XValidater() *validator.Validate {
	return validater
}

// 验证信息翻译器
func XTranslater() ut.Translator {
	return translater
}

// 验证 DTO Struct
func V(s interface{}) (err error) {
	// 开始验证并判断错误类型
	err = XValidater().Struct(s)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			// 无效数据验证错误的记录日志
			XLogger.XCommon().Error("验证错误", zap.Error(err))
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				// 错误类型翻译打印
				// log.Error(e.Translate(XTranslater()))
				e.Translate(XTranslater())
			}
		}
		return err
	}
	return nil
}
