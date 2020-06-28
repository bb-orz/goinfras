package user

import (
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
	"errors"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ services.IUserService = new(userService)
var once sync.Once

func init() {
	// 初始化该业务模块时实例化服务
	once.Do(func() {
		services.SetUserService(new(userService))
	})
}

type userService struct{}

func (*userService) CreateUser(dto services.CreateUserDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, err
	}

	// 实例user模块领域模型
	domain := NewUserDomain()

	// 验证参数是否存在
	if isExist, err := domain.IsUserExist(dto); err != nil {
		return nil, errors.New("查询错误! ")
	} else if isExist {
		return nil, errors.New("该用户已经存在! ")
	}

	// 构造user dto
	userDTO := services.UserDTO{
		Name:     dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Status:   1,
	}
	res, err := domain.CreateUser(userDTO)
	return res, err

}

func (*userService) GetUserInfo(userId string) (*services.UserDTO, error) {
	panic("implement me")
}

func (*userService) SetUserInfo(dto services.SetUserInfoDTO) error {
	panic("implement me")
}

func (*userService) BindEmail(email string) error {
	panic("implement me")
}

func (*userService) ValidateEmail(validateCode int) bool {
	panic("implement me")
}

func (*userService) BindPhone(phone string) error {
	panic("implement me")
}

func (*userService) ValidatePhone(validateCode int) bool {
	panic("implement me")
}

func (*userService) SetStatus(status int) int {
	panic("implement me")
}

func (*userService) ChangePassword(dto services.ChangePassword) bool {
	panic("implement me")
}

func (*userService) SendEmailForgetPassword() bool {
	panic("implement me")
}

func (*userService) ReSetPassword(dto services.ReSetPassword) bool {
	panic("implement me")
}

func (*userService) UploadAvatar() bool {
	panic("implement me")
}