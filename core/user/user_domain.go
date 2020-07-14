package user

import (
	"GoWebScaffold/core"
	"GoWebScaffold/infras/global"
	"GoWebScaffold/services"
	"github.com/segmentio/ksuid"
)

/*
User 领域层：实现用户相关具体业务逻辑
封装领域层的错误信息并返回给调用者
*/
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
	// 采用ksuid的ID生成策略来创建全局唯一的ID
	return ksuid.New().Next().String()
}

// 加密密码，设置密文和盐值
func (domain *UserDomain) encryptPassword(password string) (hashStr, salt string) {
	hashStr, salt = global.HashPassword(password)
	return
}

// 查找邮箱是否已存在
func (domain *UserDomain) IsEmailExist(dto services.CreateUserWithEmailDTO) (bool, error) {
	if isExist, err := domain.dao.IsEmailExist(dto.Email); err != nil {
		return false, core.WrapError(err, core.DomainErrorFormatSqlQuery, DomainName, "IsEmailExist")
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 查找手机用户是否已存在
func (domain *UserDomain) IsPhoneExist(dto services.CreateUserWithPhoneDTO) (bool, error) {
	if isExist, err := domain.dao.IsPhoneExist(dto.Phone); err != nil {
		return false, core.WrapError(err, core.DomainErrorFormatSqlQuery, DomainName, "IsPhoneExist")
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
	var user *User
	var err error
	if user, err = domain.dao.Create(userModel); err != nil {
		return nil, core.WrapError(err, core.DomainErrorFormatSqlInsert, DomainName, "Create")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

// 手机号码创建用户
func (domain *UserDomain) CreateUserForPhone(dto services.CreateUserWithPhoneDTO) (*services.UserDTO, error) {
	userModel := User{}
	userModel.Name = dto.Name
	userModel.Phone = dto.Phone
	userModel.No = domain.generateUserNo()
	userModel.Password, userModel.Salt = domain.encryptPassword(dto.Password)
	userModel.Status = UserStatusNotVerified // 初始创建时未验证状态
	var user *User
	var err error
	if user, err = domain.dao.Create(userModel); err != nil {
		return nil, core.WrapError(err, core.DomainErrorFormatSqlInsert, DomainName, "Create")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfo(uid int) (*services.UserDTO, error) {
	user, err := domain.dao.GetById(uid)
	if err != nil {
		return nil, core.WrapError(err, core.DomainErrorFormatSqlQuery, DomainName, "GetById")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfoByEmail(email string) (*services.UserDTO, error) {
	user, err := domain.dao.GetByEmail(email)
	if err != nil {
		return nil, core.WrapError(err, core.DomainErrorFormatSqlQuery, DomainName, "GetByEmail")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfoByPhone(phone string) (*services.UserDTO, error) {
	user, err := domain.dao.GetByPhone(phone)
	if err != nil {
		return nil, core.WrapError(err, core.DomainErrorFormatSqlQuery, DomainName, "GetByPhone")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

// 设置单个用户信息
func (domain *UserDomain) SetUserInfo(uid int, field string, value interface{}) error {
	err := domain.dao.SetUserInfo(uid, field, value)
	if err != nil {
		return core.WrapError(err, core.DomainErrorFormatSqlUpdate, DomainName, "SetUserInfo")
	}
	return nil
}

// 设置多个用户信息
func (domain *UserDomain) SetUserInfos(uid int, dto services.SetUserInfoDTO) error {
	err := domain.dao.SetUserInfos(uid, dto)
	if err != nil {
		return core.WrapError(err, core.DomainErrorFormatSqlUpdate, DomainName, "SetUserInfo")
	}
	return nil
}
