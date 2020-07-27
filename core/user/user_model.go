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
	UserOauths    []UserOauth `gorm:"many2many:user_oauth_binding;"` // Many-To-Many , 'user_oauth_binding'是连接表
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
	userDTO.Status = int8(model.Status)
	userDTO.Password = model.Password
	userDTO.Salt = model.Salt
	userDTO.CreatedAt = model.CreatedAt
	userDTO.UpdatedAt = model.UpdatedAt
	userDTO.DeletedAt = model.DeletedAt
	return &userDTO
}

func (model *User) ToOauthBindingDTO() *services.UserOauthInfoDTO {
	userDTO := services.UserOauthInfoDTO{}
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
	userDTO.User.Status = int8(model.Status)
	userDTO.User.Password = model.Password
	userDTO.User.Salt = model.Salt
	userDTO.User.CreatedAt = model.CreatedAt
	userDTO.User.UpdatedAt = model.UpdatedAt
	userDTO.User.DeletedAt = model.DeletedAt

	for _, item := range model.UserOauths {
		userDTO.UserOauths = append(userDTO.UserOauths, item.ToDTO())
	}

	return &userDTO
}

type UserOauth struct {
	gorm.Model
	Platform    uint   `gorm:"type:tinyint(1)"`
	AccessToken string `gorm:"type:varchar(64);unique"`
	OpenId      string `gorm:"type:varchar(20);unique"`
	UnionId     string `gorm:"type:varchar(20);unique"`
	NickName    string `gorm:"type:varchar(20)"`
	Gender      uint   `gorm:"type:tinyint(1)"`
	Avatar      string `gorm:"type:varchar(255)"`
}

func (model *UserOauth) ToDTO() *services.OauthInfoDTO {
	oauthDTO := services.OauthInfoDTO{}
	oauthDTO.Platform = model.Platform
	oauthDTO.OpenId = model.OpenId
	oauthDTO.UnionId = model.UnionId
	oauthDTO.AccessToken = model.AccessToken
	oauthDTO.NickName = model.NickName
	oauthDTO.Avatar = model.Avatar
	oauthDTO.Gender = model.Gender
	return &oauthDTO
}
