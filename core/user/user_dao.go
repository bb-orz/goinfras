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
	var err error
	var count int

	err = ormStore.GormDb().Where(uid).First(UserModel{}).Count(&count).Error
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
	return d.isExist(&UserModel{Name: name})
}

// 查找邮箱是否存在
func (d *userDAO) IsEmailExist(email string) (bool, error) {

	return d.isExist(&UserModel{Email: email})
}

// 查找手机号码是否存在
func (d *userDAO) IsPhoneExist(phone string) (bool, error) {

	return d.isExist(&UserModel{Phone: phone})
}

func (d *userDAO) isExist(where *UserModel) (bool, error) {
	var err error
	var count int
	err = ormStore.GormDb().Where(where).First(&UserModel{}).Count(&count).Error
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

// 插入单个用户信息
func (d *userDAO) Create(dto *services.UserDTO) (*services.UserDTO, error) {
	var err error
	var userDTO *services.UserDTO
	var userModel UserModel

	userModel.FromDTO(dto)
	if err = ormStore.GormDb().Create(&userModel).Error; err != nil {
		return nil, err
	}
	userDTO = userModel.ToDTO()
	return userDTO, nil
}

// 插入单个用户信息并关联三方平台账户
func (d *userDAO) CreateUserWithOAuth(dto *services.UserOAuthsDTO) (*services.UserOAuthsDTO, error) {
	var err error
	var userOAuthsDTO *services.UserOAuthsDTO
	var userOAuthsModel UserOAuthsModel

	userOAuthsModel.FromDTO(dto)
	if err = ormStore.GormDb().Create(&userOAuthsModel).Error; err != nil {
		return nil, err
	}

	userOAuthsDTO = userOAuthsModel.ToDTO()
	return userOAuthsDTO, nil
}

func (d *userDAO) GetUserOAuths(platform uint, openId, unionId string) (*services.UserOAuthsDTO, error) {
	var err error
	var oAuthResult OAuthModel
	var userResult UserModel
	var userOAuthDTO *services.UserOAuthsDTO
	var authDTOs []services.OAuthDTO

	if err = ormStore.GormDb().Where(&OAuthModel{Platform: platform, OpenId: openId, UnionId: unionId}).Find(&oAuthResult).Error; err != nil {
		return nil, err
	}

	if err = ormStore.GormDb().First(&userResult, oAuthResult.UserId).Error; err != nil {
		return nil, err
	}

	authDTOs = make([]services.OAuthDTO, 0)
	authDTOs = append(authDTOs, oAuthResult.ToDTO())

	userOAuthDTO = &services.UserOAuthsDTO{}
	userOAuthDTO.UserOAuths = authDTOs
	userOAuthDTO.User = *userResult.ToDTO()

	return userOAuthDTO, nil
}

// 通过Id查找
func (d *userDAO) GetById(id uint) (*services.UserDTO, error) {
	var err error
	var userResult UserModel
	err = ormStore.GormDb().Where(id).First(&userResult).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}
	dto := userResult.ToDTO()
	return dto, nil
}

// 通过邮箱账号查找
func (d *userDAO) GetByEmail(email string) (*services.UserDTO, error) {
	var err error
	var userResult UserModel
	err = ormStore.GormDb().Where(&UserModel{Email: email}).First(&userResult).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}

	dto := userResult.ToDTO()
	return dto, nil
}

// 通过邮箱账号查找
func (d *userDAO) GetByPhone(phone string) (*services.UserDTO, error) {
	var err error
	var userResult UserModel
	err = ormStore.GormDb().Where(&UserModel{Phone: phone}).First(&userResult).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}
	dto := userResult.ToDTO()
	return dto, nil
}

// 设置单个用户信息字段
func (d *userDAO) SetUserInfo(uid uint, field string, value interface{}) error {
	var err error
	if err = ormStore.GormDb().Model(&UserModel{}).Where("id", uid).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个用户信息字段
func (d *userDAO) SetUserInfos(uid uint, dto services.SetUserInfoDTO) error {
	var err error
	var updater UserModel
	updater.Name = dto.Name
	updater.Avatar = dto.Avatar
	updater.Age = dto.Age
	updater.Gender = dto.Gender
	updater.Status = dto.Status

	if err = ormStore.GormDb().Model(&UserModel{}).Where("id", uid).Updates(&updater).Error; err != nil {
		return err
	}
	return nil
}

// 设置用户密码和盐值
func (d *userDAO) SetPasswordAndSalt(uid uint, passHash, salt string) error {
	var err error
	if err = ormStore.GormDb().Model(&UserModel{}).Where("id", uid).Update(&UserModel{Password: passHash, Salt: salt}).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *userDAO) DeleteById(uid uint) error {
	var err error
	if err = ormStore.GormDb().Model(&UserModel{}).Delete(uid).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *userDAO) SetDeletedAtById(uid uint) error {
	var err error
	if err = ormStore.GormDb().Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(uid).Error; err != nil {
		return err
	}
	return nil
}
