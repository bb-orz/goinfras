### Global 全局配置信息及工具函数


一些工具函数
```
// 幂运算
func Powerf(x float64, n int) float64 {}

// 生成固定位随机数字
func RandomNumber(l int) (string, error) {}

// 生成固定长度随机字符串
func RandomString(l int) string {}

// 用户密码加盐生成哈希
func HashPassword(src string) (hashStr, salt string) {}

// 校验密码
func ValidatePassword(passStr, salt, passHash string) bool {}


// TODO  
```