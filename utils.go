package goinfras

import (
	"os"
	"path"
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
