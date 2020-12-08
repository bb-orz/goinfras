package services

import (
	"goinfras"
)

/*
定义服务及数据传输对象
*/
var service1 IService1

func SetService1(sv IService1) {
	service1 = sv
}

func GetService1() IService1 {
	Check(service1)
	return service1
}

type IService1 interface {
	Foo(i InDTO) error
	Bar(i InDTO) error
}

type InDTO struct {
	Email string `validate:"required,email"`
}
