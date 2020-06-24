package user

import (
	"GoWebScaffold/services"
	"github.com/segmentio/ksuid"
)

/*领域层：实现具体业务逻辑*/
type userDomain struct {
	user UserPO // 持有持久化对象
}

func NewUserDomain() *userDomain {
	return new(userDomain)
}

// 判断该用户是否已经存在
func (domain *userDomain) IsUserExist(dto services.CreateUserDTO) bool {

	return false
}

// 生成用户编号
func (domain *userDomain) generateUserNo() {
	// 采用ksuid的ID生成策略来创建No
	// 全局唯一的ID
	domain.user.UserNo = ksuid.New().Next().String()
}

// 加密密码，设置密文和盐值到po
func (domain *userDomain) encryptPassword(password string) {

}

// 创建用户
func (domain *userDomain) Create(dto services.UserDTO) (*services.UserDTO, error) {
	// 设置po
	domain.user = UserPO{}
	domain.generateUserNo()
	domain.encryptPassword(dto.Password)

	// TODO 实例DAO，执行持久化操作
	userDao := UserDAO{}
	userDao.Create(domain.user)

	return nil, nil
}
