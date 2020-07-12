package user

import (
	"GoWebScaffold/infras/store/ormStore"
	"GoWebScaffold/services"
	"github.com/jinzhu/gorm"
)

/*
数据访问层，实现具体数据持久化操作

*/

type UserDAO struct{}

func NewUserDao() *UserDAO {
	dao := new(UserDAO)
	return dao
}

// 查找用户名是否存在
func (d *UserDAO) IsUserNameExist(name string) (bool, error) {
	return d.isExist(&User{Name: name})
}

// 查找邮箱是否存在
func (d *UserDAO) IsEmailExist(email string) (bool, error) {

	return d.isExist(&User{Email: email})
}

// 查找手机号码是否存在
func (d *UserDAO) IsPhoneExist(phone string) (bool, error) {

	return d.isExist(&User{Phone: phone})
}

func (d *UserDAO) isExist(where *User) (bool, error) {
	var count int
	err := ormStore.GormDb().Where(where).First(&User{}).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return false, nil
		} else {
			// 除无记录外的错误返回
			return false, err
		}
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 插入
func (d *UserDAO) Create(model User) (*User, error) {
	if err := ormStore.GormDb().Create(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

// 通过Id查找
func (d *UserDAO) GetById(id int) (*User, error) {
	var user User
	err := ormStore.GormDb().Where(id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}

	return &user, nil
}

// 通过邮箱账号查找
func (d *UserDAO) GetByEmail(email string) (*User, error) {
	var user User
	err := ormStore.GormDb().Where(&User{Email: email}).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}

	return &user, nil
}

// 通过邮箱账号查找
func (d *UserDAO) GetByPhone(phone string) (*User, error) {
	var user User
	err := ormStore.GormDb().Where(&User{Phone: phone}).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}

	return &user, nil
}

func (d *UserDAO) SetUserInfo(uid int, field string, value interface{}) error {
	if err := ormStore.GormDb().Model(&User{}).Where("id", uid).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

func (d *UserDAO) SetUserInfos(uid int, dto services.SetUserInfoDTO) error {
	if err := ormStore.GormDb().Model(&User{}).Where("id", uid).Updates(dto).Error; err != nil {
		return err
	}
	return nil
}
