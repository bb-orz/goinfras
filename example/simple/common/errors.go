package common

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
	ErrorFormatDomainSqlQuery      = "[Domain Error]: SQL Query  Error | [LEVEL]:%s | [CALL]:dao.%s"     // sql查询错误
	ErrorFormatDomainSqlInsert     = "[Domain Error]: SQL Insert Error | [LEVEL]:%s | [CALL]:dao.%s"     // sql插入错误
	ErrorFormatDomainSqlUpdate     = "[Domain Error]: SQL Update Error | [LEVEL]:%s | [CALL]:dao.%s"     // sql更新错误
	ErrorFormatDomainSqlDelete     = "[Domain Error]: SQL Delete Error | [LEVEL]:%s | [CALL]:dao.%s"     // sql删除错误
	ErrorFormatDomainSqlShamDelete = "[Domain Error]: SQL ShamDelete Error | [LEVEL]:%s | [CALL]:dao.%s" // sql更新deleted_at字段错误，假删除

	// 领域层缓存执行错误信息格式
	ErrorFormatDomainCacheSet = "[Domain Error]: Cache Set Error | [LEVEL]:%s | [CALL]:cache.%s" // 缓存设置错误
	ErrorFormatDomainCacheGet = "[Domain Error]: Cache Get Error | [LEVEL]:%s | [CALL]:cache.%s" // 缓存获取错误

	// 网络请求报错
	ErrorFormatDomainNetRequest = "[Domain Error]: Network Request Error | [Request]:%s"   // 网络请求相关错误
	ErrorFormatDomainThirdPart  = "[Domain Error]: Network ThirdPart Error | [Request]:%s" // 第三方接口错误相关错误

	// 领域层算法逻辑类错误信息格式
	ErrorFormatDomainAlgorithm = "[Domain Error]: Algorithm Error | [LEVEL]:%s | [CALL]:%s" // 算法执行错误
)

// 服务层错误信息格式
const (
	ErrorFormatServiceDTOValidate   = "[Service Error]: Validate DTO Error"                   // DTO验证错误信息
	ErrorFormatServiceCheckInfo     = "[Service Error]: Check Info Error | [Info]:%s"         // 信息检查相关错误
	ErrorFormatServiceStorage       = "[Service Error]: Storage Data Error"                   // 存储相关错误
	ErrorFormatServiceCache         = "[Service Error]: Cache Data Error"                     // 存储相关错误
	ErrorFormatServiceBusinesslogic = "[Service Error]: Business Logic Error | [Info]:%s"     // 业务逻辑相关错误
	ErrorFormatServiceNetRequest    = "[Service Error]: Network Request Error | [Request]:%s" // 网络请求相关错误

)
