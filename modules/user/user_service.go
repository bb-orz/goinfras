package user

import (
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
	"errors"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ services.IUserService = new(UserService)
var once sync.Once

func init() {
	// 初始化该业务模块时实例化服务
	once.Do(func() {
		userService := new(UserService)
		userService.userDomain = NewUserDomain()
		services.SetUserService(userService)
	})
}

type UserService struct {
	userDomain *UserDomain
}

func (service *UserService) CreateUserWithEmail(dto services.CreateUserWithEmailDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, err
	}

	// 验证用户邮箱是否存在
	if isExist, err := service.userDomain.IsUserEmailExist(dto); err != nil {
		return nil, errors.New("查询错误! ")
	} else if isExist {
		return nil, errors.New("该用户已经存在! ")
	}

	res, err := service.userDomain.CreateUserForEmail(dto)
	return res, err

}

func (service *UserService) GetUserInfo(dto services.GetUserInfoDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, err
	}

	// 查找用户信息
	userDTO, err := service.userDomain.GetUserInfo(int(dto.ID))
	if err != nil {
		return nil, err
	}

	return userDTO, nil
}

func (service *UserService) SetUserInfos(dto services.SetUserInfoDTO) error {

	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	uid := int(dto.ID)
	return service.userDomain.SetUserInfos(uid, dto)
}

// 发起绑定邮箱操作
func (service *UserService) BindEmail(dto services.BindEmailDTO) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	return nil
}

func (service *UserService) ValidateEmail(dto services.ValidateEmailDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, err
	}

	return true, nil
}

func (service *UserService) BindPhone(dto services.BindPhoneDTO) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	return nil
}

func (service *UserService) ValidatePhone(dto services.ValidatePhoneDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, err
	}

	return true, nil
}

func (service *UserService) SetStatus(dto services.SetStatusDTO) (int, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return -1, err
	}

	return 0, nil
}

func (service *UserService) ChangePassword(dto services.ChangePassword) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	return nil
}

func (service *UserService) SendEmailForgetPassword() (bool, error) {
	return false, nil
}

func (service *UserService) ReSetPassword(dto services.ReSetPassword) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}
	return nil

}

func (service *UserService) UploadAvatar() error {
	panic("implement me")
}
