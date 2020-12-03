package XValidate

import (
	"GoWebScaffold/hub"
	ut "github.com/go-playground/universal-translator"
	"github.com/prometheus/common/log"
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

// 验证 DTO Struct
func V(s interface{}) (err error) {
	err = Validator().Struct(s)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			log.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				log.Error(e.Translate(Translator()))
			}
		}
		return err
	}
	return nil
}
