package XJwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var tku ITokenUtils

// 创建一个默认配置的TokenUtils
func CreateDefaultTku(config *Config) {
	if config == nil {
		config = DefaultConfig()
	}
	tku = NewTokenUtils(config)
}

type ITokenUtils interface {
	Encode(user UserClaim) (string, error)
	Decode(tokenString string) (*CustomerClaim, error)
	Remove(tokenString string) error
}

// 检查连接池实例
func CheckTku() bool {
	if tku != nil {
		return true
	}
	return false
}

// JWT中携带的用户个人信息
type UserClaim struct {
	Id     uint   `json:"id"`
	No     string `json:"no"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// 聚合jwt内部实现的Claims
type CustomerClaim struct {
	UserClaim
	*jwt.StandardClaims
}

// 实现token服务
type tokenUtils struct {
	privateKey []byte        // 编解码私钥，在生产环境中，该私钥请使用生成器生成，并妥善保管，此处使用简单字符串。
	expTime    time.Duration // 超时时间
}

func NewTokenUtils(config *Config) *tokenUtils {
	ts := new(tokenUtils)
	ts.privateKey = []byte(config.PrivateKey)
	ts.expTime = time.Second * time.Duration(config.ExpSeconds)
	return ts
}

// 传入用户信息编码成token
func (tks *tokenUtils) encode(user UserClaim) (string, error) {
	// privateKey, _ := base64.URLEncoding.DecodeString(string(privateKey))

	// 设置Claim
	espTime := time.Now().Add(tks.expTime).Unix()
	customer := CustomerClaim{user, &jwt.StandardClaims{ExpiresAt: espTime}}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customer)

	return token.SignedString(tks.privateKey)

}

// token字符串解码成用户信息
func (tks *tokenUtils) decode(tokenString string) (*CustomerClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaim{}, func(token *jwt.Token) (interface{}, error) {
		// return base64.URLEncoding.DecodeString(string(privateKey))
		return tks.privateKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*CustomerClaim); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}
}

func (tks *tokenUtils) Encode(user UserClaim) (string, error) {
	return tks.encode(user)
}

func (tks *tokenUtils) Decode(tokenString string) (*CustomerClaim, error) {
	return tks.decode(tokenString)
}

func (tks *tokenUtils) Remove(tokenString string) error {
	return errors.New("No Cache To Remove Token ")
}
