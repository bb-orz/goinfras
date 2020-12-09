# XSQLBuilder Starter

> 基于 https://github.com/didi/gendry 包

### didi/gendry Documentation

> Documentation https://github.com/didi/gendry/blob/master/translation/zhcn/README.md


### XSQLBuilder Starter Usage
```
goinfras.RegisterStarter(XSQLBuilder.NewStarter())

```

### XSQLBuilder Config Setting

```
DbHost                  string // 主机地址
DbPort                  int64  // 主机端口
DbUser                  string // 用户名
DbPasswd                string // 密码
DbName                  string // 数据库名
ConnMaxLifetime         int64  // 每个连接最长生命周期，单位秒
MaxIdleConns            int64  // 连接池最大闲置连接数
MaxOpenConns            int64  // 连接池最大连接数
ChartSet                string // 传输字符编码
AllowCleartextPasswords bool   // 开发环境中设置允许明文密码通信
InterpolateParams       bool   // 设置允许占位符参数
Timeout                 int64  // 连接请求的超时时间，单位秒
ReadTimeout             int64  // 读超时时间，单位秒
ParseTime               bool   // 将数据库的datetime时间格式转换为go time包数据类型
PING                    bool   // 连接时PING测试
```

### XSQLBuilder Usage

1、定义UserSchema


```
// 新增
lastedId, err := XSQLBuilder.XCommon().Insert("user", []map[string]interface{}{
    {"name": "aaaa", "age": 18, "gender": 1}, {"name": "bbbb", "age": 20, "gender": 0},
})
So(err, ShouldBeNil)
Println("Lasted Insert Id:", lastedId)

// 获取数量
count, err := XSQLBuilder.XCommon().GetCount("user", nil)
So(err, ShouldBeNil)
Println("User Count:", count)

// 查询数据到Struct
type UserSchema struct {
	Id     int    `ddb:"id"`
	Name   string `ddb:"name"`
	Age    int    `ddb:"age"`
	Gender int    `ddb:"gender"`
}
rs := UserSchema{}
err = XSQLBuilder.XCommon().GetOne("user", map[string]interface{}{"name": "joker"}, nil, &rs)
So(err, ShouldBeNil)
Println("GetOne:", rs)

// 查询多个
rsList := make([]interface{}, 0)
XSQLBuilder.XCommon().GetMulti("user", map[string]interface{}{"name": "aaaa"}, nil, rsList)
So(err, ShouldBeNil)
Println("GetMulti:", rsList)

// 更新
update, err := XSQLBuilder.XCommon().Update("user", map[string]interface{}{"age": 18}, map[string]interface{}{"age": 28})
So(err, ShouldBeNil)
Println("Update Lasted Id:", update)

// 删除
deleteId, err := XSQLBuilder.XCommon().Delete("user", map[string]interface{}{"name": "ken"})
So(err, ShouldBeNil)
Println("Delete Id:", deleteId)
```