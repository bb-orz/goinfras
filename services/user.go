package services

import "time"

// 创建用户的数据传输对象
type CreateUserDTO struct {
	Username string `validate:"required,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanumunicode"`
}

// 用户数据传输对象
type UserDTO struct {
	ID       uint
	Name     string
	Age      byte
	Avatar   string
	Gender   int8
	Email    string
	Phone    string
	Password string
	Salt     string
	Status   int8
	UpdateAt time.Time
	CreateAt time.Time
}

// 修改用户新息的数据传输对象
type SetUserInfoDTO struct {
	ID     uint   `validate:"required,alpha"`
	Name   string `validate:"alpha"`
	Age    byte   `validate:"numeric"`
	Avatar string `validate:"alphanumunicode"`
	Gender int8   `validate:"numeric"`
	Status int8   `validate:"numeric"`
}

// 定义用户服务接口
type UserService interface {
	CreateUser(dto CreateUserDTO) (*UserDTO, error) // 创建用户
	GetUserInfo(userId string) (*UserDTO, error)    // 获取用户数据
	SetUserInfo(dto SetUserInfoDTO) error           // 修改用户信息
}
