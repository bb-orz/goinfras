package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// 生成盐值
func RandomString(l int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyz"
	var rendomString []byte
	bytes := []byte(chars)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		rendomString = append(rendomString, bytes[r.Intn(len(bytes))])
	}
	return string(rendomString)
}

// 用户密码加盐生成哈希
func HashPassword(src string) (hashStr, salt string) {
	// 获取随机盐值字符串
	salt = RandomString(4)

	hash := sha1.New()

	_, _ = io.WriteString(hash, src)
	_, _ = io.WriteString(hash, salt)

	hashBytes := hash.Sum(nil)
	// 组合输出40位哈希字符
	hashStr = fmt.Sprintf("%x", hashBytes)
	return
}

// 校验密码
func ValidatePassword(passStr, salt, passHash string) bool {
	// 重新计算密码哈希，与之前的校验
	hash := sha1.New()
	_, _ = io.WriteString(hash, passStr)
	_, _ = io.WriteString(hash, salt)
	hashBytes := hash.Sum(nil)
	// 组合输出40位哈希字符
	hashStr := fmt.Sprintf("%x", hashBytes)

	return hashStr == passHash
}
