package user

import (
	"GoWebScaffold/services"
	"github.com/jinzhu/gorm"
)

/* User 单表模型 */
type UserModel struct {
	gorm.Model
	No            string `gorm:"type:char(40);unique_index"` // index是为该列创建索引
	Name          string `gorm:"type:char(12);index;not null"`
	Age           uint   `gorm:"type:tinyint"`
	Avatar        string `gorm:"type:varchar(255)"`
	Gender        uint   `gorm:"type:tinyint"`
	Email         string `gorm:"type:varchar(100);unique_index"`
	EmailVerified bool   `gorm:"type:tinyint(1)"`
	Phone         string `gorm:"type:char(11)"`
	PhoneVerified bool   `gorm:"type:tinyint(1)"`
	Password      string `gorm:"type:char(40)"`
	Salt          string `gorm:"type:char(4)"`
	Status        uint   `gorm:"type:tinyint"`
}

func (UserModel) TableName() string {
	return "user"
}

// DTO => Model
func (model *UserModel) FromDTO(dto *services.UserDTO) {
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

// Model => DTO
func (model *UserModel) ToDTO() *services.UserDTO {
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

// OAuth 第三方账号关联表
type OAuthModel struct {
	gorm.Model
	UserId      uint   `gorm:"foreignkey:UserId"`
	Platform    uint   `gorm:"type:tinyint(1)"`
	AccessToken string `gorm:"type:varchar(64);unique"`
	OpenId      string `gorm:"type:varchar(20);unique"`
	UnionId     string `gorm:"type:varchar(20);unique"`
	NickName    string `gorm:"type:varchar(20)"`
	Gender      uint   `gorm:"type:tinyint(1)"`
	Avatar      string `gorm:"type:varchar(255)"`
}

// DTO => Model
func (model *OAuthModel) FromDTO(dto services.OAuthDTO) {
	model.UserId = dto.UserId
	model.Platform = dto.Platform
	model.NickName = dto.NickName
	model.OpenId = dto.OpenId
	model.UnionId = dto.UnionId
	model.AccessToken = dto.AccessToken
	model.Gender = dto.Gender
	model.Avatar = dto.Avatar
}

// Model => DTO
func (model *OAuthModel) ToDTO() services.OAuthDTO {
	oauthDTO := services.OAuthDTO{}
	oauthDTO.UserId = model.UserId
	oauthDTO.Platform = model.Platform
	oauthDTO.OpenId = model.OpenId
	oauthDTO.UnionId = model.UnionId
	oauthDTO.AccessToken = model.AccessToken
	oauthDTO.NickName = model.NickName
	oauthDTO.Avatar = model.Avatar
	oauthDTO.Gender = model.Gender
	return oauthDTO
}

// 包含UserOAuths关联模型
type UserOAuthsModel struct {
	gorm.Model
	No            string       `gorm:"type:char(40);unique_index"` // index是为该列创建索引
	Name          string       `gorm:"type:char(12);index;not null"`
	Age           uint         `gorm:"type:tinyint"`
	Avatar        string       `gorm:"type:varchar(255)"`
	Gender        uint         `gorm:"type:tinyint"`
	Email         string       `gorm:"type:varchar(100);unique_index"`
	EmailVerified bool         `gorm:"type:tinyint(1)"`
	Phone         string       `gorm:"type:char(11)"`
	PhoneVerified bool         `gorm:"type:tinyint(1)"`
	Password      string       `gorm:"type:char(40)"`
	Salt          string       `gorm:"type:char(4)"`
	Status        uint         `gorm:"type:tinyint"`
	OAuths        []OAuthModel // Has-Many
}

func (model *UserOAuthsModel) FromDTO(dto *services.UserOAuthsDTO) {
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
	oAuths := make([]OAuthModel, 0)
	for _, item := range dto.UserOAuths {
		oAuthModel := OAuthModel{}
		oAuthModel.FromDTO(item)
		oAuths = append(oAuths, oAuthModel)
	}
	model.OAuths = oAuths
}

func (model *UserOAuthsModel) ToDTO() *services.UserOAuthsDTO {
	userOAuthsDTO := services.UserOAuthsDTO{}
	userOAuthsDTO.User.Uid = model.ID
	userOAuthsDTO.User.No = model.No
	userOAuthsDTO.User.Name = model.Name
	userOAuthsDTO.User.Age = model.Age
	userOAuthsDTO.User.Avatar = model.Avatar
	userOAuthsDTO.User.Gender = model.Gender
	userOAuthsDTO.User.Email = model.Email
	userOAuthsDTO.User.EmailVerified = model.EmailVerified
	userOAuthsDTO.User.Phone = model.Phone
	userOAuthsDTO.User.PhoneVerified = model.PhoneVerified
	userOAuthsDTO.User.Status = model.Status
	userOAuthsDTO.User.Password = model.Password
	userOAuthsDTO.User.Salt = model.Salt
	userOAuthsDTO.User.CreatedAt = model.CreatedAt
	userOAuthsDTO.User.UpdatedAt = model.UpdatedAt
	userOAuthsDTO.User.DeletedAt = model.DeletedAt
	oAuths := make([]services.OAuthDTO, 0)
	for _, item := range model.OAuths {
		oAuthDTO := item.ToDTO()
		oAuths = append(oAuths, oAuthDTO)
	}
	userOAuthsDTO.UserOAuths = oAuths
	return &userOAuthsDTO

}
