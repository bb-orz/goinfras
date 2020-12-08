package XValidate

import (
	"github.com/prometheus/common/log"
	"goinfras"
	"gopkg.in/go-playground/validator.v9"
)

// 验证 DTO Struct
func V(s interface{}) (err error) {
	// 开始验证并判断错误类型
	err = XValidater().Struct(s)
	if err != nil {
		// 无效验证打印
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			log.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				// 错误类型翻译打印
				log.Error(e.Translate(XTranslater()))
			}
		}
		return err
	}
	return nil
}
