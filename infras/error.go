package infras

import (
	"fmt"
	"runtime"
	"strconv"
)

func ErrorHandler(err error) bool {
	if err != nil {
		var path string
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + "" + strconv.Itoa(line)
		}
		_ = fmt.Errorf("[ERROR] %s; [On] %s:%d", err.Error(), path)
		return false
	}
	return true
}

func FailHandler(err error) bool {
	if err != nil {
		var path string
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + "" + strconv.Itoa(line)
		}
		panic(fmt.Sprintf("[FAIL] %s; [On] %s:%d", err.Error(), path))
	}
	return true
}
