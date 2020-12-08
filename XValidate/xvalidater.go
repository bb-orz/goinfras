package XValidate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
	"goinfras/XLogger"
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
				log.Error(e.Translate(XTranslater()))
			}
		}
		return err
	}
	return nil
}
