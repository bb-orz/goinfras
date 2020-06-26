package user

import (
	"GoWebScaffold/common"
	"GoWebScaffold/services"
	"github.com/segmentio/ksuid"
)

/*领域层：实现具体业务逻辑*/
type userDomain struct{}

func NewUserDomain() *userDomain {
	return new(userDomain)
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

// 创建用户
func (domain *userDomain) Create(dto services.UserDTO) (*services.UserDTO, error) {
	// 设置po
	userPO := UserPO{}
	userNo := domain.generateUserNo()
	encryptPassword, salt := domain.encryptPassword(dto.Password)

	// TODO 实例DAO，执行持久化操作
	userDao := UserDAO{}
	userDao.Create(userPO)

	return nil, nil
}
