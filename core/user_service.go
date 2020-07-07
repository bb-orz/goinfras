package core

import (
	"GoWebScaffold/core/user"
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
	"errors"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑

var _ services.IUserService = new(UserService)

func init() {
	// 初始化该业务模块时实例化服务
	var once sync.Once
	once.Do(func() {
		userService := new(UserService)
		userService.userDomain = user.NewUserDomain()
		services.SetUserService(userService)
	})
}

type UserService struct {
	userDomain *user.UserDomain
}

// 邮箱创建用户账号
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

// 获取用户信息
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

// 批量设置用户信息
func (service *UserService) SetUserInfos(dto services.SetUserInfoDTO) error {

	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	uid := int(dto.ID)
	return service.userDomain.SetUserInfos(uid, dto)
}

// 验证用户邮箱
func (service *UserService) ValidateEmail(dto services.ValidateEmailDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, err
	}

	return true, nil
}

// 验证手机号码
func (service *UserService) ValidatePhone(dto services.ValidatePhoneDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, err
	}

	return true, nil
}

// 设置用户账号状态
func (service *UserService) SetStatus(dto services.SetStatusDTO) (int, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return -1, err
	}

	return 0, nil
}

// 修改用户密码
func (service *UserService) ChangePassword(dto services.ChangePassword) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}

	return nil
}

// 重设密码
func (service *UserService) ReSetPassword(dto services.ReSetPassword) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return err
	}
	return nil

}

// 上传用户头像
func (service *UserService) UploadAvatar() error {
	panic("implement me")
}
