package user

import (
	"GoWebScaffold/infras/store/ormStore"
)

/* 数据访问层，实现具体数据持久化操作 */

type UserDAO struct{}

func NewUserDao() *UserDAO {
	dao := new(UserDAO)
	return dao
}

// 查找用户名是否存在
func (d *UserDAO) IsUserNameExist(name string) (bool, error) {
	var where = map[string]interface{}{
		"name": name,
	}
	return d.isExist(where)
}

// 查找邮箱是否存在
func (d *UserDAO) IsEmailExist(email string) (bool, error) {
	var where = map[string]interface{}{
		"email": email,
	}
	return d.isExist(where)
}

// 查找手机号码是否存在
func (d *UserDAO) IsPhoneExist(phone string) (bool, error) {
	var where = map[string]interface{}{
		"phone": phone,
	}
	return d.isExist(where)
}

func (d *UserDAO) isExist(where map[string]interface{}) (bool, error) {
	var count int
	if err := ormStore.GormDb().Where(where).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// 插入
func (d *UserDAO) Create(po User) error {
	if err := ormStore.GormDb().Create(&po).Error; err != nil {
		return err
	}

	return nil
}

// 通过Id查找
func (d *UserDAO) GetById(id int) (*User, error) {
	var user User
	if err := ormStore.GormDb().Where(id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
