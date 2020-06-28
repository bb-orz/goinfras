package user

import (
	"GoWebScaffold/services"
	"github.com/jinzhu/gorm"
)

/*用户模块的持久化对象，代表user表的每行数据 */
type User struct {
	gorm.Model
	Num        string      `gorm:"type:char(20);unique_index"` // index是为该列创建索引
	Name       string      `gorm:"type:char(12);index;not null"`
	Age        uint8       `gorm:"type:tinyint(3)"`
	Avatar     string      `gorm:"type:varchar(255)"`
	Gender     uint8       `gorm:"type:tinyint(1)"`
	Email      string      `gorm:"type:varchar(100);unique_index"`
	Phone      string      `gorm:"type:char(11)"`
	Password   string      `gorm:"type:char(40)"`
	Salt       string      `gorm:"type:char(4)"`
	Status     uint8       `gorm:"type:tinyint(1)"`
	UserOauths []UserOauth `gorm:"many2many:user_oauth_binding;"` // Many-To-Many , 'user_oauth_binding'是连接表
}

func (model *User) FromDTO(dto *services.UserDTO) {
	model.Email = dto.Email
	model.Name = dto.Name
	model.Age = dto.Age
	model.Avatar = dto.Avatar
	model.Gender = dto.Gender
	model.Email = dto.Email
	model.Phone = dto.Phone
	model.CreatedAt = dto.CreatedAt
	model.UpdatedAt = dto.UpdatedAt
}

func (model *User) ToDTO() *services.UserDTO {
	userDTO := services.UserDTO{}
	userDTO.Email = model.Email
	userDTO.Name = model.Name
	userDTO.Age = model.Age
	userDTO.Avatar = model.Avatar
	userDTO.Gender = model.Gender
	userDTO.Email = model.Email
	userDTO.Phone = model.Phone
	userDTO.CreatedAt = model.CreatedAt
	userDTO.UpdatedAt = model.UpdatedAt
	return &userDTO
}

type UserOauth struct {
	gorm.Model
	Platform    uint8  `gorm:"type:tinyint(1)"`
	AccessToken string `gorm:"type:varchar(64);unique"`
	OpenId      string `gorm:"type:varchar(20);unique"`
	UnionId     string `gorm:"type:varchar(20);unique"`
	NickName    string `gorm:"type:varchar(20)"`
	Gender      uint8  `gorm:"type:tinyint(1)"`
	Avatar      string `gorm:"type:varchar(255)"`
}