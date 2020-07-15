package core

import (
	"GoWebScaffold/core/user"
	"GoWebScaffold/core/verified"
	"GoWebScaffold/infras/global"
	"GoWebScaffold/infras/validate"
	"GoWebScaffold/services"
	"sync"
)

// 服务层，实现services包定义的服务并设置该服务的实例，
// 需在服务实现的方法中验证DTO传输参数并调用具体的领域层业务逻辑
// 接收领域层和dao层的错误包装处理

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
	userDomain     *user.UserDomain
	verifiedDomain *verified.VerifiedDomain
}

// 邮箱创建用户账号
func (service *UserService) CreateUserWithEmail(dto services.CreateUserWithEmailDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 验证用户邮箱是否存在
	if isExist, err := service.userDomain.IsEmailExist(dto); err != nil {
		return nil, WrapError(err, ErrorFormatServiceStorage)
	} else if isExist {
		return nil, WrapError(err, ErrorFormatServiceCheckInfo, "该用户已经存在!")
	}

	res, err := service.userDomain.CreateUserForEmail(dto)
	if err != nil {
		return nil, WrapError(err, ErrorFormatServiceBusinesslogic, "创建用户失败！")
	}
	return res, nil
}

// 邮箱创建用户账号
func (service *UserService) CreateUserWithPhone(dto services.CreateUserWithPhoneDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 验证用户邮箱是否存在
	if isExist, err := service.userDomain.IsPhoneExist(dto); err != nil {
		return nil, WrapError(err, ErrorFormatServiceStorage)
	} else if isExist {
		return nil, WrapError(err, ErrorFormatServiceCheckInfo, "该用户已经存在!")
	}

	res, err := service.userDomain.CreateUserForPhone(dto)
	if err != nil {
		return nil, WrapError(err, ErrorFormatServiceBusinesslogic, "创建用户失败！")
	}
	return res, nil
}

// 邮箱账号登录鉴权
func (service *UserService) AuthWithEmailPassword(dto services.AuthWithEmailPasswordDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	var userDTO *services.UserDTO
	var err error
	// 查找邮件账号是否存在
	if userDTO, err = service.userDomain.GetUserInfoByEmail(dto.Email); err != nil {
		return false, WrapError(err, ErrorFormatServiceStorage)
	}

	if userDTO == nil {
		return false, WrapError(err, ErrorFormatServiceCheckInfo, "该用户不存在!")
	} else if global.ValidatePassword(dto.Password, userDTO.Salt, userDTO.Password) {
		// 校验密码成功
		return true, nil
	}

	return false, nil
}

// 手机账号登录鉴权
func (service *UserService) AuthWithPhonePassword(dto services.AuthWithPhonePasswordDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	var userDTO *services.UserDTO
	var err error
	// 查找手机账号是否存在
	userDTO, err = service.userDomain.GetUserInfoByPhone(dto.Phone)
	if err != nil {
		return false, WrapError(err, ErrorFormatServiceStorage)
	}

	if userDTO == nil {
		return false, WrapError(err, ErrorFormatServiceCheckInfo, "该用户不存在!")
	} else if global.ValidatePassword(dto.Password, userDTO.Salt, userDTO.Password) {
		// 校验密码成功
		return true, nil
	}

	return false, nil
}

// 获取用户信息
func (service *UserService) GetUserInfo(dto services.GetUserInfoDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 查找用户信息
	userDTO, err := service.userDomain.GetUserInfo(int(dto.ID))
	if err != nil {
		return nil, WrapError(err, ErrorFormatServiceStorage)
	}

	return userDTO, nil
}

// 批量设置用户信息
func (service *UserService) SetUserInfos(dto services.SetUserInfoDTO) error {

	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return WrapError(err, ErrorFormatServiceDTOValidate)
	}

	uid := int(dto.ID)
	err := service.userDomain.SetUserInfos(uid, dto)
	if err != nil {
		return WrapError(err, ErrorFormatServiceStorage)
	}

	return nil
}

// 验证用户邮箱
func (service *UserService) ValidateEmail(dto services.ValidateEmailDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	return true, nil
}

// 验证手机号码
func (service *UserService) ValidatePhone(dto services.ValidatePhoneDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	return true, nil
}

// 设置用户账号状态
func (service *UserService) SetStatus(dto services.SetStatusDTO) (int, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return -1, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	return 0, nil
}

// 修改用户密码
func (service *UserService) ChangePassword(dto services.ChangePassword) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return WrapError(err, ErrorFormatServiceDTOValidate)
	}

	return nil
}

// 重设密码
func (service *UserService) ReSetPassword(dto services.ReSetPassword) error {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return WrapError(err, ErrorFormatServiceDTOValidate)
	}
	return nil

}

// 上传用户头像
func (service *UserService) UploadAvatar() error {
	panic("implement me")
}
