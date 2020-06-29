package user

import (
	"GoWebScaffold/common"
	"GoWebScaffold/services"
	"fmt"
	"github.com/segmentio/ksuid"
)

/*领域层：实现具体业务逻辑*/
type userDomain struct {
	dao   *UserDAO
	cache *userCache
}

func NewUserDomain() *userDomain {
	domain := new(userDomain)
	domain.dao = NewUserDao()
	domain.cache = NewUserCache()
	return domain
}

// 生成用户编号
func (domain *userDomain) generateUserNo() string {
	// 采用ksuid的ID生成策略来创建No
	// 全局唯一的ID
	return ksuid.New().Next().String()
}

// 加密密码，设置密文和盐值到po
func (domain *userDomain) encryptPassword(password string) (hashStr, salt string) {
	hashStr, salt = common.HashPassword(password)
	return
}

// 查找用户是否已存在
func (domain *userDomain) IsUserEmailExist(dto services.CreateUserWithEmailDTO) (bool, error) {
	if isExist, err := domain.dao.IsEmailExist(dto.Email); err != nil {
		fmt.Println("Error:", err)
		return false, err
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 邮箱账号创建用户
func (domain *userDomain) CreateUserForEmail(dto services.CreateUserWithEmailDTO) (*services.UserDTO, error) {
	userModel := User{}
	userModel.Name = dto.Name
	userModel.Email = dto.Email
	userModel.No = domain.generateUserNo()
	userModel.Password, userModel.Salt = domain.encryptPassword(dto.Password)
	userModel.Status = UserStatusNotVerified // 初始创建时未验证状态
	var user *User
	var err error
	if user, err = domain.dao.Create(userModel); err != nil {
		return nil, err
	}

	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *userDomain) GetUserInfo(uid int) (*services.UserDTO, error) {
	user, err := domain.dao.GetById(uid)
	if err != nil {
		return nil, err
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

// 生成验证码
func (domain *userDomain) GenerateVerifiedEmailCode(uid int) (string, error) {
	// 生成6位随机字符串
	code := common.RandomString(6)

	// 保存到缓存
	err := domain.cache.SetUserVerifiedEmailCode(uid, code)
	if err != nil {
		return "", err
	}

	return code, nil
}

// 验证邮箱
func (domain *userDomain) VerifiedEmail(uid int, vcode string) (bool, error) {

	// 缓存取出
	code, err := domain.cache.GetUserVerifiedEmailCode(uid)
	if err != nil {
		return false, err
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}

// 生成验证码
func (domain *userDomain) GenerateVerifiedPhoneCode(uid int) (string, error) {
	var err error
	var code string
	// 生成6位随机字符串
	code, err = common.RandomNumber(4)
	if err != nil {
		return "", nil
	}

	// 保存到缓存
	err = domain.cache.SetUserVerifiedPhoneCode(uid, code)
	if err != nil {
		return "", err
	}

	return code, nil
}

// 验证邮箱
func (domain *userDomain) VerifiedPhone(uid int, vcode string) (bool, error) {
	// 缓存取出
	code, err := domain.cache.GetUserVerifiedPhoneCode(uid)
	if err != nil {
		return false, err
	}

	// 校验
	if vcode == code {
		return true, nil
	}

	return false, nil
}
