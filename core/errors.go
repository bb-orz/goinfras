package core

import (
	"fmt"
	"runtime/debug"
)

// 自定义一个错误信息结构体，其实现了go基本错误接口
type CError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func (err CError) Error() string {
	return err.Message
}

// 工具函数：错误信息在系统各模块传递时的“错误包装器”
func WrapError(err error, messagef string, msgArgs ...interface{}) CError {
	return CError{
		Inner:      err,                               // 存储我们正在包装的错误。 如果需要调查发生的事情，我们总是希望能够查看到最低级别的错误。
		Message:    fmt.Sprintf(messagef, msgArgs...), // 格式化错误信息，第一参数为信息格式，后为内容值参数
		StackTrace: string(debug.Stack()),             // 记录了创建错误时的堆栈跟踪。
		Misc:       make(map[string]interface{}),      // 创建一个杂项信息存储字段。可以存储并发ID，堆栈跟踪的hash或可能有助于诊断错误的其他上下文信息。
	}
}

// 自定义查询错误
type ErrorQuery struct{ CError }
type ErrorUserExist struct{ CError }
type ErrorCreate struct{ CError }
type ErrorSet struct{ CError }
type ErrorGet struct {
}
