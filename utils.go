package goinfras

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func OpenFile(fileName string, flag int, perm os.FileMode) (*os.File, error) {
	var err error
	var file *os.File

	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			dir := path.Dir(fileName)
			_, err = os.Stat(dir)
			if err != nil {
				if os.IsNotExist(err) {
					err := os.MkdirAll(dir, os.ModePerm)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	file, err = os.OpenFile(fileName, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// 结构体指针检查验证，如果传入的interface为nil，就通过log.Panic函数抛出一个异常，它被用在starter中检查组件资源是否已启动
func Check(a interface{}) error {
	if a == nil {
		_, f, l, _ := runtime.Caller(1)
		strs := strings.Split(f, "/")
		size := len(strs)
		if size > 4 {
			size = 4
		}
		f = filepath.Join(strs[len(strs)-size:]...)
		return fmt.Errorf("object can't be nil, cause by: %s(%d)", f, l)
	}
	return nil
}
