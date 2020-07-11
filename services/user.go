package services

import (
	"GoWebScaffold/infras"
	"time"
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
	CreateUserWithEmail(dto CreateUserWithEmailDTO) (*UserDTO, error) // 创建邮箱账号
	CreateUserWithPhone(dto CreateUserWithPhoneDTO) (*UserDTO, error) // 创建手机号码账号

	AuthWithEmailPassword(dto AuthWithEmailPasswordDTO) (bool, error)
	AuthWithPhonePassword(dto AuthWithPhonePasswordDTO) (bool, error)

	GetUserInfo(dto GetUserInfoDTO) (*UserDTO, error) // 获取用户信息
	SetUserInfos(dto SetUserInfoDTO) error            // 修改用户信息
	ValidateEmail(dto ValidateEmailDTO) (bool, error) // 绑定邮箱，验证邮箱链接
	ValidatePhone(dto ValidatePhoneDTO) (bool, error) // 绑定手机，验证短信验证码
	SetStatus(dto SetStatusDTO) (int, error)          // 设置用户状态
	ChangePassword(dto ChangePassword) error          // 更改用户密码
	ReSetPassword(dto ReSetPassword) error            // 重设密码
	UploadAvatar() error                              // 上传头像
}

// 用户数据传输对象
type UserDTO struct {
	Uid           uint
	No            string
	Name          string
	Age           uint
	Avatar        string
	Gender        uint
	Email         string
	EmailVerified bool
	Phone         string
	PhoneVerified bool
	Password      string
	Salt          string
	Status        int8
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

// 创建用户的数据传输对象
type CreateUserWithEmailDTO struct {
	Name       string `validate:"required,alphanum"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,alphanumunicode"`
	RePassword string `validate:"required,alphanumunicode,eqfield=Password"`
}

// 创建用户的数据传输对象
type CreateUserWithPhoneDTO struct {
	Name       string `validate:"required,alphanum"`
	Phone      string `validate:"required,numeric,eq=11"`
	Password   string `validate:"required,alphanumunicode"`
	RePassword string `validate:"required,alphanumunicode,eqfield=Password"`
}

// 邮箱密码鉴权数据传输对象
type AuthWithEmailPasswordDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanumunicode"`
}

// 手机号密码鉴权数据传输对象
type AuthWithPhonePasswordDTO struct {
	Phone    string `validate:"required,numeric,eq=11"`
	Password string `validate:"required,alphanumunicode"`
}

type GetUserInfoDTO struct {
	ID uint `validate:"required,numeric"`
}

// 修改用户新息的数据传输对象
type SetUserInfoDTO struct {
	ID     uint   `validate:"required,numeric"`
	Name   string `validate:"alpha"`
	Age    byte   `validate:"numeric"`
	Avatar string `validate:"alphanumunicode"`
	Gender int8   `validate:"numeric"`
	Status int8   `validate:"numeric"`
}

type SetStatusDTO struct {
	ID     uint `validate:"required,numeric"`
	Status uint `validate:"required,numeric"` // TODO 验证枚举0/1/2
}

type ValidateEmailDTO struct {
	VerifiedCode string `validate:"required,alphanum"`
}

type ValidatePhoneDTO struct {
	VerifiedCode string `validate:"required,numeric"`
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
