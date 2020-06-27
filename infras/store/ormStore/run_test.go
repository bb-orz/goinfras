package ormStore

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/props/kvs"
	"testing"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增

	CreditCard CreditCard // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails     []Email    // One-To-Many (拥有多个 - Email表的UserID作外键)

	BillingAddress   Address // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // 忽略这个字段
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

func newOrmDb(t *testing.T) {

	config := ormConfig{}
	p := kvs.NewEmptyCompositeConfigSource()
	err := p.Unmarshal(&config)
	So(err, ShouldBeNil)
	Println("ORM Config:", config)

	gormDb, err = NewORMDb(&config)
	So(err, ShouldBeNil)

}

func TestGormDb(t *testing.T) {
	Convey("测试使用gorm：", t, func() {
		newOrmDb(t)
		// 检查模型`Address`表是否存在
		hasAddressTable := gormDb.HasTable(&Address{})
		Println("Address Table Is Exist:", hasAddressTable)
		// 表不存在则创建表
		if !hasAddressTable {
			gormDb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Address{})
		}

		// 检查模型`Language`表是否存在
		hasLanguageTable := gormDb.HasTable(&Language{})
		Println("Language Table Is Exist:", hasLanguageTable)
		// 表不存在则创建表
		if !hasLanguageTable {
			gormDb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Language{})
		}

		// 检查模型`User`表是否存在
		hasUserTable := gormDb.HasTable(&User{})
		Println("User Table Is Exist:", hasUserTable)
		// 表不存在则创建表
		if !hasUserTable {
			gormDb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
		}

	})
}

func TestGormInsert(t *testing.T) {
	Convey("测试使用 Gorm Insert：", t, func() {
		newOrmDb(t)

		// 插入
		user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		GormDb().Create(&user)

		// 查询
		GormDb().First(&user)
		Println(user)
	})
}

func TestGormFind(t *testing.T) {
	Convey("测试使用 Gorm Find：", t, func() {
		newOrmDb(t)
		var user User
		var users []User
		user = User{}
		// 查询
		// 获取第一条记录，按主键排序
		GormDb().First(&user)
		Println("First:", user)
		// SELECT * FROM users ORDER BY id LIMIT 1;

		// 获取最后一条记录，按主键排序
		user = User{}
		GormDb().Last(&user)
		Println("Last:", user)
		// SELECT * FROM users ORDER BY id DESC LIMIT 1;

		// 获取所有记录
		users = make([]User, 0)
		GormDb().Find(&users)
		Println("Find More:", users)
		// SELECT * FROM users;

		// 使用主键获取记录
		users = make([]User, 0)
		GormDb().First(&user, 10)
		Println("Find By Key:", user)
		// SELECT * FROM users WHERE id = 10;
	})
}

func TestGormSimpleWhere(t *testing.T) {
	Convey("测试使用 Gorm Simple Where：", t, func() {
		newOrmDb(t)
		var user User
		var users []User
		// 获取第一个匹配记录
		GormDb().Where("name = ?", "jinzhu").First(&user)
		// SELECT * FROM users WHERE name = 'jinzhu' limit 1;
		Println("Find By Name:", user)

		// 获取所有匹配记录
		users = nil
		GormDb().Where("name = ?", "jinzhu").Find(&users)
		// SELECT * FROM users WHERE name = 'jinzhu';
		Println("Find All By Name:", users)

		users = nil
		GormDb().Where("name <> ?", "jinzhu").Find(&users)
		Println("Find NOT By Name:", users)

		// IN
		users = nil
		GormDb().Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
		Println("Find IN By Key:", users)

		// LIKE
		users = nil
		GormDb().Where("name LIKE ?", "%jin%").Find(&users)
		Println("Find LIKE By Name:", users)

		// AND
		users = nil
		GormDb().Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
		Println("Find And By Name:", users)

		// Time
		lastWeek := time.Now().Unix() - 7*24*60*60
		now := time.Now().Unix()
		users = nil
		GormDb().Where("updated_at > ?", lastWeek).Find(&users)
		Println("Find TimeGt By UpdateAt:", users)

		users = nil
		GormDb().Where("created_at BETWEEN ? AND ?", lastWeek, now).Find(&users)
		Println("Find TimeBETWEEN By UpdateAt:", users)

	})
}

func TestGormWhereByStructOrMap(t *testing.T) {
	Convey("测试使用 Gorm Where By struct Or Map：", t, func() {
		newOrmDb(t)
		var user User
		var users []User

		// Struct
		GormDb().Where(&User{Name: "jinzhu", Age: 18}).First(&user)
		// SELECT * FROM users WHERE name = "jinzhu" AND age = 18 LIMIT 1;
		Println("Find Where By Struct:", user)

		// Map
		users = make([]User, 0)
		GormDb().Where(map[string]interface{}{"name": "jinzhu", "age": 18}).Find(&users)
		// SELECT * FROM users WHERE name = "jinzhu" AND age = 18;
		Println("Find Where By Map:", users)

		// 主键的Slice
		users = nil
		GormDb().Where([]int64{20, 21, 22}).Find(&users)
		// SELECT * FROM users WHERE id IN (20, 21, 22);
		Println("Find LIKE By slice:", users)

	})

}

// 更多CURD 查找http://gorm.book.jasperxu.com/crud.html
