package user

import (
	"GoWebScaffold/infras/store/mysqlStore"
	"time"
)

const UserTableName = "user"

/*用户模块的持久化对象，代表user表的每行数据 */
type UserPO struct {
	ID       uint      `ddb:"id" json:"id"`
	UserNo   string    `ddb:"user_no" json:"user_no"`
	Name     string    `ddb:"name" json:"name"`
	Age      byte      `ddb:"age" json:"age"`
	Avatar   string    `ddb:"avatar" json:"avatar"`
	Gender   int8      `ddb:"gender" json:"gender"`
	Email    string    `ddb:"email" json:"email"`
	Phone    string    `ddb:"phone" json:"phone"`
	Password string    `ddb:"password" json:"password"`
	Salt     string    `ddb:"salt" json:"salt"`
	Status   int8      `ddb:"status" json:"status"`
	UpdateAt time.Time `ddb:"update_at" json:"update_at"`
	CreateAt time.Time `ddb:"create_at" json:"create_at"`
}

func (d *UserPO) TableName() string {
	return UserTableName
}

/* 数据访问层，实现具体数据持久化操作 */

type UserDAO struct {
	base      *mysqlStore.BaseDao
	userTable string
}

func NewUserDao() *UserDAO {
	dao := new(UserDAO)
	dao.base = mysqlStore.NewCommonMysqlStore()
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
	count, err := d.base.GetCount(UserTableName, where)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// 插入
func (d *UserDAO) Create(po UserPO) {
	// d.dao.Insert()
}

// 通过Id查找
func (d *UserDAO) GetById() {}
