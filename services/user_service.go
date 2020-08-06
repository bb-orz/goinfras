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

	EmailAuth(dto AuthWithEmailPasswordDTO) (string, error) // 邮箱账号鉴权
	PhoneAuth(dto AuthWithPhonePasswordDTO) (string, error) // 手机号码鉴权

	QQOAuth(dto QQLoginDTO) (string, error)         // qq三方账号鉴权
	WechatOAuth(dto WechatLoginDTO) (string, error) // 微信三方账号鉴权
	WeiboOAuth(dto WeiboLoginDTO) (string, error)   // 微博三方账号鉴权

	GetUserInfo(dto GetUserInfoDTO) (*UserDTO, error) // 获取用户信息
	SetUserInfos(dto SetUserInfoDTO) error            // 修改用户信息
	ValidateEmail(dto ValidateEmailDTO) (bool, error) // 绑定邮箱，验证邮箱链接
	ValidatePhone(dto ValidatePhoneDTO) (bool, error) // 绑定手机，验证短信验证码
	SetStatus(dto SetStatusDTO) (int, error)          // 设置用户状态
	ChangePassword(dto ChangePasswordDTO) error       // 更改用户密码
	ForgetPassword(dto ForgetPasswordDTO) error       // 忘记密码重设
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
	Status        uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

// 三方平台授权信息
type OAuthInfoDTO struct {
	Platform    uint
	AccessToken string
	OpenId      string
	UnionId     string
	NickName    string
	Gender      uint
	Avatar      string
}

// 包含三方账号绑定信息的用户信息
type UserOAuthInfoDTO struct {
	User       UserDTO
	UserOAuths []OAuthInfoDTO
}

type QQLoginDTO struct {
	AccessCode string `validate:"required"`
}

type WechatLoginDTO struct {
	AccessCode string `validate:"required"`
}

type WeiboLoginDTO struct {
	AccessCode string `validate:"required"`
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
	ID           uint   `validate:"required,numeric"`
	VerifiedCode string `validate:"required,alphanum"`
}

type ValidatePhoneDTO struct {
	ID           uint   `validate:"required,numeric"`
	VerifiedCode string `validate:"required,numeric"`
}

type ChangePasswordDTO struct {
	ID    uint   `validate:"required,numeric"`
	Old   string `validate:"required,alphanumunicode"`
	New   string `validate:"required,alphanumunicode"`
	ReNew string `validate:"required,alphanumunicode"`
}

type ForgetPasswordDTO struct {
	ID    uint   `validate:"required,numeric"`
	Code  string `validate:"required,alphanum"` // 允许重设密码的key值，服务端生成后被发往邮箱，用户点击过来后接收
	New   string `validate:"required,alphanumunicode"`
	ReNew string `validate:"required,alphanumunicode"`
}
