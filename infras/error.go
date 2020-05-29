package infras

import (
	"fmt"
	"runtime"
	"strconv"
)

// 错误处理，err == nil ,return true
func ErrorHandler(err error) bool {
	if err == nil {
		return true
	} else {
		var path string
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + "" + strconv.Itoa(line)
		}
		_ = fmt.Errorf("[ERROR] %s; [On] %s:%d", err.Error(), path)
		return false
	}
}

// 致命错误处理，直接panic
func FailHandler(err error) bool {
	if err == nil {
		return true
	} else {
		var path string
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + "" + strconv.Itoa(line)
		}
		panic(fmt.Sprintf("[FAIL] %s; [On] %s:%d", err.Error(), path))
	}
}
