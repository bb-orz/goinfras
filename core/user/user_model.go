package user

import (
	"GoWebScaffold/services"
	"github.com/jinzhu/gorm"
)

/*用户模块的持久化对象，代表user表的每行数据 */
type User struct {
	gorm.Model
	No            string      `gorm:"type:char(40);unique_index"` // index是为该列创建索引
	Name          string      `gorm:"type:char(12);index;not null"`
	Age           uint        `gorm:"type:tinyint"`
	Avatar        string      `gorm:"type:varchar(255)"`
	Gender        uint        `gorm:"type:tinyint"`
	Email         string      `gorm:"type:varchar(100);unique_index"`
	EmailVerified bool        `gorm:"type:tinyint(1)"`
	Phone         string      `gorm:"type:char(11)"`
	PhoneVerified bool        `gorm:"type:tinyint(1)"`
	Password      string      `gorm:"type:char(40)"`
	Salt          string      `gorm:"type:char(4)"`
	Status        uint        `gorm:"type:tinyint"`
	UserOAuths    []UserOAuth `gorm:"many2many:user_oauth_binding;"` // Many-To-Many , 'user_oauth_binding'是连接表
}

func (User) TableName() string {
	return "user"
}

func (model *User) FromDTO(dto *services.UserDTO) {
	model.ID = dto.Uid
	model.No = dto.No
	model.Name = dto.Name
	model.Age = dto.Age
	model.Avatar = dto.Avatar
	model.Gender = dto.Gender
	model.Email = dto.Email
	model.EmailVerified = dto.EmailVerified
	model.Phone = dto.Phone
	model.PhoneVerified = dto.PhoneVerified
	model.Status = uint(dto.Status)
	model.Password = dto.Password
	model.Salt = dto.Salt
	model.CreatedAt = dto.CreatedAt
	model.UpdatedAt = dto.UpdatedAt
	model.DeletedAt = dto.DeletedAt
}

func (model *User) ToDTO() *services.UserDTO {
	userDTO := services.UserDTO{}
	userDTO.Uid = model.ID
	userDTO.No = model.No
	userDTO.Name = model.Name
	userDTO.Age = model.Age
	userDTO.Avatar = model.Avatar
	userDTO.Gender = model.Gender
	userDTO.Email = model.Email
	userDTO.EmailVerified = model.EmailVerified
	userDTO.Phone = model.Phone
	userDTO.PhoneVerified = model.PhoneVerified
	userDTO.Status = model.Status
	userDTO.Password = model.Password
	userDTO.Salt = model.Salt
	userDTO.CreatedAt = model.CreatedAt
	userDTO.UpdatedAt = model.UpdatedAt
	userDTO.DeletedAt = model.DeletedAt
	return &userDTO
}

func (model *User) FromOAuthBindingDTO(dto *services.UserOAuthInfoDTO) {
	model.ID = dto.User.Uid
	model.No = dto.User.No
	model.Name = dto.User.Name
	model.Age = dto.User.Age
	model.Avatar = dto.User.Avatar
	model.Gender = dto.User.Gender
	model.Email = dto.User.Email
	model.EmailVerified = dto.User.EmailVerified
	model.Phone = dto.User.Phone
	model.PhoneVerified = dto.User.PhoneVerified
	model.Status = uint(dto.User.Status)
	model.Password = dto.User.Password
	model.Salt = dto.User.Salt
	model.CreatedAt = dto.User.CreatedAt
	model.UpdatedAt = dto.User.UpdatedAt
	model.DeletedAt = dto.User.DeletedAt

	for _, item := range dto.UserOAuths {
		authModel := UserOAuth{}
		authModel.FromDTO(item)
		model.UserOAuths = append(model.UserOAuths, authModel)
	}
}

func (model *User) ToOAuthBindingDTO() *services.UserOAuthInfoDTO {
	userDTO := services.UserOAuthInfoDTO{}
	userDTO.User.Uid = model.ID
	userDTO.User.No = model.No
	userDTO.User.Name = model.Name
	userDTO.User.Age = model.Age
	userDTO.User.Avatar = model.Avatar
	userDTO.User.Gender = model.Gender
	userDTO.User.Email = model.Email
	userDTO.User.EmailVerified = model.EmailVerified
	userDTO.User.Phone = model.Phone
	userDTO.User.PhoneVerified = model.PhoneVerified
	userDTO.User.Status = model.Status
	userDTO.User.Password = model.Password
	userDTO.User.Salt = model.Salt
	userDTO.User.CreatedAt = model.CreatedAt
	userDTO.User.UpdatedAt = model.UpdatedAt
	userDTO.User.DeletedAt = model.DeletedAt

	for _, item := range model.UserOAuths {
		userDTO.UserOAuths = append(userDTO.UserOAuths, item.ToDTO())
	}

	return &userDTO
}

type UserOAuth struct {
	gorm.Model
	Platform    uint   `gorm:"type:tinyint(1)"`
	AccessToken string `gorm:"type:varchar(64);unique"`
	OpenId      string `gorm:"type:varchar(20);unique"`
	UnionId     string `gorm:"type:varchar(20);unique"`
	NickName    string `gorm:"type:varchar(20)"`
	Gender      uint   `gorm:"type:tinyint(1)"`
	Avatar      string `gorm:"type:varchar(255)"`
}

func (model *UserOAuth) ToDTO() services.OAuthInfoDTO {
	oauthDTO := services.OAuthInfoDTO{}
	oauthDTO.Platform = model.Platform
	oauthDTO.OpenId = model.OpenId
	oauthDTO.UnionId = model.UnionId
	oauthDTO.AccessToken = model.AccessToken
	oauthDTO.NickName = model.NickName
	oauthDTO.Avatar = model.Avatar
	oauthDTO.Gender = model.Gender
	return oauthDTO
}

func (model *UserOAuth) FromDTO(dto services.OAuthInfoDTO) {
	model.Platform = dto.Platform
	model.NickName = dto.NickName
	model.OpenId = dto.OpenId
	model.UnionId = dto.UnionId
	model.AccessToken = dto.AccessToken
	model.Gender = dto.Gender
	model.Avatar = dto.Avatar
}
