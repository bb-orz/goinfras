package user

import (
	"GoWebScaffold/infras/store/ormStore"
	"GoWebScaffold/services"
	"github.com/jinzhu/gorm"
)

/*
数据访问层，实现具体数据持久化操作
直接返回error和执行结果
*/

type userDAO struct{}

func NewUserDAO() *userDAO {
	dao := new(userDAO)
	return dao
}

// 查找用户名是否存在
func (d *userDAO) IsUserIdExist(uid uint) (bool, error) {
	var count int
	err := ormStore.GormDb().Where(uid).First(&User{}).Count(&count).Error
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

// 查找用户名是否存在
func (d *userDAO) IsUserNameExist(name string) (bool, error) {
	return d.isExist(&User{Name: name})
}

// 查找邮箱是否存在
func (d *userDAO) IsEmailExist(email string) (bool, error) {

	return d.isExist(&User{Email: email})
}

// 查找手机号码是否存在
func (d *userDAO) IsPhoneExist(phone string) (bool, error) {

	return d.isExist(&User{Phone: phone})
}

func (d *userDAO) isExist(where *User) (bool, error) {
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
func (d *userDAO) Create(model User) (*User, error) {
	if err := ormStore.GormDb().Create(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

// 通过Id查找
func (d *userDAO) GetById(id uint) (*User, error) {
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
func (d *userDAO) GetByEmail(email string) (*User, error) {
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
func (d *userDAO) GetByPhone(phone string) (*User, error) {
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

// 设置单个用户信息字段
func (d *userDAO) SetUserInfo(uid uint, field string, value interface{}) error {
	if err := ormStore.GormDb().Model(&User{}).Where("id", uid).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个用户信息字段
func (d *userDAO) SetUserInfos(uid uint, dto services.SetUserInfoDTO) error {
	if err := ormStore.GormDb().Model(&User{}).Where("id", uid).Updates(dto).Error; err != nil {
		return err
	}
	return nil
}

// 设置用户密码和盐值
func (d *userDAO) SetPasswordAndSalt(uid uint, passHash, salt string) error {
	if err := ormStore.GormDb().Model(&User{}).Where("id", uid).Update(&User{Password: passHash, Salt: salt}).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *userDAO) DeleteById(uid uint) error {
	if err := ormStore.GormDb().Model(&User{}).Delete(uid).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *userDAO) SetDeletedAtById(uid uint) error {
	if err := ormStore.GormDb().Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(uid).Error; err != nil {
		return err
	}
	return nil
}
