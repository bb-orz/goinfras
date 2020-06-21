package user

import "time"

// 用户模块的持久化对象，代表user表的每行数据
type UserPO struct {
	ID       uint      `ddb:"id" json:"id"`
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
