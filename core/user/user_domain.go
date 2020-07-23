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
	dao   *userDAO
	cache *userCache
}

func NewUserDomain() *UserDomain {
	domain := new(UserDomain)
	domain.dao = NewUserDAO()
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

// 查找用户id是否已存在
func (domain *UserDomain) IsUserExist(uid uint) (bool, error) {
	if isExist, err := domain.dao.IsUserIdExist(uid); err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "IsEmailExist")
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 查找邮箱是否已存在
func (domain *UserDomain) IsEmailExist(email string) (bool, error) {
	if isExist, err := domain.dao.IsEmailExist(email); err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "IsEmailExist")
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 查找手机用户是否已存在
func (domain *UserDomain) IsPhoneExist(phone string) (bool, error) {
	if isExist, err := domain.dao.IsPhoneExist(phone); err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "IsPhoneExist")
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
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlInsert, DomainName, "Create")
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
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlInsert, DomainName, "Create")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfo(uid uint) (*services.UserDTO, error) {
	var user *User
	var err error
	if user, err = domain.dao.GetById(uid); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "GetById")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfoByEmail(email string) (*services.UserDTO, error) {
	var user *User
	var err error
	if user, err = domain.dao.GetByEmail(email); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "GetByEmail")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfoByPhone(phone string) (*services.UserDTO, error) {
	var user *User
	var err error
	if user, err = domain.dao.GetByPhone(phone); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "GetByPhone")
	}
	userDTO := user.ToDTO()
	return userDTO, nil
}

// 设置用户状态
func (domain *UserDomain) SetStatus(uid, status uint) error {
	if err := domain.dao.SetUserInfo(uid, "status", status); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetUserInfo")
	}
	return nil
}

// 设置单个用户信息
func (domain *UserDomain) SetUserInfo(uid uint, field string, value interface{}) error {
	if err := domain.dao.SetUserInfo(uid, field, value); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetUserInfo")
	}
	return nil
}

// 设置多个用户信息
func (domain *UserDomain) SetUserInfos(uid uint, dto services.SetUserInfoDTO) error {
	if err := domain.dao.SetUserInfos(uid, dto); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetUserInfo")
	}

	return nil
}

// 改变密码
func (domain *UserDomain) ReSetPassword(uid uint, password string) error {
	hashStr, salt := domain.encryptPassword(password)
	if err := domain.dao.SetPasswordAndSalt(uid, hashStr, salt); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetPasswordAndSalt")
	}
	return nil
}

// 真删除
func (domain *UserDomain) DeleteUser(uid uint) error {
	if err := domain.dao.DeleteById(uid); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlDelete, DomainName, "DeleteById")
	}
	return nil
}

// 伪删除
func (domain *UserDomain) ShamDeleteUser(uid uint) error {
	if err := domain.dao.SetDeletedAtById(uid); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlShamDelete, DomainName, "SetDeletedAtById")
	}
	return nil
}
