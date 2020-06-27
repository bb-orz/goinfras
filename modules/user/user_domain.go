package user

import (
	"GoWebScaffold/common"
	"GoWebScaffold/services"
	"github.com/segmentio/ksuid"
)

/*领域层：实现具体业务逻辑*/
type userDomain struct {
	userDao *UserDAO
}

func NewUserDomain() *userDomain {
	domain := new(userDomain)
	domain.userDao = NewUserDao()
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
func (domain *userDomain) IsUserExist(dto services.CreateUserDTO) (bool, error) {
	if isExist, err := domain.userDao.isExist(&User{Email: dto.Email}); err != nil {
		return false, err
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 创建用户
func (domain *userDomain) CreateUser(dto services.UserDTO) (*services.UserDTO, error) {
	userModel := User{}
	userModel.FromDTO(&dto)
	var user *User
	var err error
	if user, err = domain.userDao.Create(userModel); err != nil {
		return nil, err
	}

	userDTO := user.ToDTO()
	return userDTO, nil
}
