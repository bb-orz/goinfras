package validate

import (
	"github.com/prometheus/common/log"
	"gopkg.in/go-playground/validator.v9"
)

// 验证 DTO Struct
func ValidateStruct(s interface{}) (err error) {
	err = Validater().Struct(s)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			log.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				log.Error(e.Translate(Translater()))
			}
		}
		return err
	}
	return nil
}
