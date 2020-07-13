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
func WrapError(err error, messageFormat string, msgArgs ...interface{}) CError {
	return CError{
		Inner:      err,                                    // 存储我们正在包装的错误。 如果需要调查发生的事情，我们总是希望能够查看到最低级别的错误。
		Message:    fmt.Sprintf(messageFormat, msgArgs...), // 格式化错误信息，第一参数为信息格式，后为内容值参数
		StackTrace: string(debug.Stack()),                  // 记录了创建错误时的堆栈跟踪。
		Misc:       make(map[string]interface{}),           // 创建一个杂项信息存储字段。可以存储并发ID，堆栈跟踪的hash或可能有助于诊断错误的其他上下文信息。
	}
}

// 领域层错误信息格式
const (
	// 领域层SQL数据库执行错误信息格式
	DomainSqlQueryError  = "SQL Query  Error：[LEVEL]:%s [CALL]:dao.%s"
	DomainSqlInsertError = "SQL Insert Error：[LEVEL]:%s [CALL]:dao.%s"
	DomainSqlUpdateError = "SQL Update Error：[LEVEL]:%s [CALL]:dao.%s"
	DomainSqlDeleteError = "SQL Delete Error：[LEVEL]:%s [CALL]:dao.%s"

	// 领域层缓存执行错误信息格式
	DomainCacheSetError = "Cache Set Error:[LEVEL]:%s [CALL]:cache.%s"
	DomainCacheGetError = "Cache Get Error:[LEVEL]:%s [CALL]:cache.%s"

	// 领域层算法类错误信息格式
	DomainAlgorithmError = "Algorithm Error:[LEVEL]:%s [CALL]:%s"
)

// 服务层错误信息格式
const (
	SerivceDTOValidateError = "DTO Validate Error: [SERVICE]:%s"
)
