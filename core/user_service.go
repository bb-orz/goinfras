package core

import (
	"GoWebScaffold/core/oauth2"
	"GoWebScaffold/core/user"
	"GoWebScaffold/core/verified"
	"GoWebScaffold/infras/global"
	"GoWebScaffold/infras/jwt"
	"GoWebScaffold/infras/oauth"
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
	oauthDomain    *oauth2.OauthDomain
	verifiedDomain *verified.VerifiedDomain
}

// 邮箱创建用户账号
func (service *UserService) CreateUserWithEmail(dto services.CreateUserWithEmailDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 验证用户邮箱是否存在
	if isExist, err := service.userDomain.IsEmailExist(dto.Email); err != nil {
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
	if isExist, err := service.userDomain.IsPhoneExist(dto.Phone); err != nil {
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
func (service *UserService) EmailAuth(dto services.AuthWithEmailPasswordDTO) (string, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	var userDTO *services.UserDTO
	var err error
	// 查找邮件账号是否存在
	if userDTO, err = service.userDomain.GetUserInfoByEmail(dto.Email); err != nil {
		return "", WrapError(err, ErrorFormatServiceStorage)
	}
	if userDTO == nil {
		return "", WrapError(err, ErrorFormatServiceCheckInfo, "该用户不存在!")
	} else if !global.ValidatePassword(dto.Password, userDTO.Salt, userDTO.Password) {
		// 校验密码失败
		return "", WrapError(err, ErrorFormatServiceCheckInfo, "密码错误!")
	}

	// JWT token
	token, err := jwt.TokenUtils().Encode(jwt.UserClaim{
		Id:     userDTO.No,
		Name:   userDTO.Name,
		Avatar: userDTO.Avatar,
	})

	return token, nil
}

// 手机账号登录鉴权
func (service *UserService) PhoneAuth(dto services.AuthWithPhonePasswordDTO) (string, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	var userDTO *services.UserDTO
	var err error
	// 查找手机账号是否存在
	userDTO, err = service.userDomain.GetUserInfoByPhone(dto.Phone)
	if err != nil {
		return "", WrapError(err, ErrorFormatServiceStorage)
	}

	if userDTO == nil {
		return "", WrapError(err, ErrorFormatServiceCheckInfo, "该用户不存在!")
	} else if !global.ValidatePassword(dto.Password, userDTO.Salt, userDTO.Password) {
		// 校验密码失败
		return "", WrapError(err, ErrorFormatServiceCheckInfo, "密码错误!")
	}

	// JWT token
	token, err := jwt.TokenUtils().Encode(jwt.UserClaim{
		Id:     userDTO.No,
		Name:   userDTO.Name,
		Avatar: userDTO.Avatar,
	})

	return token, nil
}

// qq oauth 鉴权
func (service *UserService) QQOAuth(dto services.QQLoginDTO) (string, error) {
	var err error
	var token string
	var qqOauthAccountInfo *oauth.OAuthAccountInfo
	var findUserBindingDTO *services.UserOauthInfoDTO
	var oauthBindingInfo *services.UserOauthInfoDTO

	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// oauth domain：使用qq回调授权码code开始鉴权流程并获取QQ用户信息
	qqOauthAccountInfo, err = service.oauthDomain.GetQQOauthUserInfo(dto.AccessCode)
	if err != nil {
		return "", WrapError(err, ErrorFormatServiceNetRequest, "GetQQUserInfo")
	}

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在
	findUserBindingDTO, err = service.userDomain.GetUserOauthBinding(user.QQOauthPlatform, qqOauthAccountInfo.OpenId, qqOauthAccountInfo.UnionId)
	if err != nil {
		return "", WrapError(err, ErrorFormatServiceStorage, "IsOauthUserExist")
	}

	// 如不存在进入创建用户流程,否则进登录流程
	if findUserBindingDTO == nil {
		oauthBindingInfo, err = service.userDomain.CreateUserOauthBinding(user.QQOauthPlatform, qqOauthAccountInfo)
		// JWT token
		if oauthBindingInfo != nil {
			token, err = jwt.TokenUtils().Encode(jwt.UserClaim{
				Id:     oauthBindingInfo.User.No,
				Name:   oauthBindingInfo.User.Name,
				Avatar: oauthBindingInfo.User.Avatar,
			})
			return token, nil
		} else {
			return "", WrapError(err, ErrorFormatServiceStorage, "CreateUserOauthBinding")
		}
	}

	token, err = jwt.TokenUtils().Encode(jwt.UserClaim{
		Id:     findUserBindingDTO.User.No,
		Name:   findUserBindingDTO.User.Name,
		Avatar: findUserBindingDTO.User.Avatar,
	})

}

// wechat Oauth 鉴权
func (service *UserService) WechatOAuth(dto services.WechatLoginDTO) (string, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// oauth domain：使用微信回调授权码code开始鉴权流程并获取微信用户信息

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在

	// 不存在进入创建用户流程，存在进入登录流程

	// 返回jwt

	return "", nil
}

// 微博 Oauth 鉴权
func (service *UserService) WeiboOAuth(dto services.WeiboLoginDTO) (string, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return "", WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// oauth domain：使用微博回调授权码code开始鉴权流程并获取微博用户信息

	// oauth domain: 使用OpenId UnionId查找user oauth表查看用户是否存在

	// 不存在进入创建用户流程，存在进入登录流程

	// 返回jwt

	return "", nil
}

// 获取用户信息
func (service *UserService) GetUserInfo(dto services.GetUserInfoDTO) (*services.UserDTO, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return nil, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 查找用户信息
	userDTO, err := service.userDomain.GetUserInfo(dto.ID)
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

	uid := dto.ID
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

	// 从cache拿出保存的邮箱验证码
	pass, err := service.verifiedDomain.VerifiedEmail(dto.ID, dto.VerifiedCode)
	if err != nil {
		return false, WrapError(err, ErrorFormatServiceCache, "缓存验证码校验错误")
	}

	if pass {
		return true, nil
	}

	return false, nil
}

// 验证手机号码
func (service *UserService) ValidatePhone(dto services.ValidatePhoneDTO) (bool, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return false, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 从cache拿出保存的短信验证码
	pass, err := service.verifiedDomain.VerifiedPhone(dto.ID, dto.VerifiedCode)
	if err != nil {
		return false, WrapError(err, ErrorFormatServiceCache, "缓存验证码校验错误")
	}

	if pass {
		return true, nil
	}

	return false, nil
}

// 设置用户账号状态
func (service *UserService) SetStatus(dto services.SetStatusDTO) (int, error) {
	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return -1, WrapError(err, ErrorFormatServiceDTOValidate)
	}

	err := service.userDomain.SetStatus(dto.ID, dto.Status)
	if err != nil {
		return -1, WrapError(err, ErrorFormatServiceStorage)
	}

	return 0, nil
}

// 修改用户密码
func (service *UserService) ChangePassword(dto services.ChangePasswordDTO) error {
	var userDTO *services.UserDTO
	var err error
	// 校验传输参数
	if err = validate.ValidateStruct(dto); err != nil {
		return WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 查找账号是否存在
	userDTO, err = service.userDomain.GetUserInfo(dto.ID)
	if err != nil {
		return WrapError(err, ErrorFormatServiceStorage)
	}

	// 校验旧密码
	if userDTO == nil {
		return WrapError(err, ErrorFormatServiceCheckInfo, "该用户不存在!")
	} else if !global.ValidatePassword(dto.Old, userDTO.Salt, userDTO.Password) {
		// 校验旧密码失败
		return WrapError(err, ErrorFormatServiceCheckInfo, "旧密码错误!")
	}

	// 设置新密码
	if err = service.userDomain.ReSetPassword(dto.ID, dto.New); err != nil {
		return WrapError(err, ErrorFormatServiceStorage)
	}

	return nil
}

// 忘记密码重设
func (service *UserService) ForgetPassword(dto services.ForgetPasswordDTO) error {

	// 校验传输参数
	if err := validate.ValidateStruct(dto); err != nil {
		return WrapError(err, ErrorFormatServiceDTOValidate)
	}

	// 查找账号是否存在
	b, err := service.userDomain.IsUserExist(dto.ID)
	if err != nil {
		return WrapError(err, ErrorFormatServiceStorage)
	}
	if !b {
		return WrapError(err, ErrorFormatServiceCheckInfo, "该用户不存在!")
	}

	// 校验Code
	b, err = service.verifiedDomain.VerifiedResetPasswordCode(dto.ID, dto.Code)
	if err != nil {
		return WrapError(err, ErrorFormatServiceCache)
	}

	if !b {
		return WrapError(nil, ErrorFormatServiceCheckInfo, "重置密码校验码错误，请重试！")
	}

	return nil

}

// 上传用户头像
func (service *UserService) UploadAvatar() error {

	return nil
}
