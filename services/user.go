package services

import (
	"GoWebScaffold/infras"
)

/* 定义用户模块的服务层方法，并定义数据传输对象DTO*/

var userService IUserService

// 用于对外暴露账户应用服务，唯一的暴露点，供接口层调用
func GetUserService() IUserService {
	infras.Check(userService)
	return userService
}

// 服务具体实现初始化时设置服务对象，供核心业务层具体实现并设置
func SetUserService(service IUserService) {
	userService = service
}

// 定义用户服务接口
type IUserService interface {
	CreateUser(dto CreateUserDTO) (*UserDTO, error) // 创建用户
	GetUserInfo(userId string) (*UserDTO, error)    // 获取用户数据
	SetUserInfo(dto SetUserInfoDTO) error           // 修改用户信息
	BindEmail(email string) error                   // 绑定邮箱，发送验证邮件到指定邮箱
	ValidateEmail(validateCode int) bool            // 绑定邮箱，验证邮箱链接
	BindPhone(phone string) error                   // 绑定手机，发送短信验证码
	ValidatePhone(validateCode int) bool            // 绑定手机，验证短信验证码
	SetStatus(status int) int                       // 设置用户锁定状态
	ChangePassword(dto ChangePassword) bool         // 更改用户密码
	SendEmailForgetPassword() bool                  // 忘记密码，发送邮件到用户绑定的邮箱
	ReSetPassword(dto ReSetPassword) bool           // 重设密码
	UploadAvatar() bool                             // 上传头像
}

// 创建用户的数据传输对象
type CreateUserDTO struct {
	Username   string `validate:"required,alphanum"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,alphanumunicode"`
	RePassword string `validate:"required,alphanumunicode,eqfield=Password"`
}

// 用户数据传输对象
type UserDTO struct {
	Name     string
	Age      byte
	Avatar   string
	Gender   int8
	Email    string
	Phone    string
	Password string
	Status   int8
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

type ChangePassword struct {
	old   string
	new   string
	reNew string
}

type ReSetPassword struct {
	code  string // 允许重设密码的key值，服务端生成后被发往邮箱，用户点击过来后接收
	new   string
	reNew string
}
