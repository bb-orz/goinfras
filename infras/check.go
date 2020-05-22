package infras

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// 结构体指针检查验证，如果传入的interface为nil，就通过log.Panic函数抛出一个异常
// 被用在starter中检查公共资源是否被实例化了
func Check(a interface{}) {
	if a == nil {
		_, f, l, _ := runtime.Caller(1)
		strs := strings.Split(f, "/")
		size := len(strs)
		if size > 4 {
			size = 4
		}
		f = filepath.Join(strs[len(strs)-size:]...)
		_ = fmt.Errorf("object can't be nil, cause by: %s(%d)", f, l)
	}
}
