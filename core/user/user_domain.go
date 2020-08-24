package user

import (
	"GoWebScaffold/core"
	"GoWebScaffold/infras/global"
	"GoWebScaffold/infras/jwt"
	"GoWebScaffold/infras/oauth"
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

// 鉴权后生成token
func (domain *UserDomain) GenToken(no, name, avatar string) (string, error) {
	var err error
	var token string
	// 生成
	token, err = jwt.TokenUtils().Encode(jwt.UserClaim{
		Id:     no,
		Name:   name,
		Avatar: avatar,
	})
	if err != nil {
		return "", core.WrapError(err, core.ErrorFormatDomainAlgorithm, DomainName, "jwt.TokenUtils().Encode")
	}

	// 缓存
	err = domain.cache.SetUserToken(no, token)
	if err != nil {
		return "", core.WrapError(err, core.ErrorFormatDomainCacheSet, DomainName, "domain.cache.SetUserToken")
	}

	return token, nil
}

// 查找用户id是否已存在
func (domain *UserDomain) IsUserExist(uid uint) (bool, error) {
	var err error
	var isExist bool

	if isExist, err = domain.dao.IsUserIdExist(uid); err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "IsEmailExist")
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 查找邮箱是否已存在
func (domain *UserDomain) IsEmailExist(email string) (bool, error) {
	var err error
	var isExist bool
	if isExist, err = domain.dao.IsEmailExist(email); err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "IsEmailExist")
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 查找手机用户是否已存在
func (domain *UserDomain) IsPhoneExist(phone string) (bool, error) {
	var err error
	var isExist bool
	if isExist, err = domain.dao.IsPhoneExist(phone); err != nil {
		return false, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "IsPhoneExist")
	} else if isExist {
		return true, nil
	}
	return false, nil
}

// 邮箱账号创建用户
func (domain *UserDomain) CreateUserForEmail(dto services.CreateUserWithEmailDTO) (*services.UserDTO, error) {
	var err error
	var userDTO *services.UserDTO

	createUserData := services.UserDTO{}
	createUserData.Name = dto.Name
	createUserData.Email = dto.Email
	createUserData.No = domain.generateUserNo()
	createUserData.Password, createUserData.Salt = domain.encryptPassword(dto.Password)
	createUserData.Status = UserStatusNotVerified // 初始创建时未验证状态

	if userDTO, err = domain.dao.Create(&createUserData); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlInsert, DomainName, "Create")
	}
	return userDTO, nil
}

// 手机号码创建用户
func (domain *UserDomain) CreateUserForPhone(dto services.CreateUserWithPhoneDTO) (*services.UserDTO, error) {
	var err error
	var userDTO *services.UserDTO

	createUserData := services.UserDTO{}
	createUserData.Name = dto.Name
	createUserData.Phone = dto.Phone
	createUserData.No = domain.generateUserNo()
	createUserData.Password, createUserData.Salt = domain.encryptPassword(dto.Password)
	createUserData.Status = UserStatusNotVerified // 初始创建时未验证状态

	if userDTO, err = domain.dao.Create(&createUserData); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlInsert, DomainName, "Create")
	}
	return userDTO, nil
}

// Oauth三方账号绑定创建用户
func (domain *UserDomain) CreateUserOAuthBinding(platform uint, oauthInfo *oauth.OAuthAccountInfo) (*services.UserOAuthsDTO, error) {
	var err error
	var userOAuthsResult *services.UserOAuthsDTO

	// 插入用户信息
	createUserData := services.UserOAuthsDTO{}
	createUserData.User.Name = oauthInfo.NickName
	createUserData.User.No = domain.generateUserNo()
	createUserData.User.Status = UserStatusNotVerified // 初始创建时未验证状态
	createUserData.UserOAuths = []services.OAuthDTO{
		{
			AccessToken: oauthInfo.AccessToken,
			UnionId:     oauthInfo.UnionId,
			OpenId:      oauthInfo.OpenId,
			NickName:    oauthInfo.NickName,
			Avatar:      oauthInfo.AvatarUrl,
			Gender:      oauthInfo.Gender,
			Platform:    platform,
		},
	}

	if userOAuthsResult, err = domain.dao.CreateUserWithOAuth(&createUserData); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlInsert, DomainName, "CreateUserWithOAuth")
	}

	return userOAuthsResult, nil
}

// 获取整个关联的用户信息和三方平台绑定信息
func (domain *UserDomain) GetUserOauths(platform uint, openId, unionId string) (*services.UserOAuthsDTO, error) {
	var err error
	var userOAuthsResult *services.UserOAuthsDTO

	if userOAuthsResult, err = domain.dao.GetUserOAuths(platform, openId, unionId); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "GetUserOAuths")
	}

	return userOAuthsResult, nil
}

func (domain *UserDomain) GetUserInfo(uid uint) (*services.UserDTO, error) {
	var err error
	var userDTO *services.UserDTO
	if userDTO, err = domain.dao.GetById(uid); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "GetById")
	}
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfoByEmail(email string) (*services.UserDTO, error) {
	var err error
	var userDTO *services.UserDTO
	if userDTO, err = domain.dao.GetByEmail(email); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "GetByEmail")
	}
	return userDTO, nil
}

func (domain *UserDomain) GetUserInfoByPhone(phone string) (*services.UserDTO, error) {
	var err error
	var userDTO *services.UserDTO
	if userDTO, err = domain.dao.GetByPhone(phone); err != nil {
		return nil, core.WrapError(err, core.ErrorFormatDomainSqlQuery, DomainName, "GetByPhone")
	}
	return userDTO, nil
}

// 设置用户状态
func (domain *UserDomain) SetStatus(uid, status uint) error {
	var err error
	if err = domain.dao.SetUserInfo(uid, "status", status); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetUserInfo")
	}
	return nil
}

// 设置单个用户信息
func (domain *UserDomain) SetUserInfo(uid uint, field string, value interface{}) error {
	var err error
	if err = domain.dao.SetUserInfo(uid, field, value); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetUserInfo")
	}
	return nil
}

// 设置多个用户信息
func (domain *UserDomain) SetUserInfos(uid uint, dto services.SetUserInfoDTO) error {
	var err error
	if err = domain.dao.SetUserInfos(uid, dto); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetUserInfo")
	}

	return nil
}

// 改变密码
func (domain *UserDomain) ReSetPassword(uid uint, password string) error {
	var err error
	var hashStr, salt string
	hashStr, salt = domain.encryptPassword(password)
	if err = domain.dao.SetPasswordAndSalt(uid, hashStr, salt); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlUpdate, DomainName, "SetPasswordAndSalt")
	}
	return nil
}

// 真删除
func (domain *UserDomain) DeleteUser(uid uint) error {
	var err error
	if err = domain.dao.DeleteById(uid); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlDelete, DomainName, "DeleteById")
	}
	return nil
}

// 伪删除
func (domain *UserDomain) ShamDeleteUser(uid uint) error {
	var err error
	if err = domain.dao.SetDeletedAtById(uid); err != nil {
		return core.WrapError(err, core.ErrorFormatDomainSqlShamDelete, DomainName, "SetDeletedAtById")
	}
	return nil
}
