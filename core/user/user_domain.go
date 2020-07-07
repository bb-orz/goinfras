package user

import (
	"GoWebScaffold/infras/global"
	"GoWebScaffold/services"
	"fmt"
	"github.com/segmentio/ksuid"
)

/*领域层：实现具体业务逻辑*/
type UserDomain struct {
	dao   *UserDAO
	cache *userCache
}

func NewUserDomain() *UserDomain {
	domain := new(UserDomain)
	domain.dao = NewUserDao()
	domain.cache = NewUserCache()
	return domain
}

// 生成用户编号
func (domain *UserDomain) generateUserNo() string {
	// 采用ksuid的ID生成策略来创建No
	// 全局唯一的ID
	return ksuid.New().Next().String()
}

// 加密密码，设置密文和盐值到po
func (domain *UserDomain) encryptPassword(password string) (hashStr, salt string) {
	hashStr, salt = global.HashPassword(password)
	return
}

// 查找用户是否已存在
func (domain *UserDomain) IsUserEmailExist(dto services.CreateUserWithEmailDTO) (bool, error) {
	if isExist, err := domain.dao.IsEmailExist(dto.Email); err != nil {
		fmt.Println("Error:", err)
		return false, err
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 邮箱账号创建用户
func (domain *UserDomain) CreateUserForEmail(dto services.CreateUserWithEmailDTO) (*services.UserDTO, error) {
	userModel := User{}
	userModel.Name = dto.Name
	userModel.Email = dto.Email
	userModel.No = domain.generateUserNo()
	userModel.Password, userModel.Salt = domain.encryptPassword(dto.Password)
	userModel.Status = UserStatusNotVerified // 初始创建时未验证状态
	fmt.Println("UserModel:", userModel)
	var user *User
	var err error
	if user, err = domain.dao.Create(userModel); err != nil {
		return nil, err
	}

	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfo(uid int) (*services.UserDTO, error) {
	user, err := domain.dao.GetById(uid)
	if err != nil {
		return nil, err
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

// 设置单个用户信息
func (domain *UserDomain) SetUserInfo(uid int, field string, value interface{}) error {
	return domain.dao.SetUserInfo(uid, field, value)
}

// 设置多个用户信息
func (domain *UserDomain) SetUserInfos(uid int, dto services.SetUserInfoDTO) error {
	return domain.dao.SetUserInfos(uid, dto)
}
